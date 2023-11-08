package uuid

type UuidHelper interface {
	Generate() string
	IsValidUUID(input string) bool
}
