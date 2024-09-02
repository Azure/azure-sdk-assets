// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

import { KeyCredential, TokenCredential, isTokenCredential } from "@azure/core-auth";
import { HttpClient, HttpMethods, Pipeline, PipelineOptions } from "@azure/core-rest-pipeline";
import { createDefaultPipeline } from "./clientHelpers.js";
import {
  Client,
  ClientOptions,
  HttpBrowserStreamResponse,
  HttpNodeStreamResponse,
  RequestParameters,
  StreamableMethod,
  TracerCallbacks,
} from "./common.js";
import { sendRequest } from "./sendRequest.js";
import { buildRequestUrl } from "./urlHelpers.js";
import { createTracingClient } from "@azure/core-tracing";

/**
 * Creates a client with a default pipeline
 * @param endpoint - Base endpoint for the client
 * @param options - Client options
 */
export function getClient(endpoint: string, options?: ClientOptions): Client;
/**
 * Creates a client with a default pipeline
 * @param endpoint - Base endpoint for the client
 * @param credentials - Credentials to authenticate the requests
 * @param options - Client options
 * @param tracer - wrapper method around operations for telemetry
 */
export function getClient(
  endpoint: string,
  credentials?: TokenCredential | KeyCredential,
  options?: ClientOptions,
  tracerCallbacks?: TracerCallbacks,
): Client;
export function getClient(
  endpoint: string,
  credentialsOrPipelineOptions?: (TokenCredential | KeyCredential) | ClientOptions,
  clientOptions: ClientOptions = {},
  tracerCallbacks?: TracerCallbacks,
): Client {
  let credentials: TokenCredential | KeyCredential | undefined;
  if (credentialsOrPipelineOptions) {
    if (isCredential(credentialsOrPipelineOptions)) {
      credentials = credentialsOrPipelineOptions;
    } else {
      clientOptions = credentialsOrPipelineOptions ?? {};
    }
  }

  const pipeline = createDefaultPipeline(endpoint, credentials, clientOptions);
  if (clientOptions.additionalPolicies?.length) {
    for (const { policy, position } of clientOptions.additionalPolicies) {
      // Sign happens after Retry and is commonly needed to occur
      // before policies that intercept post-retry.
      const afterPhase = position === "perRetry" ? "Sign" : undefined;
      pipeline.addPolicy(policy, {
        afterPhase,
      });
    }
  }

  const { allowInsecureConnection, httpClient } = clientOptions;
  const endpointUrl = clientOptions.endpoint ?? endpoint;
  const client = (path: string, ...args: Array<any>) => {
    const getUrl = (requestOptions: RequestParameters) =>
      buildRequestUrl(endpointUrl, path, args, { allowInsecureConnection, ...requestOptions });

    return {
      get: (requestOptions: RequestParameters = {}): StreamableMethod => {
        return buildOperation(
          "GET",
          getUrl(requestOptions),
          pipeline,
          requestOptions,
          allowInsecureConnection,
          httpClient,
          tracerCallbacks
        );
      },
      post: (requestOptions: RequestParameters = {}): StreamableMethod => {
        return buildOperation(
          "POST",
          getUrl(requestOptions),
          pipeline,
          requestOptions,
          allowInsecureConnection,
          httpClient,
          tracerCallbacks
        );
      },
      put: (requestOptions: RequestParameters = {}): StreamableMethod => {
        return buildOperation(
          "PUT",
          getUrl(requestOptions),
          pipeline,
          requestOptions,
          allowInsecureConnection,
          httpClient,
          tracerCallbacks
        );
      },
      patch: (requestOptions: RequestParameters = {}): StreamableMethod => {
        return buildOperation(
          "PATCH",
          getUrl(requestOptions),
          pipeline,
          requestOptions,
          allowInsecureConnection,
          httpClient,
          tracerCallbacks
        );
      },
      delete: (requestOptions: RequestParameters = {}): StreamableMethod => {
        return buildOperation(
          "DELETE",
          getUrl(requestOptions),
          pipeline,
          requestOptions,
          allowInsecureConnection,
          httpClient,
          tracerCallbacks
        );
      },
      head: (requestOptions: RequestParameters = {}): StreamableMethod => {
        return buildOperation(
          "HEAD",
          getUrl(requestOptions),
          pipeline,
          requestOptions,
          allowInsecureConnection,
          httpClient,
          tracerCallbacks
        );
      },
      options: (requestOptions: RequestParameters = {}): StreamableMethod => {
        return buildOperation(
          "OPTIONS",
          getUrl(requestOptions),
          pipeline,
          requestOptions,
          allowInsecureConnection,
          httpClient,
          tracerCallbacks
        );
      },
      trace: (requestOptions: RequestParameters = {}): StreamableMethod => {
        return buildOperation(
          "TRACE",
          getUrl(requestOptions),
          pipeline,
          requestOptions,
          allowInsecureConnection,
          httpClient,
          tracerCallbacks
        );
      },
    };
  };

  return {
    path: client,
    pathUnchecked: client,
    pipeline,
  };
}

const tracer = createTracingClient({
  packageName: "Azure.Rest",
  namespace: "Microsoft.Rest",
});

function buildOperation(
  method: HttpMethods,
  url: string,
  pipeline: Pipeline,
  options: RequestParameters,
  allowInsecureConnection?: boolean,
  httpClient?: HttpClient,
  tracerCallbacks?: TracerCallbacks,
): StreamableMethod {
  allowInsecureConnection = options.allowInsecureConnection ?? allowInsecureConnection;


  const operation: () => StreamableMethod = () => ({
    then: function (onFulfilled, onrejected) {
      return sendRequest(
        method,
        url,
        pipeline,
        { ...options, allowInsecureConnection },
        httpClient,
      ).then(onFulfilled, onrejected);
    },
    async asBrowserStream() {
      return sendRequest(
        method,
        url,
        pipeline,
        { ...options, allowInsecureConnection, responseAsStream: true },
        httpClient,
      ) as Promise<HttpBrowserStreamResponse>;
    },
    async asNodeStream() {
      return sendRequest(
        method,
        url,
        pipeline,
        { ...options, allowInsecureConnection, responseAsStream: true },
        httpClient,
      ) as Promise<HttpNodeStreamResponse>;
    },
  });

  return tracerCallbacks ?
    tracer.traceAsync(url, options, operation, tracerCallbacks.requestAttributeMapper, tracerCallbacks.responseAttributeMapper, tracerCallbacks.statusMapper) :
    operation();
}

function isCredential(
  param: (TokenCredential | KeyCredential) | PipelineOptions,
): param is TokenCredential | KeyCredential {
  if ((param as KeyCredential).key !== undefined || isTokenCredential(param)) {
    return true;
  }

  return false;
}
