package models

import (
	"golang_API/config"
	"net/http"
)

type Signin struct {
	Id     int    `json:"id"`
	UserId string `json:"userid"`
	Pass   string `json:"pass"`
	Nama   string `json:"nama"`
	Level  string `json:"level"`
	Email  string `json:"email"`
	Foto   string `json:"foto"`
	Token  string `json:"token"`
	Token2 string `json:"token2"`
}

func SigninUser() (Response, error) {
	var obj Signin
	var arrobj []Signin
	var res Response

	db := config.ConnectToDB()
	sqlStatement := "SELECT * FROM signin"
	rows, err := db.Query(sqlStatement)
	// defer rows.Close()
	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.UserId, &obj.Pass, &obj.Nama, &obj.Level, &obj.Email, &obj.Foto, &obj.Token, &obj.Token2)
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
