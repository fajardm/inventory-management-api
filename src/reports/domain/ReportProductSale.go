package domain

type ReportProductSale struct {
	SKU           string
	Name          string
	Datetime      string
	NumberShipped int
	SalePrice     float32
	TotalPrice    float32
	PurchasePrice float32
	Profit        float32
}
