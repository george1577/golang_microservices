// put the test file in the same folder of the functions we want to test.
package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotUserFound(t *testing.T) {
	user, err := UserDao.GetUser(0)

	// cleaner way to test
	assert.Nil(t, user, "we were not expecting a user with id 0") // test if equals to nil
	assert.NotNil(t, err)                                         // test if not equal to nil
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)

	// more verbose way
	// the following if statements are not expected to be entered, if enter, means the test failed, so invoke an error
	if user != nil {
		t.Error("we were not expecting a user with id 0")
	}

	if err == nil {
		t.Error("we were expecting an error when user id id 0")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Error("we were expecting 404 code when user is not found")
	}
}

func TestGetUserNoError(t *testing.T) {
	// Execution
	user, err := UserDao.GetUser(123)

	// Validation
	assert.Nil(t, err, "we are expecting err to be nil if use is found")
	assert.NotNil(t, user, "user should not be nil if ID is found in database")
	assert.EqualValues(t, 123, user.ID, "we are expecting ID to be 123")
	assert.EqualValues(t, "George", user.FirstName, "user ID 123 should be George")
	assert.EqualValues(t, "Chen", user.LastName, "user ID 123 should be Chen")
}
