package module

import (
	"testing"

	test "study/src/module/__test__"

	"github.com/stretchr/testify/assert"
)

func TestModule(t *testing.T) {
	dbSpy := &test.DatabaseSpy{}
	dbSpy.On("GetTable").Return("First Table")

	m := Module(dbSpy)
	m.DoSomething(2)
	dbSpy.AssertExpectations(t)
}

func TestShouldReturnSometingStructCorrectly(t *testing.T) {
	dbSpy := &test.DatabaseSpy{}
	dbSpy.On("GetTable").Return("First Table")

	m := Module(dbSpy)
	result := m.DoSomething(2)
	assert.Equal(t, &SomethingStruct{id: 2, name: "First Table"}, result, "Should Return SometingStruct Correctly")
}
