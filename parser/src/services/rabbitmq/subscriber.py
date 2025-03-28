import pika
import pika.adapters.blocking_connection
import pika.channel

from src.services.handlers.invoker import HandlerInvoker


class RabbitMqSubscriber:
    _self = None
    _connection: pika.BlockingConnection | None = None
    _channel: pika.adapters.blocking_connection.BlockingChannel | None = None

    def __new__(cls):
        if not cls._self:
            cls._self = super().__new__(cls)
            cls._connection = pika.BlockingConnection(
                pika.ConnectionParameters("rabbitmq"),
            )
            cls._channel = cls._connection.channel()
        return cls._self

    def subscribe(self, queue: str) -> None:
        self._channel.queue_declare(queue=queue)
        self._channel.basic_consume(
            queue=queue,
            auto_ack=True,
            on_message_callback=lambda message: HandlerInvoker().invoke_handler(
                message
            ),
        )
