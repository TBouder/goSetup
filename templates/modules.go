/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Friday 24 July 2020 - 20:01:44
** @Filename:				modules.go
**
** @Last modified by:		Tbouder
*******************************************************************************/

package templates

import (
	"os"

	"github.com/TBouder/goSetup/utils"
	"github.com/microgolang/logs"
)

//SetGoModules will create all the go module related files
func SetGoModules() {
	f, err := os.Create(utils.ProjectName + "/go.mod")
	if err != nil {
		logs.Error(err)
		return
	}
	defer f.Close()

	f.WriteString("module " + utils.ProjectRoot + "\n\n")
	f.WriteString("go 1.14\n")
}
