package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// MovilidadController for Movilidad
type MovilidadController struct {
	beego.Controller
}

// URLMapping
func (c *MovilidadController) URLMapping() {
	c.Mapping("GetMovilidad", c.GetMovilidad)
	//c.Mapping("RegistrarMovilidad", c.RegistrarMovilidad)
}

// GetMovilidad ...
// @Title Get Movilidad
// @Description get movilidad
// @Success 200 {object} models.Movilidad
// @Failure 403 not found resource
// @router GetMovilidad/ [get]
func (c *MovilidadController) GetMovilidad() {
	logs.Info("entra")
	/*logs.Info(beego.AppConfig.String("UrlMovilidadCrud") + "movilidad")
	var movilidad []models.Movilidad

	if err := request.GetJson(beego.AppConfig.String("UrlMovilidadCrud")+"movilidad", &movilidad); err == nil {
		logs.Info("retorna: ", movilidad)
		movilidadRespuesta := make([]models.Movilidad, len(movilidad))

		c.Data["json"] = movilidadRespuesta

	} else {
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
	*/
}

/*
func (c *MovilidadController) RegistrarMovilidad() {
	try.This(func() {
		var body models.Movilidad
		var res map[string]interface{}
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err == nil {
			Urlcrud := beego.AppConfig.String("UrlMovilidadCrud") + "movilidad"
			logs.Info("URL ", Urlcrud)
			if err := request.SendJson(Urlcrud, "POST", &res, &body); err == nil {
				logs.Info("Post data: ", res)
				c.Data["json"] = res

			} else {
				panic(err.Error())
			}

		} else {
			c.Data["json"] = err.Error()
		}

	}).Catch(func(e try.E) {
		logs.Info("excep: ", e)
		c.Data["json"] = models.AlertError{Code: "E_0458", Body: e, Type: "error"}
	})

}
*/
