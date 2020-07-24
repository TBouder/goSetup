/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Friday 24 July 2020 - 20:01:44
** @Filename:				cli.go
**
** @Last modified by:		Tbouder
*******************************************************************************/

package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/TBouder/goSetup/utils"
)

//Config represent the configuration file as a go struct
type Config struct {
	ProjectName string              `json:"name"`
	ProjectRoot string              `json:"root"`
	Collections []utils.TCollection `json:"collections"`
}

func getConfig() {
	configFilePath := flag.String("config", `./config.json`, "Path to the configuration file")
	flag.Parse()

	configFile, err := os.Open(*configFilePath)
	if err != nil {
		log.Fatalf(`Impossible to access config.json`)
	}
	defer configFile.Close()

	byteConfig, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Fatalf(`Impossible to read config.json`)
	}

	var config Config
	json.Unmarshal([]byte(byteConfig), &config)
	utils.Collections = config.Collections
	utils.ProjectName = strings.ToLower(config.ProjectName)
	utils.ProjectSlug = utils.ToSlug(utils.ProjectName)
	utils.ProjectRoot = config.ProjectRoot
}
