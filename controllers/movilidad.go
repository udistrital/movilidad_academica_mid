package controllers

import (
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
