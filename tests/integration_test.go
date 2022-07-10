package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/Kryniol/ports/domain"
	"github.com/Kryniol/ports/infrastructure"
)

const testInputPath = "testdata/valid.json"

func TestIntegration_FullFlow_Successful(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	ctx := context.Background()
	repo := infrastructure.NewInMemoryPortRepository()
	reader := infrastructure.NewJSONReader(testInputPath, zap.NewNop())
	svc := domain.NewPortService(reader, repo)
	err := svc.SavePorts(ctx)

	assert.NoError(t, err)

	portID, err := domain.NewPortID("AGSJO")
	require.NoError(t, err)
	port, err := repo.Get(ctx, *portID)
	require.NoError(t, err)
	assert.Equal(t, *portID, port.ID)
	assert.Equal(t, "Saint John's", port.Name)
	assert.Equal(t, "Saint John's", port.Address.City)
	assert.Equal(t, "Antigua and Barbuda", port.Address.Country)
	assert.Equal(t, "Saint John", port.Address.Province)
	assert.EqualValues(t, "America/Antigua", port.Address.Timezone)
	require.Len(t, port.Unlocs, 1)
	assert.Equal(t, *portID, port.Unlocs[0])
}
