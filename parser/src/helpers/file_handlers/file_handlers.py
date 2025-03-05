from abc import ABC, abstractmethod
from typing import Any, Optional

import pdfplumber

from src.logger.logger import Logger


class FileHandler(ABC):
    @abstractmethod
    def open(self) -> Optional[Any]:
        pass

    @abstractmethod
    def close(self):
        pass


class PDFHandler(FileHandler):
    def __init__(self, file_path: str):
        self.file_path = file_path

    def open(self) -> Optional[pdfplumber.pdf.PDF]:
        try:
            return pdfplumber.open(self.file_path)
        except Exception as e:
            Logger().log(f"Error with opening PDF file. Reason: \r\n {e}")
            return None

    def close(self, pdf: pdfplumber.pdf.PDF):
        try:
            pdf.close()
        except Exception as e:
            Logger().log(f"Error with closing PDF file. Reason: \r\n {e}")
