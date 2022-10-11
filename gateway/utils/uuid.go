package utils

import UUID "github.com/google/uuid"

func IsValidUUID(uuid string) bool {
	_, err := UUID.Parse(uuid)
	return err == nil
}
