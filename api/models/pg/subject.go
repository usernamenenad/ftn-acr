package models

import "github.com/google/uuid"

type Subject struct {
	Id          uuid.UUID
	SubjectName string
}
