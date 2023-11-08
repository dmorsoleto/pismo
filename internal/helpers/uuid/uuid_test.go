package uuid

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UuidHelperTestSuite struct {
	suite.Suite
	ctx        context.Context
	uuidHelper UuidHelper
}

func TestUuidHelperTest(t *testing.T) {
	suiteTest := &UuidHelperTestSuite{
		ctx: context.Background(),
	}
	suite.Run(t, suiteTest)
}

func (ref *UuidHelperTestSuite) SetupTest() {
	ref.uuidHelper = NewUuidHelper()
}

func (ref *UuidHelperTestSuite) TestGenerate_Success() {
	uuid := ref.uuidHelper.Generate()
	ref.NotEmpty(uuid)
	ref.True(ref.uuidHelper.IsValidUUID(uuid))
}

func (ref *UuidHelperTestSuite) TestIsValidUUID_Not_Valid_Success() {
	uuid := "123"
	isValid := ref.uuidHelper.IsValidUUID(uuid)
	ref.False(isValid)
}
