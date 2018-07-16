package appbase

import (
	"fmt"
	"testing"
)

type tstplugin struct {
}

func (impl tstplugin) Initialize() {
	fmt.Printf("tstplugin Initialize ...\n")
}

func (impl tstplugin) Startup() {
	fmt.Printf("tstplugin Startup ...\n")
}

func (impl tstplugin) Shutdown() {
	fmt.Printf("tstplugin Shutdown ...\n")
}

func (impl tstplugin) SetOptions() {
	fmt.Printf("tstplugin SetOptions ...\n")
}

func (impl tstplugin) Require(do func(p Plugin)) {
	fmt.Printf("tstplugin Require ...\n")

	PluginRequire(do,
		func() PluginImpl { return &tst2plugin{} },
		func() PluginImpl { return &tst3plugin{} })
}

type tst2plugin struct {
}

func (impl tst2plugin) Initialize() {
	fmt.Printf("tst2plugin Initialize ...\n")
}

func (impl tst2plugin) Startup() {
	fmt.Printf("tst2plugin Startup ...\n")
}

func (impl tst2plugin) Shutdown() {
	fmt.Printf("tst2plugin Shutdown ...\n")
}

func (impl tst2plugin) SetOptions() {
	fmt.Printf("tst2plugin SetOptions ...\n")
}

func (impl tst2plugin) Require(do func(p Plugin)) {
	fmt.Printf("tst2plugin Require ...\n")

	PluginRequire(do,
		func() PluginImpl { return &tst3plugin{} })
}

type tst3plugin struct {
}

func (impl tst3plugin) Initialize() {
	fmt.Printf("tst3plugin Initialize ...\n")
}

func (impl tst3plugin) Startup() {
	fmt.Printf("tst3plugin Startup ...\n")
}

func (impl tst3plugin) Shutdown() {
	fmt.Printf("tst3plugin Shutdown ...\n")
}

func (impl tst3plugin) SetOptions() {
	fmt.Printf("tst3plugin SetOptions ...\n")
}

func (impl tst3plugin) Require(do func(p Plugin)) {
	fmt.Printf("tst3plugin Require ...\n")
}

func Test_tstplugin(t *testing.T) {
	plug := App().Register(func() PluginImpl {
		return &tstplugin{}
	})
	if plug != nil {
		fmt.Printf("plug name: %s\n", plug.Name())
	} else {
		fmt.Printf("Register() error\n")
	}

	plug2 := App().Register(func() PluginImpl {
		return &tst2plugin{}
	})
	if plug2 != nil {
		fmt.Printf("plug name: %s\n", plug2.Name())
	} else {
		fmt.Printf("Register() error\n")
	}

	plug3 := App().Register(func() PluginImpl {
		return &tst3plugin{}
	})
	if plug3 != nil {
		fmt.Printf("plug name: %s\n", plug3.Name())
	} else {
		fmt.Printf("Register() error\n")
	}

	plug.Initialize()
	plug.Startup()
	plug.Shutdown()
}
