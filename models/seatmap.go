package models

type Seat struct {
	Code            string   `json:"code"`
	Available       bool     `json:"available"`
	Price           float64  `json:"price"`
	Currency        string   `json:"currency"`
	Characteristics []string `json:"characteristics"`
	RowNumber       int      `json:"rowNumber"`
	Column          string   `json:"column"`
}

type Row struct {
	RowNumber int    `json:"rowNumber"`
	Seats     []Seat `json:"seats"`
}

type Cabin struct {
	Deck       string   `json:"deck"`
	SeatRows   []Row    `json:"seatRows"`
	SeatLayout []string `json:"seatColumns"`
}

type SeatMap struct {
	Aircraft string  `json:"aircraft"`
	Cabins   []Cabin `json:"cabins"`
}
