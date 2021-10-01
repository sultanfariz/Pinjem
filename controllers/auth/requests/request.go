package requests

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	ID          uint
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Fullname    string `json:"fullname" form:"fullname"`
	NIK         string `json:"nik" form:"nik"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Birthdate   string `json:"birthdate" form:"birthdate"`
	Address     string `json:"address" form:"address"`
	Provinsi    string `json:"provinsi" form:"provinsi"`
	Kota        string `json:"kota" form:"kota"`
	Kecamatan   string `json:"kecamatan" form:"kecamatan"`
	Desa        string `json:"desa" form:"desa"`
	PostalCode  string `json:"postalCode" form:"postalCode"`
	Role        string `json:"role" form:"role"`
	Status      int    `json:"status" form:"status"`
}
