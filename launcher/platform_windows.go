//+build windows

package launcher

var (
	DefaultBinary = []string{"rundll32", "url.dll,FileProtocolHandler"}
)
