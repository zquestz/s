//+build !darwin,!windows

package launcher

var (
	DefaultBinary = []string{"xdg-open"}
)
