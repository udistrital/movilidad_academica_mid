package models

type Persona struct {
	Id                int
	Nombre            string
	Apellido          string
	Identificacion    int
	TipoDocumento     TipoDocumento
	ProgramaAcademico ProgramaAcademico
	TipoPersona       string
}

type ProgramaAcademico struct {
	Id       int
	Nombre   string
	Facultad string
}

type TipoDocumento struct {
	Id     int
	Nombre string
}
