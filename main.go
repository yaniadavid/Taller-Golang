package main

import "fmt"

func main() {
	var nombre string //"" | int 0 | float 0.00

	var edad int = 25

	salario := 5000.50

	var activo bool

	fmt.Println("Hola Mundo")

	fmt.Printf("Nombre: %s, Edad %d, Salario: %.2f, Activo: %t\n", nombre, edad, salario, activo)

	// Calculadora
	var a, b int = 12, 3
	ImprimirValores(9, 13)
	ImprimirValores(a, b)
	resultadoSuma := Sumar(a, b)
	fmt.Println(resultadoSuma)

	//Estudiante
	estudianteDominic := NuevoEstudiante("Dominic", 18, "Odontologia", 61)
	fmt.Println(estudianteDominic)
	fmt.Println(estudianteDominic.String())
}

func ImprimirValores(a int, b int) {
	fmt.Printf("valor a: %d, valor b: %d\n", a, b)
}

func Sumar(a int, b int) int {
	return a + b
}

// ------------------ Struct Estudiante
type Estudiante struct {
	Nombre       string
	Edad         int
	Carrera      string
	Calificacion float64
}

// Contructor
// Puntero *
func NuevoEstudiante(nombre string, edad int, carrera string, calificacion float64) *Estudiante {
	return &Estudiante{
		Nombre:       nombre,
		Edad:         edad,
		Carrera:      carrera,
		Calificacion: calificacion,
	}
}

func (e Estudiante) String() string {
	return fmt.Sprintf("%s - %d años. Carrera: %s, Calificación: %.2f", e.Nombre, e.Edad, e.Carrera, e.Calificacion)
}
