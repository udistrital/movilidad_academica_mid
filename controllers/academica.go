package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/movilidad_academica_mid/models"
)

// AcademicaController estructura para funciones y endpoints de información academica
type AcademicaController struct {
	beego.Controller
}

// URLMapping ...
func (a *AcademicaController) URLMapping() {
	a.Mapping("GetAcademica", a.GetAcademica)
}

// GetAcademica ...
// @Title Get Academica
// @Description get Academica
// @Success 200 {object} models.Posicion
// @Failure 403 not found resource
// @router /GetAcademica [get]
func (a *AcademicaController) GetAcademica() {
	logs.Info("entro al información academica")
	Academica := []models.Persona{
		{
			Id:             1,
			Nombre:         "Andres",
			Apellido:       "Perez",
			Identificacion: 123456,
			TipoDocumento: models.TipoDocumento{
				Id:     1,
				Nombre: "Cedula de Ciudadanía",
			},
			Telefono:  123456,
			Direccion: "calle 3 sur",
			Pais: models.Pais{
				Id:     1,
				Nombre: "Colombia",
				Activo: true,
			},
			NivelAcademico: models.NivelAcademico{
				Id:     1,
				Nombre: "Pregrado",
			},
			ProgramaAcademico: models.ProgramaAcademico{
				Id:       1,
				Nombre:   "Ingeniería de Sistemas",
				Facultad: "Facultad de Ingeniería",
			},
			TipoPersona: "Estudiante",
		},
		{
			Id:             2,
			Nombre:         "Felipe",
			Apellido:       "Rincón",
			Identificacion: 789456,
			TipoDocumento: models.TipoDocumento{
				Id:     1,
				Nombre: "Cedula de Ciudadanía",
			},
			Telefono:  789456,
			Direccion: "calle 36 sur",
			Pais: models.Pais{
				Id:     1,
				Nombre: "Colombia",
				Activo: true,
			},
			NivelAcademico: models.NivelAcademico{
				Id:     1,
				Nombre: "Postgrado",
			},
			ProgramaAcademico: models.ProgramaAcademico{
				Id:       1,
				Nombre:   "Maestría en Electronica",
				Facultad: "Facultad de Ingeniería",
			},
			TipoPersona: "Docente",
		},
		{
			Id:             3,
			Nombre:         "Rosy",
			Apellido:       "Sanchez",
			Identificacion: 147258,
			TipoDocumento: models.TipoDocumento{
				Id:     1,
				Nombre: "Cedula de Extranjería",
			},
			Telefono:  147258,
			Direccion: "calle 3 sur",
			Pais: models.Pais{
				Id:     1,
				Nombre: "Mexico",
				Activo: true,
			},
			NivelAcademico: models.NivelAcademico{
				Id:     1,
				Nombre: "Pregrado",
			},
			ProgramaAcademico: models.ProgramaAcademico{
				Id:       1,
				Nombre:   "Ingeniería Electrica",
				Facultad: "Facultad de Ingeniería",
			},
			TipoPersona: "Estudiante",
		},
		{
			Id:             4,
			Nombre:         "Mari",
			Apellido:       "León",
			Identificacion: 963852,
			TipoDocumento: models.TipoDocumento{
				Id:     1,
				Nombre: "Cedula de Extranjería",
			},
			Telefono:  963852,
			Direccion: "calle 3 sur",
			Pais: models.Pais{
				Id:     1,
				Nombre: "Argentina",
				Activo: true,
			},
			NivelAcademico: models.NivelAcademico{
				Id:     1,
				Nombre: "Pregrado",
			},
			ProgramaAcademico: models.ProgramaAcademico{
				Id:       1,
				Nombre:   "Ingeniería Carastral",
				Facultad: "Facultad de Ingeniería",
			},
			TipoPersona: "Docente",
		},
	}

	logs.Info(Academica)
	a.Data["json"] = Academica
	a.ServeJSON()

}
