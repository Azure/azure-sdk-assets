// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

/**
 * Demonstrates how to get instrumentation by open telemetry.
 *
 * @summary get instrumentation by open telemetry.
 */

import { diag, DiagConsoleLogger, DiagLogLevel } from "@opentelemetry/api";
import { registerInstrumentations } from "@opentelemetry/instrumentation";
import { createAzureSdkInstrumentation } from "@azure/opentelemetry-instrumentation-azure-sdk";
import { ConsoleSpanExporter, NodeTracerProvider, SimpleSpanProcessor } from "@opentelemetry/sdk-trace-node";
import { AzureMonitorTraceExporter } from "@azure/monitor-opentelemetry-exporter";

diag.setLogger(new DiagConsoleLogger(), DiagLogLevel.VERBOSE);

// Load the .env file if it exists
import * as dotenv from "dotenv";
dotenv.config();

// You will need to set these environment variables or edit the following values
const endpoint: string = process.env["ENDPOINT"] || "<endpoint>";
const key: string = process.env["KEY"] || "<key>";
const connectionString = process.env["APPLICATIONINSIGHTS_CONNECTION_STRING"];

// Initialize the exporter
const provider = new NodeTracerProvider();
if (connectionString) {
  const exporter = new AzureMonitorTraceExporter({
    connectionString
  });
  provider.addSpanProcessor(
    new SimpleSpanProcessor(exporter)
  );
}
provider.addSpanProcessor(
  new SimpleSpanProcessor(new ConsoleSpanExporter())
);
provider.register();

// register Azure SDK Instrumentation.
// ************** IMPORTANT SETUP STEP ***************
// this must be done before loading @azure/core-tracing
registerInstrumentations({
  instrumentations: [createAzureSdkInstrumentation()],
});

import ModelClient from "@azure-rest/ai-inference";
import { isUnexpected } from "@azure-rest/ai-inference";
import { AzureKeyCredential } from "@azure/core-auth";
import { createTracingClient } from "@azure/core-tracing";

const tracingClient = createTracingClient({
  namespace: "Microsoft.OtelSample",
  packageName: "otel-sample",
  packageVersion: "1.0.0",
});

async function main() {
  console.log("== Chat Completions Sample ==");

  // initialize a span named "main" with default options for the spans
  await tracingClient.withSpan("main", {}, async (updatedOptions) => {
    const client = ModelClient(endpoint, new AzureKeyCredential(key));

    const response = await client.path("/chat/completions").post({
      body: {
        messages: [
          { role: "system", content: "You are a helpful assistant. You will talk like a pirate." },
          { role: "user", content: "Can you help me?" },
          { role: "assistant", content: "Arrrr! Of course, me hearty! What can I do for ye?" },
          { role: "user", content: "What's the best way to train a parrot?" },
        ],
        temperature: 1.0,
        max_tokens: 1000,
        top_p: 1.0
      },
      ...updatedOptions
    });

    if (isUnexpected(response)) {
      throw response.body.error;
    }

    for (const choice of response.body.choices) {
      console.log(choice.message.content);
    }
  });
}

main().catch((err) => {
  console.error("The sample encountered an error:", err);
});

/**
 * Output of the sample:
 */
/*
== Chat Completions Sample ==
{
  resource: {
    attributes: {
      'service.name': 'howieinsight',
      'telemetry.sdk.language': 'nodejs',
      'telemetry.sdk.name': 'opentelemetry',
      'telemetry.sdk.version': '1.26.0'
    }
  },
  instrumentationScope: {
    name: '@azure/core-rest-pipeline',
    version: '1.16.4',
    schemaUrl: undefined
  },
  traceId: '6c9a1a89426e53ee27d4ac127ccaf832',
  parentId: 'f1b4a676d9b58b78',
  traceState: undefined,
  name: 'HTTP POST',
  id: '4a99e2564e3e17e4',
  kind: 2,
  timestamp: 1725578892511000,
  duration: 4317602.4,
  attributes: {
    'http.url': 'https://mistral-large-ajmih-serverless.eastus2.inference.ai.azure.com/chat/completions?api-version=2024-05-01-preview',
    'http.method': 'POST',
    'http.user_agent': 'azsdk-js-ai-inference/1.0.0-beta.2 core-rest-pipeline/1.16.4 Node/18.20.4 OS/(x64-Windows_NT-10.0.22631)',
    requestId: 'bac64077-def9-4c75-a704-0d6bcd19d3f7',
    'az.namespace': 'Microsoft.OtelSample',
    'http.status_code': 200
  },
  status: { code: 1 },
  events: [],
  links: []
}
Exporting 1 span(s). Converting to envelopes...
Exporting 2 envelope(s)
Instrumentation suppressed, returning Noop Span
{
  resource: {
    attributes: {
      'service.name': 'howieinsight',
      'telemetry.sdk.language': 'nodejs',
      'telemetry.sdk.name': 'opentelemetry',
      'telemetry.sdk.version': '1.26.0'
    }
  },
  instrumentationScope: { name: 'ai-inference-rest', version: '1.0.0', schemaUrl: undefined },
  traceId: '6c9a1a89426e53ee27d4ac127ccaf832',
  parentId: 'f1b4a676d9b58b78',
  traceState: undefined,
  name: 'chat',
  id: '97620d98f691963a',
  kind: 0,
  timestamp: 1725578892479000,
  duration: 4376479.3,
  attributes: {
    'server.address': 'mistral-large-ajmih-serverless.eastus2.inference.ai.azure.com',
    'server.port': 443,
    'gen_ai.operation.name': 'chat',
    'gen_ai.system': 'az.ai_inference',
    'gen_ai.request.max_tokens': 1000,
    'gen_ai.request.temperature': 1,
    'gen_ai.request.top_p': 1,
    'az.namespace': 'Microsoft.OtelSample',
    'gen_ai.response.id': '73b1a5ba89b545a1af217843000855aa',
    'gen_ai.response.model': 'mistral-large',
    'gen_ai.usage.input_tokens': 57,
    'gen_ai.usage.output_tokens': 98
  },
  status: { code: 0 },
  events: [],
  links: []
}
Exporting 1 span(s). Converting to envelopes...
Exporting 2 envelope(s)
Sure thing! To train a parrot, matey, first ye need to build trust with the bird. Spend time with it, feed it by hand, and speak to it softly. Then, start with simple commands, like "step up" or "come here". Reward the parrot with treats and praise when it follows the command. Be patient, and practice consistently, and soon ye will have a well-trained parrot, wise as old Captain Flint himself!
Instrumentation suppressed, returning Noop Span
{
  resource: {
    attributes: {
      'service.name': 'howieinsight',
      'telemetry.sdk.language': 'nodejs',
      'telemetry.sdk.name': 'opentelemetry',
      'telemetry.sdk.version': '1.26.0'
    }
  },
  instrumentationScope: { name: 'otel-sample', version: '1.0.0', schemaUrl: undefined },
  traceId: '6c9a1a89426e53ee27d4ac127ccaf832',
  parentId: undefined,
  traceState: undefined,
  name: 'main',
  id: 'f1b4a676d9b58b78',
  kind: 0,
  timestamp: 1725578892476000,
  duration: 4387288.6,
  attributes: { 'az.namespace': 'Microsoft.OtelSample' },
  status: { code: 1 },
  events: [],
  links: []
}
Exporting 1 span(s). Converting to envelopes...
Exporting 2 envelope(s)
*/
