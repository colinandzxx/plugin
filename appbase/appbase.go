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
			pm:		pluginManager{plugins: make(map[string]IPlugin)},
			}
	})
	return instance
}

type Application struct {
	pm pluginManager
}

// get plugin by plugin type name
func (app *Application) Get(name string) IPlugin {
	plugin := app.pm.find(name)
	if plugin == nil {
		panic("unable to find plugin: " + name)
	}
	return plugin
}

// register plugin, need input the constructor of plugin
func (app *Application) Register(constructor func() IPlugin) IPlugin {
	return app.pm.register(constructor)
}

// plugin manager, key = type.Name()
type pluginManager struct {
	plugins map[string]IPlugin
}

func (pm pluginManager) find(name string) IPlugin {
	if plugin, ok := pm.plugins[name]; ok {
		return plugin
	}
	return nil
}

func (pm pluginManager) register(constructor func() IPlugin) IPlugin {
	p := constructor()
	name := GetNameByType(p)
	plugin := pm.find(name)
	if plugin != nil {
		return plugin
	}

	pm.plugins[name] = p
	return p
}

