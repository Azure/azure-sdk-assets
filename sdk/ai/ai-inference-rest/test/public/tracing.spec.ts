// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

import { createRecorder, createModelClient } from "./utils/recordedClient.js";
import { Recorder, env } from "@azure-tools/test-recorder";
import { assert, beforeEach, afterEach, it, describe } from "vitest";
import { trace, context } from "@opentelemetry/api";
import { ChatCompletionsOutput, ModelClient } from "../../src/index.js";
import { Instrumenter, InstrumenterSpanOptions, SpanStatus, TracingContext, TracingSpan, TracingSpanOptions, useInstrumenter } from "@azure/core-tracing";

describe("chat test suite", () => {
  let recorder: Recorder;
  let client: ModelClient;
  let instrumenter: MockInstrumenter;

  const getCurrentWeather = {
    name: "get_current_weather",
    description: "Get the current weather in a given location",
    parameters: {
      type: "object",
      properties: {
        location: {
          type: "string",
          description: "The city and state, e.g. San Francisco, CA",
        },
        unit: {
          type: "string",
          enum: ["celsius", "fahrenheit"],
        },
      },
      required: ["location"],
    },
  };

  beforeEach(async (context) => {
    recorder = await createRecorder(context);
    client = await createModelClient("completions", recorder);
    instrumenter = new MockInstrumenter();
    useInstrumenter(instrumenter);
  });

  afterEach(async () => {
    await recorder.stop();
  });

  it("client test", function () {
    assert.isNotNull(client);
    assert.isNotNull(client.path);
    assert.isNotNull(client.pipeline);
  });

  it("tracing enabled", async function () {
    env.AZURE_TRACING_GEN_AI_CONTENT_RECORDING_ENABLED = "true";

    const tracer = trace.getTracer('sample', '0.1.0');

    const response = await tracer.startActiveSpan('main', async (span) => {
      return client.path("/chat/completions").post({
        body: {
          messages: [{ role: "user", content: "What's the weather like in Boston?" }],
          tools: [
            {
              type: "function",
              function: getCurrentWeather,
            },
          ],
        },
        tracingOptions: { tracingContext: context.active() }
      }).then((response) => {
        span.end();
        return response;
      });
    });

    const createdSpan = instrumenter.createdSpans.get("chat");
    if (!createdSpan) {
      assert.fail("expected span to be created");
    }

    const mockSpan = createdSpan;
    assert.isTrue(mockSpan.endCalled, "expected span to be ended");
    assert.equal(mockSpan.name, "chat");
    assert.equal(mockSpan.getAttribute("az.namespace"), "Micirsoft.CognitiveServices");
    assert.isTrue(mockSpan.getAttribute("server.address")?.toString().endsWith("azure.com"));
    assert.equal(mockSpan.getAttribute("server.port"), 443);
    assert.equal(mockSpan.getAttribute("gen_ai.operation.name"), "chat");
    assert.equal(mockSpan.getAttribute("gen_ai.system"), "az.ai.inference");
    assert.equal(mockSpan.getAttribute("gen_ai.response.id"), (response.body as ChatCompletionsOutput).id);
    assert.equal(mockSpan.getAttribute("gen_ai.response.model"), (response.body as ChatCompletionsOutput).model);
    assert.equal(mockSpan.getAttribute("gen_ai.usage.input_tokens"), (response.body as ChatCompletionsOutput).usage.prompt_tokens);
    assert.equal(mockSpan.getAttribute("gen_ai.usage.output_tokens"), (response.body as ChatCompletionsOutput).usage.completion_tokens);
    assert.equal(mockSpan.events.size, 2);

    const userMessageEvent = mockSpan.events.get("gen_ai.user.message");
    assert.isDefined(userMessageEvent);
    assert.deepEqual(userMessageEvent, {
      name: "gen_ai.user.message",
      attributesOrStartTime: {
        'gen_ai.system': 'INFERENCE_GEN_AI_SYSTEM_NAME',
        'gen_ai.event.content': `{"role":"user","content":"What's the weather like in Boston?"}`
      },
      startTime: undefined
    });
    const choiceEvent = mockSpan.events.get("gen_ai.choice");
    assert.isDefined(choiceEvent);
    assert.deepEqual(choiceEvent, {
      name: "gen_ai.choice",
      attributesOrStartTime: {
        'gen_ai.system': 'INFERENCE_GEN_AI_SYSTEM_NAME',
        'gen_ai.event.content': '{"finish_reason":"tool_calls","index":0,"message":{"content":null}}'
      },
      startTime: undefined
    });

  });

  it("tracing with errors", async function () {

    client = await createModelClient("dummy", recorder);

    env.AZURE_TRACING_GEN_AI_CONTENT_RECORDING_ENABLED = "true";

    const tracer = trace.getTracer('sample', '0.1.0');

    await tracer.startActiveSpan('main', async (span) => {
      return client.path("/chat/completions").post({
        body: {
          messages: [{ role: "user", content: "" }]
        },
        tracingOptions: { tracingContext: context.active() }
      }).then((response) => {
        span.end();
        return response;
      });
    });

    const createdSpan = instrumenter.createdSpans.get("chat");
    if (!createdSpan) {
      assert.fail("expected span to be created");
    }

    const mockSpan = createdSpan;
    assert.isTrue(mockSpan.endCalled, "expected span to be ended");
    assert.equal(mockSpan.name, "chat");
    assert.equal(mockSpan.getAttribute("az.namespace"), "Micirsoft.CognitiveServices");
    assert.isTrue(mockSpan.getAttribute("server.address")?.toString().endsWith("azure.com"));
    assert.equal(mockSpan.getAttribute("server.port"), 443);
    assert.equal(mockSpan.getAttribute("gen_ai.operation.name"), "chat");
    assert.equal(mockSpan.getAttribute("gen_ai.system"), "az.ai.inference");
    assert.deepEqual(mockSpan.status, { status: "error", error: "Unauthorized. Access token is missing, invalid, audience is incorrect (https://cognitiveservices.azure.com), or have expired." });
  });

  it("tracing disabled", async function () {
    delete env.AZURE_TRACING_GEN_AI_CONTENT_RECORDING_ENABLED;

    const tracer = trace.getTracer('sample', '0.1.0');

    await tracer.startActiveSpan('main', async (span) => {
      return client.path("/chat/completions").post({
        body: {
          messages: [{ role: "user", content: "What's the weather like in Boston?" }],
          tools: [
            {
              type: "function",
              function: getCurrentWeather,
            },
          ],
        },
        tracingOptions: { tracingContext: context.active() }
      }).then((response) => {
        span.end();
        return response;
      });
    });

    const createdSpan = instrumenter.createdSpans.get("chat");
    if (createdSpan) {
      assert.fail("expected span not to be created");
    }

  });
});



