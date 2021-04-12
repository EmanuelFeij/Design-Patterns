package main

import (
	"fmt"

	"github.com/EmanuelFeij/Design-Paterns/builder"
)

func main() {
	builder := &builder.Builder{}
	employee := builder.
		Name("Emanuel").
		Role("CEO").
		Build()

	fmt.Println(employee)
}
