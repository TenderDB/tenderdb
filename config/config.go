package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Specification struct {
	Port           string   `default:"8000"`
	UsersDB        string   `default:"../db/bolt_users.db"`
	CartsDB        string   `default:"../db/bolt_carts.db"`
	TreeDB        string   `default:"../db/bolt_tree.db"`
	MinDB          string   `default:"../db/bolt_test_final.db"`
	AutocompleteDB string   `default:"../db/bolt_autocomplete.db"`
	CodesDB        string   `default:"../db/bolt_tree.db"`
	Templates      string   `default:"../templates/*.html"`
	Assets		string   `default:"../assets"`
	MaxDB          string   `default:"../db/bolt_test_%s_%s.db"`
	Years          []string `default:"2020,2021"`
	FZs            []string `default:"44,223"`
	YandexID       string   `required:"true"`
	YandexSecret   string   `required:"true"`
	YandexRedirect   string   `required:"true"`
	GoogleID       string   `required:"true"`
	GoogleSecret   string   `required:"true"`
	GoogleRedirect   string   `required:"true"`
	Session        string   `required:"true"`
}

func NewConfig() *Specification {
	var s Specification
	err := envconfig.Process("go", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &s
}
