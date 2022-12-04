package database

type SaleItem struct {
	ID     int `json:"id"`
	SaleID int `json:"sale_id"`
	ItemID int `json:"item_id"`
	Amount int `json:"amount"`
}

type Sale struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
}

type Item struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	CreatedAt string  `json:"created_at"`
	Cost      float64 `json:"cost"`
	UserID    int     `json:"user_id"`
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type AccessToken struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}
