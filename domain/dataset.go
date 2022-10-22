package domain

type Dataset struct {
	Name string            `json:"name"`
	Data map[string]string `json:"data"`
}
