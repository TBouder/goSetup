/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Friday 24 July 2020 - 20:01:44
** @Filename:				db.go
**
** @Last modified by:		Tbouder
*******************************************************************************/

package templates

import (
	"os"

	"github.com/TBouder/goSetup/utils"
	"github.com/microgolang/logs"
)

func setInit() {
	f, err := os.Create(utils.ProjectName+"/internal/db/init.go")
	if err != nil {
		logs.Error(err)
		return
	}
	defer f.Close()
	f.WriteString(`package db

// Init initialize connection to the DB and create collection objects if needed
func Init() {
    InitClient()
    InitCollections()
}`)
}

func setInitClient() {
	f, err := os.Create(utils.ProjectName+"/internal/db/initClient.go")
	if err != nil {
		logs.Error(err)
		return
	}
	defer f.Close()
	f.WriteString(`package db

import (
    "context"
    "log"

    "`+utils.ProjectRoot+`/internal/utils"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// InitClient initializes a DB client object
func InitClient() {
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(utils.MongodbURI))
    if err != nil {
        log.Fatalf("failed to connect to MongoDB: %v\n", err)
    }
    // create DB client object
    db = client.Database(utils.MongodbDBName)
}`)
}

func setInitCollections() {
	f, err := os.Create(utils.ProjectName+"/internal/db/initCollections.go")
	if err != nil {
		logs.Error(err)
		return
	}
	defer f.Close()

	f.WriteString("package db\n\n")
	f.WriteString("import (\n")
	f.WriteString("    \"go.mongodb.org/mongo-driver/mongo\"\n")
	f.WriteString(")\n\n")
	f.WriteString("var (")
	for _, collection := range utils.Collections {
		f.WriteString("\n")
		f.WriteString(`    //`+utils.Capitalize(collection.DBName)+` is the collection used to store the `+collection.Name)
		f.WriteString("\n")
		f.WriteString("    "+utils.Capitalize(collection.DBName)+` *mongo.Collection`)
	}

	f.WriteString("\n)\n\n")
	f.WriteString("// InitCollections initializes collection objects and creates indexes\n")
	f.WriteString("func InitCollections() {")
	for _, collection := range utils.Collections {
		f.WriteString("\n")
		f.WriteString("    "+utils.Capitalize(collection.DBName)+` = db.Collection("`+collection.DBName+`")`)
	}
	f.WriteString("\n}")
}

//SetDb will create the internal/db/* files
func SetDb() {
	os.MkdirAll(utils.ProjectName+`/internal/db`, os.ModePerm)
	setInit()
	setInitClient()
	setInitCollections()
}
