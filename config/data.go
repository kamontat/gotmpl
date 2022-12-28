package config

import (
	"fmt"

	"github.com/kc-workspace/go-lib/mapper"
)

type data struct {
	underlay map[string]map[string]interface{}
}

func (d *data) GetMergedData() map[string]interface{} {
	var result = make(map[string]interface{})

	for _, mapping := range d.underlay {
		result = mapper.Merge(result, mapping, make(mapper.Mapper))
	}

	return result
}

func (d *data) String() string {
	return fmt.Sprintf(`%v`, d.underlay)
}
