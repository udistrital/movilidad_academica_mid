package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/movilidad_academica_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// MovilidadController for Movilidad
type MovilidadController struct {
	beego.Controller
}

// URLMapping ...
func (c *MovilidadController) URLMapping() {
	c.Mapping("GetMovilidad", c.GetMovilidad)

}

// GetMovilidad ...
// @Title Get Movilidad
// @Description get Movilidad
// @Success 200 {object} models.Movilidad
// @Failure 403 not found resource
// @router /GetMovilidad [get]
func (c *MovilidadController) GetMovilidad() {

	logs.Info("entro a Get Movilidad")
	queryFilter := ""
	if r := c.GetString("query"); r != "" {
		queryFilter = "?query=" + r

	}
	logs.Info(queryFilter)
	logs.Info(beego.AppConfig.String("UrlMovilidadCrud") + "/movilidad")
	var movilidad []models.Movilidad

	if err := request.GetJson(beego.AppConfig.String("UrlMovilidadCrud")+"/movilidad", &movilidad); err == nil {
		movilidadRespuesta := c.GetMovilidadByID(queryFilter)

		c.Data["json"] = movilidadRespuesta
	} else {
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

func (c *MovilidadController) GetMovilidadByID(IdMovilidad string) (res []models.Movilidad) {
	logs.Info("entro a recojer id")

	if err := request.GetJson(beego.AppConfig.String("UrlMovilidadCrud")+"/movilidad/"+IdMovilidad, &res); err == nil {
		logs.Info("Retorna movilidadId")
	} else {
		logs.Info("error Id")
	}
	logs.Info("res", res)

	return
}
