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

type LeaseID struct {
	Owner    string `json:"owner" yaml:"owner"`
	OSeq     uint64 `json:"oseq" yaml:"oseq"`
	Provider string `json:"provider" yaml:"provider"`
}

func (lid LeaseID) String() string {
	return fmt.Sprintf("%v/%v", lid.OSeq, lid.Provider)
}

func LeasePath(id LeaseID) string {
	return fmt.Sprintf("%s/%s/%s/%v/%s", leasePath, orderPath, id.Owner, id.OSeq, id.Provider)
}
func IngressHost(lid LeaseID, svcName string) string {
	uid := uuid.NewV5(uuid.NamespaceDNS, lid.String()+svcName).Bytes()
	return strings.ToLower(base32.HexEncoding.WithPadding(base32.NoPadding).EncodeToString(uid))
}
