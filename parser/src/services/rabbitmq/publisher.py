import pika
import pika.adapters.blocking_connection
import pika.channel


class RabbitMqPublisher:
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

    def publish(self, queue: str, message: any) -> None:
        self._channel.queue_declare(queue=queue)
        self._channel.basic_publish(
            exchange="",
            routing_key=queue,
            body=message,
        )
