package sortedfunction_test

import (
	"sort"
	"testing"

	"GoLeetCode/sortedfunction"
	"GoLeetCode/tools"
	"github.com/stretchr/testify/suite"
)

type SortTestSuite struct {
	suite.Suite
}

func TestSortTestSuite(t *testing.T) {
	suite.Run(t, new(SortTestSuite))
}

func (s *SortTestSuite) TestBubble() {
	nums := tools.GenerateRandomInts(10)
	sortedfunction.Bubble(nums)
	actual := tools.Copy(nums)
	sort.Ints(actual)
	s.Equal(actual, nums)
}

func (s *SortTestSuite) TestInsert() {
	nums := tools.GenerateRandomInts(10)
	sortedfunction.Insert(nums)
	actual := tools.Copy(nums)
	sort.Ints(actual)
	s.Equal(actual, nums)
}

func (s *SortTestSuite) TestSelect() {
	nums := tools.GenerateRandomInts(10)
	sortedfunction.Select(nums)
	actual := tools.Copy(nums)
	sort.Ints(actual)
	s.Equal(actual, nums)
}

func (s *SortTestSuite) TestShell() {
	nums := tools.GenerateRandomInts(10)
	sortedfunction.Shell(nums)
	actual := tools.Copy(nums)
	sort.Ints(actual)
	s.Equal(actual, nums)
}
