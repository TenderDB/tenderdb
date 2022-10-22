package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestTree(t *testing.T) {
	r := NewRepository()
	r.AddTreeDB("/usr/local/deploy/db/bolt_tree.db")

	subCatalogue, err := r.GetList("01.11.11.111")
	subCatalogueMustBe := map[string]string{"01.11.11.111":"Зерно озимой твердой пшеницы"}

	
	assert.Equal(t, subCatalogueMustBe, subCatalogue)
	assert.Nil(t, err)

	value, err := r.GetValue("01.11.11.111")
	valueMustBe := "Зерно озимой твердой пшеницы"
	
	assert.Equal(t, valueMustBe, value)
	assert.Nil(t, err)

}
