package utils

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"providerService/src/config"
	"providerService/src/util"
	"time"

	"github.com/pkg/errors"
)

// NewServerTLSConfig create
func NewServerTLSConfig(ctx context.Context, certs []tls.Certificate, cConfig *config.ProviderConfig) (*tls.Config, error) {
	// InsecureSkipVerify is set to true due to inability to use normal TLS verification
	// certificate validation and authentication performed in VerifyPeerCertificate
	cfg := &tls.Config{
		Certificates:       certs,
		ClientAuth:         tls.RequestClientCert,
		InsecureSkipVerify: true, // nolint: gosec
		MinVersion:         tls.VersionTLS13,
		VerifyPeerCertificate: func(certificates [][]byte, _ [][]*x509.Certificate) error {

			if len(certificates) > 0 {
				if len(certificates) != 1 {
					return errors.Errorf("tls: invalid certificate chain")
				}

				cert, err := x509.ParseCertificate(certificates[0])
				if err != nil {
					return errors.Wrap(err, "tls: failed to parse certificate")
				}

				// validation
				var owner common.Address
				owner = common.HexToAddress(cert.Subject.CommonName)

				// 1. CommonName in issuer and Subject must match and be as Bech32 format
				if cert.Subject.CommonName != cert.Issuer.CommonName {
					return errors.Wrap(err, "tls: invalid certificate's issuer common name")
				}

				// 2. serial number must be in
				if cert.SerialNumber == nil {
					return errors.Wrap(err, "tls: invalid certificate serial number")
				}

				// 3. look up certificate on chain
				fmt.Println("cert raw", base64.StdEncoding.EncodeToString(cert.Raw))
				//Todo open chain check cert
				conn, err := ethclient.Dial(cConfig.NodeURL)
				if false {
					defer conn.Close()
					certHandle, err := util.NewCert(common.HexToAddress(cConfig.Cert), conn)
					if err != nil {

						return errors.Wrap(err, "tls: Unable to connect to the chain")
					}
					certState, err := certHandle.UserCertState(nil, owner, string(cert.Raw))
					if err != nil {
						return errors.Wrap(err, "chain: Unable to connect to the chain")
					}
					if certState != 1 {
						return errors.New("tls: attempt to use non-existing or revoked certificate")
					}

				}
				//if err != nil {
				//
				//	return errors.Wrap(err, "chain: Unable to connect to the chain")
				//}

				clientCertPool := x509.NewCertPool()
				clientCertPool.AddCert(cert)

				opts := x509.VerifyOptions{
					Roots:                     clientCertPool,
					CurrentTime:               time.Now(),
					KeyUsages:                 []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
					MaxConstraintComparisions: 0,
				}

				if _, err = cert.Verify(opts); err != nil {
					return errors.Wrap(err, "tls: unable to verify certificate")
				}
			}
			return nil
		},
	}

	return cfg, nil
}
