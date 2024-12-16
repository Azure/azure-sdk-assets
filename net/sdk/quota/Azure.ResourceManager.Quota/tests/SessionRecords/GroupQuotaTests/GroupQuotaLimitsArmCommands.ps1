armclient GET "https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/java-sdk-test-gq/resourceProviders/Microsoft.Compute/groupQuotaLimits/cores?api-version=2023-06-01-preview&$filter=location eq westus"

armclient GET https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/java-sdk-test-gq/subscriptions?api-version=2023-06-01-preview

//Group Quota GET
armclient Get  https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/configGq?api-version=2023-06-01-preview -verbose
//Group Quota PUT
armclient PUT https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/billTejasGQ?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreateBillingId.json -verbose
//Group Quota GETLIST
armclient Get  https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas?api-version=2023-06-01-preview -verbose
//Group Quota DELETE
armclient Delete  https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/billgroupQuota?api-version=2023-06-01-preview -verbose

//SDK TEST
armclient PUT https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreate.json -verbose
armclient Delete https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota?api-version=2023-06-01-preview -verbose
armclient GET https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota?api-version=2023-06-01-preview -verbose


//GroupQuotasSubscriptions GET
armclient Get https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/gq-tejas-1-new/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview -verbose
//GroupQuotasSubscriptions PUT
armclient PUT https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/configGq/subscriptions/9f6cce51-6baf-4de5-a3c4-6f58b85315b9?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreateBillingId.json -verbose
//GroupQuotasSubscriptions GET-List
armclient Get https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/pdp-tejas-gq/subscriptions?api-version=2023-06-01-preview -verbose
//GroupQuotasSubscriptions DELETE
armclient DELETE https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/pdp-tejas-gq/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview -verbose


//Limit GET
armclient Get https://management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/gq-tejas/groupQuotaLimits/standardav2family?api-version=2023-06-01-preview&$filter=provider eq microsoft.compute and location eq westus" -verbose
//LIMIT GET V2
armclient GET https://centraluseuap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/gq-tejas-1-new/resourceProviders/Microsoft.Compute/groupQuotaLimits/standardav2family?api-version=2023-06-01-preview&$filter=location eq westus " -verbose

//Limit GETLIST
armclient Get https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/tj-gq-groupquota1/groupQuotaLimits?api-version=2023-06-01-preview&$filter=provider eq microsoft.compute and location eq westus" -verbose
//Limit PUT V1
armclient PUT https://centraluseuap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/gq-tejas-1-new/groupQuotaLimits/standardav2family/?api-version=2023-06-01-preview" @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaLimit.json -verbose

//Limit GET V1
armclient GET https://centraluseuap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/gq-tejas-1-new/resourceProviders/Microsoft.Compute/groupQuotaLimits/standardav2family?api-version=2023-06-01-preview&$filter=provider eq microsoft.compute and location eq westus" -verbose

armclient PATCH https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/gq-tejas/groupQuotaLimits/standardav2family/?api-version=2023-06-01-preview&$filter=provider eq microsoft.compute and location eq westus" @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaLimit.json -verbose

//QuotaLimitOperationStatus V1
armclient GET https://centraluseuap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/gq-tejas-1-new/groupQuotaOperationsStatus/764c9a2b-af56-4ba1-b440-cf2d295828f8?api-version=2023-06-01-preview&$filter=provider eq microsoft.compute and location eq westus"


