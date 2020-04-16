package loadcnf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type configLoad interface {
	updateLastRead()
}

type onceConfig struct{
	cnfLoad configLoad
	once *sync.Once
	onceReload *sync.Once
}

var(
	cnfMap = map[string]onceConfig{
		"allData": {cnfLoad:AllDataCnf, once: &sync.Once{}, onceReload: &sync.Once{}},
		"localDb": {cnfLoad:LocalDbCnf, once: &sync.Once{}, onceReload: &sync.Once{}}}
)


func loadConfigOnce(onceCnf onceConfig, filePath string, lastRead time.Time) (err error){

	onceCnf.onceReload.Do(func() {
		var needReload bool
		needReload, err = needTimeReload(filePath, lastRead)
		if needReload{
			refreshOnce(onceCnf.once)
		}
	})
	if err != nil{
		return err
	}

	defer refreshOnce(onceCnf.onceReload)
	onceCnf.once.Do(func() {
		err = readDBaseConfig(filePath, onceCnf.cnfLoad)
	})

	return err
}


func readDBaseConfig(filePath string, cLoad configLoad) (err error){
	file, err := ioutil.ReadFile(filePath)
	if err != nil{
		return err
	}
	err = json.Unmarshal([]byte(file), cLoad)

	cLoad.updateLastRead()
	return err
}

func refreshOnce(refOnce *sync.Once) {
	refOnce = new(sync.Once)
}

func needTimeReload(filePath string, lastRead time.Time) (needReload bool,err error) {
	info, err := os.Stat(filePath)
	if err != nil{
		return false, err
	}

	lastMod := info.ModTime()
	timeDiff := lastMod.Sub(lastRead)
	if timeDiff > 0 {
		return true, nil
	}
	return false, nil
}
