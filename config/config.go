package config

type Config struct {
	Data    *data
	Setting *Setting
}

func (c *Config) GetData() map[string]interface{} {
	return c.Data.GetMergedData()
}
