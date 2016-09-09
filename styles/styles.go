package styles

import (
	log "github.com/Sirupsen/logrus"

	"github.com/shad7/gochlog/types"
)

// StyleConfigHook provides a callback function to create a specific Sytle
type StyleConfigHook func(interface{}) (types.Styler, error)

var availableStyles = []string{}
var configHooks = map[string]StyleConfigHook{}

// RegisterStyler registers a styler config hook for a config key
func RegisterStyler(name string, hook StyleConfigHook) {
	log.Debugf("Style hook: %s", name)
	configHooks[name] = hook
	availableStyles = append(availableStyles, name)
}

// GetStylers provides a list of available Stylers
func GetStylers() []string {
	return availableStyles
}

// GetStyler provides a specific Styler instance
func GetStyler(name string) (types.Styler, error) {
	if hook, ok := configHooks[name]; ok {
		return hook(nil)
	}
	return nil, nil
}
