package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestCart(t *testing.T) {
	
	r := NewRepository()
	r.AddCartsDB("/usr/local/deploy/db/bolt_carts.db")

	//Test Updating
	err  := r.UpdateChart("test@test.ru", "One", "tenderdb.ru?okpd=222.222.22")
	assert.Nil(t, err)
	err = r.UpdateChart("test@test.ru", "Two", "tenderdb.ru?okpd=222.222.22")
	assert.Nil(t, err)

	//Test Reading
	message, err := r.ReadCharts("test@test.ru")
	messageMustBe := make(map[string]string)
	messageMustBe["One"] = "tenderdb.ru?okpd=222.222.22"
	messageMustBe["Two"] = "tenderdb.ru?okpd=222.222.22"
	assert.Equal(t, message, messageMustBe)
	assert.Nil(t, err)

	//Test Deleting
	err = r.DeleteChart("test@test.ru", "One")
	assert.Nil(t, err)
	err = r.DeleteChart("test@test.ru", "Two")
	assert.Nil(t, err)

	message, err = r.ReadCharts("test@test.ru")
	messageMustBe = make(map[string]string)
	assert.Equal(t, message, messageMustBe)
	assert.Nil(t, err)
}
