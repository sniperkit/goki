package gnosis

import {
	"sync"
	"encoding/json"
}

var staticConfig *ConfigStruct
var configLock = new(sync.RWMutex)

type GlobalConfigStruct struct {
	port     string
	hostname string
}

type ServerConfigStruct struct {
	path        string
	prefix      string
	defaultPage string
	serverType  string
	restricted  []string
}

type ConfigStruct struct {
	global     GlobalConfigStruct
	mainserver ServerConfigStruct
	server     []ServerConfigStruct
}

func GetConfig() *ConfigStruct {
	configLock.RLock()
	defer configLock.RUnlock()
	return staticConfig
}

func LoadConfig(configFile string) bool {

	if configFile == "" {
		log.Println("no configuration file specified, using ./config.json")
		// return an empty config file
		configFile = "config.json"
	}
	//temp, err := mini.LoadConfiguration(configFile)

	//var fileContents []byte
	//var err error
	// have to read in the line into a byte[] array
	fileContents, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Println("Problem loading config file: ", err)
	}

	// UnMarshal the config file that was read in
	temp := new(ConfigStruct)
	err = json.Unmarshal(fileContents, temp)

	//Make sure you were able to read it in
	if err != nil {
		log.Println("parse config error: ", err)
		return false
	}

	configLock.Lock()
	staticConfig = temp
	configLock.Unlock()

	return true
}