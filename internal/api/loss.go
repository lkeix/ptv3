package api

import (
	"encoding/json"
	"fmt"
	"ptv3/internal/dblib"
	"ptv3/internal/spline"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Result is loss result
type Result struct {
	Dates       []string
	ExplainLoss [][]int64
	WorkLoss    [][]int64
	Rng         []string
}

// CalclateLoss calclate loss from exp, work log.
func CalclateLoss(ctx *gin.Context) Result {
	divide := ctx.PostForm("divideNum")
	patientID := ctx.PostForm("patient")
	rows := dblib.ReadPatientLog(patientID)
	var res Result
	if rows == nil {
		return res
	}
	var explains [][][]interface{}
	var works [][][]interface{}
	var dates []string
	for _, row := range rows {
		var explain [][]interface{}
		var work [][]interface{}
		exp := string(row[0].([]uint8))
		wo := string(row[1].([]uint8))
		date := row[2].(string)
		json.Unmarshal(([]byte)(exp), &explain)
		json.Unmarshal(([]byte)(wo), &work)
		explains = append(explains, explain[1:len(explain)])
		works = append(works, work[1:len(work)])
		dates = append(dates, date)
	}
	for i := 0; i < len(dates); i++ {
		res.Dates = append(res.Dates, dates[i])
		res.ExplainLoss = append(res.ExplainLoss, convInt64(calclateLoss(explains[i], divide)))
		res.WorkLoss = append(res.WorkLoss, convInt64(calclateLoss(works[i], divide)))
	}
	divideNum, _ := strconv.Atoi(divide)
	sum := 0
	for i := 0; i < divideNum; i++ {
		res.Rng = append(res.Rng, strconv.Itoa(sum)+"~"+strconv.Itoa(sum+(100/divideNum)))
		sum += (100 / divideNum)
	}
	return res
}

func convInt64(bases []float64) []int64 {
	var result []int64
	for _, base := range bases {
		result = append(result, int64(base))
	}
	return result
}

func calclateLoss(log [][]interface{}, divide string) []float64 {
	var rng []int
	var losses []float64
	divideNum, _ := strconv.Atoi(divide)
	sln := spline.Init(log)
	sln.Calclate()
	min := float64(0)
	max := float64(0)
	sec := 0.01
	switch log[0][0].(type) {
	case int:
		min = float64(log[0][0].(int))
		max = float64(log[len(log)-1][0].(int))
	case int32:
		min = float64(log[0][0].(int32))
		max = float64(log[len(log)-1][0].(int32))
	case int64:
		min = float64(log[0][0].(int64))
		max = float64(log[len(log)-1][0].(int64))
	case float32:
		min = float64(log[0][0].(float32))
		max = float64(log[len(log)-1][0].(float32))
	default:
		min = log[0][0].(float64)
		max = log[len(log)-1][0].(float64)
	}
	result := sln.Interpolation(min, max, sec)
	divideItr := 0
	for i := 0; i < divideNum; i++ {
		divideItr += len(result) / divideNum
		rng = append(rng, divideItr)
	}
	var j int
	for i := 0; i < len(rng); i++ {
		var loss float64
		for j < rng[i] {
			var height float64
			height = 4 - result[j]
			if result[j] > 4 {
				height = 4
			}
			if result[j] < 0 {
				height = 0
			}
			loss += height * sec
			j++
		}
		losses = append(losses, loss)
	}
	return losses
}

func dispresult(res []float64) {
	for _, row := range res {
		fmt.Println(row)
	}
}
