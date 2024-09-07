// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

import { PathUncheckedResponse, RequestParameters, StreamableMethod } from "@azure-rest/core-client";
import { createTracingClient, SpanStatus, TracingSpan } from "@azure/core-tracing";
import { GetChatCompletionsBodyParam, GetEmbeddingsBodyParam, GetImageEmbeddingsBodyParam } from "./parameters.js";
import { ChatRequestMessage } from "./models.js";
import { ChatChoiceOutput } from "./outputModels.js";

const traceClient = createTracingClient({ namespace: "Micirsoft.CognitiveServices", packageName: "ai-inference-rest", packageVersion: "1.0.0" });

type RequestParameterWithBodyType = RequestParameters & GetImageEmbeddingsBodyParam &
  GetEmbeddingsBodyParam &
  GetChatCompletionsBodyParam &
  GetImageEmbeddingsBodyParam;

export function traceInference(
  routePath: string,
  url: string,
  args: RequestParameters,
  methodToTrace: () => StreamableMethod): StreamableMethod {

  enum TracingAttributesEnum {
    Operation_Name = "gen_ai.operation.name",
    Request_Model = "gen_ai.request.model",
    System = "gen_ai.system",
    Error_Type = "error.type",
    Server_Port = "server.port",
    Request_Frequency_Penalty = "gen_ai.request.frequency_penalty",
    Request_Max_Tokens = "gen_ai.request.max_tokens",
    Request_Presence_Penalty = "gen_ai.request.presence_penalty",
    Request_Stop_Sequences = "gen_ai.request.stop_sequences",
    Request_Temperature = "gen_ai.request.temperature",
    Request_Top_P = "gen_ai.request.top_p",
    Response_Finish_Reasons = "gen_ai.response.finish_reasons",
    Response_Id = "gen_ai.response.id",
    Response_Model = "gen_ai.response.model",
    Usage_Input_Tokens = "gen_ai.usage.input_tokens",
    Usage_Output_Tokens = "gen_ai.usage.output_tokens",
    Server_Address = "server.address"
  }

  //TODO: if model is not provided, we probably should parse from the URL
  const model = (args as RequestParameterWithBodyType).body?.model;
  const getOperationName = (path: string) => {
    switch (path) {
      case "/chat/completions":
        return "chat";
      case "/info":
        return "info";
      case "/embeddings":
        return "text_embeddings";
      case "/images/embeddings":
        return "image_embeddings";
      default:
        throw new Error(`Unknown path for span name: ${path}`);
    }
  }

  function onStartTracing(span: TracingSpan, request: RequestParameters) {

    const urlObj = new URL(url);
    const port = urlObj.port || (urlObj.protocol === 'https:' ? 443 : 80);

    span.setAttribute(TracingAttributesEnum.Server_Address, urlObj.hostname);
    span.setAttribute(TracingAttributesEnum.Server_Port, port);
    span.setAttribute(TracingAttributesEnum.Operation_Name, getOperationName(routePath));
    span.setAttribute(TracingAttributesEnum.System, "az.ai_inference");
    span.setAttribute(TracingAttributesEnum.Request_Model, model);

    const body = (request as RequestParameterWithBodyType).body;
    if (!body) return;

    span.setAttribute(TracingAttributesEnum.Request_Frequency_Penalty, body.frequency_penalty);
    span.setAttribute(TracingAttributesEnum.Request_Max_Tokens, body.max_tokens);
    span.setAttribute(TracingAttributesEnum.Request_Presence_Penalty, body.presence_penalty);
    span.setAttribute(TracingAttributesEnum.Request_Stop_Sequences, body.stop);
    span.setAttribute(TracingAttributesEnum.Request_Temperature, body.temperature);
    span.setAttribute(TracingAttributesEnum.Request_Top_P, body.top_p);

    if (body.messages) {
      addRequestChatMessageEvent(span, body.messages);
    }
  }

  function onEndTracing(span: TracingSpan, _request: RequestParameters, response?: PathUncheckedResponse, error?: unknown) {
    let status: SpanStatus = { status: "unset" };
    if (error) {
      if (error instanceof Error) {
        span.setAttribute(TracingAttributesEnum.Error_Type, error.message);
        status = { status: "error", error: error.message };
      } else {
        span.setAttribute(TracingAttributesEnum.Error_Type, error);
        status = { status: "error", error: error.toString() };
      }

    }
    if (response) {
      let body = response.body;
      span.setAttribute(TracingAttributesEnum.Response_Id, body.id);
      span.setAttribute(TracingAttributesEnum.Response_Model, body.model);
      if (body.usage) {
        span.setAttribute(TracingAttributesEnum.Usage_Input_Tokens, body.usage.prompt_tokens);
        span.setAttribute(TracingAttributesEnum.Usage_Output_Tokens, body.usage.completion_tokens);
      }
      if (body.error) {
        span.setAttribute(TracingAttributesEnum.Error_Type, body.error.code);
      }
      addResponseChatMessageEvent(span, response);
    }
    span.setStatus(status);
  }

  /*
  * Add event to span.  Sample:
      {
      name: 'gen_ai.user.message',
      attributes: {
        'gen_ai.system': 'INFERENCE_GEN_AI_SYSTEM_NAME',
        'gen_ai.event.content': `{"role":"user","content":"What's the weather like in Boston?"}`  
      },
      time: [ 1725666879, 622695900 ],
      droppedAttributesCount: 0
    },
  */
  function addRequestChatMessageEvent(span: TracingSpan, messages: Array<ChatRequestMessage>) {
    messages.forEach((message) => {

      if (message.role) {
        span.addEvent(`gen_ai.${message.role}.message`, {
          "gen_ai.system": "INFERENCE_GEN_AI_SYSTEM_NAME", // Replace with actual system name
          "gen_ai.event.content": JSON.stringify(message)
        });
      }

    });
  }

  /*
* Add event to span.  Sample:
  {
    name: 'gen_ai.choice',
    attributes: {
      'gen_ai.system': 'INFERENCE_GEN_AI_SYSTEM_NAME',
      'gen_ai.event.content': '{"finish_reason":"tool_calls","index":0,"message":{"content":""}}'
    },
    time: [ 1725666881, 780608000 ],
    droppedAttributesCount: 0
  }  
*/
  function addResponseChatMessageEvent(span: TracingSpan, response: PathUncheckedResponse) {
    response.body?.choices?.forEach((choice: ChatChoiceOutput) => {
      let response: any = {
        "finish_reason": choice.finish_reason,
        "index": choice.index,
      };

      let attributes: any = {
        "gen_ai.system": "INFERENCE_GEN_AI_SYSTEM_NAME",
      };

      response = { ...response, message: { "content": choice.message.content } };
      attributes = { ...attributes, "gen_ai.event.content": JSON.stringify(response) };

      if (choice.message.tool_calls) {
        response.message = { ...response, toolCalls: JSON.stringify(choice.message.tool_calls) };
      }

      span.addEvent("gen_ai.choice", attributes);
    });
  }


  const request = args as RequestParameterWithBodyType;

  /// TODO: the code for streaming needs to be clean up.   We will implement tracing for streaming later 
  if (request.body?.stream == true || !Boolean(process.env["AZURE_TRACING_GEN_AI_CONTENT_RECORDING_ENABLED"])) {
    return methodToTrace();
  }
  const name = `${getOperationName(routePath)} ${model ?? ""}`.trim();
  return traceClient.traceAsync(name, request, methodToTrace, onStartTracing, onEndTracing, args.tracingOptions);
}
