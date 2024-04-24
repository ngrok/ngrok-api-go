// Code generated for API Clients. DO NOT EDIT.

package ngrok

import "runtime/debug"

const modulePath = "github.com/ngrok/ngrok-api-go"

var _version string

func init() {
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range buildInfo.Deps {
			if dep.Path == modulePath {
				_version = dep.Version
			}
		}
	}
	_version = "unknown"
}
