package dblib

import (
	"strconv"
	"time"
)

// GetPatientsData get patients's data from doctorid
func GetPatientsData(doctorid string) [][]interface{} {
	query := "select distinct patient.id, patient.name, patient.age, patient.gender from doctor, patient, log where doctor.id = log.doctorid and log.patientid = patient.id and doctor.id = $1"
	args := []interface{}{
		doctorid,
	}
	if readMasterDoctor(doctorid) {
		query = "select distinct patient.id, patient.name, patient.gender, patient.age " +
			"from doctor, institute, log, patient " +
			"where doctor.instituteid in (select doctor.instituteid from doctor, institute where doctor.id = $1 and doctor.instituteid = institute.id) and " +
			"log.doctorid = doctor.id and patient.id = log.patientid"
	}
	rows, _ := GetRow(query, args)
	return rows
}

// ReadDoctorAuthorizeInfo read doctor hashed password from doctorid
func ReadDoctorAuthorizeInfo(doctorid string) string {
	query := "select password from doctor where ID = $1"
	args := []interface{}{
		doctorid,
	}
	rows, _ := GetRow(query, args)
	if len(rows) == 0 {
		return ""
	}
	return rows[0][0].(string)
}

// ReadInstitutionID read instituteID from institution name.
func ReadInstitutionID(institutionName string) int64 {
	query := "select ID from institute where name = $1"
	args := []interface{}{institutionName}
	rows, _ := GetRow(query, args)
	if len(rows) == 0 {
		return -1
	}
	return rows[0][0].(int64)
}

// InsertNewDoctor create new doctor basic data.
func InsertNewDoctor(id, password string, master bool, instituteID int64) bool {
	query := "insert into doctor (id, password, master, instituteid) values ($1, $2, $3, $4)"
	args := []interface{}{
		id,
		password,
		master,
		instituteID,
	}
	return Exec(query, args)
}

// InsertNewPatient create new Patient basic infomation. returning patientID
func InsertNewPatient(name, age, gender string) int64 {
	query := "insert into patient (name, age, gender) values ($1, $2, $3) returning id"
	args := []interface{}{
		name, age, gender,
	}
	idMat, _ := GetRow(query, args)
	return idMat[0][0].(int64)
}

// InsertPatientLog insert patient work log.
func InsertPatientLog(doctorid string, patientid int64, exp, work, check string) {
	t := time.Now()
	today := t.Format("2006/01/02")
	query := "insert into log (doctorid, patientid, exp, work, chk, date) values ($1, $2, $3, $4, $5, $6)"
	args := []interface{}{
		doctorid, patientid, exp, work, check, today,
	}
	Exec(query, args)
}

func readMasterDoctor(doctorid string) bool {
	query := "select master from doctor where id = $1"
	args := []interface{}{
		doctorid,
	}
	rows, _ := GetRow(query, args)
	if len(rows) == 0 {
		return false
	}
	return rows[0][0].(bool)
}

// ReadPatientData read patient past log
func ReadPatientData(patientID string) [][]interface{} {
	query := "select log.exp, log.work, log.chk, log.date, patient.name from log, patient where patientid = $1 and patient.id = log.patientid"
	args := []interface{}{
		patientID,
	}
	rows, _ := GetRow(query, args)
	return rows
}

// ReadPastPatients read past patients data from doctor id.
func ReadPastPatients(doctorid, searchTxt string) [][]interface{} {
	searchTxt = "%" + searchTxt + "%"
	query := "select patient.id, patient.name from patient, log where log.doctorid = $1 and log.patientid = patient.id and patient.name like $2"
	if readMasterDoctor(doctorid) {
		query = "select patient.id patient.name from doctor, institute, log, patient " +
			"where doctor.instituteid in (select doctor.instituteid from doctor, institute where doctor.id = $1 and doctor.instituteid = institute.id) and " +
			"log.doctorid = doctor.id and patient.id = log.patientid and patient.name like $2"
	}
	args := []interface{}{
		doctorid, searchTxt,
	}
	rows, _ := GetRow(query, args)
	return rows
}

// IsExistPatient check past patient existence or not.
func IsExistPatient(patientID string) bool {
	query := "select * from patient where patient.id = $1"
	patientIDInt, err := strconv.Atoi(patientID)
	if err != nil {
		return false
	}
	args := []interface{}{
		patientIDInt,
	}
	rows, _ := GetRow(query, args)
	if len(rows) > 0 {
		return true
	}
	return false
}

// DeletePatient delete patient data.
func DeletePatient(patientID string) {
	query := "delete from patient where id = $1"
	args := []interface{}{
		patientID,
	}
	Exec(query, args)
}
