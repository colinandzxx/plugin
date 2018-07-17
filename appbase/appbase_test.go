package appbase

import (
	"github.com/urfave/cli"
	"fmt"
	"testing"
	"os"
)

func Test_run(t *testing.T) {
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

	App().SetBaseInfo("this is a app", "", "", "1.0.0")
	App().SetFlags(
		func() PluginImpl {
			return &tstplugin{}
		},
		func() PluginImpl {
			return &tst2plugin{}
		},
		func() PluginImpl {
			return &tst3plugin{}
		})
	App().SetAction(func(ctx *cli.Context) error {
		App().Initialize(
			func() PluginImpl {
				return &tstplugin{}
			},
			func() PluginImpl {
				return &tst2plugin{}
			},
			func() PluginImpl {
				return &tst3plugin{}
			})

		App().Startup()

		App().Shutdown()

		return nil
	})

	fmt.Printf("====================================================================================\n")
	App().Run([]string{os.Args[0], "-h"})

	fmt.Printf("====================================================================================\n")
	//App().Run([]string{os.Args[0], "-tst3"})

	fmt.Printf("====================================================================================\n")
	//App().Run([]string{os.Args[0], "--tst3_bool"})

	fmt.Printf("====================================================================================\n")
	//App().Run([]string{os.Args[0], "-tst_str"})

	fmt.Printf("====================================================================================\n")
	App().Run([]string{os.Args[0], "-tst2_str", "AA222"})
}