//Allocation GET DV2 V1
armclient GET https://centraluseuap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/gq-tejas-1-new/quotaAllocations/standardav2family/?api-version=2023-06-01-preview&$filter=provider eq microsoft.compute and location eq westus2" -verbose
//Allocation GET LIST V1
armclient GET https://centraluseuap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/gq-tejas-1-new/quotaAllocations?api-version=2023-06-01-preview&$filter=provider eq microsoft.compute and location eq westus" -verbose
//Allocation Requests PATCH V1
armclient PUT https://centraluseuap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/gq-tejas-1-new/groupQuotaRequests/?api-version=2023-06-01-preview&$filter=provider eq microsoft.compute and location eq westus" -verbose
//Allocation Requests PUT
armclient PUT https://centraluseuap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/gq-tejas-1-new/resourceProviders/Microsoft.Compute/quotaAllocationRequests/standardav2family?api-version=2023-06-01-preview" @C:\Users\tejasma\GroupQuotaLimitPayloads\allocateQuotaRequest.json -verbose

//Allocation PUT V1
armclient Put https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/gq-tejas-1-new/quotaAllocations/standardav2family?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocation.json -verbose

//AllocationRequest PUT
armclient PUT https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/9f6cce51-6baf-4de5-a3c4-6f58b85315b9/providers/Microsoft.Quota/groupQuotas/configGq/resourceProviders/Microsoft.Compute/quotaAllocationRequests/standardav2family?api-version=2023-06-01-preview&$filter=provider eq microsoft.compute and location eq westus" @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocation.json -verbose
//AllocationRequest PATCH
armclient PATCH https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/billGroupQuota/quotaAllocations/standardav2family?api-version=2023-06-01-preview&$filter=provider eq microsoft.compute and location eq westus" @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocation.json -verbose



//Allocation Requests PATCH
armclient PUT https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/gq-tejas-1-new/groupQuotaRequests/?api-version=2023-06-01-preview&$filter=provider eq microsoft.compute and location eq westus" -verbose
//Allocation Requests PUT
armclient PUT https://centraluseuap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/gq-tejas-1-new/resourceProviders/Microsoft.Compute/quotaAllocationRequests/standardav2family?api-version=2023-06-01-preview" @C:\Users\tejasma\GroupQuotaLimitPayloads\allocateQuotaRequest.json -verbose
//Allocation Requests GET
armclient GET https://centraluseuap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/gq-tejas-1-new/resourceProviders/Microsoft.Compute/quotaAllocationRequests?api-version=2023-06-01-preview&$filter=location eq westus " -verbose


//QuotaLimitOperationStatus
armclient GET https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/gq-tejas/groupQuotaOperationsStatus/764c9a2b-af56-4ba1-b440-cf2d295828f8?api-version=2023-06-01-preview&$filter=provider eq microsoft.compute and location eq westus"
//QuotaLimitOperationStatus
armclient GET  https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/gq-tejas/quotaAllocationOperationsStatus/c2d14f44-053a-4125-9fa3-25d2c2a198d9?api-version=2023-06-01-preview&$filter=provider eq microsoft.compute and location eq westus"


armclient delete  https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/gq-tejas-1-new/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview

armclient delete  https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota?api-version=2023-06-01-preview
armclient PATCH  https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreate.json -verbose
armclient delete  https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota?api-version=2023-06-01-preview
armclient PUT  https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota-121?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreate.json -verbose

//LIMIT V2
armclient PUT  https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/groupQuotaRequests/standarddv2family?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaLimitV2.json -verbose
armclient GET  https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota?api-version=2023-06-01-preview -verbose


armclient GET  https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota?api-version=2023-06-01-preview -verbose
armclient delete  https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota-create?api-version=2023-06-01-preview
armclient GET  https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota-create?api-version=2023-06-01-preview
armclient PATCH  https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreate.json -verbose

//GroupQuotasSubscriptions PUT
armclient PUT https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreate.json -verbose
armclient DELETE https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview -verbose

//GroupQuotasSubscriptionsAllocations
armclient PUT https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/quotaAllocationRequests/standarddv2family?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocation.json -verbose

armclient PUT https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/tgProdTest/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreate.json -verbose
armclient PUT https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/tgProdTest?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreate.json -verbose

armclient DELETE https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/gq-tejas/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview -verbose

