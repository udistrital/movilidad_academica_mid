package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/movilidad_academica_mid/models"
)

// ConvenioController estructura para funciones y endpoints de convenio
type ConvenioController struct {
	beego.Controller
}

// URLMapping ...
func (c *ConvenioController) URLMapping() {
	c.Mapping("GetConvenio", c.GetConvenio)
}

// GetConvenio ...
// @Title Get Convenio
// @Description get Convenio
// @Success 200 {object} models.Posicion
// @Failure 403 not found resource
// @router /GetConvenio [get]
func (c *ConvenioController) GetConvenio() {
	logs.Info("entro al convenio")
	Convenio := []models.Convenio{
		{
			Id: 1,
			Organizacion: models.Organizacion{
				Id:     1,
				Nombre: "Universidad Distrital",
				Nit:    "222333",
			},
			TipoConvenio: models.TipoConvenio{
				Id:     1,
				Nombre: "primer tipo de convenio",
				Activo: true,
			},
			Pais: models.Pais{
				Id:     1,
				Nombre: "Colombia",
				Activo: true,
			},
		},
		{
			Id: 2,
			Organizacion: models.Organizacion{
				Id:     2,
				Nombre: "Universidad Javeriana",
				Nit:    "222333",
			},
			TipoConvenio: models.TipoConvenio{
				Id:     1,
				Nombre: "primer tipo de convenio",
				Activo: true,
			},
			Pais: models.Pais{
				Id:     1,
				Nombre: "Colombia",
				Activo: true,
			},
		},
	}

	logs.Info(Convenio)
	c.Data["json"] = Convenio
	c.ServeJSON()
}
