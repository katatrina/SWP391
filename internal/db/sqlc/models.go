// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package sqlc

import (
	"database/sql"
	"time"
)

type Blog struct {
	ID        int32     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Cart struct {
	ID        int32        `json:"id"`
	UserID    int32        `json:"user_id"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type CartItem struct {
	ID        int32 `json:"id"`
	CartID    int32 `json:"cart_id"`
	ServiceID int32 `json:"service_id"`
	Quantity  int32 `json:"quantity"`
	Price     int32 `json:"price"`
}

type Category struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	ThumbnailUrl string `json:"thumbnail_url"`
	Description  string `json:"description"`
}

type Feedback struct {
	ID        int32     `json:"id"`
	ServiceID int32     `json:"service_id"`
	UserID    int32     `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Order struct {
	ID           int32     `json:"id"`
	BuyerID      int32     `json:"buyer_id"`
	SellerID     int32     `json:"seller_id"`
	DeliveryDate time.Time `json:"delivery_date"`
	DeliveredTo  string    `json:"delivered_to"`
	Status       string    `json:"status"`
	Total        int32     `json:"total"`
	CreatedAt    time.Time `json:"created_at"`
}

type Orderdetail struct {
	ID        int32     `json:"id"`
	OrderID   int32     `json:"order_id"`
	ServiceID int32     `json:"service_id"`
	Quantity  int32     `json:"quantity"`
	Price     int32     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type ProviderDetail struct {
	ID          int32     `json:"id"`
	ProviderID  int32     `json:"provider_id"`
	CompanyName string    `json:"company_name"`
	TaxCode     string    `json:"tax_code"`
	CreatedAt   time.Time `json:"created_at"`
}

type Role struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Service struct {
	ID            int32     `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Price         int32     `json:"price"`
	CategoryID    int32     `json:"category_id"`
	ThumbnailUrl  string    `json:"thumbnail_url"`
	OwnedByUserID int32     `json:"owned_by_user_id"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

type Session struct {
	Token  string    `json:"token"`
	Data   []byte    `json:"data"`
	Expiry time.Time `json:"expiry"`
}

type User struct {
	ID        int32     `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	RoleID    int32     `json:"role_id"`
	Password  string    `json:"hashed_password"`
	CreatedAt time.Time `json:"created_at"`
}
