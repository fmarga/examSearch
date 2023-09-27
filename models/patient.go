package models

type Patient struct {
	CPF       string `json:"cpf"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	BirthDate string `json:"birthDate"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
}
