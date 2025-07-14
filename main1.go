// package main

// import (
// 	"fmt"

// 	"gonum.org/v1/gonum/mat"
// )

// func Preprocess(Xraw [][]float64) *mat.Dense {
// 	nSamples := len(Xraw)
// 	nFeatures := len(Xraw[0])

// 	// Вычисляем min и max для каждого признака
// 	mins := make([]float64, nFeatures)
// 	maxs := make([]float64, nFeatures)

// 	for j := 0; j < nFeatures; j++ {
// 		mins[j] = Xraw[0][j]
// 		maxs[j] = Xraw[0][j]
// 	}

// 	for i := 0; i < nSamples; i++ {
// 		for j := 0; j < nFeatures; j++ {
// 			if Xraw[i][j] < mins[j] {
// 				mins[j] = Xraw[i][j]
// 			}
// 			if Xraw[i][j] > maxs[j] {
// 				maxs[j] = Xraw[i][j]
// 			}
// 		}
// 	}

// 	X := mat.NewDense(nSamples, nFeatures+1, nil)

// 	for i := 0; i < nSamples; i++ {
// 		X.Set(i, 0, 1) // Bias term
// 		for j := 0; j < nFeatures; j++ {
// 			if maxs[j] == mins[j] {
// 				X.Set(i, j+1, 0)
// 			} else {
// 				xNorm := (Xraw[i][j] - mins[j]) / (maxs[j] - mins[j])
// 				X.Set(i, j+1, xNorm)
// 			}
// 		}
// 	}
// 	return X
// }

// func Train(X *mat.Dense, Y *mat.VecDense) *mat.VecDense {
// 	var XT mat.Dense
// 	XT.Mul(X.T(), X)

// 	var XTInv mat.Dense
// 	err := XTInv.Inverse(&XT)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var XTY mat.VecDense
// 	XTY.MulVec(X.T(), Y)

// 	var theta mat.VecDense
// 	theta.MulVec(&XTInv, &XTY)

// 	return &theta
// }

// func Predict(X *mat.Dense, theta *mat.VecDense) *mat.VecDense {
// 	// Используем коэффициенты θ для предсказаний
// 	var Ypred mat.VecDense
// 	Ypred.MulVec(X, theta)
// 	return &Ypred
// }

// func MAE(YTrue, YPred *mat.VecDense) float64 {
// 	n := YTrue.Len()
// 	var sumAbs float64
// 	for i := 0; i < n; i++ {
// 		diff := YTrue.AtVec(i) - YPred.AtVec(i)
// 		if diff < 0 {
// 			diff = -diff
// 		}
// 		sumAbs += diff
// 	}
// 	return sumAbs / float64(n)
// }

// func main() {
// 	Xraw := [][]float64{
// 		{50000, 5},
// 		{70000, 3},
// 		{30000, 7},
// 		{100000, 2},
// 	}
// 	Ydata := []float64{8500, 9500, 7200, 10000}

// 	X := Preprocess(Xraw)
// 	Y := mat.NewVecDense(len(Ydata), Ydata)
// 	fmt.Printf("Обработанные данные:\n%v\n", mat.Formatted(X, mat.Prefix(""), mat.Excerpt(0)))
// 	theta := Train(X, Y)
// 	fmt.Printf("Коэффициенты модели:\n%v\n", mat.Formatted(theta, mat.Prefix(""), mat.Excerpt(0)))
// 	Ypred := Predict(X, theta)

// 	fmt.Println("MAE:", MAE(Y, Ypred))
// }

// package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"image/color"
// 	"log"
// 	"os"
// 	"strconv"

// 	"gonum.org/v1/plot"
// 	"gonum.org/v1/plot/plotter"
// 	"gonum.org/v1/plot/vg"
// )

// func main() {
// 	// 1️⃣ Чтение CSV файла
// 	file, err := os.Open("Linear Regression - Sheet1.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var points plotter.XYs
// 	for i, record := range records {
// 		if i == 0 {
// 			continue // пропуск заголовков
// 		}
// 		if len(record) < 2 {
// 			continue // убедись, что в строке достаточно колонок
// 		}
// 		x, err1 := strconv.ParseFloat(record[0], 64)
// 		y, err2 := strconv.ParseFloat(record[1], 64)
// 		fmt.Println("x:", x, "y:", y)

