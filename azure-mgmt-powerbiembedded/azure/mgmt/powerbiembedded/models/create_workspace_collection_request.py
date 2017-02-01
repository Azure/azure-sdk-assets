# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for
# license information.
#
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.
# --------------------------------------------------------------------------

from .azure_sku import AzureSku
from msrest.serialization import Model


class CreateWorkspaceCollectionRequest(Model):
    """CreateWorkspaceCollectionRequest.

    Variables are only populated by the server, and will be ignored when
    sending a request.

    :param location: Azure location
    :type location: str
    :param tags:
    :type tags: dict
    :ivar sku:
    :vartype sku: :class:`AzureSku
     <azure.mgmt.powerbiembedded.models.AzureSku>`
    """

    _validation = {
        'sku': {'constant': True},
    }

    _attribute_map = {
        'location': {'key': 'location', 'type': 'str'},
        'tags': {'key': 'tags', 'type': '{str}'},
        'sku': {'key': 'sku', 'type': 'AzureSku'},
    }

    sku = AzureSku()

    def __init__(self, location=None, tags=None):
        self.location = location
        self.tags = tags
