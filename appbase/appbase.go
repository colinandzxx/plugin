package appbase

import (
	"sync"
)

var (
	instance *Application
	once sync.Once
)

// get the Application singleton
func App() *Application {
	once.Do(func() {
		instance = &Application{
			plugins: make(map[string]Plugin),
			}
	})
	return instance
}

type Application struct {
	plugins map[string]Plugin
}

// get plugin by plugin type name
func (app *Application) GetByName(name string) Plugin {
	plug := app.findByName(name)
	assertEx(plug != nil, "unable to find plugin: " + name)
	return plug
}

// register plugin, need input the constructor of plugin
func (app *Application) Register(plugin func() PluginImpl) Plugin {
	plug := newPluginObj(plugin())
	if app.findByName(plug.Name()) != nil {
		return nil
	}

	app.plugins[plug.Name()] = plug
	return plug
}

func (app Application) findByName(name string) Plugin {
	if plugin, ok := app.plugins[name]; ok {
		return plugin
	}
	return nil
}

func (app Application) find(plugin func() PluginImpl) Plugin {
	name := pluginName{}
	name.Set(plugin())
	if plugin, ok := app.plugins[name.Name()]; ok {
		return plugin
	}
	return nil
}

/*
type pluginWrapper struct {
	name string
	plugin IPlugin
	constructor func() IPlugin
	state State
}

func (pw *pluginWrapper) initialize() {
	if pw.state != Registered {
		panic(fmt.Sprintf("plugin %s state is %v(need %v).", pw.state, Registered))
	}

	if pw.plugin == nil {
		pw.plugin = pw.constructor()
		if pw.plugin == nil {
			panic(fmt.Sprintf("plugin %s constructor() failed.", pw.name))
		}
	}

	for _, name := range pw.plugin.RequireDependencies() {
		dependency := App().Get(name)
		if dependency == nil {
			panic(fmt.Sprintf("plugin %s constructor() failed.", pw.name))
		}
		dependency.Initialize()
	}

	pw.plugin.Initialize()
}

func (pw *pluginWrapper) startup() {

}

func (pw *pluginWrapper) shutdown() {

}

func (pw *pluginWrapper) getState() State {
	return pw.state
}
*/
