package main

import "plugin/appbase"

func main() {
	plugin := appbase.App().Get(appbase.GetNameByType(&myplugin{}))
	plugin.Initialize()
}