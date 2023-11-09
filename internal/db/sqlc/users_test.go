package sqlc

import (
	"context"
	"github.com/katatrina/SWP391/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createFakeCustomer(t *testing.T) {
	//fullNames := []string{
	//	"Trần Văn A",
	//	"Nguyễn Văn B",
	//	"Lê Văn C",
	//	"Phạm Văn D",
	//	"Đặng Văn E",
	//}
	//
	//addresses := []string{
	//	"89 Nguyễn Thị Minh Khai",
	//	"88 Nguyễn Trãi",
	//	"99 Nguyễn Văn Cừ",
	//	"100 Nguyễn Văn Linh",
	//	"101 Nguyễn Văn Trỗi",
	//}

	err := testStore.CreateCustomerTx(context.Background(), CreateCustomerParams{
		FullName: "",
		Email:    "",
		Phone:    util.RandomPhone(10),
		Address:  "",
		Password: "123456789",
	})
	require.NoError(t, err)

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
	//fullNames := []string{
	//	"Nguyễn Lê Văn",
	//	"Trần Văn Luyến",
	//	"Nguyễn Thị Thanh",
	//	"Trần Thị Thúy",
	//	"Nguyễn Trung Trực",
	//}

	//addresses := []string{
	//	"44 Lê Văn Lương",
	//	"55 Nguyễn Văn Cừ",
	//	"66 Nguyễn Văn Linh",
	//	"77 Nguyễn Văn Trỗi",
	//	"88 Nguyễn Thị Minh Khai",
	//}

	//companyNames := []string{
	//	"FPTU",
	//	"NEU",
	//	"GRAP",
	//	"GOOGLE",
	//	"FACEBOOK",
	//}

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
}

func createUser(t *testing.T) {
	createFakeCustomer(t)
	createFakeProvider(t)
}
