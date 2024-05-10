# coding=utf-8
# pylint: disable=too-many-lines
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) Python Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

import datetime
import sys
from typing import Any, Dict, List, Literal, Mapping, Optional, TYPE_CHECKING, Union, overload

from .. import _model_base
from .._model_base import rest_discriminator, rest_field
from ._enums import ChatRole

if sys.version_info >= (3, 9):
    from collections.abc import MutableMapping
else:
    from typing import MutableMapping  # type: ignore  # pylint: disable=ungrouped-imports

if TYPE_CHECKING:
    # pylint: disable=unused-import,ungrouped-imports
    from .. import models as _models
JSON = MutableMapping[str, Any]  # pylint: disable=unsubscriptable-object


class ChatRequestMessage(_model_base.Model):
    """An abstract representation of a chat message as provided in a request.

    You probably want to use the sub-classes and not this class directly. Known sub-classes are:
    AssistantMessage, SystemMessage, ToolMessage, UserMessage

    All required parameters must be populated in order to send to server.

    :ivar role: The chat role associated with this message. Required. Known values are: "system",
     "user", "assistant", and "tool".
    :vartype role: str or ~azure.ai.inference.models.ChatRole
    """

    __mapping__: Dict[str, _model_base.Model] = {}
    role: str = rest_discriminator(name="role")
    """The chat role associated with this message. Required. Known values are: \"system\", \"user\",
     \"assistant\", and \"tool\"."""

    @overload
    def __init__(
        self,
        *,
        role: str,
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class AssistantMessage(ChatRequestMessage, discriminator="assistant"):
    """A request chat message representing response or action from the assistant.

    All required parameters must be populated in order to send to server.

    :ivar role: The chat role associated with this message, which is always 'assistant' for
     assistant messages. Required. The role that provides responses to system-instructed,
     user-prompted input.
    :vartype role: str or ~azure.ai.inference.models.ASSISTANT
    :ivar content: The content of the message. Required.
    :vartype content: str
    :ivar tool_calls: The tool calls that must be resolved and have their outputs appended to
     subsequent input messages for the chat
     completions request to resolve as configured.
    :vartype tool_calls: list[~azure.ai.inference.models.ChatCompletionsToolCall]
    """

    role: Literal[ChatRole.ASSISTANT] = rest_discriminator(name="role")  # type: ignore
    """The chat role associated with this message, which is always 'assistant' for assistant messages.
     Required. The role that provides responses to system-instructed, user-prompted input."""
    content: str = rest_field()
    """The content of the message. Required."""
    tool_calls: Optional[List["_models.ChatCompletionsToolCall"]] = rest_field()
    """The tool calls that must be resolved and have their outputs appended to subsequent input
     messages for the chat
     completions request to resolve as configured."""

    @overload
    def __init__(
        self,
        *,
        content: str,
        tool_calls: Optional[List["_models.ChatCompletionsToolCall"]] = None,
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, role=ChatRole.ASSISTANT, **kwargs)


class ChatChoice(_model_base.Model):
    """The representation of a single prompt completion as part of an overall chat completions
    request.
    Generally, ``n`` choices are generated per provided prompt with a default value of 1.
    Token limits and other settings may limit the number of choices generated.

    All required parameters must be populated in order to send to server.

    :ivar index: The ordered index associated with this chat completions choice. Required.
    :vartype index: int
    :ivar finish_reason: The reason that this chat completions choice completed its generated.
     Required. Known values are: "stop", "length", "content_filter", and "tool_calls".
    :vartype finish_reason: str or ~azure.ai.inference.models.CompletionsFinishReason
    :ivar message: The chat message for a given chat completions prompt. Required.
    :vartype message: ~azure.ai.inference.models.ChatResponseMessage
    """

    index: int = rest_field()
    """The ordered index associated with this chat completions choice. Required."""
    finish_reason: Union[str, "_models.CompletionsFinishReason"] = rest_field()
    """The reason that this chat completions choice completed its generated. Required. Known values
     are: \"stop\", \"length\", \"content_filter\", and \"tool_calls\"."""
    message: "_models.ChatResponseMessage" = rest_field()
    """The chat message for a given chat completions prompt. Required."""

    @overload
    def __init__(
        self,
        *,
        index: int,
        finish_reason: Union[str, "_models.CompletionsFinishReason"],
        message: "_models.ChatResponseMessage",
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class ChatChoiceUpdate(_model_base.Model):
    """Represents an update to a single prompt completion when the service is streaming updates
    using Server Sent Events (SSE).
    Generally, ``n`` choices are generated per provided prompt with a default value of 1.
    Token limits and other settings may limit the number of choices generated.

    All required parameters must be populated in order to send to server.

    :ivar index: The ordered index associated with this chat completions choice. Required.
    :vartype index: int
    :ivar finish_reason: The reason that this chat completions choice completed its generated.
     Required. Known values are: "stop", "length", "content_filter", and "tool_calls".
    :vartype finish_reason: str or ~azure.ai.inference.models.CompletionsFinishReason
    :ivar delta: An update to the chat message for a given chat completions prompt. Required.
    :vartype delta: ~azure.ai.inference.models.ChatResponseMessage
    """

    index: int = rest_field()
    """The ordered index associated with this chat completions choice. Required."""
    finish_reason: Union[str, "_models.CompletionsFinishReason"] = rest_field()
    """The reason that this chat completions choice completed its generated. Required. Known values
     are: \"stop\", \"length\", \"content_filter\", and \"tool_calls\"."""
    delta: "_models.ChatResponseMessage" = rest_field()
    """An update to the chat message for a given chat completions prompt. Required."""

    @overload
    def __init__(
        self,
        *,
        index: int,
        finish_reason: Union[str, "_models.CompletionsFinishReason"],
        delta: "_models.ChatResponseMessage",
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class ChatCompletions(_model_base.Model):
    """Representation of the response data from a chat completions request.
    Completions support a wide variety of tasks and generate text that continues from or
    "completes"
    provided prompt data.

    All required parameters must be populated in order to send to server.

    :ivar id: A unique identifier associated with this chat completions response. Required.
    :vartype id: str
    :ivar object: The response object type, which is always ``chat.completion``. Required.
    :vartype object: str
    :ivar created: The first timestamp associated with generation activity for this completions
     response,
     represented as seconds since the beginning of the Unix epoch of 00:00 on 1 Jan 1970. Required.
    :vartype created: ~datetime.datetime
    :ivar model: The model used for the chat completion. Required.
    :vartype model: str
    :ivar usage: Usage information for tokens processed and generated as part of this completions
     operation. Required.
    :vartype usage: ~azure.ai.inference.models.CompletionsUsage
    :ivar choices: The collection of completions choices associated with this completions response.
     Generally, ``n`` choices are generated per provided prompt with a default value of 1.
     Token limits and other settings may limit the number of choices generated. Required.
    :vartype choices: list[~azure.ai.inference.models.ChatChoice]
    """

    id: str = rest_field()
    """A unique identifier associated with this chat completions response. Required."""
    object: str = rest_field()
    """The response object type, which is always ``chat.completion``. Required."""
    created: datetime.datetime = rest_field(format="unix-timestamp")
    """The first timestamp associated with generation activity for this completions response,
     represented as seconds since the beginning of the Unix epoch of 00:00 on 1 Jan 1970. Required."""
    model: str = rest_field()
    """The model used for the chat completion. Required."""
    usage: "_models.CompletionsUsage" = rest_field()
    """Usage information for tokens processed and generated as part of this completions operation.
     Required."""
    choices: List["_models.ChatChoice"] = rest_field()
    """The collection of completions choices associated with this completions response.
     Generally, ``n`` choices are generated per provided prompt with a default value of 1.
     Token limits and other settings may limit the number of choices generated. Required."""

    @overload
    def __init__(
        self,
        *,
        id: str,  # pylint: disable=redefined-builtin
        object: str,
        created: datetime.datetime,
        model: str,
        usage: "_models.CompletionsUsage",
        choices: List["_models.ChatChoice"],
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class ChatCompletionsToolCall(_model_base.Model):
    """An abstract representation of a tool call that must be resolved in a subsequent request to
    perform the requested
    chat completion.

    You probably want to use the sub-classes and not this class directly. Known sub-classes are:
    ChatCompletionsFunctionToolCall

    All required parameters must be populated in order to send to server.

    :ivar type: The object type. Required. Default value is None.
    :vartype type: str
    :ivar id: The ID of the tool call. Required.
    :vartype id: str
    """

    __mapping__: Dict[str, _model_base.Model] = {}
    type: str = rest_discriminator(name="type")
    """The object type. Required. Default value is None."""
    id: str = rest_field()
    """The ID of the tool call. Required."""

    @overload
    def __init__(
        self,
        *,
        type: str,
        id: str,  # pylint: disable=redefined-builtin
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class ChatCompletionsFunctionToolCall(ChatCompletionsToolCall, discriminator="function"):
    """A tool call to a function tool, issued by the model in evaluation of a configured function
    tool, that represents
    a function invocation needed for a subsequent chat completions request to resolve.

    All required parameters must be populated in order to send to server.

    :ivar id: The ID of the tool call. Required.
    :vartype id: str
    :ivar type: The type of tool call, in this case always 'function'. Required. Default value is
     "function".
    :vartype type: str
    :ivar function: The details of the function invocation requested by the tool call. Required.
    :vartype function: ~azure.ai.inference.models.FunctionCall
    """

    type: Literal["function"] = rest_discriminator(name="type")  # type: ignore
    """The type of tool call, in this case always 'function'. Required. Default value is \"function\"."""
    function: "_models.FunctionCall" = rest_field()
    """The details of the function invocation requested by the tool call. Required."""

    @overload
    def __init__(
        self,
        *,
        id: str,  # pylint: disable=redefined-builtin
        function: "_models.FunctionCall",
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, type="function", **kwargs)


class ChatCompletionsToolDefinition(_model_base.Model):
    """An abstract representation of a tool that can be used by the model to improve a chat
    completions response.

    You probably want to use the sub-classes and not this class directly. Known sub-classes are:
    ChatCompletionsFunctionToolDefinition

    All required parameters must be populated in order to send to server.

    :ivar type: The object type. Required. Default value is None.
    :vartype type: str
    """

    __mapping__: Dict[str, _model_base.Model] = {}
    type: str = rest_discriminator(name="type")
    """The object type. Required. Default value is None."""

    @overload
    def __init__(
        self,
        *,
        type: str,
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class ChatCompletionsFunctionToolDefinition(ChatCompletionsToolDefinition, discriminator="function"):
    """The definition information for a chat completions function tool that can call a function in
    response to a tool call.

    All required parameters must be populated in order to send to server.

    :ivar type: The object name, which is always 'function'. Required. Default value is "function".
    :vartype type: str
    :ivar function: The function definition details for the function tool. Required.
    :vartype function: ~azure.ai.inference.models.FunctionDefinition
    """

    type: Literal["function"] = rest_discriminator(name="type")  # type: ignore
    """The object name, which is always 'function'. Required. Default value is \"function\"."""
    function: "_models.FunctionDefinition" = rest_field()
    """The function definition details for the function tool. Required."""

    @overload
    def __init__(
        self,
        *,
        function: "_models.FunctionDefinition",
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, type="function", **kwargs)


class ChatCompletionsNamedToolSelection(_model_base.Model):
    """An abstract representation of an explicit, named tool selection to use for a chat completions
    request.

    All required parameters must be populated in order to send to server.

    :ivar type: The object type. Required.
    :vartype type: str
    """

    type: str = rest_discriminator(name="type")
    """The object type. Required."""


class ChatCompletionsUpdate(_model_base.Model):
    """Represents a response update to a chat completions request, when the service is streaming
    updates
    using Server Sent Events (SSE).
    Completions support a wide variety of tasks and generate text that continues from or
    "completes"
    provided prompt data.

    All required parameters must be populated in order to send to server.

    :ivar id: A unique identifier associated with this chat completions response. Required.
    :vartype id: str
    :ivar object: The response object type, which is always ``chat.completion``. Required.
    :vartype object: str
    :ivar created: The first timestamp associated with generation activity for this completions
     response,
     represented as seconds since the beginning of the Unix epoch of 00:00 on 1 Jan 1970. Required.
    :vartype created: ~datetime.datetime
    :ivar model: The model used for the chat completion. Required.
    :vartype model: str
    :ivar usage: Usage information for tokens processed and generated as part of this completions
     operation. Required.
    :vartype usage: ~azure.ai.inference.models.CompletionsUsage
    :ivar choices: An update to the collection of completion choices associated with this
     completions response.
     Generally, ``n`` choices are generated per provided prompt with a default value of 1.
     Token limits and other settings may limit the number of choices generated. Required.
    :vartype choices: list[~azure.ai.inference.models.ChatChoiceUpdate]
    """

    id: str = rest_field()
    """A unique identifier associated with this chat completions response. Required."""
    object: str = rest_field()
    """The response object type, which is always ``chat.completion``. Required."""
    created: datetime.datetime = rest_field(format="unix-timestamp")
    """The first timestamp associated with generation activity for this completions response,
     represented as seconds since the beginning of the Unix epoch of 00:00 on 1 Jan 1970. Required."""
    model: str = rest_field()
    """The model used for the chat completion. Required."""
    usage: "_models.CompletionsUsage" = rest_field()
    """Usage information for tokens processed and generated as part of this completions operation.
     Required."""
    choices: List["_models.ChatChoiceUpdate"] = rest_field()
    """An update to the collection of completion choices associated with this completions response.
     Generally, ``n`` choices are generated per provided prompt with a default value of 1.
     Token limits and other settings may limit the number of choices generated. Required."""

    @overload
    def __init__(
        self,
        *,
        id: str,  # pylint: disable=redefined-builtin
        object: str,
        created: datetime.datetime,
        model: str,
        usage: "_models.CompletionsUsage",
        choices: List["_models.ChatChoiceUpdate"],
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class ChatResponseMessage(_model_base.Model):
    """A representation of a chat message as received in a response.

    All required parameters must be populated in order to send to server.

    :ivar role: The chat role associated with the message. Required. Known values are: "system",
     "user", "assistant", and "tool".
    :vartype role: str or ~azure.ai.inference.models.ChatRole
    :ivar content: The content of the message. Required.
    :vartype content: str
    :ivar tool_calls: The tool calls that must be resolved and have their outputs appended to
     subsequent input messages for the chat
     completions request to resolve as configured.
    :vartype tool_calls: list[~azure.ai.inference.models.ChatCompletionsToolCall]
    """

    role: Union[str, "_models.ChatRole"] = rest_field()
    """The chat role associated with the message. Required. Known values are: \"system\", \"user\",
     \"assistant\", and \"tool\"."""
    content: str = rest_field()
    """The content of the message. Required."""
    tool_calls: Optional[List["_models.ChatCompletionsToolCall"]] = rest_field()
    """The tool calls that must be resolved and have their outputs appended to subsequent input
     messages for the chat
     completions request to resolve as configured."""

    @overload
    def __init__(
        self,
        *,
        role: Union[str, "_models.ChatRole"],
        content: str,
        tool_calls: Optional[List["_models.ChatCompletionsToolCall"]] = None,
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class CompletionsUsage(_model_base.Model):
    """Representation of the token counts processed for a completions request.
    Counts consider all tokens across prompts, choices, choice alternates, best_of generations, and
    other consumers.

    All required parameters must be populated in order to send to server.

    :ivar capacity_type: Indicates whether your capacity has been affected by the usage amount
     (token count) reported here. Required. Known values are: "usage" and "fixed".
    :vartype capacity_type: str or ~azure.ai.inference.models.CapacityType
    :ivar completion_tokens: The number of tokens generated across all completions emissions.
     Required.
    :vartype completion_tokens: int
    :ivar prompt_tokens: The number of tokens in the provided prompts for the completions request.
     Required.
    :vartype prompt_tokens: int
    :ivar total_tokens: The total number of tokens processed for the completions request and
     response. Required.
    :vartype total_tokens: int
    """

    capacity_type: Union[str, "_models.CapacityType"] = rest_field()
    """Indicates whether your capacity has been affected by the usage amount (token count) reported
     here. Required. Known values are: \"usage\" and \"fixed\"."""
    completion_tokens: int = rest_field()
    """The number of tokens generated across all completions emissions. Required."""
    prompt_tokens: int = rest_field()
    """The number of tokens in the provided prompts for the completions request. Required."""
    total_tokens: int = rest_field()
    """The total number of tokens processed for the completions request and response. Required."""

    @overload
    def __init__(
        self,
        *,
        capacity_type: Union[str, "_models.CapacityType"],
        completion_tokens: int,
        prompt_tokens: int,
        total_tokens: int,
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class EmbeddingItem(_model_base.Model):
    """Representation of a single embeddings relatedness comparison.

    All required parameters must be populated in order to send to server.

    :ivar embedding: List of embeddings value for the input prompt. These represent a measurement
     of the
     vector-based relatedness of the provided input. Required.
    :vartype embedding: list[float]
    :ivar index: Index of the prompt to which the EmbeddingItem corresponds. Required.
    :vartype index: int
    :ivar object: The object type of this embeddings item. Will always be ``embedding``. Required.
    :vartype object: str
    """

    embedding: List[float] = rest_field()
    """List of embeddings value for the input prompt. These represent a measurement of the
     vector-based relatedness of the provided input. Required."""
    index: int = rest_field()
    """Index of the prompt to which the EmbeddingItem corresponds. Required."""
    object: str = rest_field()
    """The object type of this embeddings item. Will always be ``embedding``. Required."""

    @overload
    def __init__(
        self,
        *,
        embedding: List[float],
        index: int,
        object: str,
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class EmbeddingsResult(_model_base.Model):
    """Representation of the response data from an embeddings request.
    Embeddings measure the relatedness of text strings and are commonly used for search,
    clustering,
    recommendations, and other similar scenarios.

    All required parameters must be populated in order to send to server.

    :ivar id: Unique identifier for the embeddings result. Required.
    :vartype id: str
    :ivar data: Embedding values for the prompts submitted in the request. Required.
    :vartype data: list[~azure.ai.inference.models.EmbeddingItem]
    :ivar usage: Usage counts for tokens input using the embeddings API. Required.
    :vartype usage: ~azure.ai.inference.models.EmbeddingsUsage
    :ivar object: The object type of the embeddings result. Will always be ``list``. Required.
    :vartype object: str
    :ivar model: The model ID used to generate this result. Required.
    :vartype model: str
    """

    id: str = rest_field()
    """Unique identifier for the embeddings result. Required."""
    data: List["_models.EmbeddingItem"] = rest_field()
    """Embedding values for the prompts submitted in the request. Required."""
    usage: "_models.EmbeddingsUsage" = rest_field()
    """Usage counts for tokens input using the embeddings API. Required."""
    object: str = rest_field()
    """The object type of the embeddings result. Will always be ``list``. Required."""
    model: str = rest_field()
    """The model ID used to generate this result. Required."""

    @overload
    def __init__(
        self,
        *,
        id: str,  # pylint: disable=redefined-builtin
        data: List["_models.EmbeddingItem"],
        usage: "_models.EmbeddingsUsage",
        object: str,
        model: str,
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class EmbeddingsUsage(_model_base.Model):
    """Measurement of the amount of tokens used in this request and response.

    All required parameters must be populated in order to send to server.

    :ivar capacity_type: Indicates whether your capacity has been affected by the usage amount
     (token count) reported here. Required. Known values are: "usage" and "fixed".
    :vartype capacity_type: str or ~azure.ai.inference.models.CapacityType
    :ivar input_tokens: Number of tokens in the request prompt. Required.
    :vartype input_tokens: int
    :ivar prompt_tokens: Number of tokens used for the prompt sent to the AI model. Typically
     identical to ``input_tokens``.
     However, certain AI models may add extra tokens to the input hence the number can be higher.
     (for example when input_type="query"). Required.
    :vartype prompt_tokens: int
    :ivar total_tokens: Total number of tokens transacted in this request/response. Required.
    :vartype total_tokens: int
    """

    capacity_type: Union[str, "_models.CapacityType"] = rest_field()
    """Indicates whether your capacity has been affected by the usage amount (token count) reported
     here. Required. Known values are: \"usage\" and \"fixed\"."""
    input_tokens: int = rest_field()
    """Number of tokens in the request prompt. Required."""
    prompt_tokens: int = rest_field()
    """Number of tokens used for the prompt sent to the AI model. Typically identical to
     ``input_tokens``.
     However, certain AI models may add extra tokens to the input hence the number can be higher.
     (for example when input_type=\"query\"). Required."""
    total_tokens: int = rest_field()
    """Total number of tokens transacted in this request/response. Required."""

    @overload
    def __init__(
        self,
        *,
        capacity_type: Union[str, "_models.CapacityType"],
        input_tokens: int,
        prompt_tokens: int,
        total_tokens: int,
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class FunctionCall(_model_base.Model):
    """The name and arguments of a function that should be called, as generated by the model.

    All required parameters must be populated in order to send to server.

    :ivar name: The name of the function to call. Required.
    :vartype name: str
    :ivar arguments: The arguments to call the function with, as generated by the model in JSON
     format.
     Note that the model does not always generate valid JSON, and may hallucinate parameters
     not defined by your function schema. Validate the arguments in your code before calling
     your function. Required.
    :vartype arguments: str
    """

    name: str = rest_field()
    """The name of the function to call. Required."""
    arguments: str = rest_field()
    """The arguments to call the function with, as generated by the model in JSON format.
     Note that the model does not always generate valid JSON, and may hallucinate parameters
     not defined by your function schema. Validate the arguments in your code before calling
     your function. Required."""

    @overload
    def __init__(
        self,
        *,
        name: str,
        arguments: str,
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class FunctionDefinition(_model_base.Model):
    """The definition of a caller-specified function that chat completions may invoke in response to
    matching user input.

    All required parameters must be populated in order to send to server.

    :ivar name: The name of the function to be called. Required.
    :vartype name: str
    :ivar description: A description of what the function does. The model will use this description
     when selecting the function and
     interpreting its parameters.
    :vartype description: str
    :ivar parameters: The parameters the function accepts, described as a JSON Schema object.
    :vartype parameters: any
    """

    name: str = rest_field()
    """The name of the function to be called. Required."""
    description: Optional[str] = rest_field()
    """A description of what the function does. The model will use this description when selecting the
     function and
     interpreting its parameters."""
    parameters: Optional[Any] = rest_field()
    """The parameters the function accepts, described as a JSON Schema object."""

    @overload
    def __init__(
        self,
        *,
        name: str,
        description: Optional[str] = None,
        parameters: Optional[Any] = None,
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class ModelInfo(_model_base.Model):
    """Represents some basic information about the AI model.

    All required parameters must be populated in order to send to server.

    :ivar model_name: The name of the AI model. For example: ``Phi21``. Required.
    :vartype model_name: str
    :ivar model_type: The type of the AI model. A Unique identifier for the profile. Required.
     Known values are: "embeddings", "image_generation", "text_generation", "image_embeddings",
     "audio_generation", and "chat".
    :vartype model_type: str or ~azure.ai.inference.models.ModelType
    :ivar model_provider_name: The model provider name. For example: ``Microsoft Research``.
     Required.
    :vartype model_provider_name: str
    """

    model_name: str = rest_field()
    """The name of the AI model. For example: ``Phi21``. Required."""
    model_type: Union[str, "_models.ModelType"] = rest_field()
    """The type of the AI model. A Unique identifier for the profile. Required. Known values are:
     \"embeddings\", \"image_generation\", \"text_generation\", \"image_embeddings\",
     \"audio_generation\", and \"chat\"."""
    model_provider_name: str = rest_field()
    """The model provider name. For example: ``Microsoft Research``. Required."""

    @overload
    def __init__(
        self,
        *,
        model_name: str,
        model_type: Union[str, "_models.ModelType"],
        model_provider_name: str,
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class SystemMessage(ChatRequestMessage, discriminator="system"):
    """A request chat message containing system instructions that influence how the model will
    generate a chat completions
    response.

    All required parameters must be populated in order to send to server.

    :ivar role: The chat role associated with this message, which is always 'system' for system
     messages. Required. The role that instructs or sets the behavior of the assistant.
    :vartype role: str or ~azure.ai.inference.models.SYSTEM
    :ivar content: The contents of the system message. Required.
    :vartype content: str
    """

    role: Literal[ChatRole.SYSTEM] = rest_discriminator(name="role")  # type: ignore
    """The chat role associated with this message, which is always 'system' for system messages.
     Required. The role that instructs or sets the behavior of the assistant."""
    content: str = rest_field()
    """The contents of the system message. Required."""

    @overload
    def __init__(
        self,
        *,
        content: str,
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, role=ChatRole.SYSTEM, **kwargs)


class ToolMessage(ChatRequestMessage, discriminator="tool"):
    """A request chat message representing requested output from a configured tool.

    All required parameters must be populated in order to send to server.

    :ivar role: The chat role associated with this message, which is always 'tool' for tool
     messages. Required. The role that represents extension tool activity within a chat completions
     operation.
    :vartype role: str or ~azure.ai.inference.models.TOOL
    :ivar content: The content of the message. Required.
    :vartype content: str
    :ivar tool_call_id: The ID of the tool call resolved by the provided content. Required.
    :vartype tool_call_id: str
    """

    role: Literal[ChatRole.TOOL] = rest_discriminator(name="role")  # type: ignore
    """The chat role associated with this message, which is always 'tool' for tool messages. Required.
     The role that represents extension tool activity within a chat completions operation."""
    content: str = rest_field()
    """The content of the message. Required."""
    tool_call_id: str = rest_field()
    """The ID of the tool call resolved by the provided content. Required."""

    @overload
    def __init__(
        self,
        *,
        content: str,
        tool_call_id: str,
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, role=ChatRole.TOOL, **kwargs)


class UserMessage(ChatRequestMessage, discriminator="user"):
    """A request chat message representing user input to the assistant.

    All required parameters must be populated in order to send to server.

    :ivar role: The chat role associated with this message, which is always 'user' for user
     messages. Required. The role that provides input for chat completions.
    :vartype role: str or ~azure.ai.inference.models.USER
    :ivar content: The contents of the user message, with available input types varying by selected
     model. Required.
    :vartype content: str
    """

    role: Literal[ChatRole.USER] = rest_discriminator(name="role")  # type: ignore
    """The chat role associated with this message, which is always 'user' for user messages. Required.
     The role that provides input for chat completions."""
    content: str = rest_field()
    """The contents of the user message, with available input types varying by selected model.
     Required."""

    @overload
    def __init__(
        self,
        *,
        content: str,
    ): ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, role=ChatRole.USER, **kwargs)
