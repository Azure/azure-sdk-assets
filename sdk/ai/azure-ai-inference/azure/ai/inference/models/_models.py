# coding=utf-8
# pylint: disable=too-many-lines
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) Python Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

import datetime
from typing import Any, Dict, List, Literal, Mapping, Optional, TYPE_CHECKING, Union, overload

from .. import _model_base
from .._model_base import rest_discriminator, rest_field
from ._enums import ChatRole

if TYPE_CHECKING:
    # pylint: disable=unused-import,ungrouped-imports
    from .. import models as _models


class ChatChoice(_model_base.Model):
    """The representation of a single prompt completion as part of an overall chat completions
    request.
    Generally, ``n`` choices are generated per provided prompt with a default value of 1.
    Token limits and other settings may limit the number of choices generated.

    All required parameters must be populated in order to send to server.

    :ivar message: The chat message for a given chat completions prompt.
    :vartype message: ~azure.ai.inference.models.ChatResponseMessage
    :ivar index: The ordered index associated with this chat completions choice. Required.
    :vartype index: int
    :ivar finish_reason: The reason that this chat completions choice completed its generated.
     Required. Known values are: "stop", "length", and "content_filter".
    :vartype finish_reason: str or ~azure.ai.inference.models.CompletionsFinishReason
    :ivar delta: The delta message content for a streaming response.
    :vartype delta: ~azure.ai.inference.models.ChatResponseMessage
    """

    message: Optional["_models.ChatResponseMessage"] = rest_field()
    """The chat message for a given chat completions prompt."""
    index: int = rest_field()
    """The ordered index associated with this chat completions choice. Required."""
    finish_reason: Union[str, "_models.CompletionsFinishReason"] = rest_field()
    """The reason that this chat completions choice completed its generated. Required. Known values
     are: \"stop\", \"length\", and \"content_filter\"."""
    delta: Optional["_models.ChatResponseMessage"] = rest_field()
    """The delta message content for a streaming response."""

    @overload
    def __init__(
        self,
        *,
        index: int,
        finish_reason: Union[str, "_models.CompletionsFinishReason"],
        message: Optional["_models.ChatResponseMessage"] = None,
        delta: Optional["_models.ChatResponseMessage"] = None,
    ):
        ...

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
    :ivar created: The first timestamp associated with generation activity for this completions
     response,
     represented as seconds since the beginning of the Unix epoch of 00:00 on 1 Jan 1970. Required.
    :vartype created: ~datetime.datetime
    :ivar choices: The collection of completions choices associated with this completions response.
     Generally, ``n`` choices are generated per provided prompt with a default value of 1.
     Token limits and other settings may limit the number of choices generated. Required.
    :vartype choices: list[~azure.ai.inference.models.ChatChoice]
    :ivar usage: Usage information for tokens processed and generated as part of this completions
     operation. Required.
    :vartype usage: ~azure.ai.inference.models.CompletionsUsage
    :ivar object: The response object type, which is always ``chat.completion``. Required.
    :vartype object: str
    :ivar model: The model used for the chat completion. Required.
    :vartype model: str
    """

    id: str = rest_field()
    """A unique identifier associated with this chat completions response. Required."""
    created: datetime.datetime = rest_field(format="unix-timestamp")
    """The first timestamp associated with generation activity for this completions response,
     represented as seconds since the beginning of the Unix epoch of 00:00 on 1 Jan 1970. Required."""
    choices: List["_models.ChatChoice"] = rest_field()
    """The collection of completions choices associated with this completions response.
     Generally, ``n`` choices are generated per provided prompt with a default value of 1.
     Token limits and other settings may limit the number of choices generated. Required."""
    usage: "_models.CompletionsUsage" = rest_field()
    """Usage information for tokens processed and generated as part of this completions operation.
     Required."""
    object: str = rest_field()
    """The response object type, which is always ``chat.completion``. Required."""
    model: str = rest_field()
    """The model used for the chat completion. Required."""

    @overload
    def __init__(
        self,
        *,
        id: str,  # pylint: disable=redefined-builtin
        created: datetime.datetime,
        choices: List["_models.ChatChoice"],
        usage: "_models.CompletionsUsage",
        object: str,
        model: str,
    ):
        ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class ChatCompletionsResponseFormat(_model_base.Model):
    """An abstract representation of a response format configuration usable by Chat Completions. Can
    be used to enable JSON
    mode.

    You probably want to use the sub-classes and not this class directly. Known sub-classes are:
    ChatCompletionsJsonResponseFormat, ChatCompletionsTextResponseFormat

    All required parameters must be populated in order to send to server.

    :ivar type: The discriminated type for the response format. Required. Default value is None.
    :vartype type: str
    """

    __mapping__: Dict[str, _model_base.Model] = {}
    type: str = rest_discriminator(name="type")
    """The discriminated type for the response format. Required. Default value is None."""

    @overload
    def __init__(
        self,
        *,
        type: str,
    ):
        ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class ChatCompletionsJsonResponseFormat(ChatCompletionsResponseFormat, discriminator="json_object"):
    """A response format for Chat Completions that restricts responses to emitting valid JSON objects.

    All required parameters must be populated in order to send to server.

    :ivar type: The discriminated object type, which is always 'json_object' for this format.
     Required. Default value is "json_object".
    :vartype type: str
    """

    type: Literal["json_object"] = rest_discriminator(name="type")  # type: ignore
    """The discriminated object type, which is always 'json_object' for this format. Required. Default
     value is \"json_object\"."""


