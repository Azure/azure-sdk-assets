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

// Set the logger to the console with verbose logging level
diag.setLogger(new DiagConsoleLogger(), DiagLogLevel.VERBOSE);

// Initialize the NodeTracerProvider
const provider = new NodeTracerProvider();
// Add a SimpleSpanProcessor and ConsoleSpanExporter to the provider
provider.addSpanProcessor(new SimpleSpanProcessor(new ConsoleSpanExporter()));
// Register the provider
provider.register();

// register Azure SDK Instrumentation.
// ************** IMPORTANT SETUP STEP ***************
// this must be done before loading @azure/core-tracing
registerInstrumentations({
  instrumentations: [createAzureSdkInstrumentation()],
});

import ModelClient from "@azure-rest/ai-inference";
import { isUnexpected, ChatChoiceOutput } from "@azure-rest/ai-inference";
import { AzureKeyCredential } from "@azure/core-auth";
import { createTracingClient } from "@azure/core-tracing";

// Create a tracing client
const tracingClient = createTracingClient({
  namespace: "Microsoft.OtelSample",
  packageName: "otel-sample",
  packageVersion: "1.0.0",
});

// Load the .env file if it exists
import "dotenv/config";

// You will need to set these environment variables or edit the following values
const endpoint = process.env["ENDPOINT"] || "<endpoint>";
const key = process.env["KEY"] || "<key>";

async function main() {
  console.log("== Chat Completions Sample ==");

  // Initialize a span named "main" with default options for the spans
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

    const firstChoice = getFirstChoiceToDisplay(response.body.choices);

    console.log(firstChoice);
  });
}

function getFirstChoiceToDisplay(choice: ChatChoiceOutput[]): string[] {
  const allChoices: string = choice.map((arg) => arg.message.content).join(", ");
  const impl = () => {
    return choice[0].message.content;
  };
  return trace("getChoiceToDisplay", arguments, impl, argsAttributeMapper, returnAttributeMapper);
}

function argsAttributeMapper(args: IArguments): Map<string, any> {
  const attributes = new Map<string, any>();
  attributes.set("all-choices", args[0].map((arg) => arg.message.content).join(", "));
  return attributes;
}

function returnAttributeMapper(returnVal: string): Map<string, any> {
  const attributes = new Map<string, any>();
  attributes.set("first-choice", returnVal);
  return attributes;
}


main().catch((err) => {
  console.error("The sample encountered an error:", err);
});

export { main };

/**
 * The following spans will be created:
 */
