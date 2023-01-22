package utils

import uuid "github.com/satori/go.uuid"

func GetFID() string {
	return uuid.NewV4().String()
}
