package config

import (
	"fmt"

	"github.com/kamontat/gotmpl/maps"
)

type data struct {
	underlay map[string]map[string]interface{}
}

func (d *data) GetMergedData() map[string]interface{} {
	var result = make(map[string]interface{})

	for _, mapper := range d.underlay {
		result = maps.Merge(result, mapper, make(maps.Mapper))
	}

	return result
}

func (d *data) String() string {
	return fmt.Sprintf(`data: %v
`, d.underlay)
}
