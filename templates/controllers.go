/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Friday 24 July 2020 - 20:01:44
** @Filename:				controllers.go
**
** @Last modified by:		Tbouder
*******************************************************************************/

package templates

import (
	"os"

	"github.com/TBouder/goSetup/utils"
	"github.com/microgolang/logs"
)

func setServer() {
	f, err := os.Create(utils.ProjectName+"/internal/controllers/server.go")
	if err != nil {
		logs.Error(err)
		return
	}
	defer f.Close()

	f.WriteString("package controllers\n\n")
	f.WriteString("import (\n")
	f.WriteString("    \"time\"\n")
	f.WriteString("    \"github.com/gin-contrib/cors\"\n")
	f.WriteString("    \"github.com/gin-gonic/contrib/ginrus\"\n")
	f.WriteString("    \"github.com/gin-gonic/gin\"\n")
	f.WriteString("    \"github.com/sirupsen/logrus\"\n")
	f.WriteString(")\n\n")

	f.WriteString("// NewRouter create the routes and setup the server\n")
	f.WriteString("func NewRouter() *gin.Engine {\n")
	f.WriteString("    router := gin.New()\n")
	f.WriteString("    router.Use(gin.Recovery())\n")
	f.WriteString("    logger := logrus.StandardLogger()\n")
	f.WriteString("    logger.SetFormatter(&logrus.JSONFormatter{})\n")
	f.WriteString("    router.Use(ginrus.Ginrus(logger, time.RFC3339, true))\n")
	f.WriteString("    corsConf := cors.Config{\n")
	f.WriteString("        AllowAllOrigins: true,\n")
	f.WriteString("        AllowHeaders:    []string{`Origin`, `Content-Length`, `Content-Type`, `Authorization`},\n")
	f.WriteString("    }\n")
	f.WriteString("    router.Use(cors.New(corsConf))\n")

	for _, collection := range utils.Collections {
		f.WriteString("\n")
		f.WriteString("    {\n")
		f.WriteString("        c := "+collection.DBName+"Controller{}\n")
		f.WriteString("        router.GET(`"+collection.ModelName+"/find/:publicID`, c.FindID)\n")
		f.WriteString("        router.POST(`"+collection.ModelName+"/find`, c.Find)\n")
		f.WriteString("        router.POST(`"+collection.ModelName+"s/list`, c.List)\n")
		f.WriteString("        router.POST(`"+collection.ModelName+"/post`, c.Post)\n")
		f.WriteString("        router.PUT(`"+collection.ModelName+"/update`, c.Update)\n")
		f.WriteString("        router.DELETE(`"+collection.ModelName+"/delete`, c.Delete)\n")
		f.WriteString("    }\n")
	}

	f.WriteString("\n")
	f.WriteString("    return router\n")
	f.WriteString("}\n\n")
}

