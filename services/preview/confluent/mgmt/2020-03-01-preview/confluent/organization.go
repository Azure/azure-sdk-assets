package confluent

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
	"net/http"
)

// OrganizationClient is the client for the Organization methods of the Confluent service.
type OrganizationClient struct {
	BaseClient
}

// NewOrganizationClient creates an instance of the OrganizationClient client.
func NewOrganizationClient(subscriptionID string) OrganizationClient {
	return NewOrganizationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOrganizationClientWithBaseURI creates an instance of the OrganizationClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewOrganizationClientWithBaseURI(baseURI string, subscriptionID string) OrganizationClient {
	return OrganizationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create sends the create request.
// Parameters:
// resourceGroupName - resource group name
// organizationName - organization resource name
// body - organization resource model
func (client OrganizationClient) Create(ctx context.Context, resourceGroupName string, organizationName string, body *OrganizationResource) (result OrganizationCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrganizationClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, resourceGroupName, organizationName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client OrganizationClient) CreatePreparer(ctx context.Context, resourceGroupName string, organizationName string, body *OrganizationResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"organizationName":  autorest.Encode("path", organizationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Name = nil
	body.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations/{organizationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if body != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(body))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client OrganizationClient) CreateSender(req *http.Request) (future OrganizationCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client OrganizationClient) (or OrganizationResource, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "confluent.OrganizationCreateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("confluent.OrganizationCreateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		or.Response.Response, err = future.GetResult(sender)
		if or.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "confluent.OrganizationCreateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && or.Response.Response.StatusCode != http.StatusNoContent {
			or, err = client.CreateResponder(or.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "confluent.OrganizationCreateFuture", "Result", or.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client OrganizationClient) CreateResponder(resp *http.Response) (result OrganizationResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete sends the delete request.
// Parameters:
// resourceGroupName - resource group name
// organizationName - organization resource name
func (client OrganizationClient) Delete(ctx context.Context, resourceGroupName string, organizationName string) (result OrganizationDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrganizationClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, organizationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client OrganizationClient) DeletePreparer(ctx context.Context, resourceGroupName string, organizationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"organizationName":  autorest.Encode("path", organizationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations/{organizationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client OrganizationClient) DeleteSender(req *http.Request) (future OrganizationDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client OrganizationClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "confluent.OrganizationDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("confluent.OrganizationDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client OrganizationClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get sends the get request.
// Parameters:
// resourceGroupName - resource group name
// organizationName - organization resource name
func (client OrganizationClient) Get(ctx context.Context, resourceGroupName string, organizationName string) (result OrganizationResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrganizationClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, organizationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client OrganizationClient) GetPreparer(ctx context.Context, resourceGroupName string, organizationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"organizationName":  autorest.Encode("path", organizationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations/{organizationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OrganizationClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OrganizationClient) GetResponder(resp *http.Response) (result OrganizationResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup sends the list by resource group request.
// Parameters:
// resourceGroupName - resource group name
func (client OrganizationClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result OrganizationResourceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrganizationClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.orlr.Response.Response != nil {
				sc = result.orlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.orlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.orlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.orlr.hasNextLink() && result.orlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client OrganizationClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client OrganizationClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client OrganizationClient) ListByResourceGroupResponder(resp *http.Response) (result OrganizationResourceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client OrganizationClient) listByResourceGroupNextResults(ctx context.Context, lastResults OrganizationResourceListResult) (result OrganizationResourceListResult, err error) {
	req, err := lastResults.organizationResourceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "confluent.OrganizationClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "confluent.OrganizationClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client OrganizationClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result OrganizationResourceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrganizationClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListBySubscription sends the list by subscription request.
func (client OrganizationClient) ListBySubscription(ctx context.Context) (result OrganizationResourceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrganizationClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.orlr.Response.Response != nil {
				sc = result.orlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.orlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.orlr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.orlr.hasNextLink() && result.orlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client OrganizationClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Confluent/organizations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client OrganizationClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client OrganizationClient) ListBySubscriptionResponder(resp *http.Response) (result OrganizationResourceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client OrganizationClient) listBySubscriptionNextResults(ctx context.Context, lastResults OrganizationResourceListResult) (result OrganizationResourceListResult, err error) {
	req, err := lastResults.organizationResourceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "confluent.OrganizationClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "confluent.OrganizationClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client OrganizationClient) ListBySubscriptionComplete(ctx context.Context) (result OrganizationResourceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrganizationClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx)
	return
}

// Update sends the update request.
// Parameters:
// resourceGroupName - resource group name
// organizationName - organization resource name
// body - updated Organization resource
func (client OrganizationClient) Update(ctx context.Context, resourceGroupName string, organizationName string, body *OrganizationResourceUpdate) (result OrganizationResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrganizationClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, organizationName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confluent.OrganizationClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client OrganizationClient) UpdatePreparer(ctx context.Context, resourceGroupName string, organizationName string, body *OrganizationResourceUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"organizationName":  autorest.Encode("path", organizationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations/{organizationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if body != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(body))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client OrganizationClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client OrganizationClient) UpdateResponder(resp *http.Response) (result OrganizationResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
