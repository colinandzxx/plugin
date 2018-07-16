package appbase

import (
	"testing"
	"os"
)

func Test_app(T *testing.T)  {
	app := NewApp("", "aaa oo ye")
	app.Run([]string{os.Args[0]})
}
