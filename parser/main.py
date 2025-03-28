from src.helpers.extractors.extractors import TableExtractor
from src.helpers.file_handlers.file_handlers import PDFHandler
from src.services.parsers.parsers import ClassroomsParser
# from src.services.rabbitmq.publisher import RabbitMqPublisher


def main() -> None:
    pdf_handler = PDFHandler("./.temp/pdf/rac-3.pdf")
    pdf = pdf_handler.open()
    m: set[str] = set()

    if pdf is not None:
        tables = TableExtractor().extract_all(pdf)
        for table in tables:
            m.update(ClassroomsParser().parse(table))

    print(m)


if __name__ == "__main__":
    main()
