/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Friday 24 July 2020 - 20:01:44
** @Filename:				capitalize.go
**
** @Last modified by:		Tbouder
*******************************************************************************/

package utils

import (
	"unicode"
)

// Capitalize replace the first letter with a uppercase.
// There is some bug with strings.Title() which lead to an all uppercase string
func Capitalize(str string) string {
	arrByteStr := []byte(str)
	arrByteStr[0] = byte(unicode.ToUpper(rune(arrByteStr[0])))
	return string(arrByteStr)
}

