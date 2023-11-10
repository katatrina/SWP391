package sqlc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
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

	// -----------------------------------------------------
	// --------------------Phụ Kiện-------------------------
	// -----------------------------------------------------
	// Service 1
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "Cung cấp lồng chim đa dạng",
		Description:       "Lồng chim đa dạng: Cung cấp lồng chim chất lượng cao với nhiều lựa chọn về kích thước và chất liệu để đáp ứng nhu cầu sinh hoạt và di chuyển của chim cảnh.",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[0],
		OwnedByProviderID: providers[0].ID,
	})
	require.NoError(t, err)

	// Service 2
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[0],
		OwnedByProviderID: providers[1].ID,
	})
	require.NoError(t, err)

	// Service 3
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[0],
		OwnedByProviderID: providers[2].ID,
	})
	require.NoError(t, err)

	// Service 4
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[0],
		OwnedByProviderID: providers[3].ID,
	})
	require.NoError(t, err)

	// Service 5
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[0],
		OwnedByProviderID: providers[4].ID,
	})
	require.NoError(t, err)

	// Service 6
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[0],
		OwnedByProviderID: providers[5].ID,
	})
	require.NoError(t, err)

	// -----------------------------------------------------
	// --------------------Dinh dưỡng và Thức ăn------------
	// -----------------------------------------------------
	// Service 7
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[1],
		OwnedByProviderID: providers[0].ID,
	})
	require.NoError(t, err)

	// Service 8
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[1],
		OwnedByProviderID: providers[1].ID,
	})
	require.NoError(t, err)

	// Service 9
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[1],
		OwnedByProviderID: providers[2].ID,
	})
	require.NoError(t, err)

	// Service 10
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[1],
		OwnedByProviderID: providers[3].ID,
	})
	require.NoError(t, err)

	// Service 11
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[1],
		OwnedByProviderID: providers[4].ID,
	})
	require.NoError(t, err)

	// Service 12
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[1],
		OwnedByProviderID: providers[5].ID,
	})
	require.NoError(t, err)

	// -----------------------------------------------------------
	// --------------------Y tế và Chăm sóc sức khỏe--------------
	// -----------------------------------------------------------
	// Service 13
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[2],
		OwnedByProviderID: providers[0].ID,
	})
	require.NoError(t, err)

	// Service 14
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[2],
		OwnedByProviderID: providers[1].ID,
	})
	require.NoError(t, err)

	// Service 15
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[2],
		OwnedByProviderID: providers[2].ID,
	})
	require.NoError(t, err)

	// Service 16
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[2],
		OwnedByProviderID: providers[3].ID,
	})
	require.NoError(t, err)

	// Service 17
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[2],
		OwnedByProviderID: providers[4].ID,
	})
	require.NoError(t, err)

	// Service 18
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[2],
		OwnedByProviderID: providers[5].ID,
	})
	require.NoError(t, err)

	// -----------------------------------------------------------
	// --------------------Grooming-------------------------------
	// -----------------------------------------------------------
	// Service 19
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[3],
		OwnedByProviderID: providers[0].ID,
	})
	require.NoError(t, err)

	// Service 20
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[3],
		OwnedByProviderID: providers[1].ID,
	})
	require.NoError(t, err)

	// Service 21
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[3],
		OwnedByProviderID: providers[2].ID,
	})
	require.NoError(t, err)

	// Service 22
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[3],
		OwnedByProviderID: providers[3].ID,
	})
	require.NoError(t, err)

	// Service 23
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[3],
		OwnedByProviderID: providers[4].ID,
	})
	require.NoError(t, err)

	// Service 24
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[3],
		OwnedByProviderID: providers[5].ID,
	})
	require.NoError(t, err)

	// -----------------------------------------------------------
	// --------------------Đào tạo và Huấn luyện------------------
	// -----------------------------------------------------------
	// Service 25
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[4],
		OwnedByProviderID: providers[0].ID,
	})
	require.NoError(t, err)

	// Service 26
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[4],
		OwnedByProviderID: providers[1].ID,
	})
	require.NoError(t, err)

	// Service 27
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[4],
		OwnedByProviderID: providers[2].ID,
	})
	require.NoError(t, err)

	// Service 28
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[4],
		OwnedByProviderID: providers[3].ID,
	})
	require.NoError(t, err)

	// Service 29
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[4],
		OwnedByProviderID: providers[4].ID,
	})
	require.NoError(t, err)

	// Service 30
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[4],
		OwnedByProviderID: providers[5].ID,
	})
	require.NoError(t, err)

	// ------------------------------------------------
	// --------------------Khác------------------------
	// ------------------------------------------------
	// Service 31
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[5],
		OwnedByProviderID: providers[0].ID,
	})
	require.NoError(t, err)

	// Service 32
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[5],
		OwnedByProviderID: providers[1].ID,
	})
	require.NoError(t, err)

	// Service 33
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[5],
		OwnedByProviderID: providers[2].ID,
	})
	require.NoError(t, err)

	// Service 34
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[5],
		OwnedByProviderID: providers[3].ID,
	})
	require.NoError(t, err)

	// Service 35
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[5],
		OwnedByProviderID: providers[4].ID,
	})
	require.NoError(t, err)

	// Service 36
	err = testStore.CreateService(context.Background(), CreateServiceParams{
		Title:             "",
		Description:       "",
		Price:             200_000,
		ImagePath:         "",
		CategoryID:        categoryIDs[5],
		OwnedByProviderID: providers[5].ID,
	})
	require.NoError(t, err)
	// -----------------------------------------------------
	// -------------End of adding another service-----------
	// -----------------------------------------------------
	// Create admin
	adminPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	err = testStore.CreateAdmin(context.Background(), CreateAdminParams{
		Email:    "admin@gmail.com",
		Password: string(adminPassword),
	})
}
