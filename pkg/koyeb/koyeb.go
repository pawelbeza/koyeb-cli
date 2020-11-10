package koyeb

import (
	"errors"
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Version = "develop"

	// Used for flags.
	file                 string
	cfgFile              string
	apiurl               string
	token                string
	outputFormat         string
	debug                bool
	newStackName         string
	stackRevisionMessage string

	rootCmd = &cobra.Command{
		Use:   "koyeb",
		Short: "Koyeb cli",
	}
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Init configuration",
		Run:   Init,
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Get version",
		Run:   PrintVersion,
	}

	getCmd = &cobra.Command{
		Use:     "get [resource]",
		Aliases: []string{"g", "list"},
		Short:   "Display one or many resources",
	}
	describeCmd = &cobra.Command{
		Use:     "describe [resource]",
		Aliases: []string{"d", "desc", "show"},
		Short:   "Display one resources",
	}
	updateCmd = &cobra.Command{
		Use:     "update [resource]",
		Aliases: []string{"u", "update", "edit"},
		Short:   "Update one resources",
	}
	createCmd = &cobra.Command{
		Use:     "create [resource]",
		Aliases: []string{"c", "new"},
		Short:   "Create a resource from a file",
	}
	deleteCmd = &cobra.Command{
		Use:     "delete [resource]",
		Aliases: []string{"del", "rm"},
		Short:   "Delete resources by name and id",
	}
	logCmd = &cobra.Command{
		Use:     "logs [resource]",
		Aliases: []string{"l", "log"},
		Short:   "Get the log of one resources",
	}
	invokeCmd = &cobra.Command{
		Use:     "invoke [resource]",
		Aliases: []string{"i"},
		Short:   "Invoke a function",
	}
)

func notImplemented(cmd *cobra.Command, args []string) error {
	return errors.New("Not implemented")
}

func Log(cmd *cobra.Command, args []string) {
	log.Infof("Cmd %v", cmd)
	log.Infof("Cmd has parent %v", cmd.HasParent())
	log.Infof("Cmd parent %v", cmd.Parent())
	log.Infof("Args %v", args)
}

func genericArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires a resource argument")
	}
	return nil
}

func Run() error {
	return rootCmd.Execute()
}

func er(msg interface{}) {
	log.Errorf("Error: %s", msg)
	os.Exit(1)
}

func PrintVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", Version)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.koyeb.yaml)")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "", "output format (yaml,json,table)")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug")
	rootCmd.PersistentFlags().String("url", "https://app.koyeb.com", "url of the api")
	rootCmd.PersistentFlags().String("token", "", "API token")
	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(versionCmd)

	// Create
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createStackCommand)
	createStackCommand.Flags().StringVarP(&newStackName, "name", "n", "", "Name of the stack")
	createStackCommand.Flags().StringVarP(&file, "file", "f", "", "Manifest file")
	createCmd.AddCommand(createStackRevisionCommand)
	createStackRevisionCommand.Flags().StringVarP(&file, "file", "f", "", "Manifest file")
	createStackRevisionCommand.Flags().StringVarP(&stackRevisionMessage, "message", "m", "", "Message")
	createStackRevisionCommand.MarkFlagRequired("file")
	createCmd.AddCommand(createStoreCommand)
	createStoreCommand.Flags().StringVarP(&file, "file", "f", "", "Manifest file")
	createStoreCommand.MarkFlagRequired("file")
	createCmd.AddCommand(createSecretCommand)
	createSecretCommand.Flags().StringVarP(&file, "file", "f", "", "Manifest file")
	createSecretCommand.MarkFlagRequired("file")

	// Get
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getAllCommand)
	getCmd.AddCommand(getStackCommand)
	getCmd.AddCommand(getStackRevisionCommand)
	getCmd.AddCommand(getStackFunctionCommand)
	getCmd.AddCommand(getStoreCommand)
	getCmd.AddCommand(getSecretCommand)

	// Describe
	rootCmd.AddCommand(describeCmd)
	describeCmd.AddCommand(describeStackCommand)
	describeCmd.AddCommand(describeStackRevisionCommand)
	describeCmd.AddCommand(describeStackFunctionCommand)
	describeCmd.AddCommand(describeStoreCommand)
	describeCmd.AddCommand(describeSecretCommand)

	// Update
	rootCmd.AddCommand(updateCmd)
	updateStackCommand.Flags().StringVarP(&file, "file", "f", "", "Manifest file")
	updateStackCommand.MarkFlagRequired("file")
	updateCmd.AddCommand(updateStackCommand)
	updateStoreCommand.Flags().StringVarP(&file, "file", "f", "", "Manifest file")
	updateStoreCommand.MarkFlagRequired("file")
	updateCmd.AddCommand(updateStoreCommand)
	updateSecretCommand.Flags().StringVarP(&file, "file", "f", "", "Manifest file")
	updateSecretCommand.MarkFlagRequired("file")
	updateCmd.AddCommand(updateSecretCommand)

	// Delete
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteStackCommand)
	deleteCmd.AddCommand(deleteStoreCommand)
	deleteCmd.AddCommand(deleteSecretCommand)

	// Logs
	rootCmd.AddCommand(logCmd)
	logCmd.AddCommand(logsStackFunctionCommand)

	// Run
	rootCmd.AddCommand(invokeCmd)
	invokeCmd.AddCommand(invokeStackFunctionCommand)
	invokeStackFunctionCommand.Flags().StringVarP(&file, "file", "f", "", "Event file")
}

func initConfig() {

	if debug {
		log.SetLevel(log.DebugLevel)
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".koyeb")
	}

	viper.AutomaticEnv()

	// TODO check if no .koyeb.yaml or no --config file exists, if not, ask if we want to create a new one

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Error("Configuration not found, use `koyeb init` to create a new one, or use `--config`.")
		} else {
			er(err)
		}
	} else {
		log.Debugf("Using config file: %s", viper.ConfigFileUsed())
	}

	apiurl = viper.GetString("url")
	token = viper.GetString("token")
	debug = viper.GetBool("debug")
}
