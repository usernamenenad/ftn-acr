from enum import Enum, StrEnum
from typing import List, Optional

Table = List[List[List[Optional[str]]]]


class Day(StrEnum):
    MONDAY = "P O N E D E L J A K"
    TUESDAY = "U T O R A K"
    WEDNESDAY = "S R E D A"
    THURSDAY = "Č E T V R T A K"
    FRIDAY = "P E T A K"
    SATURDAY = "S U B O T A"
    SUNDAY = "N E D E L J A"


class Block(StrEnum):
    BLOCK_CLASSES = "B L O K N A S T A V A"


class TitleType(Enum):
    Day = Day
    Block = Block
    NoType = None
