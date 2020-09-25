package main

import (
	"bytes"
	"fmt"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"
	"text/template"
)

type User struct {
	ShortName string
	FullName  string
	Email     string
}

type MessageContent struct {
	Pair       []User
	Message    string
	JiraTag    string
	JiraNumber int
}

type Config struct {
	Users   []User
	JiraTag string
}

func main() {
	var ConfigPath = "./config.yml"
	var team, jiraTag = loadConfig(ConfigPath)
	var pair string
	var FullMsg MessageContent
	FullMsg.JiraTag = jiraTag

	app := &cli.App{
		Name:  "Git-cli template ",
		Usage: "Easy to use git comment generator",
		Commands: []*cli.Command{
			{
				Name:  "commit",
				Usage: "Record changes to the repository",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "message, m",
						Aliases:     []string{"m"},
						Usage:       "commit message ",
						Destination: &FullMsg.Message,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "pair",
						Aliases:     []string{"p"},
						Usage:       "add pair(s) name(s) to commit message",
						Destination: &pair,
					},
					&cli.IntFlag{
						Name:        "jira",
						Aliases:     []string{"j"},
						Usage:       "add jira ticket number to commit message",
						Destination: &FullMsg.JiraNumber,
					},
				},
				Action: func(c *cli.Context) error {
					FullMsg.Pair = getPairDetails(pair, team)
					formattedMsg := format(FullMsg)
					commitMsg(formattedMsg)
					return nil
				},
			},
			{
				Name:    "pair-list",
				Aliases: []string{"pl"},
				Usage:   "List available pairs",
				Action: func(c *cli.Context) error {
					printPairs(team)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func printPairs(team []User) {
	tabFormat := new(tabwriter.Writer)
	tabFormat.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer tabFormat.Flush()
	fmt.Fprintf(tabFormat, "\n %s\t\t%s\t%s\t", "User", "FullName", "Email")
	fmt.Fprintf(tabFormat, "\n %s\t\t%s\t%s\t", "----", "----", "----")

	for _, user := range team {
		fmt.Fprintf(tabFormat, "\n %s\t\t%s\t%s\t", user.ShortName, user.FullName, user.Email)
	}
}

func getPairDetails(pair string, team []User) []User {
	var pairNames = strings.Split(pair, ",")
	var pairDetails []User
	for _, pairName := range pairNames {
		for _, collaborator := range team {
			if pairName == collaborator.ShortName {
				pairDetails = append(pairDetails, collaborator)
			}
		}
	}
	return pairDetails
}

func commitMsg(formattedMsg string) {
	cmd := exec.Command("git", "commit", "-m", formattedMsg)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Commit unsuccessful: " + fmt.Sprint(err) + ": " + stderr.String())
	} else {
		fmt.Println("Commit successful: " + out.String())
	}
}

func format(FullMsg MessageContent) string {
	buf := new(bytes.Buffer)
	t, err := template.New("commitMessage").Parse(msgFormat)
	if err != nil {
		panic(err)
	}
	err = t.Execute(buf, FullMsg)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func loadConfig(configPath string) ([]User, string) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Print(err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("error:", err)
	}
	return config.Users, config.JiraTag
}
