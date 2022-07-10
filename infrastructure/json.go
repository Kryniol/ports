package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/Kryniol/ports/domain"
)

type jsonReader struct {
	inputPath string
	logger    *zap.Logger
}

func NewJSONReader(path string, logger *zap.Logger) *jsonReader {
	return &jsonReader{
		inputPath: path,
		logger:    logger,
	}
}

func (r *jsonReader) Read(ctx context.Context) (<-chan domain.Port, error) {
	f, err := os.Open(r.inputPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't open input CSV file, %w", err)
	}

	d := json.NewDecoder(f)
	portsCh := make(chan domain.Port)

	go func() {
		<-ctx.Done()
		_ = f.Close()
	}()

	go func() {
		defer func() {
			_ = f.Close()
			close(portsCh)
		}()

		_, err := d.Token()
		if err != nil {
			r.logger.Error("got an error when reading first JSON token", zap.Error(err))
		}

		for d.More() {
			port, err := r.parsePort(d)
			if err != nil {
				r.logger.Error("couldn't create port entity from JSON data", zap.Error(err))
				continue
			}

			portsCh <- *port
		}

		_, err = d.Token()
		if err != nil {
			r.logger.Error("got an error when reading last JSON token", zap.Error(err))
		}
	}()

	return portsCh, nil
}

func (r *jsonReader) parsePort(d *json.Decoder) (*domain.Port, error) {
	idxToken, err := d.Token()
	if err != nil {
		return nil, fmt.Errorf("got an error when reading array index from JSON: %w", err)
	}

	portID, err := domain.NewPortID(idxToken.(string))
	if err != nil {
		return nil, fmt.Errorf("got invalid port ID from JSON data: %w", err)
	}

	var dto portDTO
	err = d.Decode(&dto)
	if err != nil {
		return nil, fmt.Errorf("couldn't unmarshall JSON data to DTO object: %w", err)
	}

	port, err := dto.toEntity(*portID)
	if err != nil {
		return nil, fmt.Errorf("couldn't create port entity from JSON object: %w", err)
	}

	return port, nil
}

type portDTO struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

func (p *portDTO) toEntity(id domain.PortID) (*domain.Port, error) {
	var unlocs []domain.PortID
	for _, u := range p.Unlocs {
		uID, err := domain.NewPortID(u)
		if err != nil {
			return nil, fmt.Errorf("got invalid port ID in unlocs JSON field: %w", err)
		}

		unlocs = append(unlocs, *uID)
	}

	var (
		coordinates *domain.Coordinates
		err         error
	)
	if len(p.Coordinates) == 2 {
		coordinates, err = domain.NewCoordinates(p.Coordinates[1], p.Coordinates[0])
		if err != nil {
			return nil, fmt.Errorf("got invalid coordinates from JSON: %w", err)
		}
	}

	tz, err := domain.NewTimezone(p.Timezone)
	if err != nil {
		return nil, fmt.Errorf("got invalid timezone from JSON: %w", err)
	}

	address := domain.NewAddress(p.City, p.Country, p.Province, coordinates, *tz)

	if len(p.Regions) > 0 {
		panic(id)
	}

	return domain.NewPort(id, p.Name, *address, p.Alias, p.Regions, unlocs, p.Code), nil
}
