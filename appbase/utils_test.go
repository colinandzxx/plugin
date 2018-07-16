package appbase

import (
	"testing"
	"bytes"
	"github.com/spf13/cobra"
	"fmt"
)


func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}
func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOutput(buf)
	root.SetArgs(args)
	c, err = root.ExecuteC()
	//root.Execute()
	return c, buf.String(), err
}

func Test_cfg(t *testing.T) {
	rootCmd := cfg()
	c, output, err := executeCommandC(rootCmd, "--help")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("output: %s\n", output)
		fmt.Printf("%v\n", c.ValidArgs)
	}
}
