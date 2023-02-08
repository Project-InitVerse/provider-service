package rest

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	dtypes "github.com/ovrclk/akash/x/deployment/types/v1beta2"

	ctypes "providerService/src/cluster/types/v1"
)

type contextKey int

const (
	leaseContextKey contextKey = iota + 1
	deploymentContextKey
	logFollowContextKey
	tailLinesContextKey
	serviceContextKey
	ownerContextKey
	providerContextKey
	servicesContextKey
)

func requestLeaseID(req *http.Request) ctypes.LeaseID {
	return context.Get(req, leaseContextKey).(ctypes.LeaseID)
}

func requestLogFollow(req *http.Request) bool {
	return context.Get(req, logFollowContextKey).(bool)
}

func requestLogTailLines(req *http.Request) *int64 {
	return context.Get(req, tailLinesContextKey).(*int64)
}

func requestService(req *http.Request) string {
	return context.Get(req, serviceContextKey).(string)
}

func requestServices(req *http.Request) string {
	return context.Get(req, servicesContextKey).(string)
}

func requestProvider(req *http.Request) common.Address {
	return context.Get(req, providerContextKey).(common.Address)
}

func requestOwner(req *http.Request) common.Address {
	return context.Get(req, ownerContextKey).(common.Address)
}

func requestDeploymentID(req *http.Request) dtypes.DeploymentID {
	return context.Get(req, deploymentContextKey).(dtypes.DeploymentID)
}

func requireOwner() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.TLS == nil || len(r.TLS.PeerCertificates) == 0 {
				http.Error(w, "", http.StatusUnauthorized)
				return
			}

			// at this point client certificate has been validated
			// so only thing left to do is get account id stored in the CommonName
			owner := common.HexToAddress(r.TLS.PeerCertificates[0].Subject.CommonName)
			fmt.Println(owner.Hex(), r.TLS.PeerCertificates[0].Subject.CommonName)
			if strings.ToLower(owner.Hex()) != strings.ToLower(r.TLS.PeerCertificates[0].Subject.CommonName) {
				http.Error(w, "invalid certificates", http.StatusUnauthorized)
				return
			}

			context.Set(r, ownerContextKey, owner)
			next.ServeHTTP(w, r)
		})
	}
}

func requireLeaseID() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			id, err := parseLeaseID(req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			context.Set(req, leaseContextKey, id)
			next.ServeHTTP(w, req)
		})
	}
}

func requireService() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			vars := mux.Vars(req)

			svc := vars["serviceName"]
			if svc == "" {
				http.Error(w, "empty service name", http.StatusBadRequest)
				return
			}

			context.Set(req, serviceContextKey, svc)
			next.ServeHTTP(w, req)
		})
	}
}

func parseDeploymentID(req *http.Request) (dtypes.DeploymentID, error) {
	var parts []string
	parts = append(parts, requestOwner(req).String())
	parts = append(parts, mux.Vars(req)["dseq"])
	return dtypes.ParseDeploymentPath(parts)
}

func parseLeaseID(req *http.Request) (ctypes.LeaseID, error) {
	vars := mux.Vars(req)

	oseq, err := strconv.ParseInt(vars["oseq"], 10, 64)
	if err != nil {
		return ctypes.LeaseID{}, err
	}

	return ctypes.LeaseID{Owner: requestOwner(req).String(), OSeq: uint64(oseq), Provider: requestProvider(req).String()}, nil
}

func requestStreamParams() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			vars := req.URL.Query()

			var err error

			defer func() {
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}()

			var tailLines *int64

			services := vars.Get("service")
			if strings.HasSuffix(services, ",") {
				err = errors.Errorf("parameter \"service\" must not contain trailing comma")
				return
			}

			follow := false

			if val := vars.Get("follow"); val != "" {
				follow, err = strconv.ParseBool(val)
				if err != nil {
					return
				}
			}

			vl := new(int64)
			if val := vars.Get("tail"); val != "" {
				*vl, err = strconv.ParseInt(val, 10, 32)
				if err != nil {
					return
				}

				if *vl < -1 {
					err = errors.Errorf("parameter \"tail\" contains invalid value")
					return
				}
			} else {
				*vl = -1
			}

			if *vl > -1 {
				tailLines = vl
			}

			context.Set(req, logFollowContextKey, follow)
			context.Set(req, tailLinesContextKey, tailLines)
			context.Set(req, servicesContextKey, services)

			next.ServeHTTP(w, req)
		})
	}
}