armclient GET https://management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/tgProdTest/resourceProviders/Microsoft.Compute/groupQuotaLimits?api-version=2023-06-01-preview&$filter=location eq westus2" -verbose

armclient Patch https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testgroupquota/resourceProviders/Microsoft.Compute/locationSettings/southcentralusstg?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\enforcementBody.json -verbose



//CANARY TESTING
//Group Quota GET
armclient Get  https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota?api-version=2023-06-01-preview -verbose
//Group Quota PUT
armclient PUT https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/gq-tejas-test-canary?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreate.json -verbose
//Group Quota GETLIST
armclient Get  https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas?api-version=2023-06-01-preview -verbose
//Group Quota DELETE
armclient Delete  https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/billgroupQuota?api-version=2023-06-01-preview -verbose

//GroupQuotasSubscriptions GET
armclient Get https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview -verbose
//GroupQuotasSubscriptions PUT
armclient PUT https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocation.json -verbose
//GroupQuotasSubscriptions GET-List
armclient Get https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/subscriptions?api-version=2023-06-01-preview -verbose
//GroupQuotasSubscriptions DELETE
armclient DELETE https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview -verbose

//GroupQuotasSubscriptionsAllocations
armclient PUT https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/quotaAllocationRequests/standarddv2family?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocation.json -verbose
armclient DELETE https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/gq-tejas/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview -verbose

armclient GET https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/quotaAllocations/standarddv2family?api-version=2023-06-01-preview&$filter=location eq westus2" -verbose

armclient Patch https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testgroupquota/resourceProviders/Microsoft.Compute/locationSettings/southcentralusstg?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\enforcementBody.json -verbose

//LIMIT Request V2
armclient POST  https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/groupQuotaRequests/standarddv2family?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaLimitV2.json -verbose
armclient GET  https://centraluseuap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/groupQuotaRequests/d28bd1d6-8554-41a6-8f08-f6406744cd07?api-version=2023-06-01-preview&$filter=location eq westus2" -verbose
//LIMIT V2
armclient GET  https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/groupQuotaLimits/standarddv4family?api-version=2023-06-01-preview&$filter=location eq westus" -verbose




//GroupQuotaEnforcement Tests
//Group Quota PUT
armclient PUT https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/gq-tejas-enforcement?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreateBillingId.json -verbose
//Enforcement PUT/PATCH
armclient PUT https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/gq-tejas-enforcement/resourceProviders/Microsoft.Compute/locationSettings/southcentralusstg?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\enforcementEnabled.json -verbose

//Group Quota DELETE
armclient Delete  https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/putGroupQuotaSubscriptionAllocationGroup?api-version=2023-06-01-preview -verbose


//SDK TESTS

//GetGroupQuota
armclient Get https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-tejas-test?api-version=2023-06-01-preview -verbose
//GetGroupQuotaLimit
armclient GET https://management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/groupQuotaLimits/standarddv4family?api-version=2023-06-01-preview&$filter=location eq westus" -verbose
//GetGroupQuotaLimitList
armclient GET https://management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/groupQuotaLimits?api-version=2023-06-01-preview&$filter=location eq westus" -verbose
//Group Quota GETLIST
armclient Get  https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas?api-version=2023-06-01-preview -verbose
//GetSubscriptionAllocation
armclient GET https://management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/quotaAllocations/standarddv2family?api-version=2023-06-01-preview&$filter=location eq westus2" -verbose
//SetGroupQuota
armclient PUT https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreateBillingId.json -verbose
//SetGroupQuotaLimit
armclient PUT  https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/groupQuotaRequests/standarddv4family?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaLimitV2.json -verbose
//SetGroupQuotasSubscriptions 
armclient PUT https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreateBillingId.json -verbose





//SDK GroupQuotasSubscriptions DELETE
armclient DELETE https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview -verbose
//Group Quota DELETE
armclient Delete  https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota-create?api-version=2023-06-01-preview -verbose


