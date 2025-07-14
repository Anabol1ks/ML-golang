package main

import "log"

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

func multiplyMatrices(a, b [][]float64) [][]float64 {
	rowsA := len(a)
	rowsB := len(b)
	colsA := len(a[0])
	colsB := len(b[0])

	if colsA != rowsB {
		log.Fatal("Несоответствие размеров матриц")
	}

	result := make([][]float64, rowsA)
	for i := range result {
		result[i] = make([]float64, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			var sum float64
			for k := 0; k < colsA; k++ {
				sum += a[i][k] * b[k][j]
			}
			result[i][j] = sum
		}
	}
	return result
}
func main() {

}
