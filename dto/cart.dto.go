package dto

import "github.com/google/uuid"

type UpsertCartInput struct {
	Products []uuid.UUID
}
