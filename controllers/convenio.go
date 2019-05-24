package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
}
