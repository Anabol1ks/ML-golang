```go
import (
	"fmt"
	"log"
	"gonum.org/v1/gonum/mat"
)
```

### 🔹 Урок 1: Линейная алгебра — векторы и скалярное произведение

#### 📘 Теория:

Вектор — это упорядоченный набор чисел, например:
`A = [1, 2, 3]` — вектор в 3D пространстве.

Скалярное произведение (dot product) двух векторов:

```text
A • B = A₁*B₁ + A₂*B₂ + A₃*B₃
```

Если угол между ними острый → скалярное произведение положительное.
Если прямой угол → ноль.
Если тупой угол → отрицательное.


```go
import (
	"fmt"
	"log"
)

func dotProduct(a, b []float64) float64 {
	if len(a) != len(b) {
		log.Fatal("Векторы должны быть равны")
	}
	var product float64
	for i := 0; i < len(a); i++ {
		product += a[i] * b[i]
	}
	return product
}
```


```go
vec1 := []float64{1, 2, 3}
vec2 := []float64{4, 5, 6}
fmt.Println("Скалярное произведение:", dotProduct(vec1, vec2))
```

### 🔹 Урок 2: Матрицы и умножение матриц

---

#### 📘 Теория:

**Матрица** — это таблица чисел, например:

```
A =  | 1  2 |
     | 3  4 |

B =  | 5  6 |
     | 7  8 |
```

Чтобы перемножить две матрицы `A` и `B`, нужно:

* Количество **столбцов в A** должно = количеству **строк в B**.
* Результат — новая матрица размером `(строки A) x (столбцы B)`.

Каждый элемент результата вычисляется как **скалярное произведение** строки `A` и столбца `B`.



```go
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

a := [][]float64{
	{1, 2},
	{3, 4},
	{12, 9},
	{11, 13},
}

b := [][]float64{
	{5, 6, 12, 9},
	{7, 8, 11, 10},
}

result := multiplyMatrices(a, b)

fmt.Println("Результат умножения:")
for _, row := range result {
	fmt.Println(row)
}
```

### 🔹 Урок 3: Работа с `gonum/mat` — базовые векторы и матрицы

---

#### 📘 Теория:

