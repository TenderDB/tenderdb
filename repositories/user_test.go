package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)


func TestUsers(t *testing.T) {
	r := NewRepository()
	r.AddUsersDB("/usr/local/deploy/db/bolt_users.db")
	
	err := r.CreateUser("test@test.ru")
	assert.Nil(t, err)

	
	err = r.UpdateUser("test@test.ru", "download")
	infoMustBe := time.Now().Format(time.RFC3339)
	assert.Nil(t, err)
	
	info, err := r.CheckUserAction("test@test.ru", "download")
	assert.Nil(t, err)
	assert.Equal(t, infoMustBe, info)
	
	
}
