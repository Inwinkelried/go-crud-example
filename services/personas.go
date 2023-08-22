package services

import "github.com/maxilovera/go-crud-example/dto"

var personas []*dto.Persona

//GET
func ObtenerPersonas() []*dto.Persona {
	return personas
}

//GET
func ObtenerPersonaPorId(id int) *dto.Persona {
	for _, persona := range personas {
		if persona.ID == id {
			return persona
		}
	}
	return nil
}

//POST
func CrearPersona(nuevaPersona dto.Persona) *dto.Persona {

	nuevaPersona.ID = len(personas) + 1

	personas = append(personas, &nuevaPersona)

	return &nuevaPersona
}

//PUT
func ActualizarPersona(personaActualizada dto.Persona) *dto.Persona {
	for _, persona := range personas {
		if persona.ID == personaActualizada.ID {
			persona.Nombre = personaActualizada.Nombre
			persona.Edad = personaActualizada.Edad

			return persona
		}
	}

	return nil
}

//DELETE
func EliminarPersona(id int) bool {
	for i, persona := range personas {
		if persona.ID == id {
			personas = append(personas[:i], personas[i+1:]...)
			return true
		}
	}
	return false
}
