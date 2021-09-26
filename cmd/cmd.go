package cmd

import (
	"log"

	"github.com/delihiros/shockv/pkg/client"
	"github.com/delihiros/shockv/pkg/jsonutil"
	"github.com/delihiros/shockv/pkg/server"

	"github.com/spf13/cobra"
)

var (
	format       bool
	serverURL    string
	port         int
	databaseName string
	key          string
	value        string
	diskless     bool
)

var (
	rootCmd = &cobra.Command{
		Use:   "shockv",
		Short: "simple RESTful key-value store",
	}

	clientCmd = &cobra.Command{
		Use:   "client",
		Short: "client command to interact with shockV server",
	}

	newCmd = &cobra.Command{
		Use:   "new",
		Short: "create new database",
		Run: func(cmd *cobra.Command, args []string) {
			c := client.New(serverURL, port)
			err := c.NewDB(databaseName, diskless)
			if err != nil {
				log.Println(err)
			}
		},
	}

	getCmd = &cobra.Command{
		Use:   "get",
		Short: "get value by key",
		Run: func(cmd *cobra.Command, args []string) {
			c := client.New(serverURL, port)
			v, err := c.Get(databaseName, key)
			if err != nil {
				log.Println(err)
			} else {
				jsonutil.PrintJSON(v, format)
			}
		},
	}

	setCmd = &cobra.Command{
		Use:   "set",
		Short: "set value by key",
		Run: func(cmd *cobra.Command, args []string) {
			c := client.New(serverURL, port)
			err := c.Set(databaseName, key, value)
			if err != nil {
				log.Println(err)
			}
		},
	}

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list keys and values",
		Run: func(cmd *cobra.Command, args []string) {
			c := client.New(serverURL, port)
			kv, err := c.List(databaseName)
			if err != nil {
				log.Println(err)
			} else {
				jsonutil.PrintJSON(kv, format)
			}
		},
	}

	delCmd = &cobra.Command{
		Use:   "delete",
		Short: "delete by key",
		Run: func(cmd *cobra.Command, args []string) {
			c := client.New(serverURL, port)
			err := c.Delete(databaseName, key)
			if err != nil {
				log.Println(err)
			}
		},
	}

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "server command to use shockV server",
	}

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "starts shockV server",
		Run: func(cmd *cobra.Command, args []string) {
			s := server.New()
			s.Execute(port)
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&format, "format", "f", false, "format output json")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "server port number")

	clientCmd.PersistentFlags().StringVarP(&databaseName, "database", "d", "", "database name you want to use")
	clientCmd.PersistentFlags().StringVarP(&serverURL, "server", "s", "http://localhost", "server you want to use")

	newCmd.Flags().BoolVar(&diskless, "diskless", false, "set if you want to use database with diskless mode")

	getCmd.Flags().StringVarP(&key, "key", "k", "", "key you want to get")

	setCmd.Flags().StringVarP(&key, "key", "k", "", "key you want to set")
	setCmd.Flags().StringVarP(&value, "value", "v", "", "value you want to set")

	delCmd.Flags().StringVarP(&key, "key", "k", "", "key you want to delete")

	clientCmd.AddCommand(newCmd)
	clientCmd.AddCommand(listCmd)
	clientCmd.AddCommand(getCmd)
	clientCmd.AddCommand(setCmd)
	clientCmd.AddCommand(delCmd)

	serverCmd.AddCommand(startCmd)

	rootCmd.AddCommand(clientCmd)
	rootCmd.AddCommand(serverCmd)
}