class ChatCompletionsOptions(_model_base.Model):
    """The configuration information for a chat completions request.
    Completions support a wide variety of tasks and generate text that continues from or
    "completes"
    provided prompt data.

    All required parameters must be populated in order to send to server.

    :ivar messages: The collection of context messages associated with this chat completions
     request.
     Typical usage begins with a chat message for the System role that provides instructions for
     the behavior of the assistant, followed by alternating messages between the User and
     Assistant roles. Required.
    :vartype messages: list[~azure.ai.inference.models.ChatRequestMessage]
    :ivar max_tokens: The maximum number of tokens to generate.
    :vartype max_tokens: int
    :ivar temperature: The sampling temperature to use that controls the apparent creativity of
     generated completions.
     Higher values will make output more random while lower values will make results more focused
     and deterministic.
     It is not recommended to modify temperature and top_p for the same completions request as the
     interaction of these two settings is difficult to predict.
    :vartype temperature: float
    :ivar top_p: An alternative to sampling with temperature called nucleus sampling. This value
     causes the
     model to consider the results of tokens with the provided probability mass. As an example, a
     value of 0.15 will cause only the tokens comprising the top 15% of probability mass to be
     considered.
     It is not recommended to modify temperature and top_p for the same completions request as the
     interaction of these two settings is difficult to predict.
    :vartype top_p: float
    :ivar stop: A collection of textual sequences that will end completions generation.
    :vartype stop: list[str]
    :ivar presence_penalty: A value that influences the probability of generated tokens appearing
     based on their existing
     presence in generated text.
     Positive values will make tokens less likely to appear when they already exist and increase
     the
     model's likelihood to output new topics.
    :vartype presence_penalty: float
    :ivar frequency_penalty: A value that influences the probability of generated tokens appearing
     based on their cumulative
     frequency in generated text.
     Positive values will make tokens less likely to appear as their frequency increases and
     decrease the likelihood of the model repeating the same statements verbatim.
    :vartype frequency_penalty: float
    :ivar stream: A value indicating whether chat completions should be streamed for this request.
    :vartype stream: bool
    :ivar seed: If specified, the system will make a best effort to sample deterministically such
     that repeated requests with the
     same seed and parameters should return the same result. Determinism is not guaranteed, and you
     should refer to the
     system_fingerprint response parameter to monitor changes in the backend.".
    :vartype seed: int
    :ivar response_format: An object specifying the format that the model must output. Used to
     enable JSON mode.
    :vartype response_format: ~azure.ai.inference.models.ChatCompletionsResponseFormat
    """

    messages: List["_models.ChatRequestMessage"] = rest_field()
    """The collection of context messages associated with this chat completions request.
     Typical usage begins with a chat message for the System role that provides instructions for
     the behavior of the assistant, followed by alternating messages between the User and
     Assistant roles. Required."""
    max_tokens: Optional[int] = rest_field()
    """The maximum number of tokens to generate."""
    temperature: Optional[float] = rest_field()
    """The sampling temperature to use that controls the apparent creativity of generated completions.
     Higher values will make output more random while lower values will make results more focused
     and deterministic.
     It is not recommended to modify temperature and top_p for the same completions request as the
     interaction of these two settings is difficult to predict."""
    top_p: Optional[float] = rest_field()
    """An alternative to sampling with temperature called nucleus sampling. This value causes the
     model to consider the results of tokens with the provided probability mass. As an example, a
     value of 0.15 will cause only the tokens comprising the top 15% of probability mass to be
     considered.
     It is not recommended to modify temperature and top_p for the same completions request as the
     interaction of these two settings is difficult to predict."""
    stop: Optional[List[str]] = rest_field()
    """A collection of textual sequences that will end completions generation."""
    presence_penalty: Optional[float] = rest_field()
    """A value that influences the probability of generated tokens appearing based on their existing
     presence in generated text.
     Positive values will make tokens less likely to appear when they already exist and increase the
     model's likelihood to output new topics."""
    frequency_penalty: Optional[float] = rest_field()
    """A value that influences the probability of generated tokens appearing based on their cumulative
     frequency in generated text.
     Positive values will make tokens less likely to appear as their frequency increases and
     decrease the likelihood of the model repeating the same statements verbatim."""
    stream: Optional[bool] = rest_field()
    """A value indicating whether chat completions should be streamed for this request."""
    seed: Optional[int] = rest_field()
    """If specified, the system will make a best effort to sample deterministically such that repeated
     requests with the
     same seed and parameters should return the same result. Determinism is not guaranteed, and you
     should refer to the
     system_fingerprint response parameter to monitor changes in the backend.\"."""
    response_format: Optional["_models.ChatCompletionsResponseFormat"] = rest_field()
    """An object specifying the format that the model must output. Used to enable JSON mode."""

    @overload
    def __init__(
        self,
        *,
        messages: List["_models.ChatRequestMessage"],
        max_tokens: Optional[int] = None,
        temperature: Optional[float] = None,
        top_p: Optional[float] = None,
        stop: Optional[List[str]] = None,
        presence_penalty: Optional[float] = None,
        frequency_penalty: Optional[float] = None,
        stream: Optional[bool] = None,
        seed: Optional[int] = None,
        response_format: Optional["_models.ChatCompletionsResponseFormat"] = None,
    ):
        ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class ChatCompletionsTextResponseFormat(ChatCompletionsResponseFormat, discriminator="text"):
    """The standard Chat Completions response format that can freely generate text and is not
    guaranteed to produce response
    content that adheres to a specific schema.

    All required parameters must be populated in order to send to server.

    :ivar type: The discriminated object type, which is always 'text' for this format. Required.
     Default value is "text".
    :vartype type: str
    """

    type: Literal["text"] = rest_discriminator(name="type")  # type: ignore
    """The discriminated object type, which is always 'text' for this format. Required. Default value
     is \"text\"."""


