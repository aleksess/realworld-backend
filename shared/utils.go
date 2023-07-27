package shared

import "github.com/google/uuid"

func UUIDGenerator() ID {
	return ID(uuid.New().String())
}