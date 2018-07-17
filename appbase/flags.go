package appbase

import (
	"io"
	"sort"

	"github.com/urfave/cli"
	"strings"
)

// flagGroup is a collection of flags belonging to a single topic.
type flagGroup struct {
	Name  string
	Flags []cli.Flag
}

// appHelpFlagGroups is the application flags, grouped by functionality.
var appHelpFlagGroups = []flagGroup{
	/*{
		Name: "TEST",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "booltest",
				Usage: "this is booltest",
			},
			cli.StringFlag{
				Name:  "stringtest",
				Usage: "this is stringtest",
				Value: "AAddaa",
			},
		},
	},
	{
		Name: "MISC",
	},*/
}

// byCategory sorts an array of flagGroup by Name in the order
// defined in appHelpFlagGroups.
type byCategory []flagGroup

func (a byCategory) Len() int      { return len(a) }
func (a byCategory) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byCategory) Less(i, j int) bool {
	iCat, jCat := a[i].Name, a[j].Name
	iIdx, jIdx := len(appHelpFlagGroups), len(appHelpFlagGroups) // ensure non categorized flags come last

	for i, group := range appHelpFlagGroups {
		if iCat == group.Name {
			iIdx = i
		}
		if jCat == group.Name {
			jIdx = i
		}
	}

	return iIdx < jIdx
}

func flagCategory(flag cli.Flag) string {
	for _, category := range appHelpFlagGroups {
		for _, flg := range category.Flags {
			if flg.GetName() == flag.GetName() {
				return category.Name
			}
		}
	}
	return "MISC"
}

func overrideHelpTemplates() {
	// Override the default app help template
	cli.AppHelpTemplate = AppHelpTemplate
	// Override the default command help template
	cli.CommandHelpTemplate = CommandHelpTemplate

	// Define a one shot struct to pass to the usage template
	type helpData struct {
		App        interface{}
		FlagGroups []flagGroup
	}

	// Override the default app help printer, but only for the global app help
	originalHelpPrinter := cli.HelpPrinter
	cli.HelpPrinter = func(w io.Writer, tmpl string, data interface{}) {
		if tmpl == AppHelpTemplate {
			// Iterate over all the flags and add any uncategorized ones
			categorized := make(map[string]struct{})
			for _, group := range appHelpFlagGroups {
				for _, flag := range group.Flags {
					categorized[flag.String()] = struct{}{}
				}
			}
			uncategorized := []cli.Flag{}
			for _, flag := range data.(*cli.App).Flags {
				if _, ok := categorized[flag.String()]; !ok {
					if strings.HasPrefix(flag.GetName(), "dashboard") {
						continue
					}
					uncategorized = append(uncategorized, flag)
				}
			}
			if len(uncategorized) > 0 {
				// Append all ungategorized options to the misc group
				miscs := len(appHelpFlagGroups[len(appHelpFlagGroups)-1].Flags)
				appHelpFlagGroups[len(appHelpFlagGroups)-1].Flags = append(appHelpFlagGroups[len(appHelpFlagGroups)-1].Flags, uncategorized...)

				// Make sure they are removed afterwards
				defer func() {
					appHelpFlagGroups[len(appHelpFlagGroups)-1].Flags = appHelpFlagGroups[len(appHelpFlagGroups)-1].Flags[:miscs]
				}()
			}
			// Render out custom usage screen
			originalHelpPrinter(w, tmpl, helpData{data, appHelpFlagGroups})
		} else if tmpl == CommandHelpTemplate {
			// Iterate over all command specific flags and categorize them
			categorized := make(map[string][]cli.Flag)
			for _, flag := range data.(cli.Command).Flags {
				if _, ok := categorized[flag.String()]; !ok {
					categorized[flagCategory(flag)] = append(categorized[flagCategory(flag)], flag)
				}
			}

			// sort to get a stable ordering
			sorted := make([]flagGroup, 0, len(categorized))
			for cat, flgs := range categorized {
				sorted = append(sorted, flagGroup{cat, flgs})
			}
			sort.Sort(byCategory(sorted))

			// add sorted array to data and render with default printer
			originalHelpPrinter(w, tmpl, map[string]interface{}{
				"cmd":              data,
				"categorizedFlags": sorted,
			})
		} else {
			originalHelpPrinter(w, tmpl, data)
		}
	}
}

func NewFlags(name string) *flagGroup {
	return &flagGroup{Name:name,}
}

func (fg *flagGroup) Add(flag cli.Flag) {
	fg.Flags = append(fg.Flags, flag)
}

func AddFlagGroup(fg flagGroup) {
	appHelpFlagGroups = append(appHelpFlagGroups, fg)
}

func AddFlagGroups(fgs []flagGroup) {
	if len(fgs) != 0 {
		appHelpFlagGroups = append(appHelpFlagGroups, fgs...)
	}
}