//SDK GroupQuotasSubscriptions SYNTHETICS SUB
armclient PUT https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/serviceTreeV2GQ/subscriptions/6159567b-f71b-4eb1-a2d4-c5280bbfdd54?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreateBillingId.json -verbose
//Group Quota DELETE
armclient Delete  https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/serviceTreeV2GQ/subscriptions/6159567b-f71b-4eb1-a2d4-c5280bbfdd54?api-version=2023-06-01-preview -verbose


//Group Script Quota GET
armclient DELETE  https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/apiGQScriptTest?api-version=2023-06-01-preview -verbose
armclient PUT https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/bashScriptOAItejas?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreateServiceTreeId.json -verbose

//Group Script Quota Subscription PUT
armclient PUT https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/apiGQScriptTest/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreateBillingId.json -verbose
armclient DELETE https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/apiGQScriptTest/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview -verbose

//Group Script Quota Limit PUT
//Limit PUT V2
armclient PUT https://management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/apiGQScriptTest/groupQuotaRequests/standardav2family?api-version=2023-06-01-preview" @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaLimitV2.json -verbose

//Group Script Allocation PUT
//AllocationRequest PUT
armclient PUT https://management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/apiGQScriptTest/resourceProviders/Microsoft.Compute/quotaAllocationRequests/standardav2family?api-version=2023-06-01-preview&$filter=location eq westus" @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocation.json -verbose



armclient PUT https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testing-gq-1?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreateBillingId.json -verbose
armclient PUT https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testing-gq/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreateBillingId.json -verbose
armclient DELETE https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testing-gq/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview -verbose


armclient PUT https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/billTejasGQ?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreateBillingId.json -verbose
armclient PUT https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/billTejasGQ/subscriptions/7dc5548e-8df6-49de-ab32-d78108bba7bf?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreateBillingId.json -verbose
armclient DELETE https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/test-ts-gq/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview -verbose


armclient PUT  https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/groupQuotaRequests/standarddv4family?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaLimitV2.json -verbose


armclient PUT https://centraluseuap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/gq-tejas/groupQuotaRequests/standardav2family?api-version=2023-06-01-preview" @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaLimitV2.json -verbose

armclient GET https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/groupQuotaRequests/e0790180-a101-4d17-af13-06a98cd49070?api-version=2023-06-01-preview
armclient GET https://management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/groupQuotaRequests?api-version=2023-06-01-preview&$filter=location eq westus" -verbose

armclient PUT https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/quotaAllocationRequests/standarddv2family?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocation.json -verbose
armclient GET https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/quotaAllocationRequests/d5d5d8e0-52ce-47b5-9a6f-c768c03d9b8a?api-version=2023-06-01-preview


//GetGroupQuotaLimit
armclient GET https://management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/groupQuotaLimits/cores?api-version=2023-06-01-preview&$filter=location eq australiacentral" -verbose


//GroupQuotaLimit PUT V3
armclient PUT https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/groupQuotaLimits/eastus?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaLimitV3.json -verbose
//GroupQuotaLimit GETLIST V3
armclient GET https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/groupQuotaLimits/eastus?api-version=2023-06-01-preview&$filter=resourceName eq cores"  -verbose
//GroupQuotaLimit GET Single V3
armclient GET https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/groupQuotaLimits/eastus?api-version=2023-06-01-preview"  -verbose
//Error in URL
armclient GET https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/groupQuotaLimits?api-version=2023-06-01-preview"  -verbose
//SubscriptionQuotaAllocation Put V3
armclient PUT https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/quotaAllocations/eastus?api-version=2023-06-01-preview&" @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocationV3.json -verbose

armclient GET https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota?api-version=2023-06-01-preview" -verbose

