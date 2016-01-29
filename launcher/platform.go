//+build !darwin,!windows

package launcher

var (
	// DefaultBinary is the default binary used for linux and unix systems.
	DefaultBinary = []string{"xdg-open"}
)
