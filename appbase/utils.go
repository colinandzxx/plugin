package appbase

import (
	"reflect"
	"github.com/spf13/cobra"
	"fmt"
	//"github.com/spf13/viper"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func GetNameByType(v interface{}) string {
	refV := reflect.ValueOf(v)
	return reflect.Indirect(refV).Type().String()
}

func cfg() *cobra.Command {
	cobra.OnInitialize()

	var subCmd1 = cobra.Command{
		Use: "sub1",
		Short: "this is a sub1",
		Long: "this is long long long sub1 description",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s %v\n", cmd.Name(), args)
			fmt.Printf("subCmd1 nothing to do\n")
		},
	}

	var rootCmd = cobra.Command{
		Use: "test",
		Short: "this is a test",
		Long: "this is long long long test description",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s %v\n", cmd.Name(), args)
			fmt.Printf("rootCmd nothing to do\n")
			//subCmd1.Run(cmd, args)
		},
	}

	rootCmd.SetVersionTemplate(`customized version: aaa.{{.Version}}`)

	//rootCmd.Flag()
	var cfgFile, projectBase, userLicense string
	rootCmd.Flags().SortFlags = false
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

	//rootCmd.AddCommand()
	fset := pflag.NewFlagSet("ttttt-tst", pflag.ContinueOnError)
	fset.SortFlags = false
	names := []string{"C", "B", "A", "D"}
	for _, name := range names {
		fset.Bool(name, false, "AAoo")
		fset.Set(name, "true")
	}
	//fset.SetAnnotation("ad", "aaab", []string{"adfsfafsa"})

	i := 0
	fset.Visit(func(f *pflag.Flag) {
		if names[i] != f.Name {
			fmt.Errorf("Incorrect order. Expected %v, got %v", names[i], f.Name)
		}
		i++
	})

	//rootCmd.MarkPersistentFlagRequired("config")
	//rootCmd.PersistentFlags().AddFlagSet(fset)
	subCmd1.PersistentFlags().AddFlagSet(fset)
	//subCmd1.InheritedFlags().AddFlagSet(fset)
	//subCmd1.di
	//rootCmd.AddCommand(&subCmd1)
	/*
	rootCmd.SetUsageFunc(func(command *cobra.Command) error {
		err := rootCmd.UsageFunc()(command)
		return err
	})
	*/
	//rootCmd.SetUsageTemplate()

	return &rootCmd
}

func initcfg() {

}

func cfg2() {
	//altsrc.NewTomlSourceFromFile()
}