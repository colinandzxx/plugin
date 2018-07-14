package appbase

import "reflect"

type State int
const (
	Registered  = iota ///< the plugin is constructed but doesn't do anything
	Initialized        ///< the plugin has initialized any state required but is idle
	Started            ///< the plugin is actively running
	Stopped            ///< the plugin is no longer running
)

type PluginImpl interface {
	Initialize()
	Startup()
	Shutdown()

	SetOptions()

	PluginDependence
}

type Plugin interface {
	GetState() State
	Name() string

	Initialize()
	Startup()
	Shutdown()

	//Register( /*p *Plugin*/ )
}

type PluginDependence interface {
	Require(do func(p Plugin))
}

type PluginObj struct {
	pImpl PluginImpl
	state State
	name pluginName
}

func newPluginObj(pImpl PluginImpl) *PluginObj {
	plugObj := &PluginObj{
		pImpl,
		Registered,
		pluginName{},
	}
	plugObj.name.Set(pImpl)
	return plugObj
}

func (obj *PluginObj) Initialize() {
	assert(obj.pImpl != nil)

	if obj.state == Registered {
		obj.state = Initialized
		obj.pImpl.Require(func(p Plugin) {
			p.Initialize()
		})
		obj.pImpl.Initialize()

		// app().Initialize
	}
	assert(obj.state == Initialized)
}

func (obj *PluginObj) Startup() {
	assert(obj.pImpl != nil)

	if obj.state == Initialized {
		obj.state = Started
		obj.pImpl.Require(func(p Plugin) {
			p.Startup()
		})
		obj.pImpl.Startup()

		// app().Start
	}
	assert(obj.state == Started)
}

func (obj *PluginObj) Shutdown() {
	//assert(obj.state == Started)

	obj.state = Stopped
	obj.pImpl.Shutdown()
}

func (obj PluginObj) GetState() State {
	return obj.state
}

func (obj PluginObj) Name() string {
	return obj.name.Name()
}

type pluginName struct {
	pImpl PluginImpl
	name string
}

func (pn *pluginName) Set(pImpl interface{})  {
	ppImpl := reflect.TypeOf(pImpl).Elem()
	required := reflect.TypeOf((*PluginImpl)(nil)).Elem()
	assert(ppImpl.Implements(required))
	pn.pImpl, _ = ppImpl.(PluginImpl)
	pn.name = ppImpl.String()
	// TODO: The name only needs the string after the last ".". i.e., main.pkgname => pkgname
}

func (pn pluginName) Name() string {
	return pn.name
}
