// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

import {
  OperationTracingOptions,
  OptionsWithTracingContext,
  Resolved,
  TracingClient,
  TracingClientOptions,
  TracingContext,
  TracingSpan,
  TracingSpanOptions,
} from "./interfaces.js";
import { getInstrumenter } from "./instrumenter.js";
import { knownContextKeys } from "./tracingContext.js";
import { createClientLogger } from "@azure/logger";
import { getErrorMessage } from "@azure/core-util";

/**
 * Creates a new tracing client.
 *
 * @param options - Options used to configure the tracing client.
 * @returns - An instance of {@link TracingClient}.
 */
export function createTracingClient(options: TracingClientOptions): TracingClient {
  const { namespace, packageName, packageVersion } = options;

  const logger = createClientLogger(namespace);

  function startSpan<Options extends { tracingOptions?: OperationTracingOptions }>(
    name: string,
    operationOptions?: Options,
    spanOptions?: TracingSpanOptions,
  ): {
    span: TracingSpan;
    updatedOptions: OptionsWithTracingContext<Options>;
  } {
    const startSpanResult = getInstrumenter().startSpan(name, {
      ...spanOptions,
      packageName: packageName,
      packageVersion: packageVersion,
      tracingContext: operationOptions?.tracingOptions?.tracingContext,
    });
    let tracingContext = startSpanResult.tracingContext;
    const span = startSpanResult.span;
    if (!tracingContext.getValue(knownContextKeys.namespace)) {
      tracingContext = tracingContext.setValue(knownContextKeys.namespace, namespace);
    }
    span.setAttribute("az.namespace", tracingContext.getValue(knownContextKeys.namespace));
    const updatedOptions: OptionsWithTracingContext<Options> = Object.assign({}, operationOptions, {
      tracingOptions: { ...operationOptions?.tracingOptions, tracingContext },
    });

    return {
      span,
      updatedOptions,
    };
  }

  async function withSpan<
    Options extends { tracingOptions?: OperationTracingOptions },
    Callback extends (
      updatedOptions: Options,
      span: Omit<TracingSpan, "end">,
    ) => ReturnType<Callback>,
  >(
    name: string,
    operationOptions: Options,
    callback: Callback,
    spanOptions?: TracingSpanOptions,
  ): Promise<Resolved<ReturnType<Callback>>> {
    const { span, updatedOptions } = startSpan(name, operationOptions, spanOptions);
    try {
      const result = await withContext(updatedOptions.tracingOptions.tracingContext, () =>
        Promise.resolve(callback(updatedOptions, span)),
      );
      span.setStatus({ status: "success" });
      return result as ReturnType<typeof withSpan>;
    } catch (err: any) {
      span.setStatus({ status: "error", error: err });
      throw err;
    } finally {
      span.end();
    }
  }

  function withContext<
    CallbackArgs extends unknown[],
    Callback extends (...args: CallbackArgs) => ReturnType<Callback>,
  >(
    context: TracingContext,
    callback: Callback,
    ...callbackArgs: CallbackArgs
  ): ReturnType<Callback> {
    return getInstrumenter().withContext(context, callback, ...callbackArgs);
  }

  /**
   * Parses a traceparent header value into a span identifier.
   *
   * @param traceparentHeader - The traceparent header to parse.
   * @returns An implementation-specific identifier for the span.
   */
  function parseTraceparentHeader(traceparentHeader: string): TracingContext | undefined {
    return getInstrumenter().parseTraceparentHeader(traceparentHeader);
  }

  /**
   * Creates a set of request headers to propagate tracing information to a backend.
   *
   * @param tracingContext - The context containing the span to serialize.
   * @returns The set of headers to add to a request.
   */
  function createRequestHeaders(tracingContext?: TracingContext): Record<string, string> {
    return getInstrumenter().createRequestHeaders(tracingContext);
  }

  /**
 * Capture the arguments and return of a function and create a span.
 * @param name - name of the span.
 * @param args - arguments of the method to be traced.  Generally, you should pass in `arguments` reserve word.
 * @param methodToTrace - function pointer of the implementation.
 * @param paramAttributeMapper - mapping function to map the arguments to span's attributes.
 * @param returnAttributeMapper - mapping function to map the return object to span's attributes.
 * @returns - return back the return from methodToTrace.
 */
  function trace<Arguments, Return>(
    name: string,
    args: Arguments,
    methodToTrace: () => Return,
    paramAttributeMapper?: (args: Arguments) => Map<string, unknown>,
    returnAttributeMapper?: (args: Arguments, rt?: Return, error?: unknown) => Map<string, unknown>,
    tracingOptions?: OperationTracingOptions): Return {

    let spanAttributesInObject = {};

    if (paramAttributeMapper) {
      const spanAttributesInMap = paramAttributeMapper(args);
      spanAttributesInObject = mapToObject(spanAttributesInMap);
    }
    const { span, tracingContext } = tryCreateSpan(name, spanAttributesInObject, tracingOptions) ?? {};

    if (!span || !tracingContext) {
      return methodToTrace();
    }

    try {
      const returnObj = withContext(tracingContext, methodToTrace)
      tryProcessReturn(span, args, returnObj, returnAttributeMapper);

      return returnObj;
    } catch (err: any) {
      tryProcessReturn(span, args, undefined, err, returnAttributeMapper);
      throw err;
    }
  }

  /**
* Capture the arguments and return of a function and create a span.
* @param name - name of the span.
* @param args - arguments of the method to be traced.  Generally, you should pass in `arguments` reserve word.
* @param methodToTrace - function pointer of the implementation.
* @param paramAttributeMapper - mapping function to map the arguments to span's attributes.
* @param returnAttributeMapper - mapping function to map the return object to span's attributes.
* @returns - return back the return from methodToTrace.
*/
  function traceAsync<Arguments, ResolvedReturn, PromiseReturn extends Promise<ResolvedReturn> | PromiseLike<ResolvedReturn>>(
    name: string,
    args: Arguments,
    methodToTrace: () => PromiseReturn,
    paramAttributeMapper?: (args: Arguments) => Map<string, unknown>,
    returnAttributeMapper?: (args: Arguments, rt?: ResolvedReturn, error?: unknown) => Map<string, unknown>,
    tracingOptions?: OperationTracingOptions): PromiseReturn {

    let spanAttributesInObject = {};

    if (paramAttributeMapper) {
      const spanAttributesInMap = paramAttributeMapper(args);
      spanAttributesInObject = mapToObject(spanAttributesInMap);
    }
    const { span, tracingContext } = tryCreateSpan(name, spanAttributesInObject, tracingOptions) ?? {};

    if (!span || !tracingContext) {
      return methodToTrace();
    }

    try {
      const returnObj = withContext(tracingContext, methodToTrace);

      returnObj.then((response) => {
        tryProcessReturn(span, args, response, undefined, returnAttributeMapper);
      });

      return returnObj;
    } catch (err) {
      tryProcessReturn(span, args, undefined, err, returnAttributeMapper);
      throw err;
    }
  }

  function mapToObject(map: Map<string, any>): any {
    const object: any = {};
    map.forEach((value, key) => {
      object[key] = value;
    });
    return object;
  }

  function tryCreateSpan(
    spanName: string,
    spanAttributes: Record<string, unknown>,
    tracingOptions?: OperationTracingOptions
  ): { span: TracingSpan; tracingContext: TracingContext } | undefined {
    try {
      const { span, updatedOptions } = startSpan(
        spanName,
        { tracingOptions },
        {
          spanAttributes,
        },
      );

      // If the span is not recording, don't do any more work.
      if (!span.isRecording()) {
        span.end();
        return undefined;
      }


      return { span, tracingContext: updatedOptions.tracingOptions.tracingContext };
    } catch (e: any) {
      logger.warning(`Skipping creating a tracing span due to an error: ${getErrorMessage(e)}`);
      return undefined;
    }
  }

  function tryProcessReturn<Arguments, Return>(
    span: TracingSpan,
    args: Arguments,
    returnObj?: Return,
    error?: unknown,
    returnAttributeMapper?: (args: Arguments, rt?: Return, error?: unknown) => Map<string, any>): void {
    try {
      if (returnAttributeMapper) {
        const returnAttributes = returnAttributeMapper(args, returnObj, error);
        returnAttributes.forEach((value, key) => {
          span.setAttribute(key, value);
        });
      }

      span.end();
    } catch (e: any) {
      logger.warning(`Skipping tracing span processing due to an error: ${getErrorMessage(e)}`);
    }
  }


  return {
    startSpan,
    withSpan,
    withContext,
    parseTraceparentHeader,
    createRequestHeaders,
    trace,
    traceAsync
  };
}
