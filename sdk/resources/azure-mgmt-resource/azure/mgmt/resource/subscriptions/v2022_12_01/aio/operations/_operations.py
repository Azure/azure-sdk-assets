# pylint: disable=too-many-lines,too-many-statements
# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------
from io import IOBase
from typing import Any, AsyncIterable, Callable, Dict, IO, Optional, TypeVar, Union, overload
import urllib.parse

from azure.core.async_paging import AsyncItemPaged, AsyncList
from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceExistsError,
    ResourceNotFoundError,
    ResourceNotModifiedError,
    map_error,
)
from azure.core.pipeline import PipelineResponse
from azure.core.pipeline.transport import AsyncHttpResponse
from azure.core.rest import HttpRequest
from azure.core.tracing.decorator import distributed_trace
from azure.core.tracing.decorator_async import distributed_trace_async
from azure.core.utils import case_insensitive_dict
from azure.mgmt.core.exceptions import ARMErrorFormat

from ... import models as _models
from ..._vendor import _convert_request
from ...operations._operations import (
    build_operations_list_request,
    build_subscription_check_resource_name_request,
    build_subscriptions_check_zone_peers_request,
    build_subscriptions_get_request,
    build_subscriptions_list_locations_request,
    build_subscriptions_list_request,
    build_tenants_list_request,
)
from .._vendor import SubscriptionClientMixinABC

T = TypeVar("T")
ClsType = Optional[Callable[[PipelineResponse[HttpRequest, AsyncHttpResponse], T, Dict[str, Any]], Any]]


class Operations:
    """
    .. warning::
        **DO NOT** instantiate this class directly.

        Instead, you should access the following operations through
        :class:`~azure.mgmt.resource.subscriptions.v2022_12_01.aio.SubscriptionClient`'s
        :attr:`operations` attribute.
    """

    models = _models

    def __init__(self, *args, **kwargs) -> None:
        input_args = list(args)
        self._client = input_args.pop(0) if input_args else kwargs.pop("client")
        self._config = input_args.pop(0) if input_args else kwargs.pop("config")
        self._serialize = input_args.pop(0) if input_args else kwargs.pop("serializer")
        self._deserialize = input_args.pop(0) if input_args else kwargs.pop("deserializer")
        self._api_version = input_args.pop(0) if input_args else kwargs.pop("api_version")

    @distributed_trace
    def list(self, **kwargs: Any) -> AsyncIterable["_models.Operation"]:
        """Lists all of the available Microsoft.Resources REST API operations.

        :return: An iterator like instance of either Operation or the result of cls(response)
        :rtype:
         ~azure.core.async_paging.AsyncItemPaged[~azure.mgmt.resource.subscriptions.v2022_12_01.models.Operation]
        :raises ~azure.core.exceptions.HttpResponseError:
        """
        _headers = kwargs.pop("headers", {}) or {}
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version: str = kwargs.pop("api_version", _params.pop("api-version", self._api_version or "2022-12-01"))
        cls: ClsType[_models.OperationListResult] = kwargs.pop("cls", None)

        error_map = {
            401: ClientAuthenticationError,
            404: ResourceNotFoundError,
            409: ResourceExistsError,
            304: ResourceNotModifiedError,
        }
        error_map.update(kwargs.pop("error_map", {}) or {})

        def prepare_request(next_link=None):
            if not next_link:

                _request = build_operations_list_request(
                    api_version=api_version,
                    headers=_headers,
                    params=_params,
                )
                _request = _convert_request(_request)
                _request.url = self._client.format_url(_request.url)

            else:
                # make call to next link with the client's api-version
                _parsed_next_link = urllib.parse.urlparse(next_link)
                _next_request_params = case_insensitive_dict(
                    {
                        key: [urllib.parse.quote(v) for v in value]
                        for key, value in urllib.parse.parse_qs(_parsed_next_link.query).items()
                    }
                )
                _next_request_params["api-version"] = self._api_version
                _request = HttpRequest(
                    "GET", urllib.parse.urljoin(next_link, _parsed_next_link.path), params=_next_request_params
                )
                _request = _convert_request(_request)
                _request.url = self._client.format_url(_request.url)
                _request.method = "GET"
            return _request

        async def extract_data(pipeline_response):
            deserialized = self._deserialize("OperationListResult", pipeline_response)
            list_of_elem = deserialized.value
            if cls:
                list_of_elem = cls(list_of_elem)  # type: ignore
            return deserialized.next_link or None, AsyncList(list_of_elem)

        async def get_next(next_link=None):
            _request = prepare_request(next_link)

            _stream = False
            pipeline_response: PipelineResponse = await self._client._pipeline.run(  # pylint: disable=protected-access
                _request, stream=_stream, **kwargs
            )
            response = pipeline_response.http_response

            if response.status_code not in [200]:
                map_error(status_code=response.status_code, response=response, error_map=error_map)
                raise HttpResponseError(response=response, error_format=ARMErrorFormat)

            return pipeline_response

        return AsyncItemPaged(get_next, extract_data)


