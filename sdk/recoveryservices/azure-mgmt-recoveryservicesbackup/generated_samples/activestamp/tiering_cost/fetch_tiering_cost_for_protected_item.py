# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicesbackup.activestamp import RecoveryServicesBackupClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesbackup
# USAGE
    python fetch_tiering_cost_for_protected_item.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RecoveryServicesBackupClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.fetch_tiering_cost.begin_post(
        resource_group_name="netsdktestrg",
        vault_name="testVault",
        parameters={
            "containerName": "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
            "objectType": "FetchTieringCostSavingsInfoForProtectedItemRequest",
            "protectedItemName": "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
            "sourceTierType": "HardenedRP",
            "targetTierType": "ArchivedRP",
        },
    ).result()
    print(response)


# x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/TieringCost/FetchTieringCostForProtectedItem.json
if __name__ == "__main__":
    main()
