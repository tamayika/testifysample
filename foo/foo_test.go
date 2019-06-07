package foo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FooTestSuite struct {
	suite.Suite
}

func (suite *FooTestSuite) SetupSuite() {
	fmt.Println("Foo SetupSuite")
}

func (suite *FooTestSuite) TearDownSuite() {
	fmt.Println("Foo TearDownSuite")
}

func (suite *FooTestSuite) SetupTest() {
	fmt.Println("Foo SetupTest")
}

func (suite *FooTestSuite) TearDownTest() {
	fmt.Println("Foo TearDownTest")
}

func (suite *FooTestSuite) TestFoo() {
	suite.Equal(10, foo())
}

func TestFooTestSuite(t *testing.T) {
	suite.Run(t, new(FooTestSuite))
}
