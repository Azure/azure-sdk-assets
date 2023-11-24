//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// TopicsClient contains the methods for the Topics group.
// Don't use this type directly, use NewTopicsClient() instead.
type TopicsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTopicsClient creates a new instance of TopicsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTopicsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TopicsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TopicsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a topic in the specified namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - topicName - The topic name.
//   - parameters - Parameters supplied to create a topic resource.
//   - options - TopicsClientCreateOrUpdateOptions contains the optional parameters for the TopicsClient.CreateOrUpdate method.
func (client *TopicsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, parameters SBTopic, options *TopicsClientCreateOrUpdateOptions) (TopicsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "TopicsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, topicName, parameters, options)
	if err != nil {
		return TopicsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TopicsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TopicsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TopicsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, parameters SBTopic, options *TopicsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TopicsClient) createOrUpdateHandleResponse(resp *http.Response) (TopicsClientCreateOrUpdateResponse, error) {
	result := TopicsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBTopic); err != nil {
		return TopicsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateAuthorizationRule - Creates an authorization rule for the specified topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - topicName - The topic name.
//   - authorizationRuleName - The authorization rule name.
//   - parameters - The shared access authorization rule.
//   - options - TopicsClientCreateOrUpdateAuthorizationRuleOptions contains the optional parameters for the TopicsClient.CreateOrUpdateAuthorizationRule
//     method.
func (client *TopicsClient) CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, parameters SBAuthorizationRule, options *TopicsClientCreateOrUpdateAuthorizationRuleOptions) (TopicsClientCreateOrUpdateAuthorizationRuleResponse, error) {
	var err error
	const operationName = "TopicsClient.CreateOrUpdateAuthorizationRule"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, topicName, authorizationRuleName, parameters, options)
	if err != nil {
		return TopicsClientCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TopicsClientCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TopicsClientCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	resp, err := client.createOrUpdateAuthorizationRuleHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateAuthorizationRuleCreateRequest creates the CreateOrUpdateAuthorizationRule request.
func (client *TopicsClient) createOrUpdateAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, parameters SBAuthorizationRule, options *TopicsClientCreateOrUpdateAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateAuthorizationRuleHandleResponse handles the CreateOrUpdateAuthorizationRule response.
func (client *TopicsClient) createOrUpdateAuthorizationRuleHandleResponse(resp *http.Response) (TopicsClientCreateOrUpdateAuthorizationRuleResponse, error) {
	result := TopicsClientCreateOrUpdateAuthorizationRuleResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBAuthorizationRule); err != nil {
		return TopicsClientCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a topic from the specified namespace and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - topicName - The topic name.
//   - options - TopicsClientDeleteOptions contains the optional parameters for the TopicsClient.Delete method.
func (client *TopicsClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, options *TopicsClientDeleteOptions) (TopicsClientDeleteResponse, error) {
	var err error
	const operationName = "TopicsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, topicName, options)
	if err != nil {
		return TopicsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TopicsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return TopicsClientDeleteResponse{}, err
	}
	return TopicsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TopicsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, options *TopicsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteAuthorizationRule - Deletes a topic authorization rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - topicName - The topic name.
//   - authorizationRuleName - The authorization rule name.
//   - options - TopicsClientDeleteAuthorizationRuleOptions contains the optional parameters for the TopicsClient.DeleteAuthorizationRule
//     method.
func (client *TopicsClient) DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, options *TopicsClientDeleteAuthorizationRuleOptions) (TopicsClientDeleteAuthorizationRuleResponse, error) {
	var err error
	const operationName = "TopicsClient.DeleteAuthorizationRule"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, topicName, authorizationRuleName, options)
	if err != nil {
		return TopicsClientDeleteAuthorizationRuleResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TopicsClientDeleteAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return TopicsClientDeleteAuthorizationRuleResponse{}, err
	}
	return TopicsClientDeleteAuthorizationRuleResponse{}, nil
}

// deleteAuthorizationRuleCreateRequest creates the DeleteAuthorizationRule request.
func (client *TopicsClient) deleteAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, options *TopicsClientDeleteAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns a description for the specified topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - topicName - The topic name.
//   - options - TopicsClientGetOptions contains the optional parameters for the TopicsClient.Get method.
func (client *TopicsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, options *TopicsClientGetOptions) (TopicsClientGetResponse, error) {
	var err error
	const operationName = "TopicsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, topicName, options)
	if err != nil {
		return TopicsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TopicsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TopicsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TopicsClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, options *TopicsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TopicsClient) getHandleResponse(resp *http.Response) (TopicsClientGetResponse, error) {
	result := TopicsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBTopic); err != nil {
		return TopicsClientGetResponse{}, err
	}
	return result, nil
}

// GetAuthorizationRule - Returns the specified authorization rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - topicName - The topic name.
//   - authorizationRuleName - The authorization rule name.
//   - options - TopicsClientGetAuthorizationRuleOptions contains the optional parameters for the TopicsClient.GetAuthorizationRule
//     method.
func (client *TopicsClient) GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, options *TopicsClientGetAuthorizationRuleOptions) (TopicsClientGetAuthorizationRuleResponse, error) {
	var err error
	const operationName = "TopicsClient.GetAuthorizationRule"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, topicName, authorizationRuleName, options)
	if err != nil {
		return TopicsClientGetAuthorizationRuleResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TopicsClientGetAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TopicsClientGetAuthorizationRuleResponse{}, err
	}
	resp, err := client.getAuthorizationRuleHandleResponse(httpResp)
	return resp, err
}

