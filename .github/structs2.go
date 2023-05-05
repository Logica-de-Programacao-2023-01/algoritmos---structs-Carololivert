package main

import "fmt"

type livro struct {
	titulo string
	autor  string
	ano    int
}

func crialivro(titulo string, autor string, ano int) livro {
	return livro{titulo: titulo, autor: autor, ano: ano}
}
func main() {
	p := crialivro("Minions", "Caroline", 2023)
	fmt.Println(p)

}
