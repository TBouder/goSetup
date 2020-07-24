/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Friday 24 July 2020 - 20:01:44
** @Filename:				main.go
**
** @Last modified by:		Tbouder
*******************************************************************************/

package main

import (
	"github.com/TBouder/goSetup/templates"
)

func main() {
	getConfig()
	templates.SetDocker()
	templates.SetMain()
	templates.SetUtils()
	templates.SetDb()
	templates.SetModels()
	templates.SetControllers()
	templates.SetGoModules()
}
