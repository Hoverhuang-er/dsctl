package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/urfave/cli/v2"
)

var (
	iwantsaysth = `# Issues Template: \n
	#### Date: yyyy-mm-dd`
)

func main() {
	ctl := &cli.App{
		Name:                 "devops ctl",
		Usage:                "AWS Developer Tools setup commands line tool",
		EnableBashCompletion: true,
		Compiled:             time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "unknown",
				Email: "unknown@root.orz",
			},
		},
		Version:   "0.0.1",
		UsageText: `devopsctl is a setup tool for AWS Developer Tools. It can running with iac, such as: terraform, AWS Cloudformation, pulumi, ansible-playbook, etc.`,
		CommandNotFound: func(ctx *cli.Context, s string) {
			fmt.Println("command not found:", s)
			// Create a new feature or issues docs here:
			pwd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			f, err := os.Create(filepath.Join(pwd, "somthing_want_to_say.md"))
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			f.WriteString(iwantsaysth)
		},
		Metadata: map[string]interface{}{
			"author": "unknown",
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config_file",
				Aliases: []string{"c"},
				Usage:   "config file path",
			},
			&cli.StringFlag{
				Name:    "load_env",
				Hidden:  true,
				EnvVars: []string{"ENV_VAR_STR"},
			},
			&cli.StringFlag{
				Name:    "log-level",
				Aliases: []string{"l"},
				Usage:   "log level",
			},
			&cli.Int64Flag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "port",
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"d", "debug"},
				Usage:   "verbose mode",
			},
			&cli.BoolFlag{
				Name:    "console_only",
				Aliases: []string{"n"},
				Usage:   "console only mode",
			},
		},
		Commands: []*cli.Command{
			// Initialize DevOps Project
			{
				Name:    "Setup",
				Aliases: []string{"s"},
				Subcommands: []*cli.Command{
					{
						Name:  "Simple",
						Usage: "Initialize DevOps Simple Project",
						Action: func(c *cli.Context) error {
							fmt.Println("TODO: Initialize DevOps Simple Project")
							return nil
						},
					},
					{
						Name:  "Full",
						Usage: "Initialize DevOps Full Project",
						Action: func(c *cli.Context) error {
							fmt.Println("TODO: Initialize DevOps Full Project")
							return nil
						},
					},
					{
						Name:  "Customazition",
						Usage: "Initialize DevOps Customazition Project",
						Action: func(c *cli.Context) error {
							fmt.Println("TODO: Initialize DevOps Customazition Project")
							return nil
						},
					},
					{
						Name:  "Migrate",
						Usage: "Migrate DevOps Project to AWS Developers toolkit",
						Action: func(c *cli.Context) error {
							fmt.Println("TODO: Migrate DevOps Project to AWS Developers toolkit")
							return nil
						},
					},
				},
			},
			// Update DevOps Project
			{
				Name:    "InitConfig",
				Aliases: []string{"f"},
				Subcommands: []*cli.Command{
					{
						Name:    "json",
						Aliases: []string{"j", "json"},
						Usage:   "Initialize Template configuration by json to DevOps Project",
						Action: func(c *cli.Context) error {
							fmt.Println("TODO: Initialize Template configuration by json to DevOps Project")
							return nil
						},
					},
					{
						Name:    "yaml",
						Aliases: []string{"y", "yaml", "yml"},
						Usage:   "Initialize Template configuration by yaml to DevOps Project",
						Action: func(c *cli.Context) error {
							fmt.Println("TODO: Initialize Template configuration by yaml to DevOps Project")
							return nil
						},
					},
					{
						Name:    "toml",
						Aliases: []string{"t", "toml"},
						Usage:   "Initialize Template configuration by toml to DevOps Project",
						Action: func(c *cli.Context) error {
							fmt.Println("TODO: Initialize Template configuration by toml to DevOps Project")
							return nil
						},
					},
					{
						Name:    "conf",
						Aliases: []string{"ini", "conf"},
						Usage:   "Initialize Template configuration by ini or conf to DevOps Project",
						Action: func(c *cli.Context) error {
							fmt.Println("TODO: Initialize Template configuration by ini or conf to DevOps Project")
							return nil
						},
					},
					{
						Name:    "env",
						Aliases: []string{"e", "env", ".env"},
						Usage:   "Initialize Template configuration load by envVar to DevOps Project",
						Action: func(c *cli.Context) error {
							fmt.Println("TODO: Initialize Template configuration load by envVar to DevOps Project")
							return nil
						},
					},
				},
			},
		},
	}
	sort.Sort(cli.FlagsByName(ctl.Flags))
	sort.Sort(cli.CommandsByName(ctl.Commands))
	if err := ctl.Run(os.Args); err != nil {
		panic(err)
	}
}
