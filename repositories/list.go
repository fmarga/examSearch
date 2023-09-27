package repositories

import (
	"examSearch/db"
	"examSearch/models"
)

func List() (exams []models.Exam, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM exams")
	if err != nil {
		return
	}

	for rows.Next() {
		var exam models.Exam

		rows.Scan(&exam.Patient.CPF, &exam.Patient.Name, &exam.Patient.Email, &exam.Patient.BirthDate, &exam.Patient.Address, &exam.Patient.City, &exam.Patient.State, &exam.Doctor.CRM, &exam.Doctor.State, &exam.Doctor.Name, &exam.Doctor.Email, &exam.Token, &exam.Date, &exam.Type, &exam.Limits, &exam.Results)
		if err != nil {
			continue
		}

		exams = append(exams, exam)
	}

	return exams, nil
}
