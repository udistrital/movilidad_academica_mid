package controllers

import (
	"strconv"

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
	c.Mapping("GetRondasLiga", c.GetRondasLiga)

}

// GetMovilidad ...
// @Title Get Movilidad
// @Description get Movilidad
// @Success 200 {object} models.Movilidad
// @Failure 403 not found resource
// @router /GetMovilidad [get]
func (c *MovilidadController) GetMovilidad() {
	logs.Info("entro")
	logs.Info(beego.AppConfig.String("UrlMovilidadCrud") + "/movilidad")
	var movilidad []models.Movilidad

	if err := request.GetJson(beego.AppConfig.String("UrlMovilidadCrud")+"/movilidad", &movilidad); err == nil {
		logs.Info("tamaño", len(movilidad))
		c.Data["json"] = movilidad
	} else {
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

// GetNumeroRondasLiga ...
// @Title Get Numero Rondas liga
// @Description get numero rondas liga
// @Success 200 {object} int
// @Failure 403 not found resource
// @router /GetRondasLiga [get]
func (c *MovilidadController) GetRondasLiga() {
	equiposTotales, _ := strconv.Atoi(beego.AppConfig.String("CantidadEquipos"))
	if equiposTotales%2 == 0 {
		c.Data["json"] = (equiposTotales - 1) * 2
	} else {
		c.Data["json"] = equiposTotales * 2
	}
	c.ServeJSON()
}
func (c *MovilidadController) GetMovilidadId(idMovilidad string) (res []models.Movilidad) {
	logs.Info("entro")
	if err := request.GetJson(beego.AppConfig.String("UrlMovilidadCrud")+"movilidad/?query=Id"+idMovilidad, &res); err == nil {
		logs.Info("Retorna movilidadId")
	} else {
		logs.Info("error")
	}
	logs.Info("res", res)
	return
}
