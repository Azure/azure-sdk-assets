package aad

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/domainservices/mgmt/2017-06-01/aad"

// ExternalAccess enumerates the values for external access.
type ExternalAccess string

const (
	// Disabled ...
	Disabled ExternalAccess = "Disabled"
	// Enabled ...
	Enabled ExternalAccess = "Enabled"
)

// PossibleExternalAccessValues returns an array of possible values for the ExternalAccess const type.
func PossibleExternalAccessValues() []ExternalAccess {
	return []ExternalAccess{Disabled, Enabled}
}

// FilteredSync enumerates the values for filtered sync.
type FilteredSync string

const (
	// FilteredSyncDisabled ...
	FilteredSyncDisabled FilteredSync = "Disabled"
	// FilteredSyncEnabled ...
	FilteredSyncEnabled FilteredSync = "Enabled"
)

// PossibleFilteredSyncValues returns an array of possible values for the FilteredSync const type.
func PossibleFilteredSyncValues() []FilteredSync {
	return []FilteredSync{FilteredSyncDisabled, FilteredSyncEnabled}
}

// Ldaps enumerates the values for ldaps.
type Ldaps string

const (
	// LdapsDisabled ...
	LdapsDisabled Ldaps = "Disabled"
	// LdapsEnabled ...
	LdapsEnabled Ldaps = "Enabled"
)

// PossibleLdapsValues returns an array of possible values for the Ldaps const type.
func PossibleLdapsValues() []Ldaps {
	return []Ldaps{LdapsDisabled, LdapsEnabled}
}

// NotifyDcAdmins enumerates the values for notify dc admins.
type NotifyDcAdmins string

const (
	// NotifyDcAdminsDisabled ...
	NotifyDcAdminsDisabled NotifyDcAdmins = "Disabled"
	// NotifyDcAdminsEnabled ...
	NotifyDcAdminsEnabled NotifyDcAdmins = "Enabled"
)

// PossibleNotifyDcAdminsValues returns an array of possible values for the NotifyDcAdmins const type.
func PossibleNotifyDcAdminsValues() []NotifyDcAdmins {
	return []NotifyDcAdmins{NotifyDcAdminsDisabled, NotifyDcAdminsEnabled}
}

// NotifyGlobalAdmins enumerates the values for notify global admins.
type NotifyGlobalAdmins string

const (
	// NotifyGlobalAdminsDisabled ...
	NotifyGlobalAdminsDisabled NotifyGlobalAdmins = "Disabled"
	// NotifyGlobalAdminsEnabled ...
	NotifyGlobalAdminsEnabled NotifyGlobalAdmins = "Enabled"
)

// PossibleNotifyGlobalAdminsValues returns an array of possible values for the NotifyGlobalAdmins const type.
func PossibleNotifyGlobalAdminsValues() []NotifyGlobalAdmins {
	return []NotifyGlobalAdmins{NotifyGlobalAdminsDisabled, NotifyGlobalAdminsEnabled}
}

// NtlmV1 enumerates the values for ntlm v1.
type NtlmV1 string

const (
	// NtlmV1Disabled ...
	NtlmV1Disabled NtlmV1 = "Disabled"
	// NtlmV1Enabled ...
	NtlmV1Enabled NtlmV1 = "Enabled"
)

// PossibleNtlmV1Values returns an array of possible values for the NtlmV1 const type.
func PossibleNtlmV1Values() []NtlmV1 {
	return []NtlmV1{NtlmV1Disabled, NtlmV1Enabled}
}

// SyncNtlmPasswords enumerates the values for sync ntlm passwords.
type SyncNtlmPasswords string

const (
	// SyncNtlmPasswordsDisabled ...
	SyncNtlmPasswordsDisabled SyncNtlmPasswords = "Disabled"
	// SyncNtlmPasswordsEnabled ...
	SyncNtlmPasswordsEnabled SyncNtlmPasswords = "Enabled"
)

// PossibleSyncNtlmPasswordsValues returns an array of possible values for the SyncNtlmPasswords const type.
func PossibleSyncNtlmPasswordsValues() []SyncNtlmPasswords {
	return []SyncNtlmPasswords{SyncNtlmPasswordsDisabled, SyncNtlmPasswordsEnabled}
}

// TLSV1 enumerates the values for tlsv1.
type TLSV1 string

const (
	// TLSV1Disabled ...
	TLSV1Disabled TLSV1 = "Disabled"
	// TLSV1Enabled ...
	TLSV1Enabled TLSV1 = "Enabled"
)

// PossibleTLSV1Values returns an array of possible values for the TLSV1 const type.
func PossibleTLSV1Values() []TLSV1 {
	return []TLSV1{TLSV1Disabled, TLSV1Enabled}
}

