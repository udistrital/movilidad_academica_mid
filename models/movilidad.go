package models

import "time"

type Movilidad struct {
	Id          int
	Presupuesto int
	FechaInicio time.Time
	FechaFinal  time.Time
	Persona     int
	Convenio    int
	IdMovilidad TipoMovilidad
	IdCategoria TipoCategoria
}