func setFuncRoutes() {
	setHeader := func(f *os.File, collection utils.TCollection){
		f.WriteString("package controllers\n\n")
		f.WriteString("import (\n")
		f.WriteString("    \"net/http\"\n")
		f.WriteString("    \""+utils.ProjectRoot+"/internal/models\"\n")
		f.WriteString("    \"github.com/gin-gonic/gin\"\n")
		f.WriteString("    \"github.com/gin-gonic/gin/binding\"\n")
		f.WriteString(")\n\n")
		f.WriteString("type "+collection.DBName+"Controller struct{}\n\n")
	}

	setFuncFind := func(f *os.File, collection utils.TCollection){
		f.WriteString("//Find will retreive one element in the database\n")
		f.WriteString("func (y "+collection.DBName+"Controller) Find(c *gin.Context) {\n")
		f.WriteString("    var requestFilter models.RequestFilter\n")
		f.WriteString("    if err := c.ShouldBindBodyWith(&requestFilter, binding.JSON); err != nil {\n")
		f.WriteString("        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{`error`: `bad request`})\n")
		f.WriteString("        return\n")
		f.WriteString("    }\n\n")
		f.WriteString("    result, err := models.New"+utils.Capitalize(collection.ModelName)+"().Find(requestFilter)\n")
		f.WriteString("    if err != nil {\n")
		f.WriteString("        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{`error`: `could not retreive element`})\n")
		f.WriteString("        return\n")
		f.WriteString("    }\n")
		f.WriteString("    c.JSON(http.StatusOK, result)\n")
		f.WriteString("}\n\n")
	}
	setFuncFindID := func(f *os.File, collection utils.TCollection){
		f.WriteString("//FindID will retreive one element in the database based on it's publicID\n")
		f.WriteString("func (y "+collection.DBName+"Controller) FindID(c *gin.Context) {\n")
		f.WriteString("    var requestFilter models.RequestFilter\n")
		f.WriteString("    var publicID = c.Param(`publicID`)\n")
		f.WriteString("    if publicID == `` {\n")
		f.WriteString("        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{`error`: `bad request`})\n")
		f.WriteString("        return\n")
		f.WriteString("    }\n\n")
		f.WriteString("    requestFilter.Filter = map[string]string{`publicID`: publicID}\n")
		f.WriteString("    result, err := models.New"+utils.Capitalize(collection.ModelName)+"().Find(requestFilter)\n")
		f.WriteString("    if err != nil {\n")
		f.WriteString("        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{`error`: `could not retreive element`})\n")
		f.WriteString("        return\n")
		f.WriteString("    }\n")
		f.WriteString("    c.JSON(http.StatusOK, result)\n")
		f.WriteString("}\n\n")
	}
	setFuncList := func(f *os.File, collection utils.TCollection){
		f.WriteString("//List will retreive multiple elements in the database\n")
		f.WriteString("func (y "+collection.DBName+"Controller) List(c *gin.Context) {\n")
		f.WriteString("    var requestFilter models.RequestFilter\n")
		f.WriteString("    if err := c.ShouldBindBodyWith(&requestFilter, binding.JSON); err != nil {\n")
		f.WriteString("        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{`error`: `bad request`})\n")
		f.WriteString("        return\n")
		f.WriteString("    }\n\n")
		f.WriteString("    result, err := models.New"+utils.Capitalize(collection.ModelName)+"().List(requestFilter)\n")
		f.WriteString("    if err != nil {\n")
		f.WriteString("        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{`error`: `could not retreive elements`})\n")
		f.WriteString("        return\n")
		f.WriteString("    }\n")
		f.WriteString("    c.JSON(http.StatusOK, result)\n")
		f.WriteString("}\n\n")
	}
	setFuncPost := func(f *os.File, collection utils.TCollection){
		f.WriteString("//Post will add one element in the database\n")
		f.WriteString("func (y "+collection.DBName+"Controller) Post(c *gin.Context) {\n")
		f.WriteString("    var modelFilter models."+utils.Capitalize(collection.ModelName)+"Filter\n")
		f.WriteString("    if err := c.ShouldBindBodyWith(&modelFilter, binding.JSON); err != nil {\n")
		f.WriteString("        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{`error`: `bad request`})\n")
		f.WriteString("        return\n")
		f.WriteString("    }\n\n")
		f.WriteString("    newElement := models.New"+utils.Capitalize(collection.ModelName)+"().Assign(modelFilter)\n")
		f.WriteString("    err := newElement.Post()\n")
		f.WriteString("    if err != nil {\n")
		f.WriteString("        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{`error`: `could not save element`})\n")
		f.WriteString("        return\n")
		f.WriteString("    }\n")
		f.WriteString("    c.JSON(http.StatusOK, newElement)\n")
		f.WriteString("}\n\n")
	}
	setFuncUpdate := func(f *os.File, collection utils.TCollection){
		f.WriteString("//Update will update one element from the database\n")
		f.WriteString("func (y "+collection.DBName+"Controller) Update(c *gin.Context) {\n")
		f.WriteString("    var modelFilter models."+utils.Capitalize(collection.ModelName)+"Filter\n")
		f.WriteString("    if err := c.ShouldBindBodyWith(&modelFilter, binding.JSON); err != nil {\n")
		f.WriteString("        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{`error`: `bad request`})\n")
		f.WriteString("        return\n")
		f.WriteString("    }\n\n")
		f.WriteString("    err := models.New"+utils.Capitalize(collection.ModelName)+"().Update(modelFilter)\n")
		f.WriteString("    if err != nil {\n")
		f.WriteString("        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{`error`: `could not update element`})\n")
		f.WriteString("        return\n")
		f.WriteString("    }\n")
		f.WriteString("    c.JSON(http.StatusOK, gin.H{`result`: `SUCCESS`})\n")
		f.WriteString("}\n\n")
	}
	setFuncDelete := func(f *os.File, collection utils.TCollection){
		f.WriteString("//Delete will delete one element from the database\n")
		f.WriteString("func (y "+collection.DBName+"Controller) Delete(c *gin.Context) {\n")
		f.WriteString("    var modelFilter models."+utils.Capitalize(collection.ModelName)+"Filter\n")
		f.WriteString("    if err := c.ShouldBindBodyWith(&modelFilter, binding.JSON); err != nil {\n")
		f.WriteString("        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{`error`: `bad request`})\n")
		f.WriteString("        return\n")
		f.WriteString("    }\n\n")
		f.WriteString("    err := models.New"+utils.Capitalize(collection.ModelName)+"().Delete(modelFilter)\n")
		f.WriteString("    if err != nil {\n")
		f.WriteString("        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{`error`: `could not delete element`})\n")
		f.WriteString("        return\n")
		f.WriteString("    }\n")
		f.WriteString("    c.JSON(http.StatusOK, gin.H{`result`: `SUCCESS`})\n")
		f.WriteString("}\n\n")
	}




	for _, collection := range utils.Collections {
		f, err := os.Create(utils.ProjectName+"/internal/controllers/"+collection.DBName+".go")
		if err != nil {
			logs.Error(err)
			continue
		}
		defer f.Close()

		setHeader(f, collection)
		setFuncFind(f, collection)
		setFuncFindID(f, collection)
		setFuncList(f, collection)
		setFuncPost(f, collection)
		setFuncUpdate(f, collection)
		setFuncDelete(f, collection)
	}
}

//SetControllers will create the internal/controllers/* files
func SetControllers() {
	os.MkdirAll(utils.ProjectName+`/internal/controllers`, os.ModePerm)
	setServer()
	setFuncRoutes()
}
