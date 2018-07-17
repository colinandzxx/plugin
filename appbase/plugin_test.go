package appbase

import (
	"fmt"
	"testing"
	"github.com/urfave/cli"
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

func (impl tstplugin) SetFlags(fg *flagGroup) {
	fmt.Printf("tstplugin SetFlags ...\n")
	fg.Add(cli.BoolFlag{
		Name:"tst_bool",
		Usage: "tst_boolllllllllllllllll desc",
	})
	fg.Add(cli.StringFlag{
		Name: "tst_str",
		Usage: "tst_strrrrrrrrrrrrrrrrrr desc",
		Value: "AAAA",
	})
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

func (impl tst2plugin) SetFlags(fg *flagGroup) {
	fmt.Printf("tst2plugin SetFlags ...\n")
	fg.Add(cli.BoolFlag{
		Name:"tst2_bool",
		Usage: "tst2_boolllllllllllllllll desc",
	})
	fg.Add(cli.StringFlag{
		Name: "tst2_str",
		Usage: "tst2_strrrrrrrrrrrrrrrrrr desc",
		Value: "AAAA",
	})
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

func (impl tst3plugin) SetFlags(fg *flagGroup) {
	fmt.Printf("tst3plugin SetFlags ...\n")
	fg.Add(cli.BoolFlag{
		Name:"tst3_bool",
		Usage: "tst3_boolllllllllllllllll desc",
	})
	fg.Add(cli.StringFlag{
		Name: "tst3_str",
		Usage: "tst3_strrrrrrrrrrrrrrrrrr desc",
		Value: "AAAA",
	})
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

	plug.SetFlags()

	plug.Initialize()
	plug.Startup()
	plug.Shutdown()
}
