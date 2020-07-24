/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Friday 24 July 2020 - 20:01:44
** @Filename:				deployment.go
**
** @Last modified by:		Tbouder
*******************************************************************************/

package templates

import (
	"os"

	"github.com/TBouder/goSetup/utils"
	"github.com/microgolang/logs"
)

func setDockerFile() {
	f, err := os.Create(utils.ProjectName + "/deployment/Dockerfile")
	if err != nil {
		logs.Error(err)
		return
	}
	defer f.Close()

	f.WriteString(
		"FROM golang:1.14\n" +
			"LABEL version=\"0.2.0\"\n" +
			"ENV PROJECT_DIR /go/src/" + utils.ProjectRoot + "\n" +
			"RUN mkdir -p ${PROJECT_DIR}\n" +
			"COPY . ${PROJECT_DIR}\n" +
			"WORKDIR ${PROJECT_DIR}\n",
	)
}

func setDockerComposeDev() {
	f, err := os.Create(utils.ProjectName + "/deployment/dev/docker-compose.yml")
	if err != nil {
		logs.Error(err)
		return
	}
	defer f.Close()

	f.WriteString(
		"version: '3'\n" +
			"services:\n" +
			"  " + utils.ProjectSlug + "_mongodb:\n" +
			"    image: mongo:latest\n" +
			"    container_name: " + utils.ProjectSlug + "_mongo\n" +
			"    ports:\n" +
			"      - 27017:27017\n" +
			"    restart: on-failure\n" +
			"  " + utils.ProjectSlug + "_api:\n" +
			"    build:\n" +
			"      context: ../..\n" +
			"      dockerfile: ./deployment/Dockerfile\n" +
			"    container_name: " + utils.ProjectSlug + "_api\n" +
			"    command:\n" +
			"      - /bin/sh\n" +
			"      - -c\n" +
			"      - |\n" +
			"        go get github.com/cespare/reflex\n" +
			"        reflex -c ./deployment/dev/reflex.conf\n" +
			"    volumes:\n" +
			"      - ../..:/go/src/" + utils.ProjectRoot + "\n" +
			"      - /var/run/docker.sock:/var/run/docker.sock\n" +
			"    ports:\n" +
			"      - 8082:8080\n" +
			"    env_file:\n" +
			"      - ../../.env\n" +
			"    restart: on-failure\n" +
			"    depends_on:\n" +
			"      - " + utils.ProjectSlug + "_mongodb\n",
	)
}

func setReflex() {
	f, err := os.Create(utils.ProjectName + "/deployment/dev/reflex.conf")
	if err != nil {
		logs.Error(err)
		return
	}
	defer f.Close()

	f.WriteString(`-r '(\.go$|go\.mod)' -s -- sh -c 'go run ./cmd/main.go start-api --localenv'`)
}

//SetDocker will create all the docker related configuration files
func SetDocker() {
	os.MkdirAll(utils.ProjectName+`/deployment/dev`, os.ModePerm)
	os.MkdirAll(utils.ProjectName+`/deployment/prod`, os.ModePerm)

	setDockerFile()
	setDockerComposeDev()
	setReflex()
}
