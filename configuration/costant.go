package configuration

const (
	ERROR_SQLITE = "No se ha podido conectar la BD"
	ERROR_EMPTY  = "Esta informacion no puede venir sin datos"
)

const (
	SECURITY_DEBUG    = "security debug mode"
	SECURITY_NORMAL   = "security debug normal"
	CONNECTION_MYSQL  = "You have successfully connected to the Mysql database"
	CONNECTION_SQLITE = "You have successfully connected to the Sqlite database"
)

const (
	EXPIRATION_TOKEN = 24 //HOURS
	NAME_SESSION     = "Authorization"
	DELETE_USER      = "Se ha eliminado el usuario"
	DELETE_CATEGORY  = "Se ha eliminado la categoria"
	DELETE_EXPENSES  = "Se ha eliminado la lista de gastos"
	DELETE_PRODUCT   = "Se ha eliminado el producto"
)

const (
	ERROR_EMPTY_FIELD       = "Se debe ingresar datos"
	ERROR_INDEX             = "Atributo no indexado para procesar"
	ERROR_DATABASE_GET      = "No se ha encontrado ningun motor de basededatos para procesar"
	ERROR_SERVICE_SECURITY  = "Service error in Security"
	ERROR_SERVICE_USER      = "Service error in user"
	ERROR_SERVICE_CATEGORY  = "Service error in Category"
	ERROR_SERVICE_EXPENSES  = "Service error in Expenses"
	ERROR_SERVICE_PRODUCT   = "Service error in Product"
	ERROR_PRIVILEGES_GEN    = "No Tienes los privilegios para acceder"
	ERROR_UPDATE_USER_EMAIL = "No se puede encontrar el usuario a actualizar"
	ERROR_UPDATE_CATEGORY   = "No se puede encontrar la categorya a actualizar"
	ERROR_UPDATE_PRODUCT    = "No se puede encontrar el producto a actualizar"
)
