package loadcnf

import (
	"sync"
)

type DatabaseConfig struct {
}

var databaseCnf *DatabaseConfig
var once sync.Once

func GetDatabaseConfig() *DatabaseConfig {
	once.Do(func() {
		databaseCnf = &DatabaseConfig{}
	})
	return databaseCnf
}
