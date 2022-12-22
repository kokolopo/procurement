package material

type MaterialFormatter struct {
	Id            int    `json:"id"`
	Supplier_id   int    `json:"supplier_id"`
	Material_name string `json:"material_name"`
	Material_type string `json:"material_type"`
	Stock         int    `json:"stock"`
}

func FormatMaterial(material Material) MaterialFormatter {
	formatter := MaterialFormatter{
		Id:            material.Id,
		Supplier_id:   material.Supplier_id,
		Material_name: material.Material_name,
		Material_type: material.Material_type,
		Stock:         material.Stock,
	}

	return formatter
}

func FormatMaterials(materials []Material) []MaterialFormatter {
	if len(materials) == 0 {
		return []MaterialFormatter{}
	}

	var materialsFormatter []MaterialFormatter

	for _, material := range materials {
		formatter := FormatMaterial(material)
		materialsFormatter = append(materialsFormatter, formatter)
	}

	return materialsFormatter
}