class MockSpan implements TracingSpan {
  spanAttributes: Record<string, unknown> = {};
  events = new Map<string, { name: string, attributesOrStartTime?: unknown, startTime?: unknown }>();
  endCalled: boolean = false;
  status?: SpanStatus;
  exceptions: Array<Error | string> = [];

  constructor(
    public name: string,
    spanOptions: TracingSpanOptions = {},
  ) {
    this.spanAttributes = spanOptions.spanAttributes ?? {};
  }
  addEvent(name: string, attributesOrStartTime?: unknown, startTime?: unknown): void {
    this.events.set(name, { name, attributesOrStartTime, startTime });
  }

  isRecording(): boolean {
    return true;
  }

  recordException(exception: Error | string): void {
    this.exceptions.push(exception);
  }

  end(): void {
    this.endCalled = true;
  }

  setStatus(status: SpanStatus): void {
    this.status = status;
  }

  setAttribute(name: string, value: unknown): void {
    this.spanAttributes[name] = value;
  }

  getAttribute(name: string): unknown {
    return this.spanAttributes[name];
  }
}

const noopTracingContext: TracingContext = {
  deleteValue() {
    return this;
  },
  getValue() {
    return undefined;
  },
  setValue() {
    return this;
  },
};

class MockInstrumenter implements Instrumenter {
  createdSpans = new Map<string, MockSpan>;
  staticSpan: MockSpan | undefined;

  setStaticSpan(span: MockSpan): void {
    this.staticSpan = span;
  }
  startSpan(
    name: string,
    spanOptions: InstrumenterSpanOptions,
  ): {
    span: TracingSpan;
    tracingContext: TracingContext;
  } {
    const tracingContext = spanOptions.tracingContext ?? noopTracingContext;
    if (this.staticSpan) {
      return { span: this.staticSpan, tracingContext };
    }
    const span = new MockSpan(name, spanOptions);
    this.createdSpans.set(name, span);
    return {
      span,
      tracingContext,
    };
  }
  withContext<
    CallbackArgs extends unknown[],
    Callback extends (...args: CallbackArgs) => ReturnType<Callback>,
  >(
    _context: TracingContext,
    callback: Callback,
    ...callbackArgs: CallbackArgs
  ): ReturnType<Callback> {
    return callback(...callbackArgs);
  }

  parseTraceparentHeader(_traceparentHeader: string): TracingContext | undefined {
    return undefined;
  }
  createRequestHeaders(_tracingContext?: TracingContext): Record<string, string> {
    return {};
  }
}
