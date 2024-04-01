package cmd

import (
    "net/url"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "build-your-own-curl",
    Short: "A brief description of your application",
    Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
    Args: cobra.ExactArgs(1),

    Run: func(cmd *cobra.Command, args []string) {
        u, err := url.Parse(args[0])

        if err != nil {
            panic(err)
        }

        host := u.Hostname()
        port := u.Port()
        path := u.Path

        println("Host:", host)
        println("Port:", port)
        println("Path:", path)
    },
}
Run: func(cmd *cobra.Command, args []string) {
    u, err := url.Parse(args[0])

    if err != nil {
      panic(err)
    }

    host := u.Hostname()
    port := u.Port()

    if port == "" {
      port = "80"
    }

    path := u.Path

    println("Host:", host)
    println("Port:", port)
    println("Path:", path)
  },
}