package domain

type Book struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Price     float64  `json:"price"`
	Currency  string   `json:"currency"`
	Count     int      `json:"count"`
	Category  []string `json:"category"`
	Author    string   `json:"author"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	DeletedAt string   `json:"deleted_at"`
}
