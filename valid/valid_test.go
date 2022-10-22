package valid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValid(t *testing.T) {
	ok := Email("test@test.com")
	assert.True(t, ok)
	ok = Email("test@test")
	assert.False(t, ok)
	ok = Okpd("0.0.0.0")
	assert.True(t, ok)
	ok = Okpd("10.20.33.44")
	assert.True(t, ok)
	ok = Okpd("10.20.33")
	assert.True(t, ok)
	ok = Okpd("99")
	assert.True(t, ok)
	ok = Okpd("AA.20.33.44")
	assert.False(t, ok)
	ok = Okpd("00.20.33.44000000000000000000000")
	assert.False(t, ok)
	ok = Rus("Проверка")
	assert.True(t, ok)
	ok = Rus("test")
	assert.False(t, ok)
	ok = Rus("123")
	assert.False(t, ok)
	ok = Rus(" Проверка")
	assert.False(t, ok)
	ok = Inn("1234567890")
	assert.True(t, ok)
	ok = Inn("1234")
	assert.False(t, ok)
	ok = Inn("")
	assert.True(t, ok)
	ok = RegionNum("01")
	assert.True(t, ok)
	ok = RegionNum("")
	assert.True(t, ok)
	ok = RegionNum("m6")
	assert.False(t, ok)
	ok = Title("Наименование")
	assert.True(t, ok)
	ok = Title(" Наименование")
	assert.False(t, ok)
	ok = Title("Закупки томографов по 44ФЗ")
	assert.True(t, ok)
	ok = Title("<h1>Наименование</h1>")
	assert.False(t, ok)
	ok = Title("ййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййййй")
	assert.False(t, ok)
	ok = Link("title=%D0%97%D0%B0%D0%BA%D1%83%D0%BF%D0%BA%D0%B8%20%D1%82%D0%BE%D0%BC%D0%BE%D0%B3%D1%80%D0%B0%D1%84%D0%BE%D0%B2%20%D0%BF%D0%BE%2044%D0%A4%D0%97&okpd=26.60.11.111&region=77")
	assert.True(t, ok)
	ok = Link("tenderdb.ru{11}")
	assert.False(t, ok)
	ok = Link("tenderdb.ru")
	assert.False(t, ok)
	ok = Space(" ")
	assert.True(t, ok)
	ok = Space("            ")
	assert.True(t, ok)
	ok = Space("1234")
	assert.False(t, ok)
}