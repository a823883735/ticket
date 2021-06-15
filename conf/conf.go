package conf

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

var conf *ini.File

// Init load conf.ini file.
func init() {
	var err error
	conf, err = ini.Load("conf.ini")
	if err != nil {
		pwd, _ := os.Getwd()
		log.Fatal("conf.ini load failed. err: ", err, ", ", pwd)
	}
}

// Conf returns this conf.
func Conf() *ini.File {
	return conf
}