/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Friday 24 July 2020 - 20:01:44
** @Filename:				cmd.go
**
** @Last modified by:		Tbouder
*******************************************************************************/

package templates

import (
	"os"

	"github.com/TBouder/goSetup/utils"
	"github.com/microgolang/logs"
)

//SetMain will create the cmd/main.go file
func SetMain() {
	os.MkdirAll(utils.ProjectName+`/cmd`, os.ModePerm)
	f, err := os.Create(utils.ProjectName + "/cmd/main.go")
	if err != nil {
		logs.Error(err)
		return
	}
	defer f.Close()

	f.WriteString(`package main

import (
    "log"
    "os"
    "` + utils.ProjectRoot + `/internal/controllers"
    "` + utils.ProjectRoot + `/internal/db"
    "` + utils.ProjectRoot + `/internal/utils"
    "github.com/joho/godotenv"
    "github.com/urfave/cli"
)

var commands = []cli.Command{
    {
        Name: "start",
        Flags: []cli.Flag{
            &cli.BoolFlag{
                Name:  "localenv, l",
                Usage: "If we should load the environment without docker.",
            },
        },
        Action: func(c *cli.Context) error {
            if c.Bool("localenv") == true {
                godotenv.Load("../.env")
            }
            utils.InitEnvironment()
            db.Init()
            return controllers.NewRouter().Run()
        },
    },
}

func main() {
    api := cli.NewApp()
    api.Commands = commands
    if err := api.Run(os.Args); err != nil {
        log.Fatalf("failed to run the command: %v\n", err)
    }
}`)
}