class ChatRequestMessage(_model_base.Model):
    """An abstract representation of a chat message as provided in a request.

    You probably want to use the sub-classes and not this class directly. Known sub-classes are:
    ChatRequestAssistantMessage, ChatRequestSystemMessage, ChatRequestUserMessage

    All required parameters must be populated in order to send to server.

    :ivar role: The chat role associated with this message. Required. Known values are: "system",
     "assistant", and "user".
    :vartype role: str or ~azure.ai.inference.models.ChatRole
    """

    __mapping__: Dict[str, _model_base.Model] = {}
    role: str = rest_discriminator(name="role")
    """The chat role associated with this message. Required. Known values are: \"system\",
     \"assistant\", and \"user\"."""

    @overload
    def __init__(
        self,
        *,
        role: str,
    ):
        ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)


class ChatRequestAssistantMessage(ChatRequestMessage, discriminator="assistant"):
    """A request chat message representing response or action from the assistant.

    All required parameters must be populated in order to send to server.

    :ivar role: The chat role associated with this message, which is always 'assistant' for
     assistant messages. Required. The role that provides responses to system-instructed,
     user-prompted input.
    :vartype role: str or ~azure.ai.inference.models.ASSISTANT
    :ivar content: The content of the message. Required.
    :vartype content: str
    :ivar name: An optional name for the participant.
    :vartype name: str
    """

    role: Literal[ChatRole.ASSISTANT] = rest_discriminator(name="role")  # type: ignore
    """The chat role associated with this message, which is always 'assistant' for assistant messages.
     Required. The role that provides responses to system-instructed, user-prompted input."""
    content: str = rest_field()
    """The content of the message. Required."""
    name: Optional[str] = rest_field()
    """An optional name for the participant."""

    @overload
    def __init__(
        self,
        *,
        content: str,
        name: Optional[str] = None,
    ):
        ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, role=ChatRole.ASSISTANT, **kwargs)


