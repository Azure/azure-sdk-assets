package automanage

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

// ConfigurationProfileAssignmentsClient is the automanage Client
type ConfigurationProfileAssignmentsClient struct {
	BaseClient
}

// NewConfigurationProfileAssignmentsClient creates an instance of the ConfigurationProfileAssignmentsClient client.
func NewConfigurationProfileAssignmentsClient(subscriptionID string) ConfigurationProfileAssignmentsClient {
	return NewConfigurationProfileAssignmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConfigurationProfileAssignmentsClientWithBaseURI creates an instance of the ConfigurationProfileAssignmentsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewConfigurationProfileAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationProfileAssignmentsClient {
	return ConfigurationProfileAssignmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates an association between a VM and Automanage configuration profile
// Parameters:
// configurationProfileAssignmentName - name of the configuration profile assignment. Only default is
// supported.
// parameters - parameters supplied to the create or update configuration profile assignment.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// VMName - the name of the virtual machine.
func (client ConfigurationProfileAssignmentsClient) CreateOrUpdate(ctx context.Context, configurationProfileAssignmentName string, parameters ConfigurationProfileAssignment, resourceGroupName string, VMName string) (result ConfigurationProfileAssignmentsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationProfileAssignmentsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automanage.ConfigurationProfileAssignmentsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, configurationProfileAssignmentName, parameters, resourceGroupName, VMName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ConfigurationProfileAssignmentsClient) CreateOrUpdatePreparer(ctx context.Context, configurationProfileAssignmentName string, parameters ConfigurationProfileAssignment, resourceGroupName string, VMName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationProfileAssignmentName": autorest.Encode("path", configurationProfileAssignmentName),
		"resourceGroupName":                  autorest.Encode("path", resourceGroupName),
		"subscriptionId":                     autorest.Encode("path", client.SubscriptionID),
		"vmName":                             autorest.Encode("path", VMName),
	}

	const APIVersion = "2020-06-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationProfileAssignmentsClient) CreateOrUpdateSender(req *http.Request) (future ConfigurationProfileAssignmentsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ConfigurationProfileAssignmentsClient) (cpa ConfigurationProfileAssignment, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("automanage.ConfigurationProfileAssignmentsCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		cpa.Response.Response, err = future.GetResult(sender)
		if cpa.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && cpa.Response.Response.StatusCode != http.StatusNoContent {
			cpa, err = client.CreateOrUpdateResponder(cpa.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsCreateOrUpdateFuture", "Result", cpa.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ConfigurationProfileAssignmentsClient) CreateOrUpdateResponder(resp *http.Response) (result ConfigurationProfileAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a configuration profile assignment
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// configurationProfileAssignmentName - name of the configuration profile assignment
// VMName - the name of the virtual machine.
func (client ConfigurationProfileAssignmentsClient) Delete(ctx context.Context, resourceGroupName string, configurationProfileAssignmentName string, VMName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationProfileAssignmentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automanage.ConfigurationProfileAssignmentsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, configurationProfileAssignmentName, VMName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ConfigurationProfileAssignmentsClient) DeletePreparer(ctx context.Context, resourceGroupName string, configurationProfileAssignmentName string, VMName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationProfileAssignmentName": autorest.Encode("path", configurationProfileAssignmentName),
		"resourceGroupName":                  autorest.Encode("path", resourceGroupName),
		"subscriptionId":                     autorest.Encode("path", client.SubscriptionID),
		"vmName":                             autorest.Encode("path", VMName),
	}

	const APIVersion = "2020-06-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationProfileAssignmentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ConfigurationProfileAssignmentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get information about a configuration profile assignment
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// configurationProfileAssignmentName - the configuration profile assignment name.
// VMName - the name of the virtual machine.
func (client ConfigurationProfileAssignmentsClient) Get(ctx context.Context, resourceGroupName string, configurationProfileAssignmentName string, VMName string) (result ConfigurationProfileAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationProfileAssignmentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automanage.ConfigurationProfileAssignmentsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, configurationProfileAssignmentName, VMName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConfigurationProfileAssignmentsClient) GetPreparer(ctx context.Context, resourceGroupName string, configurationProfileAssignmentName string, VMName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationProfileAssignmentName": autorest.Encode("path", configurationProfileAssignmentName),
		"resourceGroupName":                  autorest.Encode("path", resourceGroupName),
		"subscriptionId":                     autorest.Encode("path", client.SubscriptionID),
		"vmName":                             autorest.Encode("path", VMName),
	}

	const APIVersion = "2020-06-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationProfileAssignmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConfigurationProfileAssignmentsClient) GetResponder(resp *http.Response) (result ConfigurationProfileAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get list of configuration profile assignments
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
func (client ConfigurationProfileAssignmentsClient) List(ctx context.Context, resourceGroupName string) (result ConfigurationProfileAssignmentList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationProfileAssignmentsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automanage.ConfigurationProfileAssignmentsClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ConfigurationProfileAssignmentsClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfileAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationProfileAssignmentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ConfigurationProfileAssignmentsClient) ListResponder(resp *http.Response) (result ConfigurationProfileAssignmentList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscription get list of configuration profile assignments under a given subscription
func (client ConfigurationProfileAssignmentsClient) ListBySubscription(ctx context.Context) (result ConfigurationProfileAssignmentList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationProfileAssignmentsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automanage.ConfigurationProfileAssignmentsClient", "ListBySubscription", err.Error())
	}

	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileAssignmentsClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client ConfigurationProfileAssignmentsClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Automanage/configurationProfileAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationProfileAssignmentsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client ConfigurationProfileAssignmentsClient) ListBySubscriptionResponder(resp *http.Response) (result ConfigurationProfileAssignmentList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
