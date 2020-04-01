package loadcnf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

var once, onceReload *sync.Once


//LoadDataConfig reads information from ./config/database_init.cnf and returns AllDataConfig struct
func LoadDataConfig() *AllDataConfig {
	onceReload.Do(func() {
		if timeReload() {
			refreshOnce(once)
		}
	})

	defer refreshOnce(onceReload)
	once.Do(readDataConfig)

	return AllDataCnf
}

func refreshOnce(refOnce *sync.Once) {
	refOnce = new(sync.Once)
}

func readDataConfig() {
	file, err := ioutil.ReadFile(DataInitCnfPath)
	err = json.Unmarshal([]byte(file), AllDataCnf)

	AllDataCnf.LastRead = time.Now()
	AllDataCnf.HasCnf = true

	if err != nil {
		AllDataCnf = nil
	}
}

func timeReload() bool {
	info, _ := os.Stat(DataInitCnfPath)
	lastMod := info.ModTime()
	timeDiff := lastMod.Sub(AllDataCnf.LastRead)
	if timeDiff > 0 {
		return true
	}
	return false
}
