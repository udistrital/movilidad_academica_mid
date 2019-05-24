package models

type TipoMovilidad struct {
	Id                int
	Nombre            string
	Descripcion       string
	CodigoAbreviacion string
	Activo            bool
	NumeroOrden       int
	Entrante          bool
}
