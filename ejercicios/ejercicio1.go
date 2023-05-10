package ejercicios

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

// Selector de actividades recursivo
// Reescribir la función que resuelve el problema de seleccionar actividades en forma recursiva
func SelectorRecursivo(actividades []Actividad) []Actividad {
	if len(actividades) == 0 {
		return []Actividad{}
	}
	primera := actividades[0]
	siguientes := []Actividad{}
	for i := 1; i < len(actividades); i++ {
		if actividades[i].Inicio >= primera.Fin {
			siguientes = append(siguientes, actividades[i])
		}
	}
	seleccionadas := SelectorRecursivo(siguientes)
	seleccionadas = append(seleccionadas, primera)
	return seleccionadas
}
