package main

import (
	"plugin/appbase"
	"fmt"
)

func init() {
	appbase.App().Register(constructor)
}

func constructor() appbase.IPlugin {
	return &myplugin{
		state: appbase.Registered,
		}
}

type myplugin struct {
	id int
	state appbase.State
}

func (p *myplugin) GetState() appbase.State {
	return p.state
}

func (p *myplugin) SetOptions() {

}

func (p *myplugin) Initialize() {
	fmt.Printf("myplugin Initializing ...\n")
	p.state = appbase.Initialized
}

func (p *myplugin) Startup() {
	fmt.Printf("myplugin Starting ...\n")
	p.state = appbase.Started
}

func (p *myplugin) Shutdown() {
	fmt.Printf("myplugin Stoping ...\n")
	p.state = appbase.Stopped
}
