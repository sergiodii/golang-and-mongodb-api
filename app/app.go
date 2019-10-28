package teste

import "fmt"

type este struct {
	Name  string
	Idade int
}

func (t este) Nome() {
	fmt.Println(t.Name)
}
