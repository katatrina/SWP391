package sqlc

import (
	"context"
<<<<<<< Updated upstream
	"github.com/katatrina/SWP391/util"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
=======
>>>>>>> Stashed changes
	"testing"

	"github.com/stretchr/testify/require"
)

func createFakeCustomer(t *testing.T) {
	err := testStore.CreateCustomerTx(context.Background(), CreateCustomerParams{
		FullName: "Nguyen Huu Tung",
		Email:    "tung@gmail.com",
		Phone:    "0268101010",
		Address:  "167 lê đại hành",
		Password: "123456",
	})
	require.NoError(t, err)

	// customer 2
	err = testStore.CreateCustomerTx(context.Background(), CreateCustomerParams{
		FullName: "Nguyen Tran Lam",
		Email:    "lam@gmail.com",
		Phone:    "0312554689",
		Address:  "123 Hai Ba Trung",
		Password: "123456",
	})
	require.NoError(t, err)

	// customer 3
	err = testStore.CreateCustomerTx(context.Background(), CreateCustomerParams{
		FullName: "Le Duc Loi",
		Email:    "loi@gmail.com",
		Phone:    "0564219875",
		Address:  "519 Tran Khac Nhu",
		Password: "123456",
	})
	require.NoError(t, err)

	// customer 4
	err = testStore.CreateCustomerTx(context.Background(), CreateCustomerParams{
		FullName: "Lam Huu Chi",
		Email:    "chi@gmail.com",
		Phone:    "0164532159",
		Address:  "147 Tran Van On",
		Password: "123456",
	})
	require.NoError(t, err)

	// customer 5
	err = testStore.CreateCustomerTx(context.Background(), CreateCustomerParams{
		FullName: "To Duc Hai",
		Email:    "hai@gmail.com",
		Phone:    "0562314987",
		Address:  "156 To Duc Chi",
		Password: "123456",
	})
	require.NoError(t, err)

}

func createFakeProvider(t *testing.T) {
	err := testStore.CreateProviderTx(context.Background(), CreateProviderTxParams{
		FullName:    "Nguyen Thi Hang",
		Email:       "nguyenthihang@example.com",
		Phone:       "0394211201",
		Address:     "123 Nguyen Hue",
		TaxCode:     "1234567890",
		CompanyName: "ABC Company",
		Password:    "123456",
	})
	require.NoError(t, err)

	// provider 2
	err = testStore.CreateProviderTx(context.Background(), CreateProviderTxParams{
		FullName:    "Tran Van Anh",
		Email:       "tranvananh@example.com",
		Phone:       "0987654321",
		Address:     "456 Le Loi",
		TaxCode:     "0987654321",
		CompanyName: "XYZ Corporation",
		Password:    "123456",
	})
	require.NoError(t, err)

	// provider 3
	err = testStore.CreateProviderTx(context.Background(), CreateProviderTxParams{
		FullName:    "Le Thi Lan",
		Email:       "lethilan@example.com",
		Phone:       "0912345678",
		Address:     "789 Tran Phu",
		TaxCode:     "1357924680",
		CompanyName: "Sunshine Enterprises",
		Password:    "123456",
	})
	require.NoError(t, err)

	// provider 4
	err = testStore.CreateProviderTx(context.Background(), CreateProviderTxParams{
		FullName:    "Pham Minh Tuan",
		Email:       "phamminhtuan@example.com",
		Phone:       "0976543210",
		Address:     "321 Hoang Dieu",
		TaxCode:     "2468135790",
		CompanyName: "Galaxy Industries",
		Password:    "123456",
	})
	require.NoError(t, err)

	// provider 5
	err = testStore.CreateProviderTx(context.Background(), CreateProviderTxParams{
		FullName:    "Vu Thi Ngoc",
		Email:       "vuthingoc@example.com",
		Phone:       "0961234567",
		Address:     "654 Nguyen Trai",
		TaxCode:     "9876543210",
		CompanyName: "Starlight Co., Ltd.",
		Password:    "123456",
	})
	require.NoError(t, err)

	// provider 6
	err = testStore.CreateProviderTx(context.Background(), CreateProviderTxParams{
		FullName:    "Hoang Van Duc",
		Email:       "hoangvanduc@example.com",
		Phone:       "0943210765",
		Address:     "456 Nam Ky Khoi Nghia",
		TaxCode:     "0123456789",
		CompanyName: "Ocean Blue LLC",
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
