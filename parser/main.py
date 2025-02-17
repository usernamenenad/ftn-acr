from src.helpers.extractors.extractors import TableExtractor
from src.helpers.file_handlers.file_handlers import PDFHandler
from src.services.parsers.parsers import TimeFromParser


def main() -> None:
    pdf_handler = PDFHandler("./.temp/pdf/rac-3.pdf")
    pdf = pdf_handler.open()
    m = []

    if pdf is not None:
        tables = TableExtractor().extract_all(pdf)
        for table in tables:
            m.extend(TimeFromParser().parse(table))


if __name__ == "__main__":
    main()
