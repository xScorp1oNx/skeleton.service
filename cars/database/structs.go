package database

type Car struct {
	ID      string `json:"id,omitempty"`
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Created string `json:"created"`
}
