package main

// func dotProduct(a, b []float64) float64 {
// 	if len(a) != len(b) {
// 		log.Fatal()
// 	}
// 	var product float64
// 	for i := 0; i < len(a); i++ {
// 		product += a[i] * b[i]
// 	}
// 	return product
// }

// func multiplyMatrices(a, b [][]float64) [][]float64 {
// 	rowsA := len(a)
// 	rowsB := len(b)
// 	colsA := len(a[0])
// 	colsB := len(b[0])

// 	if colsA != rowsB {
// 		log.Fatal("Несоответствие размеров матриц")
// 	}

// 	result := make([][]float64, rowsA)
// 	for i := range result {
// 		result[i] = make([]float64, colsB)
// 	}

//		for i := 0; i < rowsA; i++ {
//			for j := 0; j < colsB; j++ {
//				var sum float64
//				for k := 0; k < colsA; k++ {
//					sum += a[i][k] * b[k][j]
//				}
//				result[i][j] = sum
//			}
//		}
//		return result
//	}
// func main() {

// 	a := mat.NewVecDense(3, []float64{1, 2, 3})
// 	b := mat.NewVecDense(3, []float64{4, 5, 6})
// 	dot := mat.Dot(a, b)
// 	fmt.Println("Скалярное произведение:", dot)

// 	m1 := mat.NewDense(2, 3, []float64{
// 		1, 2, 3,
// 		4, 5, 6,
// 	})
// 	m2 := mat.NewDense(3, 2, []float64{
// 		7, 8,
// 		9, 10,
// 		11, 12,
// 	})

// 	var result mat.Dense

// 	result.Mul(m1, m2)

// 	fmt.Println("Умножение матриц m1 * m2:")
// 	fmt.Printf("%v\n", mat.Formatted(&result, mat.Prefix(""), mat.Excerpt(0)))
// }

// import (
// 	"fmt"

// 	"gonum.org/v1/gonum/stat"
// )

// func main() {
// 	data := []float64{5, 7, 8, 9, 10, 6}

// 	mean := stat.Mean(data, nil)
// 	variance := stat.Variance(data, nil)
// 	sdt := stat.StdDev(data, nil)

// 	fmt.Println("Данные:", data)
// 	fmt.Println("Среднее:", mean)
// 	fmt.Println("Дисперсия:", variance)
// 	fmt.Println("Среднеквадратическое отклонение:", sdt)

// 	x := []float64{1, 2, 3, 4, 5}
// 	y := []float64{2, 4, 6, 8, 10}

// 	cov := stat.Covariance(x, y, nil)
// 	corr := stat.Correlation(x, y, nil)

// 	fmt.Printf("\nКовариация(x, y): %.2f\n", cov)
// 	fmt.Printf("Корреляция(x, y): %.2f\n", corr)
// }

// import (
// 	"fmt"

// 	"gonum.org/v1/gonum/mat"
// )

// func LinearRegression(xData, yData []float64) *mat.VecDense {
// 	nSamples := len(xData)

// 	x := mat.NewDense(nSamples, 2, nil)
// 	for i := 0; i < nSamples; i++ {
// 		x.Set(i, 0, 1)
// 		x.Set(i, 1, xData[i])
// 	}

// 	y := mat.NewVecDense(nSamples, yData)

// 	var xT mat.Dense
// 	xT.CloneFrom(x.T())

// 	// X^T * X
// 	var xTx mat.Dense
// 	xTx.Mul(&xT, x)

// 	// (X^T * X)^-1
// 	var xTxInv mat.Dense
// 	err := xTxInv.Inverse(&xTx)
// 	if err != nil {
// 		panic("Матрица не обратима")
// 	}

// 	// X^T * y
// 	var xTy mat.VecDense
// 	xTy.MulVec(&xT, y)

// 	// θ = (X^T X)^-1 X^T y
// 	var theta mat.VecDense
// 	theta.MulVec(&xTxInv, &xTy)

// 	return &theta
// }

// func main() {
// 	x := []float64{1, 2, 3, 4, 5}
// 	y := []float64{2, 4, 6, 8, 10} // y = 2x, ожидаем w=2, b=0

// 	theta := LinearRegression(x, y)

// 	for _, xi := range x {
// 		yPred := theta.AtVec(0) + theta.AtVec(1)*xi
// 		fmt.Printf("x=%.1f → y_hat=%.2f\n", xi, yPred)
// 	}

// 	fmt.Printf("Результат линейной регрессии:\n")
// 	fmt.Printf("b (intercept): %.2f\n", theta.AtVec(0))
// 	fmt.Printf("w (коэффициент): %.2f\n", theta.AtVec(1))
// }

// import (
// 	"fmt"

// 	"gonum.org/v1/gonum/mat"
// )

// func LinearRegression(Xdata [][]float64, Ydata []float64) *mat.VecDense {
// 	nSamples := len(Xdata)
// 	nFeatures := len(Xdata[0])

// 	X := mat.NewDense(nSamples, nFeatures+1, nil)

// 	// Формируем матрицу X с bias (единица в первом столбце)
// 	for i := 0; i < nSamples; i++ {
// 		X.Set(i, 0, 1) // Bias
// 		for j := 0; j < nFeatures; j++ {
// 			X.Set(i, j+1, Xdata[i][j])
// 		}
// 	}

// 	y := mat.NewVecDense(nSamples, Ydata)

// 	var xT mat.Dense
// 	xT.CloneFrom(X.T())

// 	var xTx mat.Dense
// 	xTx.Mul(&xT, X)

// 	var xTxInv mat.Dense
// 	err := xTxInv.Inverse(&xTx)
// 	if err != nil {
// 		panic("Матрица не обратима")
// 	}

// 	var xTy mat.VecDense
// 	xTy.MulVec(&xT, y)

// 	var theta mat.VecDense
// 	theta.MulVec(&xTxInv, &xTy)

// 	return &theta
// }

// func main() {
// 	X := [][]float64{
// 		{10, 5}, // 10 тыс. км, 5 лет
// 		{20, 3},
// 		{30, 2},
// 		{40, 1},
// 	}
// 	Y := []float64{10000, 12000, 14000, 16000} // Цена

// 	theta := LinearRegression(X, Y)

// 	var yPred []float64
// 	for _, xi := range X {
// 		pred := theta.AtVec(0) // bias
// 		for j := 0; j < len(xi); j++ {
// 			pred += theta.AtVec(j+1) * xi[j]
// 		}
// 		yPred = append(yPred, pred)
// 	}

// 	fmt.Printf("Коэффициенты (θ):\n")
// 	for i := 0; i < theta.Len(); i++ {
// 		fmt.Printf("θ[%d]: %.2f\n", i, theta.AtVec(i))
// 	}

// 	mae := MAE(Y, yPred)
// 	fmt.Printf("MAE: %.2f\n", mae)

// }

// // MAE считает среднюю абсолютную ошибку между y и y_pred
// func MAE(yTrue, yPred []float64) float64 {
// 	if len(yTrue) != len(yPred) {
// 		panic("длины y и y_pred не совпадают")
// 	}
// 	var sum float64
// 	for i := 0; i < len(yTrue); i++ {
// 		diff := yTrue[i] - yPred[i]
// 		if diff < 0 {
// 			diff = -diff
// 		}
// 		sum += diff
// 	}
// 	return sum / float64(len(yTrue))
// }
