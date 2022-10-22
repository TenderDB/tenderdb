package repositories


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAutocomplete(t *testing.T) {
	repoTest := NewRepository()
	repoTest.AddAutocompleteDB("/usr/local/deploy/db/bolt_autocomplete.db")

	list, err := repoTest.Autocomplete("Пшеница озимая твердая")
	assert.Nil(t, err)
	listMustBe := make(map[string]string)
	listMustBe["Пшеница озимая твердая"] = "01.11.11.110"
	assert.Equal(t, listMustBe, list)
}
