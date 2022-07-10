package domain

import (
	"fmt"
	"math"
	"time"
)

type PortID string

func NewPortID(val string) (*PortID, error) {
	if val == "" {
		return nil, fmt.Errorf("%w: port ID cannot be empty", ErrInvalidArgument)
	}

	id := PortID(val)

	return &id, nil
}

type Coordinates struct {
	latitude  float64
	longitude float64
}

func NewCoordinates(latitude float64, longitude float64) (*Coordinates, error) {
	if math.Abs(latitude) > 90 {
		return nil, fmt.Errorf("%w: invalid latitude: %v", ErrInvalidArgument, latitude)
	}

	if math.Abs(longitude) > 180 {
		return nil, fmt.Errorf("%w: invalid longitude: %v", ErrInvalidArgument, longitude)
	}

	return &Coordinates{
		latitude:  latitude,
		longitude: longitude,
	}, nil
}

type Timezone string

func NewTimezone(val string) (*Timezone, error) {
	_, err := time.LoadLocation(val)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid timezone: %s", ErrInvalidArgument, val)
	}

	tz := Timezone(val)

	return &tz, nil
}

type Address struct {
	City        string
	Country     string
	Province    string
	Coordinates *Coordinates
	Timezone    Timezone
}

func NewAddress(city string, country string, province string, coordinates *Coordinates, timezone Timezone) *Address {
	return &Address{
		City:        city,
		Country:     country,
		Province:    province,
		Coordinates: coordinates,
		Timezone:    timezone,
	}
}
