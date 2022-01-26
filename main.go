// Start point

package main

import (
	myfuncs "mgr/functions"
)

func main() {
	myfuncs.Colors.White = "[#FFFFFF]"
	myfuncs.Colors.Yellow = "[#FFFF00]"
	myfuncs.Colors.Red = "[#FF0000]"
	myfuncs.Colors.Green = "[#008000]"
	myfuncs.Colors.Blue = "[#1E90FF]"
	myfuncs.Colors.Orange = "[#FFA500]"
	myfuncs.Colors.Filler = "[#123456]"

	myfuncs.GetMGFiles()
}