class SubscriptionsOperations:
    """
    .. warning::
        **DO NOT** instantiate this class directly.

        Instead, you should access the following operations through
        :class:`~azure.mgmt.resource.subscriptions.v2022_12_01.aio.SubscriptionClient`'s
        :attr:`subscriptions` attribute.
    """

    models = _models

    def __init__(self, *args, **kwargs) -> None:
        input_args = list(args)
        self._client = input_args.pop(0) if input_args else kwargs.pop("client")
        self._config = input_args.pop(0) if input_args else kwargs.pop("config")
        self._serialize = input_args.pop(0) if input_args else kwargs.pop("serializer")
        self._deserialize = input_args.pop(0) if input_args else kwargs.pop("deserializer")
        self._api_version = input_args.pop(0) if input_args else kwargs.pop("api_version")

    @distributed_trace
    def list_locations(
        self, subscription_id: str, include_extended_locations: Optional[bool] = None, **kwargs: Any
    ) -> AsyncIterable["_models.Location"]:
        """Gets all available geo-locations.

        This operation provides all the locations that are available for resource providers; however,
        each resource provider may support a subset of this list.

        :param subscription_id: The ID of the target subscription. Required.
        :type subscription_id: str
        :param include_extended_locations: Whether to include extended locations. Default value is
         None.
        :type include_extended_locations: bool
        :return: An iterator like instance of either Location or the result of cls(response)
        :rtype:
         ~azure.core.async_paging.AsyncItemPaged[~azure.mgmt.resource.subscriptions.v2022_12_01.models.Location]
        :raises ~azure.core.exceptions.HttpResponseError:
        """
        _headers = kwargs.pop("headers", {}) or {}
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version: str = kwargs.pop("api_version", _params.pop("api-version", self._api_version or "2022-12-01"))
        cls: ClsType[_models.LocationListResult] = kwargs.pop("cls", None)

        error_map = {
            401: ClientAuthenticationError,
            404: ResourceNotFoundError,
            409: ResourceExistsError,
            304: ResourceNotModifiedError,
        }
        error_map.update(kwargs.pop("error_map", {}) or {})

        def prepare_request(next_link=None):
            if not next_link:

                _request = build_subscriptions_list_locations_request(
                    subscription_id=subscription_id,
                    include_extended_locations=include_extended_locations,
                    api_version=api_version,
                    headers=_headers,
                    params=_params,
                )
                _request = _convert_request(_request)
                _request.url = self._client.format_url(_request.url)

            else:
                # make call to next link with the client's api-version
                _parsed_next_link = urllib.parse.urlparse(next_link)
                _next_request_params = case_insensitive_dict(
                    {
                        key: [urllib.parse.quote(v) for v in value]
                        for key, value in urllib.parse.parse_qs(_parsed_next_link.query).items()
                    }
                )
                _next_request_params["api-version"] = self._api_version
                _request = HttpRequest(
                    "GET", urllib.parse.urljoin(next_link, _parsed_next_link.path), params=_next_request_params
                )
                _request = _convert_request(_request)
                _request.url = self._client.format_url(_request.url)
                _request.method = "GET"
            return _request

        async def extract_data(pipeline_response):
            deserialized = self._deserialize("LocationListResult", pipeline_response)
            list_of_elem = deserialized.value
            if cls:
                list_of_elem = cls(list_of_elem)  # type: ignore
            return None, AsyncList(list_of_elem)

        async def get_next(next_link=None):
            _request = prepare_request(next_link)

            _stream = False
            pipeline_response: PipelineResponse = await self._client._pipeline.run(  # pylint: disable=protected-access
                _request, stream=_stream, **kwargs
            )
            response = pipeline_response.http_response

            if response.status_code not in [200]:
                map_error(status_code=response.status_code, response=response, error_map=error_map)
                raise HttpResponseError(response=response, error_format=ARMErrorFormat)

            return pipeline_response

        return AsyncItemPaged(get_next, extract_data)

    @distributed_trace_async
    async def get(self, subscription_id: str, **kwargs: Any) -> _models.Subscription:
        """Gets details about a specified subscription.

        :param subscription_id: The ID of the target subscription. Required.
        :type subscription_id: str
        :return: Subscription or the result of cls(response)
        :rtype: ~azure.mgmt.resource.subscriptions.v2022_12_01.models.Subscription
        :raises ~azure.core.exceptions.HttpResponseError:
        """
        error_map = {
            401: ClientAuthenticationError,
            404: ResourceNotFoundError,
            409: ResourceExistsError,
            304: ResourceNotModifiedError,
        }
        error_map.update(kwargs.pop("error_map", {}) or {})

        _headers = kwargs.pop("headers", {}) or {}
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version: str = kwargs.pop("api_version", _params.pop("api-version", self._api_version or "2022-12-01"))
        cls: ClsType[_models.Subscription] = kwargs.pop("cls", None)

        _request = build_subscriptions_get_request(
            subscription_id=subscription_id,
            api_version=api_version,
            headers=_headers,
            params=_params,
        )
        _request = _convert_request(_request)
        _request.url = self._client.format_url(_request.url)

        _stream = False
        pipeline_response: PipelineResponse = await self._client._pipeline.run(  # pylint: disable=protected-access
            _request, stream=_stream, **kwargs
        )

        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            raise HttpResponseError(response=response, error_format=ARMErrorFormat)

        deserialized = self._deserialize("Subscription", pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})  # type: ignore

        return deserialized  # type: ignore

    @distributed_trace
    def list(self, **kwargs: Any) -> AsyncIterable["_models.Subscription"]:
        """Gets all subscriptions for a tenant.

        :return: An iterator like instance of either Subscription or the result of cls(response)
        :rtype:
         ~azure.core.async_paging.AsyncItemPaged[~azure.mgmt.resource.subscriptions.v2022_12_01.models.Subscription]
        :raises ~azure.core.exceptions.HttpResponseError:
        """
        _headers = kwargs.pop("headers", {}) or {}
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version: str = kwargs.pop("api_version", _params.pop("api-version", self._api_version or "2022-12-01"))
        cls: ClsType[_models.SubscriptionListResult] = kwargs.pop("cls", None)

        error_map = {
            401: ClientAuthenticationError,
            404: ResourceNotFoundError,
            409: ResourceExistsError,
            304: ResourceNotModifiedError,
        }
        error_map.update(kwargs.pop("error_map", {}) or {})

        def prepare_request(next_link=None):
            if not next_link:

                _request = build_subscriptions_list_request(
                    api_version=api_version,
                    headers=_headers,
                    params=_params,
                )
                _request = _convert_request(_request)
                _request.url = self._client.format_url(_request.url)

            else:
                # make call to next link with the client's api-version
                _parsed_next_link = urllib.parse.urlparse(next_link)
                _next_request_params = case_insensitive_dict(
                    {
                        key: [urllib.parse.quote(v) for v in value]
                        for key, value in urllib.parse.parse_qs(_parsed_next_link.query).items()
                    }
                )
                _next_request_params["api-version"] = self._api_version
                _request = HttpRequest(
                    "GET", urllib.parse.urljoin(next_link, _parsed_next_link.path), params=_next_request_params
                )
                _request = _convert_request(_request)
                _request.url = self._client.format_url(_request.url)
                _request.method = "GET"
            return _request

        async def extract_data(pipeline_response):
            deserialized = self._deserialize("SubscriptionListResult", pipeline_response)
            list_of_elem = deserialized.value
            if cls:
                list_of_elem = cls(list_of_elem)  # type: ignore
            return deserialized.next_link or None, AsyncList(list_of_elem)

        async def get_next(next_link=None):
            _request = prepare_request(next_link)

            _stream = False
            pipeline_response: PipelineResponse = await self._client._pipeline.run(  # pylint: disable=protected-access
                _request, stream=_stream, **kwargs
            )
            response = pipeline_response.http_response

            if response.status_code not in [200]:
                map_error(status_code=response.status_code, response=response, error_map=error_map)
                raise HttpResponseError(response=response, error_format=ARMErrorFormat)

            return pipeline_response

        return AsyncItemPaged(get_next, extract_data)

    @overload
    async def check_zone_peers(
        self,
        subscription_id: str,
        parameters: _models.CheckZonePeersRequest,
        *,
        content_type: str = "application/json",
        **kwargs: Any
    ) -> _models.CheckZonePeersResult:
        """Compares a subscriptions logical zone mapping.

        :param subscription_id: The ID of the target subscription. Required.
        :type subscription_id: str
        :param parameters: Parameters for checking zone peers. Required.
        :type parameters: ~azure.mgmt.resource.subscriptions.v2022_12_01.models.CheckZonePeersRequest
        :keyword content_type: Body Parameter content-type. Content type parameter for JSON body.
         Default value is "application/json".
        :paramtype content_type: str
        :return: CheckZonePeersResult or the result of cls(response)
        :rtype: ~azure.mgmt.resource.subscriptions.v2022_12_01.models.CheckZonePeersResult
        :raises ~azure.core.exceptions.HttpResponseError:
        """

    @overload
    async def check_zone_peers(
        self, subscription_id: str, parameters: IO[bytes], *, content_type: str = "application/json", **kwargs: Any
    ) -> _models.CheckZonePeersResult:
        """Compares a subscriptions logical zone mapping.

        :param subscription_id: The ID of the target subscription. Required.
        :type subscription_id: str
        :param parameters: Parameters for checking zone peers. Required.
        :type parameters: IO[bytes]
        :keyword content_type: Body Parameter content-type. Content type parameter for binary body.
         Default value is "application/json".
        :paramtype content_type: str
        :return: CheckZonePeersResult or the result of cls(response)
        :rtype: ~azure.mgmt.resource.subscriptions.v2022_12_01.models.CheckZonePeersResult
        :raises ~azure.core.exceptions.HttpResponseError:
        """

    @distributed_trace_async
    async def check_zone_peers(
        self, subscription_id: str, parameters: Union[_models.CheckZonePeersRequest, IO[bytes]], **kwargs: Any
    ) -> _models.CheckZonePeersResult:
        """Compares a subscriptions logical zone mapping.

        :param subscription_id: The ID of the target subscription. Required.
        :type subscription_id: str
        :param parameters: Parameters for checking zone peers. Is either a CheckZonePeersRequest type
         or a IO[bytes] type. Required.
        :type parameters: ~azure.mgmt.resource.subscriptions.v2022_12_01.models.CheckZonePeersRequest
         or IO[bytes]
        :return: CheckZonePeersResult or the result of cls(response)
        :rtype: ~azure.mgmt.resource.subscriptions.v2022_12_01.models.CheckZonePeersResult
        :raises ~azure.core.exceptions.HttpResponseError:
        """
        error_map = {
            401: ClientAuthenticationError,
            404: ResourceNotFoundError,
            409: ResourceExistsError,
            304: ResourceNotModifiedError,
        }
        error_map.update(kwargs.pop("error_map", {}) or {})

        _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version: str = kwargs.pop("api_version", _params.pop("api-version", self._api_version or "2022-12-01"))
        content_type: Optional[str] = kwargs.pop("content_type", _headers.pop("Content-Type", None))
        cls: ClsType[_models.CheckZonePeersResult] = kwargs.pop("cls", None)

        content_type = content_type or "application/json"
        _json = None
        _content = None
        if isinstance(parameters, (IOBase, bytes)):
            _content = parameters
        else:
            _json = self._serialize.body(parameters, "CheckZonePeersRequest")

        _request = build_subscriptions_check_zone_peers_request(
            subscription_id=subscription_id,
            api_version=api_version,
            content_type=content_type,
            json=_json,
            content=_content,
            headers=_headers,
            params=_params,
        )
        _request = _convert_request(_request)
        _request.url = self._client.format_url(_request.url)

        _stream = False
        pipeline_response: PipelineResponse = await self._client._pipeline.run(  # pylint: disable=protected-access
            _request, stream=_stream, **kwargs
        )

        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            error = self._deserialize.failsafe_deserialize(_models.ErrorResponseAutoGenerated, pipeline_response)
            raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

        deserialized = self._deserialize("CheckZonePeersResult", pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})  # type: ignore

        return deserialized  # type: ignore


