package boot

import (
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/products/domain"
	domain2 "github.com/fajardm/inventories/src/purchases/domain"
	domain3 "github.com/fajardm/inventories/src/sales/domain"
)

func Seeder() {
	product1, product2 := createProduct()
	createPurchase(product1, product2)
	createSale(product1, product2)
}

func createProduct() (string, string) {
	db := helpers.GetDatabase()

	product1 := domain.Product{
		SKU:   "SSI-D00791015-LL-BWH",
		Name:  "Zalekia Plain Casual Blouse (L,Broken White)",
		Stock: 8,
	}
	db.Create(&product1)

	product2 := domain.Product{
		SKU:   "SSI-D00864612-LL-NAV",
		Name:  "Deklia Plain Casual Blouse (L,Navy)",
		Stock: 13,
	}
	db.Create(&product2)

	return product1.ID, product2.ID
}

func createSale(product1 string, product2 string) {
	db := helpers.GetDatabase()

	sale1 := domain3.Sale{
		NumberShipped: 2,
		Price:         20000,
		TotalPrice:    40000,
		ProductId:     product1,
	}
	db.Create(&sale1)

	sale2 := domain3.Sale{
		NumberShipped: 2,
		Price:         20000,
		TotalPrice:    40000,
		ProductId:     product2,
	}
	db.Create(&sale2)
}

func createPurchase(product1 string, product2 string) {
	db := helpers.GetDatabase()

	purchase1 := domain2.Purchase{
		OrderQuantity:  5,
		NumberReceived: 5,
		Price:          10000,
		TotalPrice:     50000,
		ProductId:      product1,
	}
	db.Create(&purchase1)

	purchase2 := domain2.Purchase{
		OrderQuantity:  5,
		NumberReceived: 5,
		Price:          10000,
		TotalPrice:     50000,
		ProductId:      product1,
	}
	db.Create(&purchase2)

	purchase3 := domain2.Purchase{
		OrderQuantity:  5,
		NumberReceived: 5,
		Price:          10000,
		TotalPrice:     50000,
		ProductId:      product2,
	}
	db.Create(&purchase3)

	purchase4 := domain2.Purchase{
		OrderQuantity:  5,
		NumberReceived: 5,
		Price:          10000,
		TotalPrice:     50000,
		ProductId:      product2,
	}
	db.Create(&purchase4)

	purchase5 := domain2.Purchase{
		OrderQuantity:  5,
		NumberReceived: 5,
		Price:          10000,
		TotalPrice:     50000,
		ProductId:      product2,
	}
	db.Create(&purchase5)
}
