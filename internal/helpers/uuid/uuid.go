package uuid

import "github.com/google/uuid"

type uuidHelper struct{}

func NewUuidHelper() UuidHelper {
	return &uuidHelper{}
}

func (ref *uuidHelper) Generate() string {
	return uuid.NewString()
}

func (ref *uuidHelper) IsValidUUID(input string) bool {
	_, err := uuid.Parse(input)
	return err == nil
}
