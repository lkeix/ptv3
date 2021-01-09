package api

import (
	"fmt"
	"ptv3/internal/dblib"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoggerStep return logger step number
// 1: form
// 2: explain
// 3: work
// 4: check
func LoggerStep(ctx *gin.Context) int64 {
	session := sessions.Default(ctx)
	step := session.Get("step").(int64)
	return step
}

// NewPatientStart start taking new patient log.
func NewPatientStart(ctx *gin.Context) {
	patientName := ctx.PostForm("name")
	patientGender := ctx.PostForm("gender")
	patientAge := ctx.PostForm("age")
	genderNum := genderConvNumber(patientGender)
	fmt.Println(patientAge)
	fmt.Println(patientGender)
	patientID := dblib.InsertNewPatient(patientName, patientAge, genderNum)
	session := sessions.Default(ctx)
	step := session.Get("step").(int64)
	step++
	session.Set("step", step)
	session.Set("patientID", patientID)
	session.Save()
	ctx.Redirect(303, "/ptv3/logger")
}

// PastPatientStart start taking past patient log.
func PastPatientStart(ctx *gin.Context) {
	patientID := ctx.PostForm("patientID")
	session := sessions.Default(ctx)
	if dblib.IsExistPatient(patientID) {
		patientIDInt, _ := strconv.Atoi(patientID)
		step := session.Get("step").(int64)
		step++
		session.Set("step", step)
		session.Set("patientID", int64(patientIDInt))
		session.Save()
		ctx.Redirect(303, "/ptv3/logger")
	} else {
		ctx.Redirect(303, "/ptv3/logger")
	}
}

// StorePatientExp store explain log
func StorePatientExp(ctx *gin.Context) {
	session := sessions.Default(ctx)
	step := session.Get("step").(int64)
	step++
	session.Set("step", step)
}

func genderConvNumber(gender string) string {
	if gender == "male" {
		return "1"
	}
	return "2"
}

// NextStep increment session step
func NextStep(ctx *gin.Context) {
	session := sessions.Default(ctx)
	step := session.Get("step").(int64)
	step++
	session.Set("step", step)
	session.Save()
}

func initStep(ctx *gin.Context) {
	session := sessions.Default(ctx)
	step := session.Get("step").(int64)
	step = 1
	session.Set("step", step)
	session.Save()
}

// StoreLog store explain, work, check log
func StoreLog(ctx *gin.Context) {
	initStep(ctx)
	session := sessions.Default(ctx)
	doctorid := session.Get("doctorID").(string)
	patientid := session.Get("patientID").(int64)
	exp := ctx.PostForm("exp")
	work := ctx.PostForm("work")
	/*
		q1 := ctx.PostForm("q1")
		q2 := ctx.PostForm("q2")
		q3 := ctx.PostForm("q3")
		q4 := ctx.PostForm("q4")
		q5 := ctx.PostForm("q5")
		q6 := ctx.PostForm("q6")
		q7 := ctx.PostForm("q7")
		q8 := ctx.PostForm("q8")
		q9 := ctx.PostForm("q9")
		q10 := ctx.PostForm("q10")
		q11 := ctx.PostForm("q11")
		q12 := ctx.PostForm("q12")
		q13 := ctx.PostForm("q13")
		q14 := ctx.PostForm("q14")
		q15 := ctx.PostForm("q15")
		q16 := ctx.PostForm("q16")
		dist := map[string]string{
			"q1":  q1,
			"q2":  q2,
			"q3":  q3,
			"q4":  q4,
			"q5":  q5,
			"q6":  q6,
			"q7":  q7,
			"q8":  q8,
			"q9":  q9,
			"q10": q10,
			"q11": q11,
			"q12": q12,
			"q13": q13,
			"q14": q14,
			"q15": q15,
			"q16": q16,
		}
	*/
	// distByte, _ := json.Marshal(dist)
	dblib.InsertPatientLog(doctorid, patientid, exp, work, "{}")
	ctx.Redirect(303, "/ptv3/logger")
}
