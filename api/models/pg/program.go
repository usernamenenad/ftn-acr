package models

import "github.com/google/uuid"

type Program struct {
	Id               uuid.UUID
	ProgramName      string
	ProgramShortName string
}
