package logger

import (
	"log"
	"os"
	"github.com/gin-gonic/gin/json"
)

// Logger config
type LogConfiguration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var logConfig LogConfiguration
var logger *log.Logger

func init() {
	loadConfig()
	file, err := os.OpenFile("./logger/brem.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	file, err := os.Open("./logger/config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	logConfig = LogConfiguration{}
	err = decoder.Decode(&logConfig)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

///////////////////////////////////////////////////
/// Function for logging
//////////////////////////////////////////////////

func Info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func Danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func Warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}