/*
{
  resource: {
    attributes: {
      'service.name': 'unknown_service:C:\\Program Files\\nodejs\\node.exe',
      'telemetry.sdk.language': 'nodejs',
      'telemetry.sdk.name': 'opentelemetry',
      'telemetry.sdk.version': '1.25.1'
    }
  },
  traceId: '46ebf950b63b90aa29f8dfb2e93b82ed',
  parentId: 'eba5005694622f2d',
  traceState: undefined,
  name: '{gen_ai.operation.name} {gen_ai.request.model}',
  id: 'bf9cab807e0d4cab',
  kind: 0,
  timestamp: 1724451988970000,
  duration: 6496652.4,
  attributes: {
    'gen_ai.operation.name': 'chat; text_completion',
    'gen_ai.system': 'openai',
    'gen_ai.request.max_tokens': 1000,
    'gen_ai.request.temperature': 1,
    'gen_ai.request.top_p': 1,
    'az.namespace': 'Microsoft.OtelSample',
    'gen_ai.response.id': 'b9f3658342bb48e18a6f7d61a362b53c',
    'gen_ai.response.model': 'mistral-large',
    'gen_ai.usage.input_tokens': 57,
    'gen_ai.usage.output_tokens': 157
  },
  status: { code: 1 },
  events: [],
  links: []
}
{
  resource: {
    attributes: {
      'service.name': 'unknown_service:C:\\Program Files\\nodejs\\node.exe',
      'telemetry.sdk.language': 'nodejs',
      'telemetry.sdk.name': 'opentelemetry',
      'telemetry.sdk.version': '1.25.1'
    }
  },
  traceId: '46ebf950b63b90aa29f8dfb2e93b82ed',
  parentId: '4b718583a5c0b378',
  traceState: undefined,
  name: 'HTTP POST',
  id: 'eba5005694622f2d',
  kind: 2,
  timestamp: 1724451988968000,
  duration: 6504423.2,
  attributes: {
    'http.url': 'https://mistral-large-ajmih-serverless.eastus2.inference.ai.azure.com/chat/completions?api-version=2024-05-01-preview',
    'http.method': 'POST',
    'http.user_agent': 'azsdk-js-AiModelInference-rest/1.0.0-beta.2 core-rest-pipeline/1.16.3 Node/20.16.0 OS/(x64-Windows_NT-10.0.22631)',
    requestId: 'b82c6c23-e3c3-41a2-8846-3be66606b470',
    'az.namespace': 'Microsoft.OtelSample',
    'http.status_code': 200
  },
  status: { code: 1 },
  events: [],
  links: []
}
{
  resource: {
    attributes: {
      'service.name': 'unknown_service:C:\\Program Files\\nodejs\\node.exe',
      'telemetry.sdk.language': 'nodejs',
      'telemetry.sdk.name': 'opentelemetry',
      'telemetry.sdk.version': '1.25.1'
    }
  },
  traceId: '46ebf950b63b90aa29f8dfb2e93b82ed',
  parentId: undefined,
  traceState: undefined,
  name: 'main',
  id: '4b718583a5c0b378',
  kind: 0,
  timestamp: 1724451988934000,
  duration: 6543868.6,
  attributes: { 'az.namespace': 'Microsoft.OtelSample' },
  status: { code: 1 },
  events: [],
  links: []
}
{
  resource: {
    attributes: {
      'service.name': 'unknown_service:C:\\Program Files\\nodejs\\node.exe',
      'telemetry.sdk.language': 'nodejs',
      'telemetry.sdk.name': 'opentelemetry',
      'telemetry.sdk.version': '1.25.1'
    }
  },
  traceId: '46ebf950b63b90aa29f8dfb2e93b82ed',
  parentId: '4b718583a5c0b378',
  traceState: undefined,
  name: 'HTTP POST',
  id: 'eba5005694622f2d',
  kind: 2,
  timestamp: 1724451988968000,
  duration: 6504423.2,
  attributes: {
    'http.url': 'https://mistral-large-ajmih-serverless.eastus2.inference.ai.azure.com/chat/completions?api-version=2024-05-01-preview',
    'http.method': 'POST',
    'http.user_agent': 'azsdk-js-AiModelInference-rest/1.0.0-beta.2 core-rest-pipeline/1.16.3 Node/20.16.0 OS/(x64-Windows_NT-10.0.22631)',
    requestId: 'b82c6c23-e3c3-41a2-8846-3be66606b470',
    'az.namespace': 'Microsoft.OtelSample',
    'http.status_code': 200
  },
  status: { code: 1 },
  events: [],
  links: []
}
{
  resource: {
    attributes: {
      'service.name': 'unknown_service:C:\\Program Files\\nodejs\\node.exe',
      'telemetry.sdk.language': 'nodejs',
      'telemetry.sdk.name': 'opentelemetry',
      'telemetry.sdk.version': '1.25.1'
    }
  },
  traceId: '46ebf950b63b90aa29f8dfb2e93b82ed',
  parentId: undefined,
  traceState: undefined,
  name: 'main',
  id: '4b718583a5c0b378',
  kind: 0,
  timestamp: 1724451988934000,
  duration: 6543868.6,
  attributes: { 
    'az.namespace': 'Microsoft.OtelSample',
    'first-choice': 'There are 5,280 feet in a mile. This is a standard unit conversion in the imperial system, where 1 mile is equivalent to 5,280 feet, 1,760 yards, or 8 furlongs.'
    'all-choices': 'There are 5,280 feet in a mile. This is a standard unit conversion in the imperial system, where 1 mile is equivalent to 5,280 feet, 1,760 yards, or 8 furlongs.'
  },
  status: { code: 1 },
  events: [],
  links: []
}

*/
