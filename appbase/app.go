package appbase

import (
	"fmt"
	"path/filepath"
	"os"
	"github.com/urfave/cli"
)

var (
	CommandHelpTemplate = `
{{.cmd.Name}}{{if .cmd.Subcommands}} command{{end}}{{if .cmd.Flags}} [command options]{{end}} [arguments...]
{{if .cmd.Description}}{{.cmd.Description}}
{{end}}{{if .cmd.Subcommands}}
SUBCOMMANDS:
	{{range .cmd.Subcommands}}{{.cmd.Name}}{{with .cmd.ShortName}}, {{.cmd}}{{end}}{{ "\t" }}{{.cmd.Usage}}
	{{end}}{{end}}{{if .categorizedFlags}}
{{range $idx, $categorized := .categorizedFlags}}{{$categorized.Name}} OPTIONS:
{{range $categorized.Flags}}{{"\t"}}{{.}}
{{end}}
{{end}}{{end}}
`

	// AppHelpTemplate is the test template for the default, global app help topic.
	AppHelpTemplate = `
NAME:
{{.App.Name}} - {{.App.Usage}}
   {{.App.Copyright}}

USAGE:
   {{.App.HelpName}} [options]{{if .App.Commands}} command [command options]{{end}} {{if .App.ArgsUsage}}{{.App.ArgsUsage}}{{else}}[arguments...]{{end}}
   {{if .App.Version}}
VERSION:
   {{.App.Version}}
   {{end}}{{if len .App.Authors}}
AUTHOR(S):
   {{range .App.Authors}}{{ . }}{{end}}
   {{end}}{{if .App.Commands}}
COMMANDS:
   {{range .App.Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end}}{{end}}{{if .FlagGroups}}
{{range .FlagGroups}}{{.Name}} OPTIONS:
  {{range .Flags}}{{.}}
  {{end}}
{{end}}{{end}}{{if .App.Copyright }}
COPYRIGHT:
   {{.App.Copyright}}
   {{end}}
`
)

// create an app with defaults.
func create() *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Author = ""
	app.Email = ""
	app.Version = ""
	app.Version = "unknown"
	app.Usage = ""
	app.Copyright = ""
	app.HideVersion = true
	app.Commands = []cli.Command{
		versionCommand,
	}

	return app
}

var versionCommand = cli.Command{
	Action:    MigrateFlags(version),
	Name:      "version",
	Usage:     "print version info",
	ArgsUsage: " ",
	Category:  "MISCELLANEOUS COMMANDS",
	Description: `The output of this command is supposed to be machine-readable.`,
}

func version(ctx *cli.Context) error {
	fmt.Printf("version: %s\n", ctx.App.Version)
	return nil
}

func MigrateFlags(action func(ctx *cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		for _, name := range ctx.FlagNames() {
			if ctx.IsSet(name) {
				ctx.GlobalSet(name, ctx.String(name))
			}
		}
		return action(ctx)
	}
}

func (app *Application) Run(args []string) {
	assert(app.appObj != nil)
	if err := app.appObj.Run(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

/*
func defaultAction(ctx *cli.Context, plugins... func() PluginImpl) error {
	App().Initialize(plugins...)
	return nil
}*/
