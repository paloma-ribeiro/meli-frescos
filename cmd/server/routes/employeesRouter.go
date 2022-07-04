package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/paloma-ribeiro/meli-frescos/internal/employee/controller"
	"github.com/paloma-ribeiro/meli-frescos/internal/employee/repository/mariadb"
	"github.com/paloma-ribeiro/meli-frescos/internal/employee/service"
)

func employeesRouter(superRouter *gin.RouterGroup, conn *sql.DB) {
	repository := mariadb.NewMariaDBRepository(conn)
	service := service.NewService(repository)
	controller := controller.NewEmployeeController(service)

	pr := superRouter.Group("/employees")
	{
		pr.GET("/", controller.GetAll())
		pr.GET("/:id", controller.GetById())
		pr.POST("/", controller.Create())
		pr.PATCH("/:id", controller.Update())
		pr.DELETE("/:id", controller.Delete())
	}
}
