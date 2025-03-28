from abc import ABC, abstractmethod
from typing import Any, List, Set

from src.services.parsers.types import Block, Day, Table, TitleType


class Parser(ABC):
    @abstractmethod
    def parse(self) -> Any:
        pass


class TitleParser(Parser):
    def parse(self, table: Table) -> TitleType:
        title = table[0][0]
        if title in Block.__members__.values():
            return TitleType.Block
        elif title in Day.__members__.values():
            return TitleType.Day
        else:
            return TitleType.NoType


class TimeFromParser(Parser):
    def parse(self, table: Table) -> List[str]:
        times = []
        time_from_index = self.__get_time_from_index(table)

        for row in table[2:]:
            times.append(row[time_from_index])

        return times

    def __get_time_from_index(self, table: Table) -> int:
        match TitleParser().parse(table):
            case TitleType.Block:
                return 2
            case TitleType.Day:
                return 1
            case TitleType.NoType:
                return 2


class TimeToParser(Parser):
    def parse(self, table: Table) -> List[str]:
        times: List[str] = []
        time_from_index = self.__get_time_to_index(table)

        for row in table[2:]:
            times.append(row[time_from_index])

        return times

    def __get_time_to_index(self, table: Table) -> int:
        match TitleParser().parse(table):
            case TitleType.Block:
                return 3
            case TitleType.Day:
                return 2
            case TitleType.NoType:
                return 3


class ClassroomsParser(Parser):
    def parse(self, table: Table) -> List[str]:
        classrooms: List[str] = []
        classrooms_index = self.__get_classroom_index(table)

        for row in table[2:]:
            classrooms.append(row[classrooms_index])

        return classrooms

    def __get_classroom_index(self, table: Table) -> int:
        match TitleParser().parse(table):
            case TitleType.Block:
                return 4
            case TitleType.Day:
                return 3
            case TitleType.NoType:
                return 4
