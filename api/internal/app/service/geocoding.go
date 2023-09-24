package service

import (
	"context"
	"otomo/internal/app/interfaces/svc"
	"otomo/internal/pkg/errs"

	"googlemaps.github.io/maps"
)

// TODO: Add tests

var _ svc.GeocodingService = (*GeocodingService)(nil)

type GeocodingService struct {
	client *maps.Client
}

func NewGeocodingService(client *maps.Client) *GeocodingService {
	return &GeocodingService{
		client: client,
	}
}

func (gs *GeocodingService) One(
	ctx context.Context,
	req *maps.GeocodingRequest,
) (*maps.GeocodingResult, error) {
	resp, err := gs.client.Geocode(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(resp) > 0 {
		return &resp[0], nil
	}

	return nil, &errs.Error{
		Cause:  errs.CauseNotFound,
		Domain: errs.DomainGeo,
		Field:  errs.FieldNone,
	}
}

func (gs *GeocodingService) List(
	ctx context.Context,
	reqs []*maps.GeocodingRequest,
) ([]*maps.GeocodingResult, error) {
	var (
		results = []*maps.GeocodingResult{}
	)
	for _, req := range reqs {
		result, err := gs.One(ctx, req)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}
