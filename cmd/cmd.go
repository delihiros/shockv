package cmd

import (
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
	ttl          int
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
			r, _ := c.NewDB(databaseName, diskless)
			jsonutil.PrintJSON(r, format)
		},
	}

	getCmd = &cobra.Command{
		Use:   "get",
		Short: "get value by key",
		Run: func(cmd *cobra.Command, args []string) {
			c := client.New(serverURL, port)
			r, _ := c.Get(databaseName, key)
			jsonutil.PrintJSON(r, format)
		},
	}

	setCmd = &cobra.Command{
		Use:   "set",
		Short: "set value by key",
		Run: func(cmd *cobra.Command, args []string) {
			c := client.New(serverURL, port)
			r, _ := c.Set(databaseName, key, value, ttl)
			jsonutil.PrintJSON(r, format)
		},
	}

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list keys and values",
		Run: func(cmd *cobra.Command, args []string) {
			c := client.New(serverURL, port)
			r, _ := c.List(databaseName)
			jsonutil.PrintJSON(r, format)
		},
	}

	delCmd = &cobra.Command{
		Use:   "delete",
		Short: "delete by key",
		Run: func(cmd *cobra.Command, args []string) {
			c := client.New(serverURL, port)
			r, _ := c.Delete(databaseName, key)
			jsonutil.PrintJSON(r, format)
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
	setCmd.Flags().IntVarP(&ttl, "ttl", "t", 0, "ttl second you want to set")

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
