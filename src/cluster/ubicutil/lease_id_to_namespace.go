package ubicUtil

import (
	"crypto/sha256"
	"encoding/base32"
	ctype "providerService/src/cluster/types/v1"
	"strings"
)

// LeaseIDToNamespace generates a unique sha256 sum for identifying a provider's object name.
func LeaseIDToNamespace(lid ctype.LeaseID) string {
	path := lid.String()
	// DNS-1123 label must consist of lower case alphanumeric characters or '-',
	// and must start and end with an alphanumeric character
	// (e.g. 'my-name',  or '123-abc', regex used for validation
	// is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')
	sha := sha256.Sum224([]byte(path))
	return strings.ToLower(base32.HexEncoding.WithPadding(base32.NoPadding).EncodeToString(sha[:]))
}
