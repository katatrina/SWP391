// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package sqlc

import (
	"time"
)

type Cart struct {
	ID         int32 `json:"id"`
	UserID     int32 `json:"user_id"`
	GrandTotal int32 `json:"grand_total"`
}

type CartItem struct {
	UUID      string `json:"uuid"`
	CartID    int32  `json:"cart_id"`
	ServiceID int32  `json:"service_id"`
	Quantity  int32  `json:"quantity"`
	SubTotal  int32  `json:"sub_total"`
}

type Category struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	ImagePath   string `json:"image_path"`
	Description string `json:"description"`
}

type Order struct {
	UUID          string    `json:"uuid"`
	BuyerID       int32     `json:"buyer_id"`
	SellerID      int32     `json:"seller_id"`
	Status        int32     `json:"status"`
	PaymentMethod string    `json:"payment_method"`
	GrandTotal    int32     `json:"grand_total"`
	CreatedAt     time.Time `json:"created_at"`
}

type OrderItem struct {
	UUID      string    `json:"uuid"`
	OrderID   string    `json:"order_id"`
	ServiceID int32     `json:"service_id"`
	Quantity  int32     `json:"quantity"`
	SubTotal  int32     `json:"sub_total"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderItemDetail struct {
	ID          int32     `json:"id"`
	OrderItemID string    `json:"order_item_id"`
	Title       string    `json:"title"`
	Price       int32     `json:"price"`
	ImagePath   string    `json:"image_path"`
	CreatedAt   time.Time `json:"created_at"`
}

type OrderStatusCategory struct {
	ID     int32  `json:"id"`
	Code   string `json:"code"`
	Detail string `json:"detail"`
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
	ID                int32     `json:"id"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Price             int32     `json:"price"`
	ImagePath         string    `json:"image_path"`
	CategoryID        int32     `json:"category_id"`
	OwnedByProviderID int32     `json:"owned_by_provider_id"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
}

type ServiceFeedback struct {
	ID        int32     `json:"id"`
	ServiceID int32     `json:"service_id"`
	UserID    int32     `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
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
