package domain

import "github.com/fajardm/inventories/src/commons/helpers"

func CreateViewReportProductSale() {
	db := helpers.GetDatabase()

	db.Exec(`CREATE VIEW report_product_sales AS
  SELECT
    products.sku                                                   AS sku,
    products.name                                                  AS name,
    products.created_at                                            AS datetime,
    sales.number_shipped                                           AS number_shipped,
    sales.price                                                    AS sale_price,
    sales.total_price                                              AS total_price,
    purchases.price                                                AS purchase_price,
    (sales.total_price - (purchases.price * sales.number_shipped)) AS profit
  FROM sales
    INNER JOIN products ON sales.product_id = products.id
    INNER JOIN purchases ON products.id = purchases.product_id
  GROUP BY sales.id`)
}
