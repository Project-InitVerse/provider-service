package rest

import (
	"fmt"
	v1 "providerService/src/ubicpkg/api/ubicnet/v1"
)

const (
	deploymentPathPrefix = "/deployment/{dseq}"
	leasePathPrefix      = "/lease/{oseq}"
	hostnamePrefix       = "/hostname"
	endpointPrefix       = "/endpoint"
	migratePathPrefix    = "/migrate"
)

func versionPath() string {
	return "version"
}

func statusPath() string {
	return "status"
}

func validatePath() string {
	return "validate"
}

func leasePath(id v1.LeaseID) string {
	return fmt.Sprintf("lease/%d", id.OSeq)
}

func submitManifestPath(dseq uint64) string {
	return fmt.Sprintf("deployment/%d/manifest", dseq)
}

func leaseStatusPath(id v1.LeaseID) string {
	return fmt.Sprintf("%s/status", leasePath(id))
}

func leaseShellPath(lID v1.LeaseID) string {
	return fmt.Sprintf("%s/shell", leasePath(lID))
}

func leaseEventsPath(id v1.LeaseID) string {
	return fmt.Sprintf("%s/kubeevents", leasePath(id))
}

func serviceStatusPath(id v1.LeaseID, service string) string {
	return fmt.Sprintf("%s/service/%s/status", leasePath(id), service)
}

func serviceLogsPath(id v1.LeaseID) string {
	return fmt.Sprintf("%s/logs", leasePath(id))
}
