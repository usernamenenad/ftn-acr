from abc import ABC
from typing import List, Optional

import pdfplumber
import pdfplumber.page

from src.services.parsers.types import Table


class Extractor(ABC):
    pass


class TableExtractor(Extractor):
    def extract_all(
        self, pdf: pdfplumber.pdf.PDF, from_page: Optional[int] = 1
    ) -> List[Table]:
        tables: List[Table] = []

        for page in pdf.pages[from_page:]:
            tables.extend(self.extract_from_page(page))

        return tables

    def extract_from_page(self, page: pdfplumber.page.Page) -> List[Table]:
        return page.extract_tables()
