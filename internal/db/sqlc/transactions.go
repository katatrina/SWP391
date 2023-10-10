package sqlc

import "context"

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

		providerID, err := qtx.CreateProvider(ctx, CreateProviderParams{
			FullName: arg.FullName,
			Email:    arg.Email,
			Phone:    arg.Phone,
			Address:  arg.Phone,
			Password: arg.Password,
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
