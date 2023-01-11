package models

import (
	"golang_API/config"
	"net/http"
)

type Kategori struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
}

func DataKategori() (Response, error) {
	var obj Kategori
	var arrobj []Kategori
	var res Response

	db := config.ConnectToDB()
	sqlStatement := "SELECT * FROM kategori"
	rows, err := db.Query(sqlStatement)
	// defer rows.Close()
	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama)
		// err = rows.Scan(&obj.Nomor, &obj.Pengirim, &obj.Penerima, &obj.Alamat_Penerima, &obj.Product, &obj.Product_Type, &obj.Status_Barang, &obj.Estimasi)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "SUKSES"
	res.Data = arrobj

	return res, nil

}
