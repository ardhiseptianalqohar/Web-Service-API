package models

import (
	"golang_API/config"
	"net/http"
)

type Produk struct {
	Id            int     `json:"id"`
	IdKategori    int     `json:"idkategori"`
	IdSubKategori int     `json:"idsubkategori"`
	Kategori      string  `json:"kategori"`
	SubKategori   string  `json:"subkategori"`
	Judul         string  `json:"judul"`
	Deskripsi     string  `json:"deskripsi"`
	Harga         float64 `json:"harga"`
	Thumbnail     string  `json:"thumbnail"`
	St            string  `json:"st"`
}

func DataProduk() (Response, error) {
	var obj Produk
	var arrobj []Produk
	var res Response

	db := config.ConnectToDB()
	sqlStatement := "SELECT * FROM produk"
	rows, err := db.Query(sqlStatement)
	// defer rows.Close()
	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdKategori, &obj.IdSubKategori, &obj.Kategori, &obj.SubKategori, &obj.Judul, &obj.Deskripsi, &obj.Harga, &obj.Thumbnail, &obj.St)
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
