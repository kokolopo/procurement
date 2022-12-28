package purchase_request

import (
	"net/http"
	"procurement/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PurchaseRequestHandler struct {
	purchaseRequestService IService
}

func NewPurchaseRequestHandler(purchaseRequestService IService) *PurchaseRequestHandler {
	return &PurchaseRequestHandler{purchaseRequestService}
}

func (h *PurchaseRequestHandler) CreateNewPurchaseRequest(c *gin.Context) {
	var input RequestForm
	// tangkap input body
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "cannot proses request")
		return
	}

	// record data PR
	newRequest, err := h.purchaseRequestService.Save(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, "cannot proses request")
	}

	formatter := FormatPR(newRequest)
	res := helper.ApiResponse("Request Successfully Accepted", 202, "Accepted", formatter)

	c.JSON(http.StatusOK, res)
}

func (h *PurchaseRequestHandler) GetAll(c *gin.Context) {
	requests, err := h.purchaseRequestService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadGateway, "server error")
		return
	}

	formatter := FormatPRs(requests)
	res := helper.ApiResponse("Purchase Requests Data", 200, "Ok", formatter)
	c.JSON(http.StatusOK, res)
}

func (h *PurchaseRequestHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, "id bukan angka")
		return
	}

	request, err := h.purchaseRequestService.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, "server error")
		return
	}

	formatter := FormatPR(request)
	res := helper.ApiResponse("Purchase Request Data", 200, "Ok", formatter)
	c.JSON(http.StatusOK, res)
}
