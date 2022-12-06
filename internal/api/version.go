package api

import "runtime/debug"

const modulePath = "github.com/ngrok/ngrok-api-go"

var Version string

func init() {
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range buildInfo.Deps {
			if dep.Path == modulePath {
				Version = dep.Version
			}
		}
	}
	Version = "unknown"
}
