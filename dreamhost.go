package dreamhost

import (
	"context"
	"fmt"

	"github.com/libdns/dreamhost"
	"github.com/libdns/libdns"
)

type CompatProvider struct {
	*dreamhost.Provider
}

// Make sure our CompatProvider implements these interfaces
var (
	_ libdns.RecordGetter   = (*CompatProvider)(nil)
	_ libdns.RecordAppender = (*CompatProvider)(nil)
	_ libdns.RecordSetter   = (*CompatProvider)(nil)
	_ libdns.RecordDeleter  = (*CompatProvider)(nil)
)

func (p *CompatProvider) GetRecords(ctx context.Context, zone string) ([]libdns.Record, error) {
	origProvider := p.Provider
	if origProvider == nil {
		return nil, fmt.Errorf("provider is nil")
	}

	records, err := origProvider.GetRecords(ctx, zone)
	return records, err
}

func (p *CompatProvider) AppendRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	origProvider := p.Provider
	if origProvider == nil {
		return nil, fmt.Errorf("provider is nil")
	}

	fixedRecords := make([]libdns.Record, len(records))
	for i, rec := range records {
		fixedRecords[i] = libdns.Record{
			ID:    rec.ID,
			Type:  rec.Type,
			Name:  rec.Name,
			Value: rec.Value,
			TTL:   rec.TTL,
		}
	}

	return origProvider.AppendRecords(ctx, zone, fixedRecords)
}

func (p *CompatProvider) SetRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	origProvider := p.Provider
	if origProvider == nil {
		return nil, fmt.Errorf("provider is nil")
	}

	fixedRecords := make([]libdns.Record, len(records))
	for i, rec := range records {
		fixedRecords[i] = libdns.Record{
			ID:    rec.ID,
			Type:  rec.Type,
			Name:  rec.Name,
			Value: rec.Value,
			TTL:   rec.TTL,
		}
	}

	return origProvider.SetRecords(ctx, zone, fixedRecords)
}

func (p *CompatProvider) DeleteRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	origProvider := p.Provider
	if origProvider == nil {
		return nil, fmt.Errorf("provider is nil")
	}

	fixedRecords := make([]libdns.Record, len(records))
	for i, rec := range records {
		fixedRecords[i] = libdns.Record{
			ID:    rec.ID,
			Type:  rec.Type,
			Name:  rec.Name,
			Value: rec.Value,
			TTL:   rec.TTL,
		}
	}

	return origProvider.DeleteRecords(ctx, zone, fixedRecords)
}
