from src.services.handlers.handlers import ClassroomHandler
from src.services.rabbitmq.message import Message, MessageType


class HandlerInvoker:
    def invoke_handler(self, message: Message) -> None:
        match message.message_type:
            case MessageType.CLASSROOM_MESSAGE:
                ClassroomHandler.handle(message)
