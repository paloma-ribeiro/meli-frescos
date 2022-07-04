package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paloma-ribeiro/meli-frescos/internal/employee/domain"
)

type requestEmployee struct {
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int    `json:"warehouse_id"`
}

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

		id, err := strconv.Atoi(ctx.Param("id"))
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

// @Summary Create employee
// @Tags Employees
// @Description Add a new employee to the list
// @Accept json
// @Produce json
// @Param employee body requestEmployee true "Employee to create"
// @Success 201 {object} schemes.Employee
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Failure 422 {object} schemes.JSONBadReqResult{error=string}
// @Router /employees [post]

func (c EmployeeController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requestEmployee
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		employee, err := c.service.Create(ctx, &domain.Employee{
			CardNumberId: req.CardNumberId,
			FirstName:    req.FirstName,
			LastName:     req.LastName,
			WarehouseId:  req.WarehouseId,
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, employee)
	}
}

// @Summary Update employee
// @Tags Employees
// @Description Update existing employee in list
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body requestEmployee true "Employee to update"
// @Success 200 {object} schemes.Employee
// @Failure 400 {object} schemes.JSONBadReqResult{error=string}
// @Failure 404 {object} schemes.JSONBadReqResult{error=string}
// @Router /employees/{id} [patch]

func (c EmployeeController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var req requestEmployee
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		employee, err := c.service.Update(ctx, &domain.Employee{
			ID:           id,
			CardNumberId: req.CardNumberId,
			FirstName:    req.FirstName,
			LastName:     req.LastName,
			WarehouseId:  req.WarehouseId,
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, employee)
	}
}
