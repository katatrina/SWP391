package sqlc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	qtx := New(tx)

	err = fn(qtx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}

		return err
	}

	return tx.Commit()
}

type CreateProviderTxParams struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	CompanyName string `json:"company_name"`
	TaxCode     string `json:"tax_code"`
	Password    string `json:"hashed_password"`
}

func (store *Store) CreateProviderTx(ctx context.Context, arg CreateProviderTxParams) error {
	err := store.execTx(ctx, func(qtx *Queries) error {
		var err error

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(arg.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		providerID, err := qtx.CreateProvider(ctx, CreateProviderParams{
			FullName: arg.FullName,
			Email:    arg.Email,
			Phone:    arg.Phone,
			Address:  arg.Address,
			Password: string(hashedPassword),
		})
		if err != nil {
			return err
		}

		err = qtx.CreateProviderDetails(ctx, CreateProviderDetailsParams{
			ProviderID:  providerID,
			CompanyName: arg.CompanyName,
			TaxCode:     arg.TaxCode,
		})
		if err != nil {
			return err
		}

		// Create an empty cart for the provider.
		err = qtx.CreateCart(ctx, providerID)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (store *Store) CreateCustomerTx(ctx context.Context, arg CreateCustomerParams) error {
	err := store.execTx(ctx, func(qtx *Queries) error {
		var err error

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(arg.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		customerID, err := qtx.CreateCustomer(ctx, CreateCustomerParams{
			FullName: arg.FullName,
			Email:    arg.Email,
			Phone:    arg.Phone,
			Address:  arg.Address,
			Password: string(hashedPassword),
		})
		if err != nil {
			return err
		}

		// Create an empty cart for the customer.
		err = qtx.CreateCart(ctx, customerID)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

type CreateOrderTxParams struct {
	BuyerID       int32
	SellerID      int32
	PaymentMethod string
	CartItems     []GetCartItemsByCartIDRow
}

func (store *Store) CreateOrderTx(ctx context.Context, arg CreateOrderTxParams) error {
	err := store.execTx(ctx, func(qtx *Queries) error {
		var err error

		randomOrderID, _ := uuid.NewRandom()

		// Create order.
		order, err := qtx.CreateOrder(ctx, CreateOrderParams{
			UUID:          randomOrderID.String(),
			BuyerID:       arg.BuyerID,
			SellerID:      arg.SellerID,
			PaymentMethod: arg.PaymentMethod,
		})
		if err != nil {
			fmt.Println("create order error")
			return err
		}

		fmt.Printf("Order ID %s, Total %d\n", order.UUID, order.GrandTotal)

		for i, cartItem := range arg.CartItems {
			randomOrderItemID := randomOrderID.String() + "-" + fmt.Sprintf("%d", i+1)

			// Create order item.
			orderItem, err := qtx.CreateOrderItem(ctx, CreateOrderItemParams{
				UUID:      randomOrderItemID,
				OrderID:   order.UUID,
				ServiceID: cartItem.ServiceID,
				Quantity:  cartItem.Quantity,
				SubTotal:  cartItem.SubTotal,
			})
			if err != nil {
				fmt.Println("create order item error")
				return err
			}

			// Create order item details.
			err = qtx.CreateOrderItemDetails(ctx, CreateOrderItemDetailsParams{
				OrderItemID: orderItem.UUID,
				Title:       cartItem.Title,
				Price:       cartItem.Price,
				ImagePath:   cartItem.ImagePath,
			})
			if err != nil {
				fmt.Println("create order item details error")
				return err
			}

			// Retrieve order again to get the latest grand total.
			updatedOrder, err := qtx.GetOrderByOrderItemID(ctx, orderItem.UUID)
			if err != nil {
				fmt.Println("get updated order error")
				return err
			}

			// Update order total.
			err = qtx.UpdateOrderTotal(ctx, UpdateOrderTotalParams{
				GrandTotal: updatedOrder.GrandTotal + cartItem.SubTotal,
				UUID:       order.UUID,
			})
			if err != nil {
				fmt.Println("update order total error")
				return err
			}

			// Remove item from cart.
			err = qtx.RemoveItemFromCart(ctx, cartItem.UUID)
			if err != nil {
				fmt.Println("remove item from cart error")
				return err
			}
		}

		return nil
	})

	return err
}
