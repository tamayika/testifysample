package foo

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type BarTestSuite struct {
	suite.Suite
}

func (suite *BarTestSuite) SetupSuite() {
	fmt.Println("Bar SetupSuite")
	time.Sleep(1 * time.Minute)
}

func (suite *BarTestSuite) TeardownSuite() {
	fmt.Println("Bar TeardownSuite")
	time.Sleep(1 * time.Minute)
}

func (suite *BarTestSuite) SetupTest() {
	fmt.Println("Bar SetupTest")
	time.Sleep(1 * time.Minute)
}

func (suite *BarTestSuite) TeardownTest() {
	fmt.Println("Bar TeardownTest")
	time.Sleep(1 * time.Minute)
}

func (suite *BarTestSuite) TestBar() {
	suite.Equal(10, bar())
}

func TestBarTestSuite(t *testing.T) {
	suite.Run(t, new(BarTestSuite))
}
