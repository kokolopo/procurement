package purchase_request

import "time"

type PRFormatter struct {
	Id          int       `json:"id"`
	Material_id int       `json:"material_id"`
	Stock_need  int       `json:"stock_need"`
	CreatedAt   time.Time `json:"created_at"`
}

func FormatPR(data PurchaseRequest) PRFormatter {
	formatter := PRFormatter{
		Id:          data.Id,
		Material_id: data.Material_id,
		Stock_need:  data.Stock_need,
		CreatedAt:   data.CreatedAt,
	}

	return formatter
}

func FormatPRs(requests []PurchaseRequest) []PRFormatter {
	if len(requests) == 0 {
		return []PRFormatter{}
	}

	var purchaseFormatter []PRFormatter

	for _, request := range requests {
		formatter := FormatPR(request)
		purchaseFormatter = append(purchaseFormatter, formatter)
	}

	return purchaseFormatter
}
