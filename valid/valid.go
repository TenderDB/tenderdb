package valid

import (
	"regexp"
)

func Space(str string) bool {
	matched, _ := regexp.MatchString("^\\s.*$", str)
	return matched
}

func Inn(str string) bool {
	matched, _ := regexp.MatchString("^\\d{10,12}$|^$", str)
	return matched
}
func RegionNum(str string) bool {
	matched, _ := regexp.MatchString("^\\d{2}$|^$", str)
	return matched
}

func Rus(word string) bool {
	if Space(word) {
		return false
	}
	matched, _ := regexp.MatchString("^[АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя ]{1,250}$", word)
	return matched
}

func Email(email string) bool {
	matched, _ := regexp.MatchString("^[a-z0-9._%+-]{1,250}@[a-z0-9.-]{1,250}\\.[a-z]{2,20}$", email)
	return matched
}

func OkpdForChart(okpd string) bool {
	matched, _ := regexp.MatchString("^(\\d{2}\\.){3}\\d{3}$", okpd)
	return matched
}
func Okpd(okpd string) bool {
	matched, _ := regexp.MatchString("^[\\d\\.]{2,12}$", okpd)
	return matched
}
func ABC(word string) bool {
	matched, _ := regexp.MatchString("^[ABCDEFGHIJKLMNOPQRSTU]$", word)
	return matched
}
func Title(title string) bool {

	if Space(title) {
		return false
	}
	matched, _ := regexp.MatchString("^[a-zA-Z АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя\\s\\d\\-\\:\\(\\)\\\\]{1,200}$", title)
	return matched
}
func Link(title string) bool {
	matched, _ := regexp.MatchString("^([a-z]+\\=[\\da-zA-Z\\%\\.]+)(\\&[a-z]+\\=[\\da-zA-Z\\%i\\.]+){1,5}$", title)
	return matched
}
