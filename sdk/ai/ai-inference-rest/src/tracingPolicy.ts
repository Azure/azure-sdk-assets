// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

import { PipelinePolicy, PipelineRequest, PipelineResponse, SendRequest } from "@azure/core-rest-pipeline";
import {
  type TracingClient,
  type TracingContext,
  type TracingSpan,
  createTracingClient,
} from "@azure/core-tracing";


import { getErrorMessage, isError } from "@azure/core-util";
import { SDK_VERSION } from "./constants.js";
import { logger } from "./logger.js";

/**
 * The programmatic identifier of the tracingPolicy.
 */
export const tracingPolicyName = "Inference_Tracing";

const tracingClient = tryCreateTracingClient();

/**
 * Options to configure the tracing policy.
 */
export interface TracingPolicyOptions {
  /**
   * Query string names whose values will be logged when logging is enabled. By default no
   * query string values are logged.
   */
  additionalAllowedQueryParameters?: string[];
}


export enum TracingAttributesEnum {
  Operation_Name = "gen_ai.operation.name",
  Request_Model = "gen_ai.request.model",
  System = "gen_ai.system",
  Error_Type = "error.type",
  Server_Port = "server.port", // not use it for now
  Request_Frequency_Penalty = "gen_ai.request.frequency_penalty",
  Request_Max_Tokens = "gen_ai.request.max_tokens",
  Request_Presence_Penalty = "gen_ai.request.presence_penalty",
  Request_Stop_Sequences = "gen_ai.request.stop_sequences",
  Request_Temperature = "gen_ai.request.temperature",
  Request_Top_K = "gen_ai.request.top_k",
  Request_Top_P = "gen_ai.request.top_p",
  Response_Finish_Reasons = "gen_ai.response.finish_reasons",
  Response_Id = "gen_ai.response.id",
  Response_Model = "gen_ai.response.model",
  Usage_Input_Tokens = "gen_ai.usage.input_tokens",
  Usage_Output_Tokens = "gen_ai.usage.output_tokens",
  Server_Address = "server.address"
}

export type TracingAttributes = { [key in TracingAttributesEnum]?: string | number };

/**
 * A simple policy to create OpenTelemetry Spans for each request made by the pipeline
 * that has SpanOptions with a parent.
 * Requests made without a parent Span will not be recorded.
 */
export function tracingPolicy(): PipelinePolicy {

  return {
    name: tracingPolicyName,
    async sendRequest(request: PipelineRequest, next: SendRequest): Promise<PipelineResponse> {
      if (!tracingClient) {
        return next(request);
      }

      /* sample of request
      request = {
        url: "https://models.inference.ai.azure.com/chat/completions?api-version=2024-05-01-preview",
        body: "{\"messages\":[{\"role\":\"system\",\"content\":\"You are a helpful assistant.\"},{\"role\":\"user\",\"content\":\"What is the capital of France?\"}],\"model\":\"Phi-3-medium-128k-instruct\",\"temperature\":1,\"max_tokens\":1000,\"top_p\":1}",
        headers: {
          _headersMap: {
          },
        },
        method: "POST",
        timeout: 0,
        multipartBody: undefined,
        formData: undefined,
        disableKeepAlive: false,
        proxySettings: undefined,
        streamResponseStatusCodes: undefined,
        withCredentials: false,
        abortSignal: undefined,
        tracingOptions: {
        },
        onUploadProgress: undefined,
        onDownloadProgress: undefined,
        requestId: "e89d3930-257d-4d96-a60b-216e5ecc8040",
        allowInsecureConnection: false,
        enableBrowserStreams: true,
      }

       */
      if (!request.body)
        return next(request);
      const body = JSON.parse(request.body as string);

      const spanAttributes: TracingAttributes = {
        [TracingAttributesEnum.Operation_Name]: "chat; text_completion", ///TODD: don't know how to determine yet
        [TracingAttributesEnum.System]: "openai",
        [TracingAttributesEnum.Request_Model]: body.model,
        [TracingAttributesEnum.Request_Frequency_Penalty]: body.frequency_penalty,
        [TracingAttributesEnum.Request_Max_Tokens]: body.max_tokens,
        [TracingAttributesEnum.Request_Presence_Penalty]: body.presence_penalty,
        [TracingAttributesEnum.Request_Stop_Sequences]: body.stop,
        [TracingAttributesEnum.Request_Temperature]: body.temperature,
        [TracingAttributesEnum.Request_Top_K]: body.top_k,
        [TracingAttributesEnum.Request_Top_P]: body.top_p
      };

      const { span, tracingContext } = tryCreateSpan(tracingClient, request, spanAttributes) ?? {};

      if (!span || !tracingContext) {
        return next(request);
      }

      try {
        const response = await tracingClient.withContext(tracingContext, next, request);
        tryProcessResponse(span, response);
        return response;
      } catch (err: any) {
        tryProcessError(span, err);
        throw err;
      }
    },
  };
}

