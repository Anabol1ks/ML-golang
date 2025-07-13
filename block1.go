package main

import (
	"log"
)

func dotProduct(a, b []float64) float64 {
	if len(a) != len(b) {
		log.Fatal()
	}
	var product float64
	for i := 0; i < len(a); i++ {
		product += a[i] * b[i]
	}
	return product
}
