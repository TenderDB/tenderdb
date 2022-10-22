package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/tenderdb/tenderdb/domain"
)


func TestFetch(t *testing.T) {
	r := NewRepository()
	r.AddMinDB("/usr/local/deploy/db/bolt_final.db")
	r.AddYears([]string{"2018"})
	r.AddFZs([]string{"44"})

	dataset, err := r.GetChartOnPrefix("01.11.11.111")
	recordedDataset := []domain.Dataset{domain.Dataset{Name:"01.11.11.111(223)", Data:map[string]string{"2019":"131496670", "2020":"142364358", "2021":"158785998", "2022":"1422613"}}, domain.Dataset{Name:"01.11.11.111(44)", Data:map[string]string{"2018":"222690443", "2019":"259820190", "2020":"102641005", "2021":"1034814458", "2022":"91779686"}}}


	assert.Equal(t, recordedDataset, dataset)
	assert.Nil(t, err)

}
