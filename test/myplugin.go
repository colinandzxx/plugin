package main

import (
	"plugin/appbase"
	"fmt"
)

func init() {
	appbase.App().Register(constructor)
}

func constructor() appbase.IPlugin {
	return &myplugin{}
}

type myplugin struct {
	id int
}

func (*myplugin) GetState() appbase.State {
	return appbase.Registered
}

func (*myplugin) SetOptions() {

}

func (*myplugin) Initialize() {
	fmt.Printf("myplugin Initialize ...\n")
}

func (*myplugin) Startup() {

}

func (*myplugin) Shutdown() {

}
