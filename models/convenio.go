package models

type Convenio struct {
	Id           int
	Organizacion *Organizacion
	TipoConvenio *TipoConvenio
	Pais         *Pais
}

type Organizacion struct {
	Id     int
	Nombre int
	Nit    string
}

type Pais struct {
	Id     int
	Nombre string
	Activo bool
}

type TipoConvenio struct {
	Id     int
	Nombre string
	Activo bool
}
