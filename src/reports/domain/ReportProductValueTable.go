package domain

import (
	"github.com/fajardm/inventories/src/commons/helpers"
)

func CreateViewReportProductValue() {
	db := helpers.GetDatabase()

	db.Exec(`CREATE VIEW report_product_values AS
  SELECT
    products.sku               AS sku,
    products.name              AS name,
    products.stock             AS stock,
    AVG(purchases.price)       AS avg_price,
    SUM(purchases.total_price) AS total_price
  FROM products
    INNER JOIN purchases ON products.id = purchases.product_id
  GROUP BY products.sku`)
}
