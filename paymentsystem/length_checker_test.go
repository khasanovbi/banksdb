package paymentsystem

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExactLengthCheckerTestSuite struct {
	suite.Suite
	length        int
	lengthChecker lengthChecker
}

func (suite *ExactLengthCheckerTestSuite) SetupSuite() {
	suite.length = 10
	suite.lengthChecker = &exactLengthChecker{Exact: suite.length}
}

func (suite *ExactLengthCheckerTestSuite) TestValidLength() {
	suite.Require().True(suite.lengthChecker.CheckLength(suite.length))
}

func (suite *ExactLengthCheckerTestSuite) TestInvalidLength() {
	const oversizeLengthAddition = 1

	suite.Require().False(suite.lengthChecker.CheckLength(suite.length + oversizeLengthAddition))
}

func TestExactLengthCheckerTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ExactLengthCheckerTestSuite))
}

type RangeLengthCheckerTestSuite struct {
	suite.Suite
	lengthFrom    int
	lengthTo      int
	lengthMiddle  int
	lengthOut     int
	lengthChecker lengthChecker
}

func (suite *RangeLengthCheckerTestSuite) SetupSuite() {
	suite.lengthFrom = 10
	suite.lengthTo = 20
	suite.lengthMiddle = 15
	suite.lengthOut = 21
	suite.lengthChecker = &rangeLengthChecker{From: suite.lengthFrom, To: suite.lengthTo}
}

func (suite *RangeLengthCheckerTestSuite) TestValidLeft() {
	suite.Require().True(suite.lengthChecker.CheckLength(suite.lengthFrom))
}

func (suite *RangeLengthCheckerTestSuite) TestValidRight() {
	suite.Require().True(suite.lengthChecker.CheckLength(suite.lengthTo))
}

func (suite *RangeLengthCheckerTestSuite) TestValidMiddle() {
	suite.Require().True(suite.lengthChecker.CheckLength(suite.lengthMiddle))
}

func (suite *RangeLengthCheckerTestSuite) TestInvalidOut() {
	suite.Require().False(suite.lengthChecker.CheckLength(suite.lengthOut))
}

func TestRangeLengthCheckerTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(RangeLengthCheckerTestSuite))
}

type OneOfLengthCheckerTestSuite struct {
	suite.Suite
	lengthFirst   int
	lengthSecond  int
	lengthInvalid int
	lengthChecker lengthChecker
}

func (suite *OneOfLengthCheckerTestSuite) SetupSuite() {
	suite.lengthFirst = 10
	suite.lengthSecond = 20
	suite.lengthInvalid = 30
	suite.lengthChecker = &oneOfLengthChecker{suite.lengthFirst, suite.lengthSecond}
}

func (suite *OneOfLengthCheckerTestSuite) TestValidOne() {
	suite.Require().True(suite.lengthChecker.CheckLength(suite.lengthFirst))
}

func (suite *OneOfLengthCheckerTestSuite) TestInvalid() {
	suite.Require().False(suite.lengthChecker.CheckLength(suite.lengthInvalid))
}

func TestOneOfLengthCheckerTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(OneOfLengthCheckerTestSuite))
}
