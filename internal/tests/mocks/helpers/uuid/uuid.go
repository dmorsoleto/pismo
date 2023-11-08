package uuidmocks

import "github.com/stretchr/testify/mock"

type UuidHelperMock struct {
	mock.Mock
}

func (ref *UuidHelperMock) Generate() string {
	args := ref.Called()
	return args.String(0)
}

func (ref *UuidHelperMock) IsValidUUID(input string) bool {
	args := ref.Called(input)
	return args.Bool(0)
}
