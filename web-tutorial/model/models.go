package model

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"web-tutorial/config"
)

// Fields first letter must be uppercase so we can export them to templates.
type Patient struct {
	Patient_id string
	Name       string
	Lastname   string
	Gender     string
	Age        int
}

// Show all patients
func ShowAllPatients() ([]Patient, error) {

	rows, err := config.DB.Query("SELECT * FROM tbl_patientinfo")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	patients := make([]Patient, 0)
	for rows.Next() {
		patient := Patient{}
		err := rows.Scan(&patient.Patient_id, &patient.Name, &patient.Lastname, &patient.Gender, &patient.Age)
		if err != nil {

			return nil, err
		}
		patients = append(patients, patient)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return patients, nil
}

func SinglePatient(r *http.Request) (Patient, error) {
	patient := Patient{}
	patient_id := r.FormValue("id")
	if patient_id == "" {
		var bk Patient
		return bk, errors.New("400. Bad Request.")
	}
	fmt.Println("id:" + patient_id)

	row := config.DB.QueryRow("SELECT * FROM tbl_patientinfo WHERE patient_id = ?", patient_id)

	err := row.Scan(&patient.Patient_id, &patient.Name, &patient.Lastname, &patient.Gender, &patient.Age)
	if err != nil {
		log.Println(err)
		return patient, err
	}
	return patient, nil
}

func InsertPatient(r *http.Request) (Patient, error) {
	p := Patient{}
	p.Patient_id = r.FormValue("patientID")
	p.Lastname = r.FormValue("lastname")
	p.Gender = r.FormValue("gender")
	p.Name = r.FormValue("name")
	pID := r.FormValue("age")
	age, err := strconv.Atoi(pID)
	p.Age = age
	// validate form values
	if p.Patient_id == "" || p.Name == "" || p.Lastname == "" || p.Gender == "" || pID == "" {
		return p, errors.New("400. Bad request. All fields must be complete.")
	}

	// insert values
	_, err = config.DB.Exec("INSERT INTO tbl_patientinfo (patient_id, name, lastname, gender, age) VALUES (?, ?, ?, ?, ?)", p.Patient_id, p.Name, p.Lastname, p.Gender, p.Age)
	if err != nil {
		log.Println(err)
		return p, errors.New("500. Internal Server Error." + err.Error())
	}
	return p, nil
}

func UpdatePatient(r *http.Request) (Patient, error) {
	// get form values
	p := Patient{}
	p.Patient_id = r.FormValue("patientID")
	p.Name = r.FormValue("name")
	p.Lastname = r.FormValue("lastname")
	p.Gender = r.FormValue("gender")
	// converst the string value to integer
	pID := r.FormValue("age")
	age, err := strconv.Atoi(pID)
	p.Age = age
	// validate form values
	if p.Patient_id == "" || p.Name == "" || p.Lastname == "" || p.Gender == "" || pID == "" {
		return p, errors.New("400. Bad request. All fields must be complete.")
	}

	// insert values
	_, err = config.DB.Exec("UPDATE tbl_patientinfo SET `name`=?, lastname=?, gender=?, age=? WHERE patient_id=?;", p.Name, p.Lastname, p.Gender, p.Age, p.Patient_id)
	if err != nil {
		return p, err
	}
	return p, nil
}

func DeletePatient(r *http.Request) error {
	pID := r.FormValue("id")
	if pID == "" {
		return errors.New("400. Bad Request.")
	}

	_, err := config.DB.Exec("DELETE FROM tbl_patientinfo WHERE patient_id=?;", pID)
	if err != nil {
		log.Println(err)
		return errors.New("500. Internal Server Error")
	}
	return nil
}