(QMS BUG TESTING)
//Group Quota PUT
armclient PUT https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroupWrong?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreateBillingId.json -verbose
//Group Quota GET
armclient GET https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup?api-version=2023-06-01-preview -verbose
//Group Quota DELETE
armclient DELETE https:management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup?api-version=2023-06-01-preview -verbose
//GroupQuotasSubscriptions PUT
armclient PUT https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocation.json -verbose
//GroupQuotasSubscriptions GET
armclient GET https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview -verbose
//GroupQuotasSubscriptions DELETE
armclient DELETE https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/subscriptions/7dc5548e-8df6-49de-ab32-d78108bba7bf?api-version=2023-06-01-preview -verbose
//GroupQuotasSubscriptionAllocaions GET
armclient GET https://management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/testQmsGroup/resourceProviders/Microsoft.Compute/quotaAllocations/standarddv4family?api-version=2023-06-01-preview&$filter=location eq westus" -verbose
//GroupQuotasSubscriptionAllocations PUT
armclient PUT https://management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/testQmsGroup/resourceProviders/Microsoft.Compute/quotaAllocationRequests/standarddv4family?api-version=2023-06-01-preview" @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocation.json -verbose
//GroupQuotaLimit GET 
armclient GET https://management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/resourceProviders/Microsoft.Compute/groupQuotaLimits/standarddv4family?api-version=2023-06-01-preview&$filter=location eq westus" -verbose
//GroupQuotaLimit PUT
armclient PUT https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/resourceProviders/Microsoft.Compute/groupQuotaRequests/standarddv4family?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaLimitV2.json -verbose

armclient GET https://centraluseuap.management.azure.com/subscriptions/7dc5548e-8df6-49de-ab32-d78108bba7bf/providers/Microsoft.Compute/locations/westus/usages?api-version=2024-07-01

//GroupQuotasSubscriptions PUT
armclient PUT https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/subscriptions/7dc5548e-8df6-49de-ab32-d78108bba7bf?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocation.json -verbose
//GroupQuotasSubscriptions DELETE
armclient DELETE https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/subscriptions/7dc5548e-8df6-49de-ab32-d78108bba7bf?api-version=2023-06-01-preview -verbose



//GroupQuotaAllocation POST
armclient POST https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/testQmsGroup/resourceProviders/Microsoft.Compute/quotaAllocationRequests/standarddv3family?api-version=2023-06-01-preview" @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocation.json -verbose


armclient DELETE https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testqmsgroup/subscriptions/7dc5548e-8df6-49de-ab32-d78108bba7bf?api-version=2023-06-01-preview -verbose
armclient  https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/groupquotasubaddsuballocationd7f4239b-115c-4097-928f-0218ef76994c?api-version=2023-06-01-preview -verbose


armclient POST https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/resourceProviders/Microsoft.Compute/groupQuotaRequests/standarddv3family/update ?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaLimitV2.json -verbose


armclient GET https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/resourceProviders/Microsoft.Compute/groupQuotaRequests/standarddv3family/update ?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaLimitV2.json -verbose


armclient POST https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/resourceProviders/Microsoft.Compute/createLimitRequest?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaLimitV2.json -verbose
armclient POST https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/testQmsGroup/resourceProviders/Microsoft.Compute/createAllocationRequest?api-version=2023-06-01-preview" @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocation.json -verbose
armclient GET https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/resourceProviders/Microsoft.Compute/groupQuotaLimits/cores?api-version=2023-06-01-preview&$filter=location eq westus" -verbose


armclient PUT https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/te?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreateBillingId.json -verbose
armclient DELETE https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/test_Qms_Group?api-version=2023-06-01-preview -verbose
armclient GET https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/test_Qms_Group?api-version=2023-06-01-preview -verbose



