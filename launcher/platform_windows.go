//+build windows

package launcher

var (
	// DefaultBinary is the default binary used for windows.
	DefaultBinary = []string{"rundll32", "url.dll,FileProtocolHandler"}
)
