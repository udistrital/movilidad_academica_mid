package models

type Persona struct {
	Id                int
	Nombre            string
	Apellido          string
	Identificacion    int
	TipoDocumento     TipoDocumento
	Telefono          int
	Direccion         string
	Pais              Pais
	NivelAcademico    NivelAcademico
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

type NivelAcademico struct {
	Id     int
	Nombre string
}
