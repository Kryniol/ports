package domain

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testFirstPort = Port{
		ID:   "test_port_1",
		Name: "Test port 1",
		Address: Address{
			City:     "City 1",
			Country:  "Country 1",
			Province: "Province 1",
			Coordinates: &Coordinates{
				latitude:  12,
				longitude: 34,
			},
			Timezone: "Europe/Warsaw",
		},
		Unlocs: []PortID{
			"test_port_1",
		},
		Code: "test_code_1",
	}
	testSecondPort = Port{
		ID:   "test_port_2",
		Name: "Test port 2",
		Address: Address{
			City:     "City 2",
			Country:  "Country 2",
			Province: "Province 2",
			Coordinates: &Coordinates{
				latitude:  56,
				longitude: 78,
			},
			Timezone: "Europe/Warsaw",
		},
		Unlocs: []PortID{
			"test_port_2",
		},
		Code: "test_code_2",
	}
)

func TestPortService_SavePorts_Successful(t *testing.T) {
	ctx := context.Background()
	ch := make(chan Port, 2)
	ch <- testFirstPort
	ch <- testSecondPort
	close(ch)

	reader := &mockPortReader{}
	reader.On("Read", ctx).Return(ch, nil)

	repo := &mockPortRepository{}
	repo.On("Save", ctx, testFirstPort).Return(nil)
	repo.On("Save", ctx, testSecondPort).Return(nil)

	svc := NewPortService(reader, repo)
	err := svc.SavePorts(ctx)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestPortService_SavePorts_ReaderReturnsError_WrappedErrorIsReturned(t *testing.T) {
	ctx := context.Background()
	rErr := errors.New("test error")

	reader := &mockPortReader{}
	reader.On("Read", ctx).Return(nil, rErr)

	svc := NewPortService(reader, &mockPortRepository{})
	err := svc.SavePorts(ctx)

	assert.ErrorIs(t, err, rErr)
}

func TestPortService_SavePorts_RepositoryReturnsError_WrappedErrorIsReturned(t *testing.T) {
	ctx := context.Background()
	rErr := errors.New("test error")
	ch := make(chan Port, 2)
	ch <- testFirstPort
	ch <- testSecondPort
	close(ch)

	reader := &mockPortReader{}
	reader.On("Read", ctx).Return(ch, nil)

	repo := &mockPortRepository{}
	repo.On("Save", ctx, testFirstPort).Return(nil)
	repo.On("Save", ctx, testSecondPort).Return(rErr)

	svc := NewPortService(reader, repo)
	err := svc.SavePorts(ctx)

	assert.ErrorIs(t, err, rErr)
}
