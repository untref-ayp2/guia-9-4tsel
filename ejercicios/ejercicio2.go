package ejercicios

import (
	"sort"
	"fmt"
)

type Tarea struct {
	Nombre string
	Tiempo int
}

func Ejercicio2(tareas []Tarea) {
	
	// Ordenar las tareas en orden ascendente por su tiempo de ejecución
	sort.Slice(tareas, func(i, j int) bool {
		return tareas[i].Tiempo < tareas[j].Tiempo
	})

	// Ejecutar las tareas en ese orden
	tiempo := 0
	for _, tarea := range tareas {
		inicio := tiempo
		tiempo += tarea.Tiempo
		fin := tiempo
		fmt.Printf("%s: [%d, %d)\n", tarea.Nombre, inicio, fin)
	}

	// Calcular el tiempo promedio de finalización
	tiempoPromedio := float64(tiempo) / float64(len(tareas))
	fmt.Printf("Tiempo promedio de finalización: %f\n", tiempoPromedio)
}
