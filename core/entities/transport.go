package entities

type DataContent struct {
	Content string `json:"content"`
}
type CountBook struct {
	Index    int    `json:"index"`
	Quality  int    `json:"quality"`
	NameBook string `json:"name_book"`
}
type PriceBook struct {
	Index    int     `json:"index"`
	Price    float64 `json:"price"`
	NameBook string  `json:"name_book"`
}
type AuthorBook struct {
	Index       int    `json:"index"`
	Author      string `json:"author"`
	Nationality string `json:"nationality"`
	BirthDate   string `json:"birth_date"`
}

type OrderDetails struct {
	Index       int    `json:"index"`
	NameBook    string `json:"name_book"`
	Quantity    int    `json:"quantity"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type TypeBookResp struct {
	Index              int                   `json:"index"`
	ListBookByTypeBook []*ListBookByTypeBook `json:"type_books"`
}

type ListBookByTypeBook struct {
	Name string `json:"name"` // Tên của thể loại
}