function tryCreateTracingClient(): TracingClient | undefined {
  try {
    return createTracingClient({
      namespace: "",
      packageName: "@azure-rest/ai-inference",
      packageVersion: SDK_VERSION,
    });
  } catch (e: unknown) {
    logger.warning(`Error when creating the TracingClient: ${getErrorMessage(e)}`);
    return undefined;
  }
}

function tryCreateSpan(
  tracingClient: TracingClient,
  request: PipelineRequest,
  spanAttributes: Record<string, unknown>,
): { span: TracingSpan; tracingContext: TracingContext } | undefined {
  try {
    // As per spec, we do not need to differentiate between HTTP and HTTPS in span name.
    const { span, updatedOptions } = tracingClient.startSpan(
      "{gen_ai.operation.name} {gen_ai.request.model}", ///TODO: not sure how to retrieve the name yet
      { tracingOptions: {} },  ///TODO: should I pass tracing options?
      {
        spanAttributes,
      },
    );

    // If the span is not recording, don't do any more work.
    if (!span.isRecording()) {
      span.end();
      return undefined;
    }

    // set headers
    const headers = tracingClient.createRequestHeaders(
      updatedOptions.tracingOptions.tracingContext,
    );
    for (const [key, value] of Object.entries(headers)) {
      request.headers.set(key, value);
    }
    return { span, tracingContext: updatedOptions.tracingOptions.tracingContext };
  } catch (e: any) {
    logger.warning(`Skipping creating a tracing span due to an error: ${getErrorMessage(e)}`);
    return undefined;
  }
}

function tryProcessError(span: TracingSpan, error: unknown): void {
  try {
    span.setStatus({
      status: "error",
      error: isError(error) ? error : undefined,
    });
    span.end();
  } catch (e: any) {
    logger.warning(`Skipping tracing span processing due to an error: ${getErrorMessage(e)}`);
  }
}

function tryProcessResponse(span: TracingSpan, response: PipelineResponse): void {
  try {
    const body = JSON.parse(response?.bodyAsText || "{}");
    /* sample of body for success
    const body = {
      choices: [
        {
          finish_reason: "stop",
          index: 0,
          message: {
            content: " The capital of France is Paris.",
            role: "assistant",
            tool_calls: [
            ],
          },
        },
      ],
      created: 1723767112,
      id: "cmpl-53909b0bebf94dab83215ffd99acfd95",
      model: "phi3-medium-128k",
      object: "chat.completion",
      usage: {
        completion_tokens: 8,
        prompt_tokens: 10,
        total_tokens: 18,
      },
    };
     */
    /* sample of body for bad model name
    const body = {
      error: {
        code: "unknown_model",
        message: "Unknown model: chat-gpt-large",
        details: null,
      },
    };
     */
    span.setStatus({
      status: "success",
    });
    span.setAttribute(TracingAttributesEnum.Response_Id, body.id);
    span.setAttribute(TracingAttributesEnum.Response_Model, body.model);
    if (body.usage) {
      span.setAttribute(TracingAttributesEnum.Usage_Input_Tokens, body.usage.prompt_tokens);
      span.setAttribute(TracingAttributesEnum.Usage_Output_Tokens, body.usage.completion_tokens);
    }
    if (body.error) {
      span.setAttribute(TracingAttributesEnum.Error_Type, body.error.code);
      span.setStatus({
        status: "error",
        error: body.error.message,
      });
    }
    span.end();
  } catch (e: any) {
    logger.warning(`Skipping tracing span processing due to an error: ${getErrorMessage(e)}`);
  }
}
