package loadcnf

import(
	"../utilities"
	"strings"
)


// allDatabaseConfig includes database_init file and config information in it
// including directory path which all databases are and directory of ledger databases
// and local databases and their config file path
type allDatabaseConfig struct {
	DatabaseInitPath string
	AllDatabaseInfo  map[string]string
	LedgerUseDB      string
	LocalUseDB       string
	LocalDBList      []string
	LedgerDBList     []string
	EachDBCnf        map[string]string
 	// if directory for storing is has not been defined yet HasCnf is false
	HasCnf bool
}

var allDataCnf = new(allDatabaseConfig)
var lineTypePlace = 0
var linePathPlace = 2
var minCnfLines = 1

// LoadDatabaseConfig reads information from ./config/database_init.cnf and returns allDatabaseConfig struct
// if refresh is True reload data
func LoadDatabaseConfig(refresh bool) *allDatabaseConfig {
	if allDataCnf != nil && !refresh{
		return allDataCnf
	}
	setConstantConfigs()
	cnfLines := utilities.ReturnFileLines(allDataCnf.DatabaseInitPath)
	allDataCnf.AllDatabaseInfo = make(map[string]string)

	if len(cnfLines) < minCnfLines {
		allDataCnf.HasCnf = false
	}else{
		for _, i := range cnfLines{
			lineType := strings.Fields(i)[lineTypePlace]
			linePath := strings.Fields(i)[linePathPlace]
			allDataCnf.AllDatabaseInfo[lineType] = linePath
		}
		allDataCnf.LocalUseDB, allDataCnf.LocalDBList = returnDBLists(allDataCnf.AllDatabaseInfo["loc_cnf"])
		allDataCnf.LedgerUseDB, allDataCnf.LedgerDBList = returnDBLists(allDataCnf.AllDatabaseInfo["leg_cnf"])
	}

	return allDataCnf
}

func setConstantConfigs(){
	allDataCnf.EachDBCnf = make(map[string]string)
	allDataCnf.DatabaseInitPath = "config/database_init.cnf"
	allDataCnf.EachDBCnf["configDir"] = "/config"
	allDataCnf.EachDBCnf["dataDir"] = "/data"
	allDataCnf.EachDBCnf["collectionDir"] = "/data/collection"
	allDataCnf.EachDBCnf["logDir"] = "/logs"
	allDataCnf.EachDBCnf["ConfigDatabaseFile"] = "/configdatabase_config.cnf"
	allDataCnf.EachDBCnf["collectionListFile"] = "/datacollection_list.cnf"
}

func returnDBLists(DBPath string)(useDB string, DBLists []string){
	DBLines := utilities.ReturnFileLines(DBPath)
	useDatabasePlace := 1
	DataBaseListPlace := 3

	useDB = DBLines[useDatabasePlace]
	for _, i := range DBLines[DataBaseListPlace:]{
		if i != ""{
			DBLists = append(DBLists, i)
		}
	}
	return useDB, DBLists
}
