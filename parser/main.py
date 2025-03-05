from src.helpers.extractors.extractors import TableExtractor
from src.helpers.file_handlers.file_handlers import PDFHandler
from src.services.parsers.parsers import TimeFromParser
from src.services.rabbitmq.publisher import RabbitMqPublisher


def main() -> None:
    # pdf_handler = PDFHandler("./.temp/pdf/rac-3.pdf")
    # pdf = pdf_handler.open()
    # m = []

    # if pdf is not None:
    #     tables = TableExtractor().extract_all(pdf)
    #     for table in tables:
    #         m.extend(TimeFromParser().parse(table))
    pub = RabbitMqPublisher()
    pub.publish("hello", "world")


if __name__ == "__main__":
    main()
