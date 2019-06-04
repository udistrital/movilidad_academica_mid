package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/movilidad_academica_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// AcademicaController estructura para funciones y endpoints de información academica
type AcademicaController struct {
	beego.Controller
}

// URLMapping ...
func (c *AcademicaController) URLMapping() {
	c.Mapping("GetAcademica", c.GetAcademica)
	c.Mapping("GetPersona", c.GetPersona)
}

// GetAcademica ...
// @Title Get Academica
// @Description get Academica
// @Success 200 {object} models.Posicion
// @Failure 403 not found resource
// @router /GetAcademica [get]
func (c *AcademicaController) GetAcademica() {
	logs.Info("entro al información academica")
	Academica := []models.Persona{
		{
			Id:             1,
			Nombre:         "Andres",
			Apellido:       "Perez",
			Identificacion: 123456,
			TipoDocumento: models.TipoDocumento{
				Id:     1,
				Nombre: "Cedula de Ciudadanía",
			},
			Telefono:  123456,
			Direccion: "calle 3 sur",
			Pais: models.Pais{
				Id:     1,
				Nombre: "Colombia",
				Activo: true,
			},
			NivelAcademico: models.NivelAcademico{
				Id:     1,
				Nombre: "Pregrado",
			},
			ProgramaAcademico: models.ProgramaAcademico{
				Id:       1,
				Nombre:   "Ingeniería de Sistemas",
				Facultad: "Facultad de Ingeniería",
			},
			TipoPersona: "Estudiante",
		},
		{
			Id:             2,
			Nombre:         "Felipe",
			Apellido:       "Rincón",
			Identificacion: 789456,
			TipoDocumento: models.TipoDocumento{
				Id:     1,
				Nombre: "Cedula de Ciudadanía",
			},
			Telefono:  789456,
			Direccion: "calle 36 sur",
			Pais: models.Pais{
				Id:     1,
				Nombre: "Colombia",
				Activo: true,
			},
			NivelAcademico: models.NivelAcademico{
				Id:     1,
				Nombre: "Postgrado",
			},
			ProgramaAcademico: models.ProgramaAcademico{
				Id:       1,
				Nombre:   "Maestría en Electronica",
				Facultad: "Facultad de Ingeniería",
			},
			TipoPersona: "Docente",
		},
		{
			Id:             3,
			Nombre:         "Rosy",
			Apellido:       "Sanchez",
			Identificacion: 147258,
			TipoDocumento: models.TipoDocumento{
				Id:     1,
				Nombre: "Cedula de Extranjería",
			},
			Telefono:  147258,
			Direccion: "calle 3 sur",
			Pais: models.Pais{
				Id:     1,
				Nombre: "Mexico",
				Activo: true,
			},
			NivelAcademico: models.NivelAcademico{
				Id:     1,
				Nombre: "Pregrado",
			},
			ProgramaAcademico: models.ProgramaAcademico{
				Id:       1,
				Nombre:   "Ingeniería Electrica",
				Facultad: "Facultad de Ingeniería",
			},
			TipoPersona: "Estudiante",
		},
		{
			Id:             4,
			Nombre:         "Mari",
			Apellido:       "León",
			Identificacion: 963852,
			TipoDocumento: models.TipoDocumento{
				Id:     1,
				Nombre: "Cedula de Extranjería",
			},
			Telefono:  963852,
			Direccion: "calle 3 sur",
			Pais: models.Pais{
				Id:     1,
				Nombre: "Argentina",
				Activo: true,
			},
			NivelAcademico: models.NivelAcademico{
				Id:     1,
				Nombre: "Pregrado",
			},
			ProgramaAcademico: models.ProgramaAcademico{
				Id:       1,
				Nombre:   "Ingeniería Carastral",
				Facultad: "Facultad de Ingeniería",
			},
			TipoPersona: "Docente",
		},
	}

	logs.Info(Academica)
	c.Data["json"] = Academica
	c.ServeJSON()

}

// GetPersona ...
// @Title Get One
// @@Description get Persona by numIdent
// @Param	numIdent	query	string	false	"buscar persona por numero de identificacion"
// @Success 200 {object} interface{}
// @Failure 403 :numIdent is empty
// @router /GetPersona/ [get]
func (c *AcademicaController) GetPersona() {
	// logs.Info("ha entrado al get persona")
	numeroIdentificacion := c.GetString("numIdent")
	// logs.Info(numeroIdentificacion)

	var resultado []map[string]interface{}
	// var resultado2 []models.PersonaCampus
	// var newpersona map[string]interface{}
	// newpersona = make(map[string]interface{})
	// var NumeroIdentificacion map[string]interface{}
	// var ElEnte map[string]interface{}

	var errNumero error
	// var errPersona error

	errNumero = request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/?query=NumeroIdentificacion:"+numeroIdentificacion, &resultado)
	if resultado != nil {
		logs.Info(resultado)
		logs.Info(errNumero)
		NumeroIdentificacion := resultado[0]["NumeroIdentificacion"]
		logs.Info("numero identificacion = ", NumeroIdentificacion)
		var Ente map[string]interface{}
		Ente = make(map[string]interface{})
		Ente = map[string]interface{}{"Id": resultado[0]["Ente"].(map[string]interface{})["Id"]}
		var EnteID string
		EnteID = fmt.Sprintf("%.f", Ente["Id"].(float64))
		// logs.Info(Ente["Id"])

		clienteHttp := &http.Client{}
		url := "http://" + beego.AppConfig.String("CampusMid") + "/persona/ConsultaPersona/?id=" + EnteID
		peticion, err := http.NewRequest("GET", url, nil)
		if err != nil {
			// Maneja el error de acuerdo a tu situación
			log.Fatalf("Error creando petición: %v", err)

		}

		peticion.Header.Add("Content-Type", "application/json")
		respuesta, err := clienteHttp.Do(peticion)
		if err != nil {
			// Maneja el error de acuerdo a tu situación
			log.Fatalf("Error haciendo petición: %v", err)
		}
		// No olvides cerrar el cuerpo al terminar
		defer respuesta.Body.Close()

		cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
		if err != nil {
			log.Fatalf("Error leyendo respuesta: %v", err)
		}
		personaMid := new(models.PersonaCampus)
		json.Unmarshal(cuerpoRespuesta, &personaMid)
		// logs.Info(t)

		// respuestaString := string(cuerpoRespuesta)
		// Aquí puedes decodificar la respuesta si es un JSON, o convertirla a cadena
		// log.Printf("Cuerpo de respuesta del servidor: '%s'", respuestaString)

		c.Data["json"] = personaMid

	}
	c.ServeJSON()

}
