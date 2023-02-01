package rest

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	manifest "github.com/ovrclk/akash/manifest/v2beta1"
	dtypes "github.com/ovrclk/akash/x/deployment/types/v1beta2"
	"github.com/pkg/errors"
	"io"
	"k8s.io/client-go/tools/remotecommand"
	"providerService/src"
	cltypes "providerService/src/cluster/types/v1"
	v1 "providerService/src/ubicpkg/api/ubicnet/v1"
)

const (
	schemeWSS   = "wss"
	schemeHTTPS = "https"
)

// Client defines the methods available for connecting to the gateway server.
type Client interface {
	Status(ctx context.Context) (*provider.Status, error)
	Validate(ctx context.Context, gspec dtypes.GroupSpec) (provider.ValidateGroupSpecResult, error)
	SubmitManifest(ctx context.Context, dseq uint64, mani manifest.Manifest) error
	LeaseStatus(ctx context.Context, id v1.LeaseID) (LeaseStatus, error)
	LeaseEvents(ctx context.Context, id v1.LeaseID, services string, follow bool) (*LeaseKubeEvents, error)
	LeaseLogs(ctx context.Context, id v1.LeaseID, services string, follow bool, tailLines int64) (*ServiceLogs, error)
	ServiceStatus(ctx context.Context, id v1.LeaseID, service string) (*cltypes.ServiceStatus, error)
	LeaseShell(ctx context.Context, id v1.LeaseID, service string, podIndex uint, cmd []string,
		stdin io.ReadCloser,
		stdout io.Writer,
		stderr io.Writer,
		tty bool,
		tsq <-chan remotecommand.TerminalSize) error
	MigrateHostnames(ctx context.Context, hostnames []string, dseq uint64, gseq uint32) error
	MigrateEndpoints(ctx context.Context, endpoints []string, dseq uint64, gseq uint32) error
}

// JwtClient is interface
type JwtClient interface {
	GetJWT(ctx context.Context) (*jwt.Token, error)
}

// LeaseKubeEvent stores action and msg
type LeaseKubeEvent struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

// ServiceLogMessage stores name and msg
type ServiceLogMessage struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// LeaseKubeEvents is struct
type LeaseKubeEvents struct {
	Stream  <-chan cltypes.LeaseEvent
	OnClose <-chan string
}

// ServiceLogs is struct
type ServiceLogs struct {
	Stream  <-chan ServiceLogMessage
	OnClose <-chan string
}

var errRequiredCertSerialNum = errors.New("cert_serial_number must be present in claims")
var errNonNumericCertSerialNum = errors.New("cert_serial_number must be numeric in claims")

//ClientResponseError stores status msg
type ClientResponseError struct {
	Status  int
	Message string
}

func (err ClientResponseError) Error() string {
	return fmt.Sprintf("remote server returned %d", err.Status)
}

//ClientError get client error
func (err ClientResponseError) ClientError() string {
	return fmt.Sprintf("Remote Server returned %d\n%s", err.Status, err.Message)
}
