package services

import (
	"github.com/fajardm/inventories/src/commons/helpers"
	"bytes"
	"encoding/csv"
)

type ReportProductSaleService struct {
	helpers.Response
	Report [][]string
	Buffer bytes.Buffer
}

func (res *ReportProductSaleService) ReportProductSale() *ReportProductSaleService {
	db := helpers.GetDatabase()

	rows, err := db.Raw("SELECT sku, name, datetime, number_shipped, total_price, purchase_price, profit FROM report_product_sales").Rows()

	if err != nil {
		res.Error = err.Error()
		return res
	}

	b := &bytes.Buffer{}
	w := csv.NewWriter(b)

	var data = [][]string{
		{"Laporan Penjualan"},
		{},
		{"Tanggal Cetak", ""},
		{"Tanggal", ""},
		{"Total Omzet", ""},
		{"Total Laba Kotor", ""},
		{"Total Penjualan", ""},
		{"Total Barang", ""},
		{"SKU", "Name"},
		{},
		{"Waktu", "SKU", "Nama Barang", "Jumlah", "Harga Jual", "Total", "Harga Beli", "Laba"},
	}

	for _, value := range data {
		if err := w.Write(value); err != nil {
			res.Error = err.Error()
			return res
		}
	}

	for rows.Next() {
		var sku string
		var name string
		var datetime string
		var numberShipped string
		var totalPrice string
		var purchasePrice string
		var profit string

		rows.Scan(&sku, &name, &datetime, &numberShipped, &totalPrice, &purchasePrice, &profit)
		temp := []string{sku, name, datetime, numberShipped, totalPrice, purchasePrice, profit}

		if err := w.Write(temp); err != nil {
			res.Error = err.Error()
			return res
		}

		data = append(data, temp)
	}

	w.Flush()

	res.Report = data
	res.Buffer = *b

	return res
}
