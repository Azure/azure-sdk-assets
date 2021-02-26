package authoring

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"github.com/gofrs/uuid"
	"net/http"
)

// TrainClient is the client for the Train methods of the Authoring service.
type TrainClient struct {
	BaseClient
}

// NewTrainClient creates an instance of the TrainClient client.
func NewTrainClient(endpoint string) TrainClient {
	return TrainClient{New(endpoint)}
}

// GetStatus gets the training status of all models (intents and entities) for the specified LUIS app. You must call
// the train API to train the LUIS app before you call this API to get training status. "appID" specifies the LUIS app
// ID. "versionId" specifies the version number of the LUIS app. For example, "0.1".
// Parameters:
// appID - the application ID.
// versionID - the version ID.
func (client TrainClient) GetStatus(ctx context.Context, appID uuid.UUID, versionID string) (result ListModelTrainingInfo, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TrainClient.GetStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetStatusPreparer(ctx, appID, versionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.TrainClient", "GetStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.TrainClient", "GetStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.TrainClient", "GetStatus", resp, "Failure responding to request")
		return
	}

	return
}

// GetStatusPreparer prepares the GetStatus request.
func (client TrainClient) GetStatusPreparer(ctx context.Context, appID uuid.UUID, versionID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/train", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetStatusSender sends the GetStatus request. The method will close the
// http.Response Body if it receives an error.
func (client TrainClient) GetStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetStatusResponder handles the response to the GetStatus request. The method always
// closes the http.Response Body.
func (client TrainClient) GetStatusResponder(resp *http.Response) (result ListModelTrainingInfo, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// TrainVersion sends a training request for a version of a specified LUIS app. This POST request initiates a request
// asynchronously. To determine whether the training request is successful, submit a GET request to get training
// status. Note: The application version is not fully trained unless all the models (intents and entities) are trained
// successfully or are up to date. To verify training success, get the training status at least once after training is
// complete.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
func (client TrainClient) TrainVersion(ctx context.Context, appID uuid.UUID, versionID string) (result EnqueueTrainingResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TrainClient.TrainVersion")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.TrainVersionPreparer(ctx, appID, versionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.TrainClient", "TrainVersion", nil, "Failure preparing request")
		return
	}

	resp, err := client.TrainVersionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.TrainClient", "TrainVersion", resp, "Failure sending request")
		return
	}

	result, err = client.TrainVersionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.TrainClient", "TrainVersion", resp, "Failure responding to request")
		return
	}

	return
}

// TrainVersionPreparer prepares the TrainVersion request.
func (client TrainClient) TrainVersionPreparer(ctx context.Context, appID uuid.UUID, versionID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/train", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TrainVersionSender sends the TrainVersion request. The method will close the
// http.Response Body if it receives an error.
func (client TrainClient) TrainVersionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// TrainVersionResponder handles the response to the TrainVersion request. The method always
// closes the http.Response Body.
func (client TrainClient) TrainVersionResponder(resp *http.Response) (result EnqueueTrainingResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