class TenantsOperations:
    """
    .. warning::
        **DO NOT** instantiate this class directly.

        Instead, you should access the following operations through
        :class:`~azure.mgmt.resource.subscriptions.v2022_12_01.aio.SubscriptionClient`'s
        :attr:`tenants` attribute.
    """

    models = _models

    def __init__(self, *args, **kwargs) -> None:
        input_args = list(args)
        self._client = input_args.pop(0) if input_args else kwargs.pop("client")
        self._config = input_args.pop(0) if input_args else kwargs.pop("config")
        self._serialize = input_args.pop(0) if input_args else kwargs.pop("serializer")
        self._deserialize = input_args.pop(0) if input_args else kwargs.pop("deserializer")
        self._api_version = input_args.pop(0) if input_args else kwargs.pop("api_version")

    @distributed_trace
    def list(self, **kwargs: Any) -> AsyncIterable["_models.TenantIdDescription"]:
        """Gets the tenants for your account.

        :return: An iterator like instance of either TenantIdDescription or the result of cls(response)
        :rtype:
         ~azure.core.async_paging.AsyncItemPaged[~azure.mgmt.resource.subscriptions.v2022_12_01.models.TenantIdDescription]
        :raises ~azure.core.exceptions.HttpResponseError:
        """
        _headers = kwargs.pop("headers", {}) or {}
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version: str = kwargs.pop("api_version", _params.pop("api-version", self._api_version or "2022-12-01"))
        cls: ClsType[_models.TenantListResult] = kwargs.pop("cls", None)

        error_map = {
            401: ClientAuthenticationError,
            404: ResourceNotFoundError,
            409: ResourceExistsError,
            304: ResourceNotModifiedError,
        }
        error_map.update(kwargs.pop("error_map", {}) or {})

        def prepare_request(next_link=None):
            if not next_link:

                _request = build_tenants_list_request(
                    api_version=api_version,
                    headers=_headers,
                    params=_params,
                )
                _request = _convert_request(_request)
                _request.url = self._client.format_url(_request.url)

            else:
                # make call to next link with the client's api-version
                _parsed_next_link = urllib.parse.urlparse(next_link)
                _next_request_params = case_insensitive_dict(
                    {
                        key: [urllib.parse.quote(v) for v in value]
                        for key, value in urllib.parse.parse_qs(_parsed_next_link.query).items()
                    }
                )
                _next_request_params["api-version"] = self._api_version
                _request = HttpRequest(
                    "GET", urllib.parse.urljoin(next_link, _parsed_next_link.path), params=_next_request_params
                )
                _request = _convert_request(_request)
                _request.url = self._client.format_url(_request.url)
                _request.method = "GET"
            return _request

        async def extract_data(pipeline_response):
            deserialized = self._deserialize("TenantListResult", pipeline_response)
            list_of_elem = deserialized.value
            if cls:
                list_of_elem = cls(list_of_elem)  # type: ignore
            return deserialized.next_link or None, AsyncList(list_of_elem)

        async def get_next(next_link=None):
            _request = prepare_request(next_link)

            _stream = False
            pipeline_response: PipelineResponse = await self._client._pipeline.run(  # pylint: disable=protected-access
                _request, stream=_stream, **kwargs
            )
            response = pipeline_response.http_response

            if response.status_code not in [200]:
                map_error(status_code=response.status_code, response=response, error_map=error_map)
                raise HttpResponseError(response=response, error_format=ARMErrorFormat)

            return pipeline_response

        return AsyncItemPaged(get_next, extract_data)


