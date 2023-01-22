package utils

import uuid "github.com/satori/go.uuid"

func GetBID() string {
	return uuid.NewV4().String()
}
