package charger

import (
	"fmt"

	"github.com/evcc-io/evcc/api"
)

func init() {
	registry.Add("nrgkick", NewNRGKickFromConfig)
}

func NewNRGKickFromConfig(other map[string]interface{}) (api.Charger, error) {
	return nil, fmt.Errorf("not implemented")
}