class SubscriptionClientOperationsMixin(SubscriptionClientMixinABC):
    def _api_version(self, op_name: str) -> str:  # pylint: disable=unused-argument
        try:
            return self._config.api_version
        except:  # pylint: disable=bare-except
            return ""

    @overload
    async def check_resource_name(
        self,
        resource_name_definition: Optional[_models.ResourceName] = None,
        *,
        content_type: str = "application/json",
        **kwargs: Any
    ) -> _models.CheckResourceNameResult:
        """Checks resource name validity.

        A resource name is valid if it is not a reserved word, does not contains a reserved word and
        does not start with a reserved word.

        :param resource_name_definition: Resource object with values for resource name and resource
         type. Default value is None.
        :type resource_name_definition:
         ~azure.mgmt.resource.subscriptions.v2022_12_01.models.ResourceName
        :keyword content_type: Body Parameter content-type. Content type parameter for JSON body.
         Default value is "application/json".
        :paramtype content_type: str
        :return: CheckResourceNameResult or the result of cls(response)
        :rtype: ~azure.mgmt.resource.subscriptions.v2022_12_01.models.CheckResourceNameResult
        :raises ~azure.core.exceptions.HttpResponseError:
        """

    @overload
    async def check_resource_name(
        self,
        resource_name_definition: Optional[IO[bytes]] = None,
        *,
        content_type: str = "application/json",
        **kwargs: Any
    ) -> _models.CheckResourceNameResult:
        """Checks resource name validity.

        A resource name is valid if it is not a reserved word, does not contains a reserved word and
        does not start with a reserved word.

        :param resource_name_definition: Resource object with values for resource name and resource
         type. Default value is None.
        :type resource_name_definition: IO[bytes]
        :keyword content_type: Body Parameter content-type. Content type parameter for binary body.
         Default value is "application/json".
        :paramtype content_type: str
        :return: CheckResourceNameResult or the result of cls(response)
        :rtype: ~azure.mgmt.resource.subscriptions.v2022_12_01.models.CheckResourceNameResult
        :raises ~azure.core.exceptions.HttpResponseError:
        """

    @distributed_trace_async
    async def check_resource_name(
        self, resource_name_definition: Optional[Union[_models.ResourceName, IO[bytes]]] = None, **kwargs: Any
    ) -> _models.CheckResourceNameResult:
        """Checks resource name validity.

        A resource name is valid if it is not a reserved word, does not contains a reserved word and
        does not start with a reserved word.

        :param resource_name_definition: Resource object with values for resource name and resource
         type. Is either a ResourceName type or a IO[bytes] type. Default value is None.
        :type resource_name_definition:
         ~azure.mgmt.resource.subscriptions.v2022_12_01.models.ResourceName or IO[bytes]
        :return: CheckResourceNameResult or the result of cls(response)
        :rtype: ~azure.mgmt.resource.subscriptions.v2022_12_01.models.CheckResourceNameResult
        :raises ~azure.core.exceptions.HttpResponseError:
        """
        error_map = {
            401: ClientAuthenticationError,
            404: ResourceNotFoundError,
            409: ResourceExistsError,
            304: ResourceNotModifiedError,
        }
        error_map.update(kwargs.pop("error_map", {}) or {})

        _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version: str = kwargs.pop(
            "api_version", _params.pop("api-version", self._api_version("check_resource_name") or "2022-12-01")
        )
        content_type: Optional[str] = kwargs.pop("content_type", _headers.pop("Content-Type", None))
        cls: ClsType[_models.CheckResourceNameResult] = kwargs.pop("cls", None)

        content_type = content_type or "application/json"
        _json = None
        _content = None
        if isinstance(resource_name_definition, (IOBase, bytes)):
            _content = resource_name_definition
        else:
            if resource_name_definition is not None:
                _json = self._serialize.body(resource_name_definition, "ResourceName")
            else:
                _json = None

        _request = build_subscription_check_resource_name_request(
            api_version=api_version,
            content_type=content_type,
            json=_json,
            content=_content,
            headers=_headers,
            params=_params,
        )
        _request = _convert_request(_request)
        _request.url = self._client.format_url(_request.url)

        _stream = False
        pipeline_response: PipelineResponse = await self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
            _request, stream=_stream, **kwargs
        )

        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            raise HttpResponseError(response=response, error_format=ARMErrorFormat)

        deserialized = self._deserialize("CheckResourceNameResult", pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})  # type: ignore

        return deserialized  # type: ignore
