package repositories


import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/tenderdb/tenderdb/domain"
)


func TestFilter(t *testing.T) {
	r := NewRepository()
	r.AddMaxDB("/usr/local/deploy/db/bolt_%s_%s.db")
	r.AddYears([]string{"2018"})
	r.AddFZs([]string{"44"})

	dataset, err := r.GetChartOnPrefixFilter("01.11.11.111", "", "", "")
	datasetMustBe := []domain.Dataset{{Name:"01.11.11.111(44)", Data:map[string]string{"2018":"222690443"}}}

	assert.Equal(t, datasetMustBe, dataset)
	assert.Nil(t, err)

}
