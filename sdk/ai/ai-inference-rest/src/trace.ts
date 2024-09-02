// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

import { PathUncheckedResponse, RequestParameters } from "@azure-rest/core-client";
import { SpanStatus } from "@azure/core-tracing";
import { GetChatCompletionsBodyParam, GetEmbeddingsBodyParam, GetImageEmbeddingsBodyParam } from "./parameters.js";


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

export type TracingAttributes = { [key in TracingAttributesEnum]?: string[] | string | number };

export function requestAttributeMapper(params: RequestParameters): Map<string, unknown> {
  const parameters = new Map<string, unknown>();
  const body = (params as GetImageEmbeddingsBodyParam &
    GetEmbeddingsBodyParam &
    GetChatCompletionsBodyParam &
    GetImageEmbeddingsBodyParam).body;

  parameters.set(TracingAttributesEnum.Operation_Name, "chat; text_completion"); ///TODD: don't know how to determine yet
  parameters.set(TracingAttributesEnum.System, "openai");
  parameters.set(TracingAttributesEnum.Request_Model, body?.model);
  parameters.set(TracingAttributesEnum.Request_Frequency_Penalty, body?.frequency_penalty);
  parameters.set(TracingAttributesEnum.Request_Max_Tokens, body?.max_tokens);
  parameters.set(TracingAttributesEnum.Request_Presence_Penalty, body?.presence_penalty);
  parameters.set(TracingAttributesEnum.Request_Stop_Sequences, body?.stop);
  parameters.set(TracingAttributesEnum.Request_Temperature, body?.temperature);
  parameters.set(TracingAttributesEnum.Request_Top_P, body?.top_p);
  return parameters;
}

export function responseAttributeMapper(response: PathUncheckedResponse): Map<string, unknown> {
  const body = response.body;
  const parameters = new Map<string, unknown>();
  parameters.set(TracingAttributesEnum.Response_Id, body.id);
  parameters.set(TracingAttributesEnum.Response_Model, body.model);
  if (body.usage) {
    parameters.set(TracingAttributesEnum.Usage_Input_Tokens, body.usage.prompt_tokens);
    parameters.set(TracingAttributesEnum.Usage_Output_Tokens, body.usage.completion_tokens);
  }
  if (body.error) {
    parameters.set(TracingAttributesEnum.Error_Type, body.error.code);
  }
  return parameters;
}

export function statusMapper({ body }: PathUncheckedResponse): SpanStatus {
  let status: SpanStatus = {
    status: "success",
  };

  if (body.error) {
    status = {
      status: "error",
      error: body.error.message,
    };
  }
  return status;
}