class ChatRequestSystemMessage(ChatRequestMessage, discriminator="system"):
    """A request chat message containing system instructions that influence how the model will
    generate a chat completions
    response.

    All required parameters must be populated in order to send to server.

    :ivar role: The chat role associated with this message, which is always 'system' for system
     messages. Required. The role that instructs or sets the behavior of the assistant.
    :vartype role: str or ~azure.ai.inference.models.SYSTEM
    :ivar content: The contents of the system message. Required.
    :vartype content: str
    :ivar name: An optional name for the participant.
    :vartype name: str
    """

    role: Literal[ChatRole.SYSTEM] = rest_discriminator(name="role")  # type: ignore
    """The chat role associated with this message, which is always 'system' for system messages.
     Required. The role that instructs or sets the behavior of the assistant."""
    content: str = rest_field()
    """The contents of the system message. Required."""
    name: Optional[str] = rest_field()
    """An optional name for the participant."""

    @overload
    def __init__(
        self,
        *,
        content: str,
        name: Optional[str] = None,
    ):
        ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, role=ChatRole.SYSTEM, **kwargs)


class ChatRequestUserMessage(ChatRequestMessage, discriminator="user"):
    """A request chat message representing user input to the assistant.

    All required parameters must be populated in order to send to server.

    :ivar role: The chat role associated with this message, which is always 'user' for user
     messages. Required. The role that provides input for chat completions.
    :vartype role: str or ~azure.ai.inference.models.USER
    :ivar content: The contents of the user message, with available input types varying by selected
     model. Required.
    :vartype content: str
    :ivar name: An optional name for the participant.
    :vartype name: str
    """

    role: Literal[ChatRole.USER] = rest_discriminator(name="role")  # type: ignore
    """The chat role associated with this message, which is always 'user' for user messages. Required.
     The role that provides input for chat completions."""
    content: str = rest_field()
    """The contents of the user message, with available input types varying by selected model.
     Required."""
    name: Optional[str] = rest_field()
    """An optional name for the participant."""

    @overload
    def __init__(
        self,
        *,
        content: str,
        name: Optional[str] = None,
    ):
        ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, role=ChatRole.USER, **kwargs)


class ChatResponseMessage(_model_base.Model):
    """A representation of a chat message as received in a response.

    All required parameters must be populated in order to send to server.

    :ivar role: The chat role associated with the message. Required. Known values are: "system",
     "assistant", and "user".
    :vartype role: str or ~azure.ai.inference.models.ChatRole
    :ivar content: The content of the message. Required.
    :vartype content: str
    """

    role: Union[str, "_models.ChatRole"] = rest_field()
    """The chat role associated with the message. Required. Known values are: \"system\",
     \"assistant\", and \"user\"."""
    content: str = rest_field()
    """The content of the message. Required."""

    @overload
    def __init__(
        self,
        *,
        role: Union[str, "_models.ChatRole"],
        content: str,
    ):
        ...

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
        completion_tokens: int,
        prompt_tokens: int,
        total_tokens: int,
    ):
        ...

    @overload
    def __init__(self, mapping: Mapping[str, Any]):
        """
        :param mapping: raw JSON to initialize the model.
        :type mapping: Mapping[str, Any]
        """

    def __init__(self, *args: Any, **kwargs: Any) -> None:  # pylint: disable=useless-super-delegation
        super().__init__(*args, **kwargs)
