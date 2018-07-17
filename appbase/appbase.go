package appbase

import (
	"sync"
	"github.com/urfave/cli"
)

var (
	instance *Application
	once     sync.Once
)

// get the Application singleton
func App() *Application {
	once.Do(func() {
		instance = &Application{
			plugins: make(map[string]Plugin),
			initializedPlugins: make([]Plugin, 0),
			runningPlugins: make([]Plugin, 0),
			appObj: create(),
		}
	})
	return instance
}

type Application struct {
	plugins            map[string]Plugin ///< all registered plugins
	initializedPlugins []Plugin
	runningPlugins     []Plugin

	appObj *cli.App
}

func (app *Application) SetBaseInfo(usage, copyright, gitCommit, version string) {
	app.appObj.Version = version
	if len(gitCommit) >= 8 {
		app.appObj.Version += "-" + gitCommit[:8]
	}
	if app.appObj.Version == "" {
		app.appObj.Version = "unknown"
	}
	app.appObj.Usage = usage
	app.appObj.Copyright = copyright
}

func (app *Application) SetAction(action cli.ActionFunc) {
	app.appObj.Action = action
}

// register plugin, need input the constructor of plugin
func (app *Application) Register(plugin func() PluginImpl) Plugin {
	plug := newPluginObj(plugin())
	if app.FindByName(plug.Name()) != nil {
		return nil
	}

	app.plugins[plug.Name()] = plug
	return plug
}

func (app *Application) SetFlags(plugins... func() PluginImpl) {
	fgs := []flagGroup{}
	for _, plugin := range plugins {
		plug := app.Find(plugin)
		assert(plug != nil)
		fgs = append(fgs, plug.SetFlags()...)
	}
	AddFlagGroups(fgs)
	// add MISC options
	fg := NewFlags("MISC")
	AddFlagGroup(*fg)

	for _, fg := range appHelpFlagGroups {
		for _, f := range fg.Flags {
			app.appObj.Flags = append(app.appObj.Flags, f)
		}
	}

	overrideHelpTemplates()
}

func (app *Application) Initialize(plugins... func() PluginImpl) bool {
	// TODO: options

	// Initialize plugins
	for _, plugin := range plugins {
		plug := app.Find(plugin)
		assert(plug != nil)
		plug.Initialize()
	}

	return true
}

func (app *Application) Startup() {
	defer func() {
		if err := recover(); err != nil {
			app.Shutdown()
			assert(false)
		}
	}()

	for _, plug := range app.initializedPlugins {
		plug.Startup()
	}
}

func (app *Application) Shutdown() {
	for _, plug := range app.runningPlugins {
		plug.Shutdown()
	}
	app.runningPlugins = nil
	app.initializedPlugins = nil
	app.plugins = nil
}

func (app *Application) pluginInitialized(plug Plugin) {
	app.initializedPlugins = append(app.initializedPlugins, plug)
}

func (app *Application) pluginStarted(plug Plugin) {
	app.runningPlugins = append(app.runningPlugins, plug)
}

// get plugin by plugin type name
func (app *Application) GetByName(name string) Plugin {
	plug := app.FindByName(name)
	assertEx(plug != nil, "unable to get plugin: "+name)
	return plug
}

func (app *Application) Get(plugin func() PluginImpl) Plugin {
	name := pluginName{}
	name.Set(plugin())
	plug := app.FindByName(name.Name())
	assertEx(plug != nil, "unable to get plugin: "+name.Name())
	return plug
}

func (app Application) FindByName(name string) Plugin {
	if plugin, ok := app.plugins[name]; ok {
		return plugin
	}
	return nil
}

func (app Application) Find(plugin func() PluginImpl) Plugin {
	name := pluginName{}
	name.Set(plugin())
	return app.FindByName(name.Name())
}
