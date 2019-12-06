package loadcnf

import(
	"../utilities"
	"strings"
)


// allDatabaseConfig includes database_init file and config information in it
// including directory path which all databases are and directory of ledger databases
// and local databases and their config file path
type allDatabaseConfig struct {
	databaseInitPath string
	allDatabaseInfo map[string]string
	ledgerUseDB string
	localUseDB string
	localDBList []string
	ledgerDBList []string
	eachDBCnf map[string]string
 	// if directory for storing is has not been defined yet hasCnf is false
	hasCnf bool
}

var AllDataCnf *allDatabaseConfig
var lineTypePlace = 0
var linePathPlace = 2
var minCnfLines = 1

// LoadDatabaseConfig reads information from ./config/database_init.cnf and returns allDatabaseConfig struct
// if refresh is True reload data
func LoadDatabaseConfig(refresh bool) *allDatabaseConfig {
	if AllDataCnf != nil && !refresh{
		return AllDataCnf
	}
	setConstantConfigs()
	cnfLines := utilities.ReturnFileLines(AllDataCnf.databaseInitPath)
	AllDataCnf.allDatabaseInfo = make(map[string]string)

	if len(cnfLines) < minCnfLines {
		AllDataCnf.hasCnf = false
	}else{
		for _, i := range cnfLines{
			lineType := strings.Fields(i)[lineTypePlace]
			linePath := strings.Fields(i)[linePathPlace]
			AllDataCnf.allDatabaseInfo[lineType] = linePath
		}
		AllDataCnf.localUseDB, AllDataCnf.localDBList = returnDBLists(AllDataCnf.allDatabaseInfo["loc_cnf"])
		AllDataCnf.ledgerUseDB, AllDataCnf.ledgerDBList = returnDBLists(AllDataCnf.allDatabaseInfo["leg_cnf"])
	}

	return AllDataCnf
}

func setConstantConfigs(){
	AllDataCnf.eachDBCnf = make(map[string]string)
	AllDataCnf.databaseInitPath = "/config/database_init.cnf"
	AllDataCnf.eachDBCnf["configDir"] = "/config"
	AllDataCnf.eachDBCnf["dataDir"] = "/data"
	AllDataCnf.eachDBCnf["collectionDir"] = "/data/collection"
	AllDataCnf.eachDBCnf["logDir"] = "/logs"
	AllDataCnf.eachDBCnf["ConfigDatabaseFile"] = "/configdatabase_config.cnf"
	AllDataCnf.eachDBCnf["collectionListFile"] = "/datacollection_list.cnf"
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
