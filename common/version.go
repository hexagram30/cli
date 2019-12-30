package common

import (
	"encoding/json"

	"github.com/geomyidia/util/version"
	rasterapi "github.com/hexagram30/raster/api"
	log "github.com/sirupsen/logrus"
)

// Versioning data
var (
	vsn        string
	buildDate  string
	gitCommit  string
	gitBranch  string
	gitSummary string
)

// VersionData stuff for things
func VersionData() *version.ProjectVersion {
	return &version.ProjectVersion{
		Semantic:   vsn,
		BuildDate:  buildDate,
		GitCommit:  gitCommit,
		GitBranch:  gitBranch,
		GitSummary: gitSummary,
	}
}

// BuildString ...
func BuildString() string {
	return version.BuildString(VersionData())
}

// VersionString ...
func VersionString() string {
	return version.String(VersionData())
}

// VersionedBuildString ...
func VersionedBuildString() string {
	return version.VersionedBuildString(VersionData())
}

// RPCVersionToJSON ...
func RPCVersionToJSON(structData *rasterapi.VersionReply) string {
	jsonData, err := json.Marshal(structData)
	if err != nil {
		log.Error(err)
		log.Fatalf("Couldn't marshal: %v", structData)
	}
	return string(jsonData)
}