// DomainSecuritySettings domain Security Settings
type DomainSecuritySettings struct {
	// NtlmV1 - A flag to determine whether or not NtlmV1 is enabled or disabled. Possible values include: 'NtlmV1Enabled', 'NtlmV1Disabled'
	NtlmV1 NtlmV1 `json:"ntlmV1,omitempty"`
	// TLSV1 - A flag to determine whether or not TlsV1 is enabled or disabled. Possible values include: 'TLSV1Enabled', 'TLSV1Disabled'
	TLSV1 TLSV1 `json:"tlsV1,omitempty"`
	// SyncNtlmPasswords - A flag to determine whether or not SyncNtlmPasswords is enabled or disabled. Possible values include: 'SyncNtlmPasswordsEnabled', 'SyncNtlmPasswordsDisabled'
	SyncNtlmPasswords SyncNtlmPasswords `json:"syncNtlmPasswords,omitempty"`
}

// DomainService domain service.
type DomainService struct {
	autorest.Response `json:"-"`
	// DomainServiceProperties - Domain service properties
	*DomainServiceProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
	// Etag - Resource etag
	Etag *string `json:"etag,omitempty"`
}

// MarshalJSON is the custom marshaler for DomainService.
func (ds DomainService) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ds.DomainServiceProperties != nil {
		objectMap["properties"] = ds.DomainServiceProperties
	}
	if ds.Location != nil {
		objectMap["location"] = ds.Location
	}
	if ds.Tags != nil {
		objectMap["tags"] = ds.Tags
	}
	if ds.Etag != nil {
		objectMap["etag"] = ds.Etag
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for DomainService struct.
func (ds *DomainService) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var domainServiceProperties DomainServiceProperties
				err = json.Unmarshal(*v, &domainServiceProperties)
				if err != nil {
					return err
				}
				ds.DomainServiceProperties = &domainServiceProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ds.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ds.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ds.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				ds.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				ds.Tags = tags
			}
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				ds.Etag = &etag
			}
		}
	}

	return nil
}

// DomainServiceListResult the response from the List Domain Services operation.
type DomainServiceListResult struct {
	autorest.Response `json:"-"`
	// Value - the list of domain services.
	Value *[]DomainService `json:"value,omitempty"`
	// NextLink - READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// DomainServiceListResultIterator provides access to a complete listing of DomainService values.
type DomainServiceListResultIterator struct {
	i    int
	page DomainServiceListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DomainServiceListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainServiceListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *DomainServiceListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DomainServiceListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter DomainServiceListResultIterator) Response() DomainServiceListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DomainServiceListResultIterator) Value() DomainService {
	if !iter.page.NotDone() {
		return DomainService{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the DomainServiceListResultIterator type.
func NewDomainServiceListResultIterator(page DomainServiceListResultPage) DomainServiceListResultIterator {
	return DomainServiceListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (dslr DomainServiceListResult) IsEmpty() bool {
	return dslr.Value == nil || len(*dslr.Value) == 0
}

// domainServiceListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dslr DomainServiceListResult) domainServiceListResultPreparer(ctx context.Context) (*http.Request, error) {
	if dslr.NextLink == nil || len(to.String(dslr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dslr.NextLink)))
}

// DomainServiceListResultPage contains a page of DomainService values.
type DomainServiceListResultPage struct {
	fn   func(context.Context, DomainServiceListResult) (DomainServiceListResult, error)
	dslr DomainServiceListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DomainServiceListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainServiceListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.dslr)
	if err != nil {
		return err
	}
	page.dslr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *DomainServiceListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DomainServiceListResultPage) NotDone() bool {
	return !page.dslr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page DomainServiceListResultPage) Response() DomainServiceListResult {
	return page.dslr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page DomainServiceListResultPage) Values() []DomainService {
	if page.dslr.IsEmpty() {
		return nil
	}
	return *page.dslr.Value
}

// Creates a new instance of the DomainServiceListResultPage type.
func NewDomainServiceListResultPage(getNextPage func(context.Context, DomainServiceListResult) (DomainServiceListResult, error)) DomainServiceListResultPage {
	return DomainServiceListResultPage{fn: getNextPage}
}

// DomainServiceProperties properties of the Domain Service.
type DomainServiceProperties struct {
	// TenantID - READ-ONLY; Azure Active Directory tenant id
	TenantID *string `json:"tenantId,omitempty"`
	// DomainName - The name of the Azure domain that the user would like to deploy Domain Services to.
	DomainName *string `json:"domainName,omitempty"`
	// VnetSiteID - READ-ONLY; Virtual network site id
	VnetSiteID *string `json:"vnetSiteId,omitempty"`
	// SubnetID - The name of the virtual network that Domain Services will be deployed on. The id of the subnet that Domain Services will be deployed on. /virtualNetwork/vnetName/subnets/subnetName.
	SubnetID *string `json:"subnetId,omitempty"`
	// LdapsSettings - Secure LDAP Settings
	LdapsSettings *LdapsSettings `json:"ldapsSettings,omitempty"`
	// HealthLastEvaluated - READ-ONLY; Last domain evaluation run DateTime
	HealthLastEvaluated *date.Time `json:"healthLastEvaluated,omitempty"`
	// HealthMonitors - READ-ONLY; List of Domain Health Monitors
	HealthMonitors *[]HealthMonitor `json:"healthMonitors,omitempty"`
	// HealthAlerts - READ-ONLY; List of Domain Health Alerts
	HealthAlerts *[]HealthAlert `json:"healthAlerts,omitempty"`
	// NotificationSettings - Notification Settings
	NotificationSettings *NotificationSettings `json:"notificationSettings,omitempty"`
	// DomainSecuritySettings - DomainSecurity Settings
	DomainSecuritySettings *DomainSecuritySettings `json:"domainSecuritySettings,omitempty"`
	// FilteredSync - Enabled or Disabled flag to turn on Group-based filtered sync. Possible values include: 'FilteredSyncEnabled', 'FilteredSyncDisabled'
	FilteredSync FilteredSync `json:"filteredSync,omitempty"`
	// DomainControllerIPAddress - READ-ONLY; List of Domain Controller IP Address
	DomainControllerIPAddress *[]string `json:"domainControllerIpAddress,omitempty"`
	// ServiceStatus - READ-ONLY; Status of Domain Service instance
	ServiceStatus *string `json:"serviceStatus,omitempty"`
	// ProvisioningState - READ-ONLY; the current deployment or provisioning state, which only appears in the response.
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// DomainServicesCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type DomainServicesCreateOrUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *DomainServicesCreateOrUpdateFuture) Result(client DomainServicesClient) (ds DomainService, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("aad.DomainServicesCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if ds.Response.Response, err = future.GetResult(sender); err == nil && ds.Response.Response.StatusCode != http.StatusNoContent {
		ds, err = client.CreateOrUpdateResponder(ds.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "aad.DomainServicesCreateOrUpdateFuture", "Result", ds.Response.Response, "Failure responding to request")
		}
	}
	return
}

// DomainServicesDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type DomainServicesDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *DomainServicesDeleteFuture) Result(client DomainServicesClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("aad.DomainServicesDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// DomainServicesUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type DomainServicesUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *DomainServicesUpdateFuture) Result(client DomainServicesClient) (ds DomainService, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("aad.DomainServicesUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if ds.Response.Response, err = future.GetResult(sender); err == nil && ds.Response.Response.StatusCode != http.StatusNoContent {
		ds, err = client.UpdateResponder(ds.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "aad.DomainServicesUpdateFuture", "Result", ds.Response.Response, "Failure responding to request")
		}
	}
	return
}

// HealthAlert health Alert Description
type HealthAlert struct {
	// ID - READ-ONLY; Health Alert Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Health Alert Name
	Name *string `json:"name,omitempty"`
	// Issue - READ-ONLY; Health Alert Issue
	Issue *string `json:"issue,omitempty"`
	// Severity - READ-ONLY; Health Alert Severity
	Severity *string `json:"severity,omitempty"`
	// Raised - READ-ONLY; Health Alert Raised DateTime
	Raised *date.Time `json:"raised,omitempty"`
	// LastDetected - READ-ONLY; Health Alert Last Detected DateTime
	LastDetected *date.Time `json:"lastDetected,omitempty"`
	// ResolutionURI - READ-ONLY; Health Alert TSG Link
	ResolutionURI *string `json:"resolutionUri,omitempty"`
}

// HealthMonitor health Monitor Description
type HealthMonitor struct {
	// ID - READ-ONLY; Health Monitor Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Health Monitor Name
	Name *string `json:"name,omitempty"`
	// Details - READ-ONLY; Health Monitor Details
	Details *string `json:"details,omitempty"`
}

