package main

import "fmt"

type Engine struct {
	Displacement int
	Cylinders    int
	FuelType     string
	Inline       bool
	Horsepower   int
	Torque       int
}

type Model struct {
	ModelName string
	Variant   string
	Engine    Engine
	Color     string
	Doors     int
}

type Make struct {
	MakeName       string
	Model          Model
	ProductionYear int
	Price          float32
}

func main() {
	// embedded structs
	// aka structs within structs

	passat := Make{
		MakeName: "Volkswagen",
		Model: Model{
			ModelName: "Passat",
			Variant:   "Break",
			Engine: Engine{
				Displacement: 2153,
				Cylinders:    6,
				FuelType:     "diesel",
				Inline:       false,
				Horsepower:   180,
				Torque:       410,
			},
			Color: "black",
			Doors: 5,
		},
		ProductionYear: 2003,
		Price:          3100,
	}

	bmw := Make{
		MakeName: "BMW",
		Model: Model{
			ModelName: "X5",
			Variant:   "E53",
			Engine: Engine{
				Displacement: 3000,
				Cylinders:    6,
				FuelType:     "diesel",
				Inline:       true,
				Horsepower:   218,
				Torque:       550,
			},
			Color: "black",
			Doors: 5,
		},
		ProductionYear: 2005,
		Price:          10500,
	}

	fmt.Println("First car:")

	fmt.Printf("This is a %d %s with %d horsepower and a %d %s engine\n", passat.ProductionYear, passat.MakeName, passat.Model.Engine.Horsepower, passat.Model.Engine.Displacement, passat.Model.Engine.FuelType)

	fmt.Println("Second car: ")

	fmt.Printf("This is a %d %s with %d horsepower and a %d %s engine\n", bmw.ProductionYear, bmw.MakeName, bmw.Model.Engine.Horsepower, bmw.Model.Engine.Displacement, bmw.Model.Engine.FuelType)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
