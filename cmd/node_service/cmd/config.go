/**
*  @file
*  @copyright defined in scan-api/LICENSE
 */

package cmd

import (
	"encoding/json"
	"io/ioutil"

	"github.com/seeleteam/scan-api/node"
)

var (
	configFile *string
)

// LoadConfigFromFile unmarshal config from a file
func LoadConfigFromFile(filepath string) (node.Config, error) {
	var config node.Config
	buff, err := ioutil.ReadFile(filepath)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(buff, &config)

	return config, err
}
