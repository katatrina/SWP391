package sqlc

import (
	"context"
	"github.com/katatrina/SWP391/util"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCreateService(t *testing.T) {
	createUser(t)

	providers, err := testStore.ListProviders(context.Background())
	require.NoError(t, err)

	categoryIDs, err := testStore.ListCategoryIDs(context.Background())
	require.NoError(t, err)

	/*
		Table categories:
		id | title
		---+----------------
		0  | Phụ kiện
		1  | Dinh dưỡng và thức ăn
		2  | Y tế và chăm sóc sức khỏe
		3  | Grooming
		4  | Đào tạo và huấn luyện
		5  | Khác
	*/
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Tên dịch vụ",
		Description:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod, nisl eget ultricies aliquam, nunc nisl aliquet nunc, vitae aliqua",
		Price:             util.RandomPrice(),
		ImagePath:         "/static/img/sadasd/asdasd.jpg",
		CategoryID:        categoryIDs[1],
		OwnedByProviderID: providers[0].ID,
	})
	require.NoError(t, err)

	// Add another service ...
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Tên dịch vụ",
		Description:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod, nisl eget ultricies aliquam, nunc nisl aliquet nunc, vitae aliqua",
		Price:             util.RandomPrice(),
		ImagePath:         "/static/img/sadasd/asdasd.jpg",
		CategoryID:        categoryIDs[1],
		OwnedByProviderID: providers[0].ID,
	})
	require.NoError(t, err)

	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Tên dịch vụ",
		Description:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod, nisl eget ultricies aliquam, nunc nisl aliquet nunc, vitae aliqua",
		Price:             util.RandomPrice(),
		ImagePath:         "/static/img/sadasd/asdasd.jpg",
		CategoryID:        categoryIDs[1],
		OwnedByProviderID: providers[0].ID,
	})
	require.NoError(t, err)
	// End of adding another service

	// Create admin
	adminPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	err = testStore.CreateAdmin(context.Background(), CreateAdminParams{
		Email:    "admin@gmail.com",
		Password: string(adminPassword),
	})
}
