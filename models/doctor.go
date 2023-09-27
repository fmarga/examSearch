package models

type Doctor struct {
	CRM   string `json:"crm"`
	State string `json:"state"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
