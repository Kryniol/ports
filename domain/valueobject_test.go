package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type coordinatesTestCase struct {
	latVal float64
	lonVal float64
}

func TestNewCoordinates_CorrectValues_Successful(t *testing.T) {

	suite := map[string]coordinatesTestCase{
		"min": {
			-90.,
			-180.,
		},
		"max": {
			90.,
			180.,
		},
		"zero": {
			0.,
			0.,
		},
	}

	for name, tc := range suite {
		tc := tc
		t.Run(
			name, func(t *testing.T) {
				t.Parallel()
				coords, err := NewCoordinates(tc.latVal, tc.lonVal)
				assert.NoError(t, err)
				require.NotNil(t, coords)
				assert.Equal(t, tc.latVal, coords.latitude)
				assert.Equal(t, tc.lonVal, coords.longitude)
			},
		)
	}
}

func TestNewCoordinates_Invalid_InvalidArgumentErrorIsReturned(t *testing.T) {
	suite := map[string]coordinatesTestCase{
		"latitude below min": {
			-90.01,
			0,
		},
		"longitude below min": {
			0,
			-180.01,
		},
		"latitude exceeding max": {
			90.01,
			0,
		},
		"longitude exceeding max": {
			0,
			180.01,
		},
	}

	for name, tc := range suite {
		tc := tc
		t.Run(
			name, func(t *testing.T) {
				t.Parallel()
				coords, err := NewCoordinates(tc.latVal, tc.lonVal)
				assert.ErrorIs(t, err, ErrInvalidArgument)
				assert.Nil(t, coords)
			},
		)
	}
}

func TestNewTimezone_CorrectValue_Successful(t *testing.T) {
	suite := []string{
		"Europe/Warsaw",
		"Asia/Dubai",
		"America/New_York",
	}
	for _, val := range suite {
		val := val
		t.Run(
			val, func(t *testing.T) {
				t.Parallel()
				tz, err := NewTimezone(val)
				assert.NoError(t, err)
				require.NotNil(t, tz)
				assert.EqualValues(t, val, *tz)
			},
		)
	}
}

func TestNewTimezone_InvalidValue_InvalidArgumentErrorIsReturned(t *testing.T) {
	suite := []string{
		"Europe/Warsaw",
		"Asia/Dubai",
		"America/New_York",
	}
	for _, val := range suite {
		val := val
		t.Run(
			val, func(t *testing.T) {
				t.Parallel()
				tz, err := NewTimezone(val)
				assert.NoError(t, err)
				require.NotNil(t, tz)
				assert.EqualValues(t, val, *tz)
			},
		)
	}
}