// getAuthorizationRuleCreateRequest creates the GetAuthorizationRule request.
func (client *TopicsClient) getAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, options *TopicsClientGetAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAuthorizationRuleHandleResponse handles the GetAuthorizationRule response.
func (client *TopicsClient) getAuthorizationRuleHandleResponse(resp *http.Response) (TopicsClientGetAuthorizationRuleResponse, error) {
	result := TopicsClientGetAuthorizationRuleResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBAuthorizationRule); err != nil {
		return TopicsClientGetAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// NewListAuthorizationRulesPager - Gets authorization rules for a topic.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - topicName - The topic name.
//   - options - TopicsClientListAuthorizationRulesOptions contains the optional parameters for the TopicsClient.NewListAuthorizationRulesPager
//     method.
func (client *TopicsClient) NewListAuthorizationRulesPager(resourceGroupName string, namespaceName string, topicName string, options *TopicsClientListAuthorizationRulesOptions) *runtime.Pager[TopicsClientListAuthorizationRulesResponse] {
	return runtime.NewPager(runtime.PagingHandler[TopicsClientListAuthorizationRulesResponse]{
		More: func(page TopicsClientListAuthorizationRulesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TopicsClientListAuthorizationRulesResponse) (TopicsClientListAuthorizationRulesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TopicsClient.NewListAuthorizationRulesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listAuthorizationRulesCreateRequest(ctx, resourceGroupName, namespaceName, topicName, options)
			}, nil)
			if err != nil {
				return TopicsClientListAuthorizationRulesResponse{}, err
			}
			return client.listAuthorizationRulesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAuthorizationRulesCreateRequest creates the ListAuthorizationRules request.
func (client *TopicsClient) listAuthorizationRulesCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, options *TopicsClientListAuthorizationRulesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAuthorizationRulesHandleResponse handles the ListAuthorizationRules response.
func (client *TopicsClient) listAuthorizationRulesHandleResponse(resp *http.Response) (TopicsClientListAuthorizationRulesResponse, error) {
	result := TopicsClientListAuthorizationRulesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBAuthorizationRuleListResult); err != nil {
		return TopicsClientListAuthorizationRulesResponse{}, err
	}
	return result, nil
}

// NewListByNamespacePager - Gets all the topics in a namespace.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - options - TopicsClientListByNamespaceOptions contains the optional parameters for the TopicsClient.NewListByNamespacePager
//     method.
func (client *TopicsClient) NewListByNamespacePager(resourceGroupName string, namespaceName string, options *TopicsClientListByNamespaceOptions) *runtime.Pager[TopicsClientListByNamespaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[TopicsClientListByNamespaceResponse]{
		More: func(page TopicsClientListByNamespaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TopicsClientListByNamespaceResponse) (TopicsClientListByNamespaceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TopicsClient.NewListByNamespacePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByNamespaceCreateRequest(ctx, resourceGroupName, namespaceName, options)
			}, nil)
			if err != nil {
				return TopicsClientListByNamespaceResponse{}, err
			}
			return client.listByNamespaceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByNamespaceCreateRequest creates the ListByNamespace request.
func (client *TopicsClient) listByNamespaceCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *TopicsClientListByNamespaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByNamespaceHandleResponse handles the ListByNamespace response.
func (client *TopicsClient) listByNamespaceHandleResponse(resp *http.Response) (TopicsClientListByNamespaceResponse, error) {
	result := TopicsClientListByNamespaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBTopicListResult); err != nil {
		return TopicsClientListByNamespaceResponse{}, err
	}
	return result, nil
}

// ListKeys - Gets the primary and secondary connection strings for the topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - topicName - The topic name.
//   - authorizationRuleName - The authorization rule name.
//   - options - TopicsClientListKeysOptions contains the optional parameters for the TopicsClient.ListKeys method.
func (client *TopicsClient) ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, options *TopicsClientListKeysOptions) (TopicsClientListKeysResponse, error) {
	var err error
	const operationName = "TopicsClient.ListKeys"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, namespaceName, topicName, authorizationRuleName, options)
	if err != nil {
		return TopicsClientListKeysResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TopicsClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TopicsClientListKeysResponse{}, err
	}
	resp, err := client.listKeysHandleResponse(httpResp)
	return resp, err
}

