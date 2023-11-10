package sqlc

import (
	"context"
	"github.com/katatrina/SWP391/util"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func createFakeCustomer(t *testing.T) {
	err := testStore.CreateCustomerTx(context.Background(), CreateCustomerParams{
		FullName: "",
		Email:    "",
		Phone:    "0268101010",
		Address:  "167 lê đại hành",
		Password: "123456",
	})
	require.NoError(t, err)

	// customer 2
	err = testStore.CreateCustomerTx(context.Background(), CreateCustomerParams{
		FullName: "",
		Email:    "",
		Phone:    util.RandomPhone(10),
		Address:  "",
		Password: "123456789",
	})
	require.NoError(t, err)

	// customer 3
	err = testStore.CreateCustomerTx(context.Background(), CreateCustomerParams{
		FullName: "",
		Email:    "",
		Phone:    util.RandomPhone(10),
		Address:  "",
		Password: "123456789",
	})
	require.NoError(t, err)

}

func createFakeProvider(t *testing.T) {
	err := testStore.CreateProviderTx(context.Background(), CreateProviderTxParams{
		FullName:    "",
		Email:       "",
		Phone:       "0394211201",
		Address:     "",
		TaxCode:     "",
		CompanyName: "",
		Password:    "123456",
	})
	require.NoError(t, err)

	err = testStore.CreateProviderTx(context.Background(), CreateProviderTxParams{
		FullName:    "",
		Email:       "",
		Phone:       "",
		Address:     "",
		TaxCode:     "",
		CompanyName: "",
		Password:    "123456",
	})
	require.NoError(t, err)

	// provider 3
	err = testStore.CreateProviderTx(context.Background(), CreateProviderTxParams{
		FullName:    "",
		Email:       "",
		Phone:       "",
		Address:     "",
		TaxCode:     "",
		CompanyName: "",
		Password:    "123456",
	})
	require.NoError(t, err)
}

func createUser(t *testing.T) {
	createFakeCustomer(t)
	createFakeProvider(t)
}

func TestCreateAdmin(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)

	err := testStore.CreateAdmin(context.Background(), CreateAdminParams{
		Email:    "admin@gmail.com",
		Password: string(hashedPassword),
	})
	require.NoError(t, err)
}
