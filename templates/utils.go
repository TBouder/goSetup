/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Friday 24 July 2020 - 20:01:44
** @Filename:				utils.go
**
** @Last modified by:		Tbouder
*******************************************************************************/

package templates

import (
	"os"

	"github.com/TBouder/goSetup/utils"
	"github.com/microgolang/logs"
)

func setPublicID() {
	f, err := os.Create(utils.ProjectName+"/internal/utils/publicID.go")
	if err != nil {
		logs.Error(err)
		return
	}
	defer f.Close()
	f.WriteString(`package utils

import (
    "crypto/sha256"
    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

//GetPublicIDFromID is an helper function used to generate an UUID based on the
//ObjectID
func GetPublicIDFromID(id primitive.ObjectID) string {
    hash := sha256.Sum256([]byte(id.String()))
    trimmedHash := hash[:16]
    finalUUID, _ := uuid.FromBytes(trimmedHash)
    return finalUUID.String()
}`)
}
func setToPtr() {
	f, err := os.Create(utils.ProjectName+"/internal/utils/toPtr.go")
	if err != nil {
		logs.Error(err)
		return
	}
	defer f.Close()
	f.WriteString(`package utils

import (
    "time"
)

//StrToPtr is an helper function used get the ptr of a string
func StrToPtr(s string) *string {
    return &s
}

//BoolToPtr is an helper function used get the ptr of a bool
func BoolToPtr(b bool) *bool {
    return &b
}

//IntToPtr is an helper function used get the ptr of an int
func IntToPtr(i int) *int {
    return &i
}

//TimeToPtr is an helper function used get the ptr of a time
func TimeToPtr(t time.Time) *time.Time {
    return &t
}`)
}
func setEnvironment() {
	f, err := os.Create(utils.ProjectName+"/internal/utils/environment.go")
	if err != nil {
		logs.Error(err)
		return
	}
	defer f.Close()
	f.WriteString(`package utils

import (
    "log"
    "os"
)

var (
    // MongodbURI : URI of the MongoDB
    MongodbURI string
    // MongodbDBName : name of the DB
    MongodbDBName string
)

// InitEnvironment initializes variables from environment
func InitEnvironment() {
    var exists bool

    /* MongoDB */
    MongodbURI, exists = os.LookupEnv("MONGODB_URI")
    if !exists {
        log.Fatal("MONGODB_URI environment variable not set")
    }
    MongodbDBName, exists = os.LookupEnv("MONGODB_DB_NAME")
    if !exists {
        log.Fatal("MONGODB_DB_NAME environment variable not set")
    }
}`)
}

//SetUtils will create the internal/utils/* files 
func SetUtils() {
	os.MkdirAll(utils.ProjectName+`/internal/utils`, os.ModePerm)
	setPublicID()
	setToPtr()
	setEnvironment()
}