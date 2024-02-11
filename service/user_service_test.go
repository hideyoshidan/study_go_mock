package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"mock.com/repository/mocks"
)

func TestCreate(t *testing.T) {
	ur := mocks.NewMockIUserRepository(t)
	ur.On("CreateUser").Return(nil)
	us := NewUserService(ur)
	err := us.Create()
	if err != nil {
		t.Error("Error has been occurred.")
	}
	assert.NoError(t, err)
}
