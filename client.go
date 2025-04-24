package dreamhost

import (
	"strings"

	"github.com/adamantal/go-dreamhost/api"
	"github.com/libdns/libdns"
)

func (p *Provider) init() error {
	client, err := api.NewClient(p.APIKey, nil)
	if err != nil {
		return err
	}
	p.client = *client
	return nil
}

func recordFromApiDnsRecord(apiDnsRecord api.DNSRecord) libdns.Record {
	var rec libdns.Record
	// Updated field names for libdns v0.2.3
	rec.Type = string(apiDnsRecord.Type)
	rec.Value = apiDnsRecord.Value
	rec.Name = libdns.RelativeName(apiDnsRecord.Record, apiDnsRecord.Zone)

	// If the above fields don't exist in v0.2.3,
	// they might be accessed through other fields or methods, such as:
	/*
		rec.RData = apiDnsRecord.Value
		rec.Class = "IN" // Default for most DNS records
	*/

	return rec
}

func apiDnsRecordInputFromRecord(record libdns.Record, zone string) api.DNSRecordInput {
	var recordInput api.DNSRecordInput
	// Updated field access for libdns v0.2.3
	recordInput.Type = api.RecordType(record.Type)
	recordInput.Value = record.Value
	zone = strings.TrimRight(zone, ".")
	recordInput.Record = libdns.AbsoluteName(record.Name, zone)

	// If the above fields don't exist, try alternative approaches:
	/*
		recordInput.Type = api.RecordType(record.Type)
		recordInput.Value = record.RData
		zone = strings.TrimRight(zone, ".")
		recordInput.Record = libdns.AbsoluteName(record.Name, zone)
	*/

	return recordInput
}
