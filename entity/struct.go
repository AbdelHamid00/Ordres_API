package entity 

// type food struct {
// 	Gateaux string
// 	boisson string
// 	Tacos	string
// 	Sandwich string
// 	plat	string
// 	entree_chaud string
// 	entree_froide string
// 	au_kg	string
// }

type Client struct {
    Name string `"json": "name" binding:"required"`
	Phone string `"json": "phone" binding:"required"`
	Adress string `"json": "adress" binding:"required"`
	Ordre string `"json": "ordre" binding:"required"`
	State uint32 `"json": "state"`
}

type  IDClient struct {
	Id uint32 `"json": "state"`
    Name string `"json": "name" `
	Phone string `"json": "phone"`
	Adress string `"json": "adress" `
	Ordre string `"json": "ordre"`
	State uint32 `"json": "state"`
}

// POST /Commande HTTP/1.1
// Host:localhost
// Accept:application/json
// Content-Type: application/json
// Content-Length: 78

//Authorization: Basic 

// {"name":"Hamid","phone":"0623313463","adress":"Rabat","commande":"Tacos kbir"}