from abc import ABC, abstractmethod

from src.services.rabbitmq.message import Message


class Handler(ABC):
    @abstractmethod
    def handle(self, message: Message) -> None:
        pass


class ClassroomHandler(Handler):
    def handle(self, message: Message) -> None:
        pass
