// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

import { PathUncheckedResponse, RequestParameters, StreamableMethod } from "@azure-rest/core-client";
import { createTracingClient, OperationTracingOptions } from "@azure/core-tracing";
import { GetChatCompletionsBodyParam, GetEmbeddingsBodyParam, GetImageEmbeddingsBodyParam } from "./parameters.js";

const traceCLient = createTracingClient({ namespace: "Micirsoft.CognitiveServices", packageName: "ai-inference-rest", packageVersion: "1.0.0" });

type RequestParameterWithBodyType = RequestParameters & GetImageEmbeddingsBodyParam &
  GetEmbeddingsBodyParam &
  GetChatCompletionsBodyParam &
  GetImageEmbeddingsBodyParam;

export function traceInference(
  path: string,
  args: RequestParameters,
  methodToTrace: () => StreamableMethod,
  options?: OperationTracingOptions): StreamableMethod {

  enum TracingAttributesEnum {
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

  const requestAttributeMapper = (request: RequestParameterWithBodyType) => {
    const map = new Map<string, unknown>();
    const body = (request as RequestParameterWithBodyType).body;

    map.set(TracingAttributesEnum.Operation_Name, getOperationName(path));
    map.set(TracingAttributesEnum.System, "az.ai_inference");
    map.set(TracingAttributesEnum.Request_Model, body?.model);
    map.set(TracingAttributesEnum.Request_Frequency_Penalty, body?.frequency_penalty);
    map.set(TracingAttributesEnum.Request_Max_Tokens, body?.max_tokens);
    map.set(TracingAttributesEnum.Request_Presence_Penalty, body?.presence_penalty);
    map.set(TracingAttributesEnum.Request_Stop_Sequences, body?.stop);
    map.set(TracingAttributesEnum.Request_Temperature, body?.temperature);
    map.set(TracingAttributesEnum.Request_Top_P, body?.top_p);
    return map;
  }

  const responseAttributeMapper = (_: RequestParameters, response?: PathUncheckedResponse, error?: unknown) => {
    const map = new Map<string, unknown>();
    if (error) {
      if (error instanceof Error) {
        map.set(TracingAttributesEnum.Error_Type, error.message);
      } else {
        map.set(TracingAttributesEnum.Error_Type, error);
      }
    }
    if (response) {
      let body = response.body;
      map.set(TracingAttributesEnum.Response_Id, body.id);
      map.set(TracingAttributesEnum.Response_Model, body.model);
      if (body.usage) {
        map.set(TracingAttributesEnum.Usage_Input_Tokens, body.usage.prompt_tokens);
        map.set(TracingAttributesEnum.Usage_Output_Tokens, body.usage.completion_tokens);
      }
      if (body.error) {
        map.set(TracingAttributesEnum.Error_Type, body.error.code);
      }
    }
    return map;
  }

  const request = args as RequestParameterWithBodyType;

  const name = `${getOperationName(path)} ${request.body?.model ?? ""}`;
  return traceCLient.traceAsync(name, request, methodToTrace, requestAttributeMapper, responseAttributeMapper, options);
}
