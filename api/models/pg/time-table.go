package models

import (
	"github.com/google/uuid"
)

type TimeTable struct {
	Id          uuid.UUID
	SubjectId   uuid.UUID
	ClassroomId uuid.UUID
	Class       string
	HoldingDay  string
	TimeTo      string
	TimeFrom    string
	ProgramId   string
	StudyGroups []string
	HoldingDate string
}
