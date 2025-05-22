package box

import cb "github.com/sagernet/sing-box/experimental/commonbox"

// SetupOptions defines parameters for setting up global paths and user info.
type SetupOptions = cb.SetupOptions

// Setup initializes package-level variables for base, working, and temp paths.
// It delegates to experimental/commonbox.Setup for the actual implementation.
func Setup(options *SetupOptions) error {
	return cb.Setup(options)
}
