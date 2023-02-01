package v1

import (
	"encoding/base32"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
)

const (
	ordersPath = "orders"
	orderPath  = "order"
	bidsPath   = "bids"
	bidPath    = "bid"
	leasesPath = "leases"
	leasePath  = "lease"
)

// LeaseID is struct
type LeaseID struct {
	Owner    string `json:"owner" yaml:"owner"`
	OSeq     uint64 `json:"oseq" yaml:"oseq"`
	Provider string `json:"provider" yaml:"provider"`
}

// String is function get lease string type
func (lid LeaseID) String() string {
	return fmt.Sprintf("%v/%v", lid.OSeq, lid.Provider)
}

// LeasePath is function get LeasePath
func LeasePath(id LeaseID) string {
	return fmt.Sprintf("%s/%s/%s/%v/%s", leasePath, orderPath, id.Owner, id.OSeq, id.Provider)
}

// IngressHost is function get IngressHost
func IngressHost(lid LeaseID, svcName string) string {
	uid := uuid.NewV5(uuid.NamespaceDNS, lid.String()+svcName).Bytes()
	return strings.ToLower(base32.HexEncoding.WithPadding(base32.NoPadding).EncodeToString(uid))
}

// Equals is function get lease is equal
func (lid LeaseID) Equals(olid LeaseID) bool {
	return lid.OSeq == olid.OSeq && strings.ToLower(lid.Owner) == strings.ToLower(olid.Owner) && strings.ToLower(lid.Provider) == strings.ToLower(lid.Provider)
}
