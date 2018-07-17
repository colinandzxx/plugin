package appbase

import (
	"testing"
	"os"
	"github.com/urfave/cli"
	"fmt"
)

func Test_flags(T *testing.T)  {
	app := create()

	app.Action = act
	app.Commands = []cli.Command{
		versionCommand,
	}

	fg := NewFlags("TEST")
	fg.Add(&cli.BoolFlag{
		Name:  "booltest",
		Usage: "this is booltest",
		})
	fg.Add(&cli.StringFlag{
		Name:  "stringtest",
		Usage: "this is stringtest",
		Value: "AAddaa",
		})
	AddFlagGroup(*fg)
	fg = NewFlags("MISC")
	AddFlagGroup(*fg)
	overrideHelpTemplates()

	for _, fg := range appHelpFlagGroups {
		for _, f := range fg.Flags {
			app.Flags = append(app.Flags, f)
		}
	}

	app.Run([]string{os.Args[0], "-asd"})
	app.Run([]string{os.Args[0], "version", "-h"})
	app.Run([]string{os.Args[0], "--booltest"})
	app.Run([]string{os.Args[0], ""})
}

func act(ctx *cli.Context) error {
	fmt.Printf("args==========: %v\n", ctx.Bool("booltest"))
	fmt.Printf("args==========: %v\n", ctx.String("stringtest"))
	return nil
}