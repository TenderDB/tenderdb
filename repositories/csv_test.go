package repositories

import (
	"bytes"
	"encoding/csv"
	"github.com/stretchr/testify/assert"
	
	"testing"
)

func TestCsv(t *testing.T) {
	r := NewRepository()
	r.AddMaxDB("/usr/local/deploy/db/bolt_%s_%s.db")
	r.AddYears([]string{"2018"})
	r.AddFZs([]string{"44"})

	buffer, err := r.CsvDownload("01.11.11.111", "", "", "")
	assert.Nil(t, err)


	reader := csv.NewReader(bytes.NewReader(buffer))
	record, err := reader.Read()
	assert.Nil(t, err)
	recordMustBe := []string{"Год", "ФЗ", "Регион", "ИНН поставщика", "ИНН заказчика", "Номер контракта", "Номер лота", "Сумма лота", "Комментарий"} 
	assert.Equal(t, recordMustBe, record)
	record, err = reader.Read()
	assert.Nil(t, err)
	recordMustBe = []string{"2018", "44", "02", "0273034305", "0224000247", "1027303430518000014", "01", "140000.00", "-"}
	assert.Equal(t, recordMustBe, record)
}