[`gonum/mat`](https://pkg.go.dev/gonum.org/v1/gonum/mat) — это основной пакет для линейной алгебры в Go.

В нём есть:

* `mat.NewVecDense` — создание вектора.
* `mat.NewDense` — создание матрицы.
* Методы вроде `Dot`, `Mul`, `T()` — для операций.





```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewVecDense(3, []float64{1, 2, 3})
	b := mat.NewVecDense(3, []float64{4, 5, 6})
	dot := mat.Dot(a, b)
	fmt.Println("Скалярное произведение:", dot)

	m1 := mat.NewDense(2, 3, []float64{
		1, 2, 3,
		4, 5, 6,
	})
	m2 := mat.NewDense(3, 2, []float64{
		7, 8,
		9, 10,
		11, 12,
	})

	var result mat.Dense
	result.Mul(m1, m2)

	fmt.Println("Умножение матриц m1 * m2:")
	fmt.Printf("%v\n", mat.Formatted(&result, mat.Prefix(""), mat.Excerpt(0)))
}

main()

```
---

### 🔹 Урок 4: Статистика и предобработка с `gonum/stat` 

---

#### 📘 Теория: основные метрики 

Для машинного обучения нам нужны:

| Метрика         | Что означает                   |
| --------------- | ------------------------------ |
| Среднее (mean)  | Усреднённое значение данных    |
| Дисперсия (var) | Насколько данные разбросаны    |
| Ст. отклонение  | Корень из дисперсии            |
| Ковариация      | Зависимость между двумя фичами |
| Корреляция      | Нормированная ковариация       |

Это всё — базовые кирпичики регрессии и кластеризации.

```go

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
)

func main() {
	data := []float64{5, 7, 8, 9, 10, 6}

	mean := stat.Mean(data, nil)
	variance := stat.Variance(data, nil)
	sdt := stat.StdDev(data, nil)

	fmt.Println("Данные:", data)
	fmt.Println("Среднее:", mean)
	fmt.Println("Дисперсия:", variance)
	fmt.Println("Среднеквадратическое отклонение:", sdt)

	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10}

	cov := stat.Covariance(x, y, nil)
	corr := stat.Correlation(x, y, nil)

	fmt.Printf("\nКовариация(x, y): %.2f\n", cov)
	fmt.Printf("Корреляция(x, y): %.2f\n", corr)
}
```
---

### 🔹 Урок 5: Линейная регрессия — теория и реализация с нуля

---

#### 📘 Теория:

**Линейная регрессия** ищет **прямую**, которая **лучше всего описывает зависимость** между входом `x` и выходом `y`:

$$
y = w \cdot x + b
$$

Где:

* `w` — коэффициент наклона (вес),
* `b` — смещение (bias, интерсепт).

---

### 📐 Цель: найти такие `w` и `b`, чтобы минимизировать **сумму квадратов ошибок**:

$$
\text{Loss} = \frac{1}{n} \sum_{i=1}^{n} (y_i - (w \cdot x_i + b))^2
$$

Мы можем вычислить параметры **аналитически** (без градиентного спуска), используя **матричную формулу**:

$$
\theta = (X^T X)^{-1} X^T y
$$

Где:

* `X` — матрица входов (добавим столбец единиц для `b`),
* `y` — вектор целевых значений,
* `θ` — вектор коэффициентов (`[b, w1, w2, ...]`).

---

### 💻 Пример кода: линейная регрессия вручную

```go
import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func LinearRegression(xData, yData []float64) *mat.VecDense {
	nSamples := len(xData)

	x := mat.NewDense(nSamples, 2, nil)
	for i := 0; i < nSamples; i++ {
		x.Set(i, 0, 1)
		x.Set(i, 1, xData[i])
	}

	y := mat.NewVecDense(nSamples, yData)

	var xT mat.Dense
	xT.CloneFrom(x.T())

	// X^T * X
	var xTx mat.Dense
	xTx.Mul(&xT, x)

	// (X^T * X)^-1
	var xTxInv mat.Dense
	err := xTxInv.Inverse(&xTx)
	if err != nil {
		panic("Матрица не обратима")
	}

	// X^T * y
	var xTy mat.VecDense
	xTy.MulVec(&xT, y)

	// θ = (X^T X)^-1 X^T y
	var theta mat.VecDense
	theta.MulVec(&xTxInv, &xTy)

	return &theta
}

func main() {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10} // y = 2x, ожидаем w=2, b=0

	theta := LinearRegression(x, y)

	for _, xi := range x {
		yPred := theta.AtVec(0) + theta.AtVec(1)*xi
		fmt.Printf("x=%.1f → y_hat=%.2f\n", xi, yPred)
	}

	fmt.Printf("Результат линейной регрессии:\n")
	fmt.Printf("b (intercept): %.2f\n", theta.AtVec(0))
	fmt.Printf("w (коэффициент): %.2f\n", theta.AtVec(1))
}

```
---

## 🔹 Урок 6: Многомерная линейная регрессия + Мини-проект

---

### 📘 Теория:

До этого ты использовал только один признак `x`. А если у нас есть несколько признаков, например:

```
y = b + w1*x1 + w2*x2 + ... + wn*xn
```

Мы можем использовать ту же формулу:

$$
\theta = (X^T X)^{-1} X^T y
$$

где `X` — теперь матрица с **n признаками + 1 столбец 1 для смещения**.

---

### 💻 Пример: многомерная линейная регрессия

Допустим, мы хотим предсказывать цену машины по:

* пробегу,
* возрасту.

```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func LinearRegression(Xdata [][]float64, Ydata []float64) *mat.VecDense {
	nSamples := len(Xdata)
	nFeatures := len(Xdata[0])

	X := mat.NewDense(nSamples, nFeatures+1, nil)

	// Формируем матрицу X с bias (единица в первом столбце)
	for i := 0; i < nSamples; i++ {
		X.Set(i, 0, 1) // Bias
		for j := 0; j < nFeatures; j++ {
			X.Set(i, j+1, Xdata[i][j])
		}
	}

	y := mat.NewVecDense(nSamples, Ydata)

	var xT mat.Dense
	xT.CloneFrom(X.T())

	var xTx mat.Dense
	xTx.Mul(&xT, X)

	var xTxInv mat.Dense
	err := xTxInv.Inverse(&xTx)
	if err != nil {
		panic("Матрица не обратима")
	}

	var xTy mat.VecDense
	xTy.MulVec(&xT, y)

	var theta mat.VecDense
	theta.MulVec(&xTxInv, &xTy)

	return &theta
}

func main() {
	X := [][]float64{
		{10, 5}, // 10 тыс. км, 5 лет
		{20, 3},
		{30, 2},
		{40, 1},
	}
	Y := []float64{10000, 12000, 14000, 16000} // Цена

	theta := LinearRegression(X, Y)

	fmt.Printf("Коэффициенты (θ):\n")
	for i := 0; i < theta.Len(); i++ {
		fmt.Printf("θ[%d]: %.2f\n", i, theta.AtVec(i))
	}
}
```
