package sqlc

import (
	"context"
	"github.com/katatrina/SWP391/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createFakeCustomer(t *testing.T) {
	fullNames := []string{
		"Trần Văn A",
		"Nguyễn Văn B",
		"Lê Văn C",
		"Phạm Văn D",
		"Đặng Văn E",
	}

	addresses := []string{
		"89 Nguyễn Thị Minh Khai",
		"88 Nguyễn Trãi",
		"99 Nguyễn Văn Cừ",
		"100 Nguyễn Văn Linh",
		"101 Nguyễn Văn Trỗi",
	}

	for i := 0; i < 5; i++ {
		err := testStore.CreateCustomerTx(context.Background(), CreateCustomerParams{
			FullName: fullNames[i],
			Email:    util.RandomString(7) + "@gmail.com",
			Phone:    util.RandomPhone(10),
			Address:  addresses[i],
			Password: "123456789",
		})
		require.NoError(t, err)
	}
}

func createFakeProvider(t *testing.T) {
	fullNames := []string{
		"Nguyễn Lê Văn",
		"Trần Văn Luyến",
		"Nguyễn Thị Thanh",
		"Trần Thị Thúy",
		"Nguyễn Trung Trực",
	}

	addresses := []string{
		"44 Lê Văn Lương",
		"55 Nguyễn Văn Cừ",
		"66 Nguyễn Văn Linh",
		"77 Nguyễn Văn Trỗi",
		"88 Nguyễn Thị Minh Khai",
	}

	companyNames := []string{
		"FPTU",
		"NEU",
		"GRAP",
		"GOOGLE",
		"FACEBOOK",
	}

	for i := 0; i < 5; i++ {
		err := testStore.CreateProviderTx(context.Background(), CreateProviderTxParams{
			FullName:    fullNames[i],
			Email:       util.RandomString(7) + "@gmail.com",
			Phone:       util.RandomPhone(10),
			Address:     addresses[i],
			TaxCode:     util.RandomString(4),
			CompanyName: companyNames[i],
			Password:    "123456789",
		})
		require.NoError(t, err)
	}
}

func createUser(t *testing.T) {
	createFakeCustomer(t)
	createFakeProvider(t)
}
