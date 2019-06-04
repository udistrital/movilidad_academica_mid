package models

type PersonaCampus struct {
	Ente               int
	EstadoCivil        EstadoCivil
	FechaNacimiento    string
	Foto               int
	Genero             Genero
	Id                 int
	NumeroDocumento    string
	PrimerApellido     string
	PrimerNombre       string
	SegundoApellido    string
	SegundoNombre      string
	SoporteDocumento   int
	TipoIdentificacion TipoIdentificacion
	Usuario            string
}

type EstadoCivil struct {
	Activo            bool
	CodigoAbreviacion string
	Descripcion       string
	Id                int
	Nombre            string
	NumeroOrden       int
}

type Genero struct {
	Activo            bool
	CodigoAbreviacion string
	Descripcion       string
	Id                int
	Nombre            string
	NumeroOrden       int
}

type TipoIdentificacion struct {
	Activo            bool
	CodigoAbreviacion string
	Descripcion       string
	Id                int
	Nombre            string
	NumeroOrden       int
}
