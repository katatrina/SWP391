package sqlc

import (
	"context"
	"github.com/katatrina/SWP391/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createFakeCustomer(t *testing.T) {
	for i := 0; i < 5; i++ {
		fullName := util.RandomString(10)
		err := testStore.CreateCustomerTx(context.Background(), CreateCustomerParams{
			FullName: fullName,
			Email:    fullName + "@gmail.com",
			Phone:    util.RandomPhone(10),
			Address:  "89 Nguyễn Thị Minh Khai",
			Password: "123456",
		})
		require.NoError(t, err)
	}
	//err := testStore.CreateCustomerTx(context.Background(), CreateCustomerParams{
	//	FullName: "Nguyễn Văn A",
	//	Email:    "nguyenvanA@gmail.com",
	//	Phone:    "0123456789",
	//	Address:  "123 Lê Lợi",
	//	Password: "123456",
	//})
	//require.NoError(t, err)
	//
	//err = testStore.CreateCustomerTx(context.Background(), CreateCustomerParams{
	//	FullName: "Trần Văn B",
	//	Email:    "tranvanB@gmail.com",
	//	Phone:    "0123456788",
	//	Address:  "67 Nguyễn Huệ",
	//	Password: "123456",
	//})
	//require.NoError(t, err)

}

func createFakeProvider(t *testing.T) {
	for i := 0; i < 5; i++ {
		err := testStore.CreateProviderTx(context.Background(), CreateProviderTxParams{
			FullName:    util.RandomString(10),
			Email:       util.RandomString(10) + "@gmail.com",
			Phone:       util.RandomPhone(10),
			Address:     "88 Nguyễn Trãi",
			TaxCode:     "444",
			CompanyName: "FPTU",
			Password:    "123456",
		})
		require.NoError(t, err)
	}

	//err = testStore.CreateProviderTx(context.Background(), CreateProviderTxParams{
	//	FullName:    "Trần Văn E",
	//	Email:       "tranvanE@gmail.com",
	//	Phone:       "0123456785",
	//	Address:     "99 Nguyễn Văn Cừ",
	//	TaxCode:     "555",
	//	CompanyName: "Công ty E",
	//	Password:    "123456",
	//})
	//require.NoError(t, err)
	//
	//err = testStore.CreateProviderTx(context.Background(), CreateProviderTxParams{
	//	FullName:    "Lê Văn F",
	//	Email:       "levanF@gmail.com",
	//	Phone:       "0123456784",
	//	Address:     "100 Nguyễn Văn Linh",
	//	TaxCode:     "666",
	//	CompanyName: "Công ty F",
	//	Password:    "123456",
	//})
	//require.NoError(t, err)
}

func createUser(t *testing.T) {
	createFakeCustomer(t)
	createFakeProvider(t)
}
