package utils

import (
	"github.com/cosmos/cosmos-sdk/version"
)

//UbicVersionInfo is struct
type UbicVersionInfo struct {
	Version          string `json:"version"`
	GitCommit        string `json:"commit"`
	BuildTags        string `json:"buildTags"`
	GoVersion        string `json:"go"`
	CosmosSdkVersion string `json:"cosmosSdkVersion"`
}

//NewUbicVersionInfo create UbicVersionInfo
func NewUbicVersionInfo() UbicVersionInfo {
	verInfo := version.NewInfo()
	return UbicVersionInfo{
		Version:          verInfo.Version,
		GitCommit:        verInfo.GitCommit,
		BuildTags:        verInfo.BuildTags,
		GoVersion:        verInfo.GoVersion,
		CosmosSdkVersion: verInfo.CosmosSdkVersion,
	}
}
