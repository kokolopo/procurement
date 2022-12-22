package database

import (
	"procurement/domain/material"
	"procurement/domain/purchase_order"
	"procurement/domain/purchase_request"
	"procurement/domain/supplier"
)

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: material.Material{}},
		{Model: supplier.Supplier{}},
		{Model: purchase_order.PurchaseOrder{}},
		{Model: purchase_request.PurchaseRequest{}},
	}
}
