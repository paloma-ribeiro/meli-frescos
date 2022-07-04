package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func AddRoutes(superRouter *gin.RouterGroup, dbConnection *sql.DB) {
	employeesRouter(superRouter)
}