V3 TESTING:
//Group Quota PUT
armclient PUT https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroupWrong?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreateBillingId.json -verbose
//Group Quota GET
armclient GET https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup?api-version=2023-06-01-preview -verbose
//GroupQuotasSubscriptions PUT
armclient PUT https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/subscriptions/7dc5548e-8df6-49de-ab32-d78108bba7bf?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocation.json -verbose
//GroupQuotasSubscriptions GET
armclient GET https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/subscriptions/7dc5548e-8df6-49de-ab32-d78108bba7bf?api-version=2023-06-01-preview -verbose
//GroupQuotaLimitV3 PUT
armclient PUT https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/resourceProviders/Microsoft.Compute/groupQuotaRequests/standarddv3family?api-version=2023-06-01-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaLimitV2.json -verbose
//GroupQuotaLimit GET 
armclient GET https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/resourceProviders/Microsoft.Compute/groupQuotaLimits/standarddv3family?api-version=2023-06-01-preview&$filter=location eq westus" -verbose
//GroupQuotasSubscriptionAllocaions GET
armclient GET https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/testQmsGroup/resourceProviders/Microsoft.Compute/quotaAllocations/westus?api-version=2023-06-01-preview&$filter=resourceName eq standardAv2Family" -verbose
//GroupQuotasSubscriptionAllocationsV3 PATCH
armclient PATCH https://eastus2euap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/testQmsGroup/resourceProviders/Microsoft.Compute/quotaAllocations/westus?api-version=2023-06-01-preview" @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocationV3.json -verbose
//GroupQuotasSubscriptions GET
armclient DELETE https://eastus2euap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/subscriptions/7dc5548e-8df6-49de-ab32-d78108bba7bf?api-version=2023-06-01-preview -verbose


//Group Quota PUT
armclient PUT https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/armPpe2Group?api-version=2023-06-01-alpha @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaCreateBillingId.json -verbose
armclient GET https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/armPpe2Group?api-version=2023-06-01-alpha -verbose

armclient GET https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup/subscriptions/7dc5548e-8df6-49de-ab32-d78108bba7bf?api-version=2023-06-01-preview -verbose


//AlphaTests
//Group Quota GET
armclient GET https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testQmsGroup?api-version=2023-06-01-alpha -verbose
//GroupQuotasSubscriptionAllocaions GET
armclient GET https://centraluseuap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/quotaAllocations/westus?api-version=2024-10-15-preview&$filter=resourceName eq standarddv4family" -verbose
//GroupQuotasSubscriptionAllocaions GET LIST
armclient GET https://centraluseuap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/quotaAllocations/westus?api-version=2024-10-15-preview" -verbose
//GroupQuotasSubscriptionAllocations PATCH
armclient PATCH https://centraluseuap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/quotaAllocations/westus?api-version=2024-10-15-preview" @C:\Users\tejasma\GroupQuotaLimitPayloads\groupQuotaAllocationV3.json -verbose
//GroupQuotaSubscriptions DELETE
armclient DELETE https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/testqmsgroup/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2024-10-15-preview -verbose
//GroupQuotaLimit GET
armclient GET https://centraluseuap.management.azure.com/"providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/resourceProviders/Microsoft.Compute/groupQuotaLimits/westus?api-version=2024-10-15-preview&$filter=resourceName eq standarddv4family" -verbose


//SDK CleanUP
//GroupQuotasSubscriptions DELETE
armclient DELETE https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/groupquota1-go-sdk-test/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2023-06-01-preview -verbose
//Group Quota DELETE
armclient DELETE https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/groupquota1-go-sdk-test?api-version=2023-06-01-preview -verbose



armclient PUT https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2024-10-15-preview @C:\Users\tejasma\GroupQuotaLimitPayloads\blank.txt -verbose
armclient DELETE https://centraluseuap.management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2024-10-15-preview -verbose
armclient DELETE https://management.azure.com/providers/Microsoft.Management/managementGroups/testMgIdRoot/providers/Microsoft.Quota/groupQuotas/sdk-test-group-quota-v3-create/subscriptions/65a85478-2333-4bbd-981b-1a818c944faf?api-version=2024-10-15-preview -verbose
