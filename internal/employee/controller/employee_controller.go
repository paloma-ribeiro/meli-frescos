package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paloma-ribeiro/meli-frescos/internal/employee/domain"
)

type EmployeeController struct {
	service domain.EmployeeService
}

func NewEmployeeController(service domain.EmployeeService) (*EmployeeController, error) {
	if service == nil {
		return nil, domain.ErrInvalidService
	}

	return &EmployeeController{
		service: service,
	}, nil
}

// @Summary List employees
// @Tags Employees
// @Description get all employees
// @Accept json
// @Produce json
// @Success 200 {object} schemes.JSONSuccessResult{data=schemes.Employee}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /employees [get]
func (c EmployeeController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employees, err := c.service.GetAll(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": employees})
	}
}

// @Summary Employee by id
// @Tags Employees
// @Description get employee by id
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} schemes.Employee
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /employees/{id} [get]
func (c EmployeeController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		employee, err := c.service.GetById(ctx, id)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, employee)
	}
}
