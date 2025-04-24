package dreamhost

import (
	"context"
	"fmt"
	"github.com/libdns/libdns"
)

// CompatProvider wraps the main Provider to ensure interface compatibility
type CompatProvider struct {
	*Provider
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

	return origProvider.AppendRecords(ctx, zone, records)
}

func (p *CompatProvider) SetRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	origProvider := p.Provider
	if origProvider == nil {
		return nil, fmt.Errorf("provider is nil")
	}

	return origProvider.SetRecords(ctx, zone, records)
}

func (p *CompatProvider) DeleteRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	origProvider := p.Provider
	if origProvider == nil {
		return nil, fmt.Errorf("provider is nil")
	}

	return origProvider.DeleteRecords(ctx, zone, records)
}
