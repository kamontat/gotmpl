package config

import (
	"fmt"

	"github.com/kamontat/gotmpl/utils"
)

type data struct {
	underlay map[string]map[interface{}]interface{}
}

func (d *data) GetMergedData() map[interface{}]interface{} {
	var result = make(map[interface{}]interface{})

	for _, mapper := range d.underlay {
		result = utils.DeepMergeMap(result, mapper)
	}

	return result
}

func (d *data) String() string {
	return fmt.Sprintf(`data: %v
`, d.underlay)
}