// LdapsSettings secure LDAP Settings
type LdapsSettings struct {
	// Ldaps - A flag to determine whether or not Secure LDAP is enabled or disabled. Possible values include: 'LdapsEnabled', 'LdapsDisabled'
	Ldaps Ldaps `json:"ldaps,omitempty"`
	// PfxCertificate - The certificate required to configure Secure LDAP. The parameter passed here should be a base64encoded representation of the certificate pfx file.
	PfxCertificate *string `json:"pfxCertificate,omitempty"`
	// PfxCertificatePassword - The password to decrypt the provided Secure LDAP certificate pfx file.
	PfxCertificatePassword *string `json:"pfxCertificatePassword,omitempty"`
	// PublicCertificate - READ-ONLY; Public certificate used to configure secure ldap.
	PublicCertificate *string `json:"publicCertificate,omitempty"`
	// CertificateThumbprint - READ-ONLY; Thumbprint of configure ldaps certificate.
	CertificateThumbprint *string `json:"certificateThumbprint,omitempty"`
	// CertificateNotAfter - READ-ONLY; NotAfter DateTime of configure ldaps certificate.
	CertificateNotAfter *date.Time `json:"certificateNotAfter,omitempty"`
	// ExternalAccess - A flag to determine whether or not Secure LDAP access over the internet is enabled or disabled. Possible values include: 'Enabled', 'Disabled'
	ExternalAccess ExternalAccess `json:"externalAccess,omitempty"`
	// ExternalAccessIPAddress - READ-ONLY; External access ip address.
	ExternalAccessIPAddress *string `json:"externalAccessIpAddress,omitempty"`
}

// NotificationSettings settings for notification
type NotificationSettings struct {
	// NotifyGlobalAdmins - Should global admins be notified. Possible values include: 'NotifyGlobalAdminsEnabled', 'NotifyGlobalAdminsDisabled'
	NotifyGlobalAdmins NotifyGlobalAdmins `json:"notifyGlobalAdmins,omitempty"`
	// NotifyDcAdmins - Should domain controller admins be notified. Possible values include: 'NotifyDcAdminsEnabled', 'NotifyDcAdminsDisabled'
	NotifyDcAdmins NotifyDcAdmins `json:"notifyDcAdmins,omitempty"`
	// AdditionalRecipients - The list of additional recipients
	AdditionalRecipients *[]string `json:"additionalRecipients,omitempty"`
}

// OperationDisplayInfo the operation supported by Domain Services.
type OperationDisplayInfo struct {
	// Description - The description of the operation.
	Description *string `json:"description,omitempty"`
	// Operation - The action that users can perform, based on their permission level.
	Operation *string `json:"operation,omitempty"`
	// Provider - Service provider: Domain Services.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
}

// OperationEntity the operation supported by Domain Services.
type OperationEntity struct {
	// Name - Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The operation supported by Domain Services.
	Display *OperationDisplayInfo `json:"display,omitempty"`
	// Origin - The origin of the operation.
	Origin *string `json:"origin,omitempty"`
}

// OperationEntityListResult the list of domain service operation response.
type OperationEntityListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of operations.
	Value *[]OperationEntity `json:"value,omitempty"`
	// NextLink - READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationEntityListResultIterator provides access to a complete listing of OperationEntity values.
type OperationEntityListResultIterator struct {
	i    int
	page OperationEntityListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationEntityListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationEntityListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationEntityListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationEntityListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationEntityListResultIterator) Response() OperationEntityListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationEntityListResultIterator) Value() OperationEntity {
	if !iter.page.NotDone() {
		return OperationEntity{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationEntityListResultIterator type.
func NewOperationEntityListResultIterator(page OperationEntityListResultPage) OperationEntityListResultIterator {
	return OperationEntityListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (oelr OperationEntityListResult) IsEmpty() bool {
	return oelr.Value == nil || len(*oelr.Value) == 0
}

// operationEntityListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (oelr OperationEntityListResult) operationEntityListResultPreparer(ctx context.Context) (*http.Request, error) {
	if oelr.NextLink == nil || len(to.String(oelr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(oelr.NextLink)))
}

// OperationEntityListResultPage contains a page of OperationEntity values.
type OperationEntityListResultPage struct {
	fn   func(context.Context, OperationEntityListResult) (OperationEntityListResult, error)
	oelr OperationEntityListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationEntityListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationEntityListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.oelr)
	if err != nil {
		return err
	}
	page.oelr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationEntityListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationEntityListResultPage) NotDone() bool {
	return !page.oelr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationEntityListResultPage) Response() OperationEntityListResult {
	return page.oelr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationEntityListResultPage) Values() []OperationEntity {
	if page.oelr.IsEmpty() {
		return nil
	}
	return *page.oelr.Value
}

// Creates a new instance of the OperationEntityListResultPage type.
func NewOperationEntityListResultPage(getNextPage func(context.Context, OperationEntityListResult) (OperationEntityListResult, error)) OperationEntityListResultPage {
	return OperationEntityListResultPage{fn: getNextPage}
}

// Resource the Resource model definition.
type Resource struct {
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
	// Etag - Resource etag
	Etag *string `json:"etag,omitempty"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	if r.Etag != nil {
		objectMap["etag"] = r.Etag
	}
	return json.Marshal(objectMap)
}
