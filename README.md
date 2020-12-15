# movilidad_academica_mid

MID del proyecto movilidad academica.  
(Ejercicio introduccion 2019)


## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
# parametros de api
MOVILIDAD_ACADEMICA_MID_HTTP_PORT=[Puerto de exposición del API]
MOVILIDAD_ACADEMICA_MID_RUN_MODE=[Modo de ejecución del API]
# EndPoints externos
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con MOVILIDAD_ACADEMICA_MID_...


### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/movilidad_academica_mid

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/movilidad_academica_mid

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
movilidad_academica_mid_PORT=8080 movilidad_academica_mid_SOME_VARIABLE bee run
```

### Ejecución Dockerfile
```shell
# implementado para el sistema de integración continua (CI)
```

### Ejecución docker-compose
```shell
# 1. Obtener repositorio
git clone https://github.com/udistrital/movilidad_academica_mid.git
# 2. Ir a la carpeta del repositorio
cd $GOPATH/src/github.com/movilidad_academica_mid
# 3. Cambiar a la rama develop
git checkout develop
# 4. Crear red back_end
docker network create back_end
# 5. Ejecutar docker compose
docker-compose up --build
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# En Proceso
```

### Arquitectura del proyecto
[Arquitectura movilidad_academica_mid](arquitectura.png)

Otros repositorios relacionados:  
**[movilidad_academica_crud](https://github.com/udistrital/movilidad_academica_crud)**  
**[convenios_crud](https://github.com/udistrital/convenios_crud)**


### Puertos
El servidor se expone en el puerto: 127.0.0.1:8081  
Para ver la documentación de swagger: [127.0.0.1:8081/swagger/](http://127.0.0.1:8081/swagger/)

### EndPoints
|                |link de prueba                  |End Point|
|----------------|-------------------------------|------------------------|
| **Obtiene datos asociados a la información academica** | [GetAcademica](http://127.0.0.1:8081/v1/academica/GetAcademica) |`127.0.0.1:8081/v1/academica/GetAcademica`|
| **Obtiene datos asociados a los convenios** |[GetConvenio](http://127.0.0.1:8081/v1/convenio/GetConvenio)| `127.0.0.1:8081/v1/convenio/GetConvenio` |
| **Obtiene datos asociados a la movilidad estudiantil** |[GetMovilidad](http://127.0.0.1:8081/v1/movilidad/GetMovilidad)| `127.0.0.1:8081/v1/movilidad/GetMovilidad` |


## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| - | - | - |


## Licencia

This file is part of movilidad_academica_mid.

movilidad_academica_mid is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

movilidad_academica_mid is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with movilidad_academica_mid. If not, see https://www.gnu.org/licenses/.
