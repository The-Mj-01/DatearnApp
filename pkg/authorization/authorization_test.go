package authorization

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// testingType for testing purpose
type testingType struct {
	Field1 string
	Field2 int
}

// customErr for testing purpose
var customErr error = errors.New("Custom testing error")

func TestSimpleFieldAuthorization(t *testing.T) {
	testObj := testingType{
		Field1: "title",
		Field2: 2,
	}

	err := SimpleFieldAuthorization(testObj, "unkownField", 2, customErr)
	assert.Error(t, err, "Detection test of passing wrong field failed")
	assert.ErrorIs(t, err, wrongFieldError, "Detection test of passing wrong field failed")

	err = SimpleFieldAuthorization(testObj, "Field2", 3, customErr)
	assert.Error(t, err, "Testing two fields are not equal failed")
	assert.ErrorIs(t, err, customErr, "Testing two fields are not equal failed")

	err = SimpleFieldAuthorization(testObj, "Field2", 2, customErr)
	assert.NoError(t, err, "Testing every thing being ok failed")
}
