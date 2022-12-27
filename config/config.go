package config

import "fmt"

type Config struct {
	data    *data
	Setting *Setting
}

func (c *Config) GetData() map[string]interface{} {
	return c.data.GetMergedData()
}

func (c *Config) String() string {
	return fmt.Sprintf(`%v%v`, c.data, c.Setting)
}
