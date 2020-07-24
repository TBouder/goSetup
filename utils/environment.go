/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Friday 24 July 2020 - 20:01:44
** @Filename:				environment.go
**
** @Last modified by:		Tbouder
*******************************************************************************/

package utils

//TModel is the structure used in the config file to indicate which and how to create
//the fields for a collection.
type TModel struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Json string `json:"json"`
	Bson string `json:"bson"`
}

//TCollection is the structure used in the config file to indicate which and how to create
//the mongodb collections (or the XXX tables)
type TCollection struct {
	DBName    string   `json:"dbName"`
	Name      string   `json:"name"`
	ModelName string   `json:"modelName"`
	Model     []TModel `json:"model"`
}

//ProjectName represent the name of the projet. Here, it's `goSetup`
var ProjectName string

//ProjectSlug is the slug representation of the name of the projet. Here, it's `gosetup`
var ProjectSlug string

//ProjectRoot is the go path for the projet. Here, it's `github.com/TBouder/goSetup`
var ProjectRoot string

//Collections represent the defined in the config file
var Collections []TCollection
