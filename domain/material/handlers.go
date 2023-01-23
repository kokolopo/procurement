package material

import (
	"net/http"
	"procurement/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MaterialHandler struct {
	materialService IService
}

func NewMaterialHandler(materialService IService) *MaterialHandler {
	return &MaterialHandler{materialService}
}

func (h *MaterialHandler) GetAll(c *gin.Context) {
	materials, err := h.materialService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadGateway, "server error")
		return
	}

	formatter := FormatMaterials(materials)
	res := helper.ApiResponse("Materials Data", 200, "Ok", formatter)
	c.JSON(http.StatusOK, res)
}

func (h *MaterialHandler) GetById(c *gin.Context) {
	// tangkap id yg diberikan dari http
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, "id bukan angka")
		return
	}

	// ambil data berdasarkan id
	material, err := h.materialService.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, "server error")
		return
	}

	formatter := FormatMaterial(material)
	res := helper.ApiResponse("Material Data", 200, "Ok", formatter)
	c.JSON(http.StatusOK, res)
}