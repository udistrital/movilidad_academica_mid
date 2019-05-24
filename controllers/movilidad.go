package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
	//idStr := c.Ctx.Input.Param(":id")
	// logs.Info("entra")
	// logs.Info(beego.AppConfig.String("UrlLigaCrud") + "Movilidad")

	// c.ServeJSON()
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
