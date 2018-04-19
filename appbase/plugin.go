package appbase

type State int
const (
	Registered = iota 	///< the plugin is constructed but doesn't do anything
	Initialized 		///< the plugin has initialized any state required but is idle
	Started 			///< the plugin is actively running
	Stopped 			///< the plugin is no longer running
)

type IPlugin interface {
	GetState() State
	SetOptions()
	Initialize()
	Startup()
	Shutdown()
}

type BasePlugin struct {

}