package services

import (
	"github.com/fajardm/inventories/src/commons/helpers"
	"bytes"
	"encoding/csv"
)

type ReportProductValueService struct {
	helpers.Response
	Report [][]string
	Buffer bytes.Buffer
}

func (res *ReportProductValueService) ReportProductValue() *ReportProductValueService {
	db := helpers.GetDatabase()

	rows, err := db.Raw("SELECT sku, name, stock, avg_price, total_price FROM report_product_values").Rows()

	if err != nil {
		res.Error = err.Error()
		return res
	}

	b := &bytes.Buffer{}
	w := csv.NewWriter(b)

	var data = [][]string{
		{"LAPORAN NILAI BARANG"},
		{},
		{"Tanggal Cetak", ""},
		{"Jumlah SKU", ""},
		{"Jumlah Total Barang", ""},
		{"Total Nilai", ""},
		{},
		{"SKU", "Nama Item", "Jumlah", "Rata-Rata Harga Beli", "Total"},
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
		var stock string
		var avgPrice string
		var totalPrice string

		rows.Scan(&sku, &name, &stock, &avgPrice, &totalPrice)
		temp := []string{sku, name, stock, avgPrice, totalPrice}

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
