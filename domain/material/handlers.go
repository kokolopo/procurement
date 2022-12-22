package material

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MaterialHandler struct {
	materialService IService
}

func NewClientHandler(materialService IService) *MaterialHandler {
	return &MaterialHandler{materialService}
}

func (h *MaterialHandler) GetAll(c *gin.Context) {
	materials, err := h.materialService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadGateway, "server error")
		return
	}

	res := FormatMaterials(materials)
	c.JSON(http.StatusOK, res)
}

func (h *MaterialHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, "id bukan angka")
		return
	}

	material, err := h.materialService.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, "server error")
		return
	}

	res := FormatMaterial(material)
	c.JSON(http.StatusOK, res)
}
