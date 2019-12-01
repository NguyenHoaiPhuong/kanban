package mdb

// MongoConfig includes configurations for Mongo
type MongoConfig struct {
	Host     *string
	Port     *string
	UserName *string
	Password *string
	DBName   *string
}

// Copy : copy values of all fields in the existing config and return a new config
func (cfg *MongoConfig) Copy() *MongoConfig {
	newCfg := &MongoConfig{
		Host:     new(string),
		Port:     new(string),
		UserName: new(string),
		Password: new(string),
		DBName:   new(string),
	}
	*newCfg.Host = *cfg.Host
	*newCfg.Port = *cfg.Port
	*newCfg.UserName = *cfg.UserName
	*newCfg.Password = *cfg.Password
	*newCfg.DBName = *cfg.DBName

	return newCfg
}
