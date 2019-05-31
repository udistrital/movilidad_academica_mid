package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {


    beego.GlobalControllerRouter["github.com/udistrital/movilidad_academica_mid/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movilidad_academica_mid/controllers:ConvenioController"],
        beego.ControllerComments{
            Method: "GetConvenio",
            Router: `/GetConvenio`,

    beego.GlobalControllerRouter["github.com/udistrital/movilidad_academica_mid/controllers:AcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movilidad_academica_mid/controllers:AcademicaController"],
        beego.ControllerComments{
            Method: "GetAcademica",
            Router: `/GetAcademica`,

            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})


    beego.GlobalControllerRouter["github.com/udistrital/movilidad_academica_mid/controllers:MovilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movilidad_academica_mid/controllers:MovilidadController"],
        beego.ControllerComments{
            Method: "GetMovilidad",
            Router: `/GetMovilidad`,

    beego.GlobalControllerRouter["github.com/udistrital/movilidad_academica_mid/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movilidad_academica_mid/controllers:ConvenioController"],
        beego.ControllerComments{
            Method: "GetConvenio",
            Router: `/GetConvenio`,

            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movilidad_academica_mid/controllers:MovilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movilidad_academica_mid/controllers:MovilidadController"],
        beego.ControllerComments{

            Method: "RegistrarMovilidad",
            Router: `/RegistrarMovilidad`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(
				param.New("movidad"),
			),

            Method: "GetMovilidad",
            Router: `/GetMovilidad`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),

            Filters: nil,
            Params: nil})

}
