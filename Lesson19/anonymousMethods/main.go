package main

import "fmt"

type Factory struct {
	city string
	name string
}

type Robot struct {
	fullName string
	age      int
	Factory
}

func (f *Factory) FullFactoryInfo() {
	fmt.Printf("Factory name: %s and city: %s\n", f.name, f.city)
}

func main() {

	r := Robot{
		fullName: "Robot2000",
		age:      1500,
		Factory: Factory{
			city: "Vulkan",
			name: "UralMash",
		},
	}

	r.FullFactoryInfo()

}
