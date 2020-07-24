/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Friday 24 July 2020 - 20:01:44
** @Filename:				models.go
**
** @Last modified by:		Tbouder
*******************************************************************************/

package templates

import (
	"os"

	"github.com/TBouder/goSetup/utils"
	"github.com/microgolang/logs"
)

func setRequestFilter() {
	f, err := os.Create(utils.ProjectName+"/internal/models/requestFilter.go")
	if err != nil {
		logs.Error(err)
		return
	}
	defer f.Close()
	f.WriteString("package models\n\n")
	f.WriteString("//RequestFilter is the struct used as request to query the mongodb database\n")
	f.WriteString("type RequestFilter struct {")
	f.WriteString("\n")
	f.WriteString("    Filter     map[string]string  `json:\"filter,omitempty\"`\n")
	f.WriteString("    Projection map[string]int     `json:\"projection,omitempty\"`\n")
	f.WriteString("    Sort       map[string]int     `json:\"sort,omitempty\"`\n")
	f.WriteString("    Skip       int64              `json:\"skip,omitempty\"`\n")
	f.WriteString("    Limit      int64              `json:\"limit,omitempty\"`\n")
	f.WriteString("}")
}

func setCollections() {
	setHeader := func(f *os.File){
		f.WriteString("package models\n\n")
		f.WriteString("import (\n")
		f.WriteString("    \"context\"\n")
		f.WriteString("    \"errors\"\n")
		f.WriteString("    \"time\"\n")
		f.WriteString("    \""+utils.ProjectRoot+"/internal/db\"\n")
		f.WriteString("    \""+utils.ProjectRoot+"/internal/utils\"\n")
		f.WriteString("    \"go.mongodb.org/mongo-driver/bson\"\n")
		f.WriteString("    \"go.mongodb.org/mongo-driver/bson/primitive\"\n")
		f.WriteString("    \"go.mongodb.org/mongo-driver/mongo/options\"\n")
		f.WriteString(")\n\n")
	}
	setStruct := func(f *os.File, collection utils.TCollection){
		f.WriteString("//"+utils.Capitalize(collection.ModelName)+" represent a element of type "+collection.ModelName+"\n")
		f.WriteString("type "+utils.Capitalize(collection.ModelName)+" struct {\n")
		f.WriteString("    ID *primitive.ObjectID `json:\"-\" bson:\"_id,omitempty\"`\n")
		f.WriteString("    PublicID *string `json:\"publicID\" bson:\"publicID\"`\n")
		f.WriteString("    CreatedAt *time.Time `json:\"-\" bson:\"createdAt\"`\n")
		f.WriteString("    UpdatedAt *time.Time `json:\"-\" bson:\"updatedAt\"`\n")
		f.WriteString("    DeletedAt *time.Time `json:\"-\" bson:\"deletedAt\"`\n")
		for _, field := range collection.Model {
			name := utils.Capitalize(field.Name)
			Type := field.Type
			Json := field.Json
			if (Json == ``) {
				Json = field.Name
			}
			Bson := field.Bson
			if (Bson == ``) {
				Bson = field.Name
			}
			f.WriteString("    "+name+" "+Type+" `json:\""+Json+"\" bson:\""+Bson+"\"`\n")
		}
		f.WriteString("}\n\n")
	}
	setStructFilter := func(f *os.File, collection utils.TCollection){
		f.WriteString("//"+utils.Capitalize(collection.ModelName)+"Filter represent the filter version of the "+utils.Capitalize(collection.ModelName)+" element\n")
		f.WriteString("    type "+utils.Capitalize(collection.ModelName)+"Filter struct {\n")
		for _, field := range collection.Model {
			name := utils.Capitalize(field.Name)
			Type := field.Type
			Json := field.Json
			if (Json == ``) {
				Json = field.Name
			}
			Bson := field.Bson
			if (Bson == ``) {
				Bson = field.Name
			}
			f.WriteString("    "+name+" "+Type+" `json:\""+Json+",omitempty\" bson:\""+Bson+",omitempty\"`\n")
		}
		f.WriteString("}\n\n")
	}
	setFuncNewElem := func(f *os.File, collection utils.TCollection){
		f.WriteString("// New"+utils.Capitalize(collection.ModelName)+" create a new "+collection.ModelName+" Object\n")
		f.WriteString("func New"+utils.Capitalize(collection.ModelName)+"() (x *"+utils.Capitalize(collection.ModelName)+") {\n")
		f.WriteString("    return &"+utils.Capitalize(collection.ModelName)+"{}\n")
		f.WriteString("}\n\n")
	}
	setFuncAssign := func(f *os.File, collection utils.TCollection){
		f.WriteString("//Assign will assign the "+utils.Capitalize(collection.ModelName)+"Filter element to x\n")
		f.WriteString("func (x *"+utils.Capitalize(collection.ModelName)+") Assign(filter "+utils.Capitalize(collection.ModelName)+"Filter) (*"+utils.Capitalize(collection.ModelName)+") {\n")
		f.WriteString("    id := primitive.NewObjectID()\n")
		f.WriteString("    publicID := utils.GetPublicIDFromID(id)\n")
		f.WriteString("    newElement := &"+utils.Capitalize(collection.ModelName)+"{\n")
		f.WriteString("        ID:             &id,\n")
		f.WriteString("        PublicID:       utils.StrToPtr(publicID),\n")
		f.WriteString("        CreatedAt:      utils.TimeToPtr(time.Now()),\n")
		f.WriteString("        UpdatedAt:      utils.TimeToPtr(time.Now()),\n")
		f.WriteString("        DeletedAt:      utils.TimeToPtr(time.Time{}),\n")
		f.WriteString("    }\n\n")

		for _, field := range collection.Model {
			name := utils.Capitalize(field.Name)
			f.WriteString("    if (filter."+name+" != nil) {\n")
			f.WriteString("        newElement."+name+" = filter."+name+"\n")
			f.WriteString("    }\n")
		}

		f.WriteString("\n    return newElement\n")
		f.WriteString("}\n\n")
	}
	setFuncPost := func(f *os.File, collection utils.TCollection){
		f.WriteString("//Post will save the "+utils.Capitalize(collection.ModelName)+" element in the database\n")
		f.WriteString("func (x *"+utils.Capitalize(collection.ModelName)+") Post() error {\n")
		f.WriteString("    var err error\n\n")
		f.WriteString("    if db."+utils.Capitalize(collection.DBName)+" != nil {\n")
		f.WriteString("        _, err = db."+utils.Capitalize(collection.DBName)+".InsertOne(context.Background(), x)\n")
		f.WriteString("    } else {\n")
		f.WriteString("        return errors.New(`database not initialized`)\n")
		f.WriteString("    }\n")
		f.WriteString("    return err\n")
		f.WriteString("}\n\n")
	}
	setFuncFind := func(f *os.File, collection utils.TCollection){
		f.WriteString("//Find will perform a search in the "+utils.Capitalize(collection.DBName)+" collection to find the element matching the filters\n")
		f.WriteString("func (x *"+utils.Capitalize(collection.ModelName)+") Find(rf RequestFilter) (*"+utils.Capitalize(collection.ModelName)+", error) {\n")
		f.WriteString("    document := db."+utils.Capitalize(collection.DBName)+".FindOne(\n")
		f.WriteString("        context.Background(),\n")
		f.WriteString("        rf.Filter,\n")
		f.WriteString("    )\n")
		f.WriteString("    element := "+utils.Capitalize(collection.ModelName)+"{}\n")
		f.WriteString("    err := document.Decode(&element)\n")
		f.WriteString("    return &element, err\n")
		f.WriteString("}\n\n")
	}
	setFuncList := func(f *os.File, collection utils.TCollection){
		f.WriteString("//List will perform a search in the "+utils.Capitalize(collection.DBName)+" collection to find the elements matching the filters\n")
		f.WriteString("func (x *"+utils.Capitalize(collection.ModelName)+") List(rf RequestFilter) ([]"+utils.Capitalize(collection.ModelName)+", error) {\n")
		f.WriteString("    query, err := db."+utils.Capitalize(collection.DBName)+".Find(\n")
		f.WriteString("        context.Background(),\n")
		f.WriteString("        rf.Filter,\n")
		f.WriteString("        &options.FindOptions{\n")
		f.WriteString("            Projection: rf.Projection,\n")
		f.WriteString("            Sort:       rf.Sort,\n")
		f.WriteString("            Skip:       &rf.Skip,\n")
		f.WriteString("            Limit:      &rf.Limit,\n")
		f.WriteString("        },\n")
		f.WriteString("    )\n")
		f.WriteString("    if err != nil {\n")
		f.WriteString("        return nil, err\n")
		f.WriteString("    }\n")
		f.WriteString("    defer query.Close(context.Background())\n\n")
		f.WriteString("    elements := make([]"+utils.Capitalize(collection.ModelName)+", 0)\n")
		f.WriteString("    if err := query.All(context.Background(), &elements); err != nil {\n")
		f.WriteString("        return nil, err\n")
		f.WriteString("    }\n")
		f.WriteString("    return elements, query.Err()\n")
		f.WriteString("}\n\n")
	}
	setFuncUpdate := func(f *os.File, collection utils.TCollection){
        f.WriteString("//Update will perform an update on one "+utils.Capitalize(collection.ModelName)+"\n")
        f.WriteString("func (x *"+utils.Capitalize(collection.ModelName)+") Update(upd "+utils.Capitalize(collection.ModelName)+"Filter) error {\n")
        f.WriteString("    var err error\n\n")
        f.WriteString("    if db."+utils.Capitalize(collection.DBName)+" != nil {\n")
        f.WriteString("        _, err = db."+utils.Capitalize(collection.DBName)+".UpdateOne(\n")
        f.WriteString("            context.Background(),\n")
        f.WriteString("            x,\n")
        f.WriteString("            bson.M{`$set`: upd},\n")
        f.WriteString("        )\n")
        f.WriteString("    } else {\n")
        f.WriteString("        return errors.New(`database not initialized`)\n")
        f.WriteString("    }\n")
        f.WriteString("    return err\n")
		f.WriteString("}\n\n")
	}
	setFuncDelete := func(f *os.File, collection utils.TCollection){
        f.WriteString("//Delete will remove one "+utils.Capitalize(collection.ModelName)+" from the database\n")
        f.WriteString("func (x *"+utils.Capitalize(collection.ModelName)+") Delete(del "+utils.Capitalize(collection.ModelName)+"Filter) error {\n")
        f.WriteString("    var err error\n\n")
        f.WriteString("    if db."+utils.Capitalize(collection.DBName)+" != nil {\n")
        f.WriteString("        _, err = db."+utils.Capitalize(collection.DBName)+".DeleteOne(\n")
        f.WriteString("            context.Background(),\n")
        f.WriteString("            del,\n")
        f.WriteString("        )\n")
        f.WriteString("    } else {\n")
        f.WriteString("        return errors.New(`database not initialized`)\n")
        f.WriteString("    }\n")
        f.WriteString("    return err\n")
		f.WriteString("}\n\n")
	}

	for _, collection := range utils.Collections {
		f, err := os.Create(utils.ProjectName+"/internal/models/"+collection.DBName+".go")
		if err != nil {
			logs.Error(err)
			continue
		}
		defer f.Close()

		/******************************************************************
		**	Setup package name & the different imports needed
		******************************************************************/
		setHeader(f)
		
		/******************************************************************
		**	Setup the model structure, used to query the database and
		**	define the element, based on the model object
		******************************************************************/
		setStruct(f, collection)

		/******************************************************************
		**	Setup the filter version of the structure, which will be used
		**	for update/delete
		******************************************************************/
		setStructFilter(f, collection)

		/******************************************************************
		**	Setup the NewXXX function to initialize the XXX object
		******************************************************************/
		setFuncNewElem(f, collection)

		/******************************************************************
		**	Assign will take a XXXFilter and assign the values to a new
		**	XXX object, and setup seamless informations
		******************************************************************/
		setFuncAssign(f, collection)

		/******************************************************************
		**	Setup the Post function to add a new element to the database
		******************************************************************/
		setFuncPost(f, collection)

		/******************************************************************
		**	Setup the Find function to retreive an element from the db
		******************************************************************/
		setFuncFind(f, collection)

		/******************************************************************
		**	Setup the List function to retreive multiple elements from the
		**	db
		******************************************************************/
		setFuncList(f, collection)

		/******************************************************************
		**	Setup the Update function to update an element
		******************************************************************/
		setFuncUpdate(f, collection)

		/******************************************************************
		**	Setup the Delete function to remove an element
		******************************************************************/
		setFuncDelete(f, collection)
	}
}

//SetModels will create the internal/models/* files
func SetModels() {
	os.MkdirAll(utils.ProjectName+`/internal/models`, os.ModePerm)
	setRequestFilter()
	setCollections()
}