// listKeysCreateRequest creates the ListKeys request.
func (client *TopicsClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, options *TopicsClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}/ListKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *TopicsClient) listKeysHandleResponse(resp *http.Response) (TopicsClientListKeysResponse, error) {
	result := TopicsClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessKeys); err != nil {
		return TopicsClientListKeysResponse{}, err
	}
	return result, nil
}

// RegenerateKeys - Regenerates primary or secondary connection strings for the topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - topicName - The topic name.
//   - authorizationRuleName - The authorization rule name.
//   - parameters - Parameters supplied to regenerate the authorization rule.
//   - options - TopicsClientRegenerateKeysOptions contains the optional parameters for the TopicsClient.RegenerateKeys method.
func (client *TopicsClient) RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *TopicsClientRegenerateKeysOptions) (TopicsClientRegenerateKeysResponse, error) {
	var err error
	const operationName = "TopicsClient.RegenerateKeys"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.regenerateKeysCreateRequest(ctx, resourceGroupName, namespaceName, topicName, authorizationRuleName, parameters, options)
	if err != nil {
		return TopicsClientRegenerateKeysResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TopicsClientRegenerateKeysResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TopicsClientRegenerateKeysResponse{}, err
	}
	resp, err := client.regenerateKeysHandleResponse(httpResp)
	return resp, err
}

// regenerateKeysCreateRequest creates the RegenerateKeys request.
func (client *TopicsClient) regenerateKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *TopicsClientRegenerateKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}/regenerateKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// regenerateKeysHandleResponse handles the RegenerateKeys response.
func (client *TopicsClient) regenerateKeysHandleResponse(resp *http.Response) (TopicsClientRegenerateKeysResponse, error) {
	result := TopicsClientRegenerateKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessKeys); err != nil {
		return TopicsClientRegenerateKeysResponse{}, err
	}
	return result, nil
}
