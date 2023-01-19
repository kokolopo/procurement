package supplier

import (
	"net/http"
	"procurement/helper"

	"github.com/gin-gonic/gin"
)

type SupplierHandler struct {
	supplierService IService
}

func NewSupplierHandler(supplierService IService) *SupplierHandler {
	return &SupplierHandler{supplierService}
}

func (h *SupplierHandler) CreateNewSupplier(c *gin.Context) {
	var input AddSuplier
	// tangkap input body
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "cannot proses request")
		return
	}

	// record data
	newSupplier, err := h.supplierService.Save(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, "cannot proses request")
	}

	// formatter := FormatPR(newRequest)
	res := helper.ApiResponse("Request Successfully Accepted", 202, "Accepted", newSupplier)

	c.JSON(http.StatusOK, res)
}

func (h *SupplierHandler) GetAll(c *gin.Context) {
	suppliers, err := h.supplierService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadGateway, "server error")
		return
	}

	// formatter := FormatMaterials(materials)
	res := helper.ApiResponse("Materials Data", 200, "Ok", suppliers)
	c.JSON(http.StatusOK, res)
}