// 		if err1 != nil || err2 != nil {
// 			continue
// 		}
// 		points = append(points, plotter.XY{X: x, Y: y})
// 	}

// 	// 2️⃣ Создание scatter plot
// 	p := plot.New()
// 	p.Title.Text = "Scatter Plot of x vs y"
// 	p.X.Label.Text = "x"
// 	p.Y.Label.Text = "y"

// 	scatter, err := plotter.NewScatter(points)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	scatter.GlyphStyle.Radius = vg.Points(2)
// 	scatter.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // красный

// 	p.Add(scatter)

// 	// 3️⃣ Сохранение в PNG
// 	if err := p.Save(6*vg.Inch, 4*vg.Inch, "scatter.png"); err != nil {
// 		log.Fatal(err)
// 	}
// }

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

func Preprocess(Xraw [][]float64) *mat.Dense {
	nSamples := len(Xraw)
	nFeatures := len(Xraw[0])

	// Вычисляем min и max для каждого признака
	mins := make([]float64, nFeatures)
	maxs := make([]float64, nFeatures)

	for j := 0; j < nFeatures; j++ {
		mins[j] = Xraw[0][j]
		maxs[j] = Xraw[0][j]
	}

	for i := 0; i < nSamples; i++ {
		for j := 0; j < nFeatures; j++ {
			if Xraw[i][j] < mins[j] {
				mins[j] = Xraw[i][j]
			}
			if Xraw[i][j] > maxs[j] {
				maxs[j] = Xraw[i][j]
			}
		}
	}

	X := mat.NewDense(nSamples, nFeatures+1, nil)

	for i := 0; i < nSamples; i++ {
		X.Set(i, 0, 1) // Bias term
		for j := 0; j < nFeatures; j++ {
			if maxs[j] == mins[j] {
				X.Set(i, j+1, 0)
			} else {
				xNorm := (Xraw[i][j] - mins[j]) / (maxs[j] - mins[j])
				X.Set(i, j+1, xNorm)
			}
		}
	}
	return X
}

func Train(X *mat.Dense, Y *mat.VecDense) *mat.VecDense {
	var XT mat.Dense
	XT.Mul(X.T(), X)

	var XTInv mat.Dense
	err := XTInv.Inverse(&XT)
	if err != nil {
		panic(err)
	}

	var XTY mat.VecDense
	XTY.MulVec(X.T(), Y)

	var theta mat.VecDense
	theta.MulVec(&XTInv, &XTY)

	return &theta
}

func Predict(X *mat.Dense, theta *mat.VecDense) *mat.VecDense {
	// Используем коэффициенты θ для предсказаний
	var Ypred mat.VecDense
	Ypred.MulVec(X, theta)
	return &Ypred
}

func MAE(YTrue, YPred *mat.VecDense) float64 {
	n := YTrue.Len()
	var sumAbs float64
	for i := 0; i < n; i++ {
		diff := YTrue.AtVec(i) - YPred.AtVec(i)
		if diff < 0 {
			diff = -diff
		}
		sumAbs += diff
	}
	return sumAbs / float64(n)
}

func main() {
	Xraw, Ydata, err := LoadCSV("data1.csv")
	if err != nil {
		panic(err)
	}

	X := Preprocess(Xraw)
	Y := mat.NewVecDense(len(Ydata), Ydata)
	fmt.Printf("Обработанные данные:\n%v\n", mat.Formatted(X, mat.Prefix(""), mat.Excerpt(0)))
	theta := Train(X, Y)
	fmt.Printf("Коэффициенты модели:\n%v\n", mat.Formatted(theta, mat.Prefix(""), mat.Excerpt(0)))
	Ypred := Predict(X, theta)

	fmt.Println("MAE:", MAE(Y, Ypred))
}

func LoadCSV(path string) ([][]float64, []float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()

	if err != nil {
		return nil, nil, err
	}

	rows = rows[1:]

	var X [][]float64
	var Y []float64

	for _, row := range rows {
		if len(row) != 3 {
			continue
		}

		area, _ := strconv.ParseFloat(row[0], 64)
		rooms, _ := strconv.ParseFloat(row[1], 64)
		price, _ := strconv.ParseFloat(row[2], 64)

		X = append(X, []float64{area, rooms})
		Y = append(Y, price)
	}

	return X, Y, nil
}
