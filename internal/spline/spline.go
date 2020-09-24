package spline

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

// CubicSpline is class of spline interpolation
type CubicSpline struct {
	X      []float64
	Y      []float64
	S      []float64
	A      [][]float64
	Coef   []float64
	Result []float64
	N      int
}

// Init create Spline instance, return it.
func Init(data [][]interface{}) CubicSpline {
	var sln CubicSpline
	if len(data) == 0 {
		return sln
	}
	switch data[0][0].(type) {
	case int:
		sln.convInttoFloat64(data)
		break
	case int32:
		sln.convInt32toFloat64(data)
		break
	case int64:
		sln.convInt64toFloat64(data)
		break
	case float32:
		sln.convFloat32toFloat64(data)
		break
	default:
		sln.convFloat64Slice(data)
		break
	}
	sln.N = len(sln.X)
	sln.A = make([][]float64, (sln.N-1)*4)
	sln.S = make([]float64, (sln.N-1)*4)
	sln.Coef = make([]float64, (sln.N-1)*4)
	for i := 0; i < (sln.N-1)*4; i++ {
		sln.A[i] = make([]float64, (sln.N-1)*4)
	}
	return sln
}

func pow(x float64, num int) float64 {
	var res float64
	res = 1
	for i := 0; i < num; i++ {
		res *= x
	}
	return res
}

func (sln *CubicSpline) convInttoFloat64(base [][]interface{}) {
	for _, row := range base {
		sln.X = append(sln.X, float64(row[0].(int)))
		sln.Y = append(sln.Y, float64(row[1].(int)))
	}
}

func (sln *CubicSpline) convInt32toFloat64(base [][]interface{}) {
	for _, row := range base {
		sln.X = append(sln.X, float64(row[0].(int32)))
		sln.Y = append(sln.Y, float64(row[1].(int32)))
	}
}

func (sln *CubicSpline) convInt64toFloat64(base [][]interface{}) {
	for _, row := range base {
		sln.X = append(sln.X, float64(row[0].(int64)))
		sln.Y = append(sln.Y, float64(row[1].(int64)))
	}
}

func (sln *CubicSpline) convFloat32toFloat64(base [][]interface{}) {
	for _, row := range base {
		sln.X = append(sln.X, float64(row[0].(float32)))
		sln.Y = append(sln.Y, float64(row[1].(float32)))
	}
}

func (sln *CubicSpline) convFloat64Slice(base [][]interface{}) {
	for _, row := range base {
		sln.X = append(sln.X, row[0].(float64))
		sln.Y = append(sln.Y, row[1].(float64))
	}
}

// Condition1 create a part of matrix A
func (sln *CubicSpline) Condition1() {
	for i := 0; i < sln.N-1; i++ {
		for j := 0; j < 4; j++ {
			sln.A[i][4*i+j] = pow(sln.X[i], j)
		}
		sln.S[i] = sln.Y[i]
	}
}

// Condition2 create a part of matrix A
func (sln *CubicSpline) Condition2() {
	for i := 0; i < sln.N-1; i++ {
		for j := 0; j < 4; j++ {
			sln.A[i+(sln.N-1)][4*i+j] = pow(sln.X[i+1], j)
		}
		sln.S[i+(sln.N-1)] = sln.Y[i+1]
	}
}

// Condition3 create a part of matrix A
func (sln *CubicSpline) Condition3() {
	for i := 0; i < sln.N-2; i++ {
		for j := 0; j < 4; j++ {
			sln.A[i+(sln.N-1)*2][4*i+j] = float64(j) * pow(sln.X[i+1], j-1)
			sln.A[i+(sln.N-1)*2][4*(i+1)+j] = float64(-j) * pow(sln.X[i+1], j-1)
		}
	}
}

// Condition4 create a part of matrix A
func (sln *CubicSpline) Condition4() {
	for i := 0; i < sln.N-2; i++ {
		sln.A[i+(sln.N-1)*3-1][4*i+3] = float64(3) * sln.X[i+1]
		sln.A[i+(sln.N-1)*3-1][4*i+2] = float64(1)
		sln.A[i+(sln.N-1)*3-1][4*(i+1)+3] = float64(-3) * sln.X[i+1]
		sln.A[i+(sln.N-1)*3-1][4*(i+1)+2] = float64(-1)
	}
}

// Condition5 create a part of matrix A
func (sln *CubicSpline) Condition5() {
	sln.A[(sln.N-1)*4-2][3] = float64(3) * sln.X[0]
	sln.A[(sln.N-1)*4-2][2] = float64(1)
	sln.A[4*(sln.N-1)-1][4*(sln.N-1)-1] = float64(3) * sln.X[(sln.N-1)]
	sln.A[4*(sln.N-1)-1][4*(sln.N-1)-2] = float64(1)
}

func (sln *CubicSpline) calclateA() {
	sln.Condition1()
	sln.Condition2()
	sln.Condition3()
	sln.Condition4()
	sln.Condition5()
}

func (sln *CubicSpline) fit(x float64) float64 {
	var val int
	for _, xi := range sln.X {
		if xi > x {
			val--
			break
		}
		if val == -1 {
			val++
		} else if val == sln.N-1 {
			val--
		}
		val++
	}
	a := sln.Coef[4*val+3]
	b := sln.Coef[4*val+2]
	c := sln.Coef[4*val+1]
	d := sln.Coef[4*val]
	return a*pow(x, 3) + b*pow(x, 2) + c*x + d
}

// Interpolation do interpolation from calclate result.
func (sln *CubicSpline) Interpolation(min, max, sec float64) []float64 {
	x := min
	for x <= max {
		val := sln.fit(x)
		sln.Result = append(sln.Result, val)
		x += sec
	}
	return sln.Result
}

// Calclate execute calclation of cubicspline base matrix
func (sln *CubicSpline) Calclate() {
	sln.calclateA()
	inverseMat := mat64.NewDense((sln.N-1)*4, (sln.N-1)*4, nil)
	baseMat := mat64.NewDense((sln.N-1)*4, (sln.N-1)*4, flatten(sln.A))
	result := mat64.NewDense((sln.N-1)*4, 1, nil)
	inverseMat.Inverse(baseMat)
	yMat := mat64.NewDense((sln.N-1)*4, 1, sln.S)
	result.Mul(inverseMat, yMat)
	for i := 0; i < (sln.N-1)*4; i++ {
		sln.Coef[i] = result.At(i, 0)
	}
}

func flatten(base [][]float64) []float64 {
	var res []float64
	for _, row := range base {
		res = append(res, row...)
	}
	return res
}

func matPrint(X mat64.Matrix) {
	fa := mat64.Formatted(X, mat64.Prefix(""), mat64.Squeeze())
	fmt.Printf("%v\n", fa)
}
