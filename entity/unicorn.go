package entity

type Unicorn struct {
	Name         string   `json:"name"`
	Id           int      `json:"id"`
	Capabilities []string `json:"capabilities"`
}
