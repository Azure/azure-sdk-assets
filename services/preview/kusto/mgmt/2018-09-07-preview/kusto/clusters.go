package kusto

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ClustersClient is the the Azure Kusto management API provides a RESTful set of web services that interact with Azure
// Kusto services to manage your clusters and databases. The API enables you to create, update, and delete clusters and
// databases.
type ClustersClient struct {
	BaseClient
}

// NewClustersClient creates an instance of the ClustersClient client.
func NewClustersClient(subscriptionID string) ClustersClient {
	return NewClustersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewClustersClientWithBaseURI creates an instance of the ClustersClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return ClustersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability checks that the cluster name is valid and is not already in use.
// Parameters:
// location - azure location.
// clusterName - the name of the cluster.
func (client ClustersClient) CheckNameAvailability(ctx context.Context, location string, clusterName ClusterCheckNameRequest) (result CheckNameResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.CheckNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "clusterName.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kusto.ClustersClient", "CheckNameAvailability", err.Error())
	}

	req, err := client.CheckNameAvailabilityPreparer(ctx, location, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "CheckNameAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client ClustersClient) CheckNameAvailabilityPreparer(ctx context.Context, location string, clusterName ClusterCheckNameRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Kusto/locations/{location}/checkNameAvailability", pathParameters),
		autorest.WithJSON(clusterName),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client ClustersClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate create or update a Kusto cluster.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// parameters - the Kusto cluster parameters supplied to the CreateOrUpdate operation.
func (client ClustersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, parameters Cluster) (result ClustersCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Sku", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.Sku.Tier", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("kusto.ClustersClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, clusterName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ClustersClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, parameters Cluster) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) CreateOrUpdateSender(req *http.Request) (future ClustersCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ClustersClient) (c Cluster, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "kusto.ClustersCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("kusto.ClustersCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		c.Response.Response, err = future.GetResult(sender)
		if c.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "kusto.ClustersCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && c.Response.Response.StatusCode != http.StatusNoContent {
			c, err = client.CreateOrUpdateResponder(c.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "kusto.ClustersCreateOrUpdateFuture", "Result", c.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ClustersClient) CreateOrUpdateResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Kusto cluster.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
func (client ClustersClient) Delete(ctx context.Context, resourceGroupName string, clusterName string) (result ClustersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ClustersClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) DeleteSender(req *http.Request) (future ClustersDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ClustersClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "kusto.ClustersDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("kusto.ClustersDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ClustersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a Kusto cluster.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
func (client ClustersClient) Get(ctx context.Context, resourceGroupName string, clusterName string) (result Cluster, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ClustersClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ClustersClient) GetResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all Kusto clusters within a subscription.
func (client ClustersClient) List(ctx context.Context) (result ClusterListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ClustersClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Kusto/clusters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ClustersClient) ListResponder(resp *http.Response) (result ClusterListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup lists all Kusto clusters within a resource group.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
func (client ClustersClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ClusterListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ClustersClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ClustersClient) ListByResourceGroupResponder(resp *http.Response) (result ClusterListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListSkus lists eligible SKUs for Kusto resource provider.
func (client ClustersClient) ListSkus(ctx context.Context) (result ListSkusResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.ListSkus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListSkusPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "ListSkus", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSkusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "ListSkus", resp, "Failure sending request")
		return
	}

	result, err = client.ListSkusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "ListSkus", resp, "Failure responding to request")
		return
	}

	return
}

// ListSkusPreparer prepares the ListSkus request.
func (client ClustersClient) ListSkusPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Kusto/skus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSkusSender sends the ListSkus request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) ListSkusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListSkusResponder handles the response to the ListSkus request. The method always
// closes the http.Response Body.
func (client ClustersClient) ListSkusResponder(resp *http.Response) (result ListSkusResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListSkusByResource returns the SKUs available for the provided resource.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
func (client ClustersClient) ListSkusByResource(ctx context.Context, resourceGroupName string, clusterName string) (result ListResourceSkusResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.ListSkusByResource")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListSkusByResourcePreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "ListSkusByResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSkusByResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "ListSkusByResource", resp, "Failure sending request")
		return
	}

	result, err = client.ListSkusByResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "ListSkusByResource", resp, "Failure responding to request")
		return
	}

	return
}

// ListSkusByResourcePreparer prepares the ListSkusByResource request.
func (client ClustersClient) ListSkusByResourcePreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/skus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSkusByResourceSender sends the ListSkusByResource request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) ListSkusByResourceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListSkusByResourceResponder handles the response to the ListSkusByResource request. The method always
// closes the http.Response Body.
func (client ClustersClient) ListSkusByResourceResponder(resp *http.Response) (result ListResourceSkusResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Start starts a Kusto cluster.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
func (client ClustersClient) Start(ctx context.Context, resourceGroupName string, clusterName string) (result ClustersStartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.Start")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StartPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "Start", nil, "Failure preparing request")
		return
	}

	result, err = client.StartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "Start", nil, "Failure sending request")
		return
	}

	return
}

// StartPreparer prepares the Start request.
func (client ClustersClient) StartPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) StartSender(req *http.Request) (future ClustersStartFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ClustersClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "kusto.ClustersStartFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("kusto.ClustersStartFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client ClustersClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop stops a Kusto cluster.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
func (client ClustersClient) Stop(ctx context.Context, resourceGroupName string, clusterName string) (result ClustersStopFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.Stop")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StopPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "Stop", nil, "Failure preparing request")
		return
	}

	result, err = client.StopSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "Stop", nil, "Failure sending request")
		return
	}

	return
}

// StopPreparer prepares the Stop request.
func (client ClustersClient) StopPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) StopSender(req *http.Request) (future ClustersStopFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ClustersClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "kusto.ClustersStopFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("kusto.ClustersStopFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client ClustersClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update update a Kusto cluster.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// parameters - the Kusto cluster parameters supplied to the Update operation.
func (client ClustersClient) Update(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterUpdate) (result ClustersUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClustersClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, clusterName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClustersClient", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ClustersClient) UpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) UpdateSender(req *http.Request) (future ClustersUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ClustersClient) (c Cluster, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "kusto.ClustersUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("kusto.ClustersUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		c.Response.Response, err = future.GetResult(sender)
		if c.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "kusto.ClustersUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && c.Response.Response.StatusCode != http.StatusNoContent {
			c, err = client.UpdateResponder(c.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "kusto.ClustersUpdateFuture", "Result", c.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ClustersClient) UpdateResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
