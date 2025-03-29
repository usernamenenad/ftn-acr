from dataclasses import dataclass
from enum import Enum


class MessageType(Enum):
    CLASSROOM_MESSAGE = 0


@dataclass
class Message:
    message_type: MessageType
    message_load: bytearray


@dataclass
class ClassroomsMessage(Message):
    pass
