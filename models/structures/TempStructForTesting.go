package structures

import (
	"time"
)

type User struct {
	ID       int
	Username string
	Email    string
}

type Product struct {
	ID            int
	Name          string
	Description   string
	Price         float64
	StockQuantity int
}

type Order struct {
	ID        int
	UserID    int
	OrderDate time.Time
	Amount    float64
	Status    string
}

type OrderItem struct {
	ID        int
	OrderID   int
	ProductID int
	Quantity  int
	Price     float64
}

// gorm.Model includes default fields like ID, CreatedAt, UpdatedAt, and DeletedAt, which you don't want if you're not using primary keys.
