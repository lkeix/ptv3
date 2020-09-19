package api

import (
	"encoding/json"
	"ptv3/internal/dblib"
	"ptv3/internal/errorlib"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// PatientData is patient's data structure
type PatientData struct {
	PatientID int64
	Name      string
	Gender    string
	Age       int64
}

// PatientLogData is pattient's log data
type PatientLogData struct {
	Name string
	Date string
	Exp  string
	Work string
	Chk  string
}

// PastPatientData is past patient name, id data structure
type PastPatientData struct {
	Name string
	ID   int64
}

// DeletePatients is selected patients data to delete.
type DeletePatients struct {
	Patient []string `form:"patient[]"`
}

// GetPatientsData get doctor's patients data
func GetPatientsData(ctx *gin.Context, page interface{}) ([]PatientData, []int64) {
	session := sessions.Default(ctx)
	doctorID := ""
	if session.Get("doctorID") != nil {
		doctorID = session.Get("doctorID").(string)
	}
	base := dblib.GetPatientsData(doctorID)
	patiensData, pages, patientNum := convPatientData(base)
	pageitr := page.(int)
	pageLastitr := min(pageitr*10, int(patientNum))
	if pageLastitr < page.(int)-1 {
		ctx.AbortWithError(400, errorlib.ErrorStruct["400"])
		return nil, nil
	}
	return patiensData[(pageitr-1)*10 : pageLastitr], pages[max(pageitr-3, 0):min(pageitr+2, len(pages))]
}

// GetAllPatientsData get all patients data
func GetAllPatientsData(ctx *gin.Context) {
	session := sessions.Default(ctx)
	doctorID := ""
	if session.Get("doctorID") != nil {
		doctorID = session.Get("doctorID").(string)
	}
	base := dblib.GetPatientsData(doctorID)
	patiensData, _, _ := convPatientData(base)
	jsonByte, _ := json.Marshal(patiensData)
	ctx.Writer.Write(jsonByte)
}

// GetAllPatients get all patients data
func GetAllPatients(ctx *gin.Context) []PatientData {
	session := sessions.Default(ctx)
	doctorID := ""
	if session.Get("doctorID") != nil {
		doctorID = session.Get("doctorID").(string)
	}
	base := dblib.GetPatientsData(doctorID)
	patiensData, _, _ := convPatientData(base)
	return patiensData
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func convPatientData(base [][]interface{}) ([]PatientData, []int64, int64) {
	var patientsData []PatientData
	patientsData = make([]PatientData, len(base))
	for i := 0; i < len(base); i++ {
		patientsData[i].PatientID = base[i][0].(int64)
		patientsData[i].Name = base[i][1].(string)
		patientsData[i].Age = base[i][2].(int64)
		patientsData[i].Gender = genderNumConv(base[i][3].(int64))
	}
	pages := make([]int64, int64((len(patientsData)-1)/10)+1)
	for i := 0; i < len(pages); i++ {
		pages[i] = int64(i + 1)
	}
	return patientsData, pages, int64(len(patientsData))
}

func genderNumConv(gender int64) string {
	if gender == 1 {
		return "男性"
	}
	return "女性"
}

// ReadPatientData is api point to get patient data.
func ReadPatientData(ctx *gin.Context) {
	var res []PatientLogData
	req := ctx.Request
	req.ParseForm()
	patientid := req.PostFormValue("patientID")
	pastLog := dblib.ReadPatientData(patientid)
	for i := 0; i < len(pastLog); i++ {
		var tmp PatientLogData
		tmp.Exp = string(pastLog[i][0].([]uint8))
		tmp.Work = string(pastLog[i][1].([]uint8))
		tmp.Chk = string(pastLog[i][2].([]uint8))
		tmp.Date = pastLog[i][3].(string)
		tmp.Name = pastLog[i][4].(string)
		res = append(res, tmp)
	}
	jsonByte, _ := json.Marshal(res)
	ctx.Writer.Write(jsonByte)
}

// ReadPastPatients read past patients name, id
func ReadPastPatients(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	searchTxt := req.PostFormValue("searchText")
	session := sessions.Default(ctx)
	doctorid := session.Get("doctorID").(string)
	rows := dblib.ReadPastPatients(doctorid, searchTxt)
	var res []PastPatientData
	for _, row := range rows {
		var tmp PastPatientData
		tmp.ID = row[0].(int64)
		tmp.Name = row[1].(string)
		res = append(res, tmp)
	}
	jsonRes, _ := json.Marshal(res)
	ctx.Writer.Write(jsonRes)
}

// DeletePatient delete patient data.
func DeletePatient(ctx *gin.Context) {
	var patients DeletePatients
	ctx.Bind(&patients)
	for _, patient := range patients.Patient {
		dblib.DeletePatient(patient)
	}
	ctx.Redirect(303, "/home")
}
