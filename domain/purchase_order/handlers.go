package purchase_order

import (
	"net/http"
	"procurement/domain/purchase_request"
	"procurement/domain/supplier"
	"procurement/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PurchaseOrderHandler struct {
	purchaseOrderService IService
	PurchaseRequest      purchase_request.IService
	Supplier             supplier.IService
}

func NewPurchaseOrderHandler(purchaseOrderService IService, PurchaseRequest purchase_request.IService, Supplier supplier.IService) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{purchaseOrderService, PurchaseRequest, Supplier}
}

func (h *PurchaseOrderHandler) CreateNewPurchaseOrder(c *gin.Context) {
	var input InputPO
	// tangkap input body
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "cannot proses request")
		return
	}

	// cek apakan id PR ada
	pr, err := h.PurchaseRequest.GetById(input.PurchaseRequest_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "cannot proses request")
		return
	}
	if pr.Id != input.PurchaseRequest_id {
		c.JSON(http.StatusBadRequest, "cannot proses request")
		return
	}

	// cek apakan id Supplier ada
	supplier, err := h.Supplier.GetById(input.Supplier_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "cannot proses request")
		return
	}
	if supplier.Id != input.Supplier_id {
		c.JSON(http.StatusBadRequest, "cannot proses request")
		return
	}

	// record data PR
	newOrder, err := h.purchaseOrderService.Save(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, "cannot proses request")
		return
	}
	newOrder.PurchaseRequest.Id = pr.Id
	newOrder.PurchaseRequest.MaterialName = pr.MaterialName
	newOrder.PurchaseRequest.Stock_need = pr.Stock_need
	newOrder.Supplier.Id = supplier.Id
	newOrder.Supplier.SupplierName = supplier.SupplierName
	newOrder.Supplier.Email = supplier.Email
	newOrder.Supplier.Phone = supplier.Phone
	newOrder.Supplier.Address = supplier.Address

	// formatter := FormatPR(newRequest)
	res := helper.ApiResponse("Request Successfully Accepted", 202, "Accepted", newOrder)

	c.JSON(http.StatusOK, res)
}

func (h *PurchaseOrderHandler) GetAll(c *gin.Context) {
	order, err := h.purchaseOrderService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadGateway, "server error")
		return
	}

	// formatter := FormatPRs(requests)
	res := helper.ApiResponse("Purchase Order Data", 200, "Ok", order)
	c.JSON(http.StatusOK, res)
}

func (h *PurchaseOrderHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, "id bukan angka")
		return
	}

	order, err := h.purchaseOrderService.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, "server error")
		return
	}

	// formatter := FormatPR(request)
	res := helper.ApiResponse("Purchase Order Data", 200, "Ok", order)
	c.JSON(http.StatusOK, res)
}

func (h *PurchaseOrderHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input InputEditPO

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "cannot proses request")
		return
	}

	_, err = h.purchaseOrderService.Update(id, input)
	if err != nil {
		res := helper.ApiResponse("Update Data Has Been Failed", http.StatusUnprocessableEntity, "failed", err)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	data := gin.H{"is_uploaded": true}
	res := helper.ApiResponse("Update Data Has Been Success", http.StatusCreated, "success", data)

	c.JSON(http.StatusCreated, res)

}

func (h *PurchaseOrderHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.purchaseOrderService.Delete(id)
	if err != nil {
		res := helper.ApiResponse("User Not Found", http.StatusBadRequest, "failed", err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	data := gin.H{"is_deleted": true}
	res := helper.ApiResponse("Purchase Order Data Deleted", 200, "Ok", data)

	c.JSON(http.StatusCreated, res)
}
