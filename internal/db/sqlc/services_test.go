package sqlc

import (
	"context"
	"fmt"
	"github.com/katatrina/SWP391/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateService(t *testing.T) {
	createUser(t)

	providers, err := testStore.ListProviders(context.Background())
	require.NoError(t, err)

	categoryIDs, err := testStore.ListCategoryIDs(context.Background())
	require.NoError(t, err)

	for _, provider := range providers {
		err = testStore.CreateService(context.Background(), CreateServiceParams{
			Title:             "This is a title",
			Description:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod, nisl eget ultricies aliquam, nunc nisl aliquet nunc, vitae aliqua",
			Price:             100000,
			ImagePath:         fmt.Sprintf("https://picsum.photos/id/%d/5000/3333", util.RandomInt(1, 1000)),
			CategoryID:        categoryIDs[util.RandomInt(0, len(categoryIDs))-1],
			OwnedByProviderID: provider.ID,
		})
		require.NoError(t, err)
	}
}