package main

import "fmt"

type retangulo struct {
	largura float64
	altura  float64
}

func calculararea(ret retangulo) float64 {
	return ret.largura * ret.altura
}

func main() {
	var ret retangulo
	fmt.Print("altura:")
	fmt.Scan(&ret.altura)
	fmt.Print("largura:")
	fmt.Scan(&ret.largura)
	a := calculararea(ret)
	fmt.Println("area:", a)
}
