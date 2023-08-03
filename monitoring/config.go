package monitoring

import "github.com/neyromanser/wrapperapp/monitoring/adapters"

type ConfigMinotiring struct {
	Use        bool                     `yaml:"use"`
	Compontens []string                 `yaml:"components"`
	Events     []string                 `yaml:"events"`
	Adapters   []adapters.ConfigAdapter `yaml:"adapters"`
}
