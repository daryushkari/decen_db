package loadcnf

import(
	"../utilities"
	"strings"
)

// databaseConfig includes database_init file and config information in it
// including directory path which all databases are and directory of ledger databases
// and local databases and their config file path
type databaseConfig struct {
	databaseInitPath string
	allDatabaseInfo map[string]string
	ledgerReadyDB string
	localReadyDB string
	localDBList []string
	ledgerDBList []string
 	// if directory for storing is has not been defined yet hasCnf is false
	hasCnf bool
}

var DataCnf *databaseConfig
var lineTypePlace = 0
var linePathPlace = 2

// LoadDatabaseConfig reads information from ./config/database_init.cnf and returns databaseConfig struct
// if refresh is True reload data
func LoadDatabaseConfig(refresh bool) *databaseConfig {
	if DataCnf != nil && !refresh{
		return DataCnf
	}
	DataCnf.databaseInitPath = "config/database_init.cnf"
	cnfLines := utilities.ReturnFileLines("config/database_init.cnf")
	if len(cnfLines) < 1 {
		DataCnf.hasCnf = false
	}else{
		for _, i := range cnfLines{
			lineType := strings.Fields(i)[lineTypePlace]
			linePath := strings.Fields(i)[linePathPlace]
			DataCnf.allDatabaseInfo[lineType] = linePath
		}
	}
	return DataCnf
}
