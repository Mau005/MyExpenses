package configuration

const (
	ERROR_SQLITE = "No se ha podido conectar la BD"
	ERROR_EMPTY  = "Esta informacion no puede venir sin datos"
)

const (
	EXPIRATION_TOKEN = 24 //HOURS
	NAME_SESSION     = "Authorization"
	DELETE_USER      = "Se ha eliminado el usuario"
	DELETE_CATEGORY  = "Se ha eliminado la categoria"
)

const (
	ERROR_EMPTY_FIELD       = "Se debe ingresar datos"
	ERROR_SERVICE_USER      = "Service error in user"
	ERROR_SERVICE_CATEGORY  = "Service error in Category"
	ERROR_PRIVILEGES_GEN    = "No Tienes los privilegios para acceder"
	ERROR_UPDATE_USER_EMAIL = "No se puede encontrar el usuario a actualizar"
	ERROR_UPDATE_CATEGORY   = "No se puede encontrar la categorya a actualizar"
)
