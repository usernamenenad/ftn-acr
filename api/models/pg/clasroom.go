package models

import "github.com/google/uuid"

type Classroom struct {
	Id                uuid.UUID
	ClassroomName     string
	ClassroomBuilding string
}
