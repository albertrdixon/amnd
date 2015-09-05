// Package amnd defines logic for updating Plex
package amnd

import "fmt"

var (
	version = "v0.0.1"

	// SHA is the build SHA
	SHA string
)

// Version returns the current software version
func Version() string {
	if SHA != "" {
		return fmt.Sprintf("%s-%s", version, SHA)
	}
	return version
}
