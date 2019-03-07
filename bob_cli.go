package main

import (
	"bob-cli/commands_bob"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Command struct {
	name string
	help string
}

type Commands struct {
	status  Command
	connect Command
}

func show(commands Commands) {
	fmt.Printf("\n%s : %s", commands.status.name, commands.status.help)
	fmt.Printf("\n%s : %s \n\n", commands.connect.name, commands.connect.help)
}

func main() {
	statusCommand := Command{
		name: "can-we-build-it",
		help: "running status of Bob",
	}
	connectCommand := Command{
		name: "connect-with",
		help: "provide the connection info of the bob",
	}
	commands := Commands{
		status:  statusCommand,
		connect: connectCommand,
	}
	args := os.Args
	if len(args) == 1 {
		fmt.Println("\nUsage of wendy: \n\nCommands")
		show(commands)
		os.Exit(0)
	}

	switch args[1] {
	case statusCommand.name:
		{
			fileContent, err := ioutil.ReadFile(".bob_config")
			if err != nil {
				fmt.Println("No connection information found for Bob. Try wendy connect-with command to connect first")
			}
			if fileContent!= nil {
				api := commands_bob.API{
					Client:http.DefaultClient,
					BaseURL:string(fileContent),
				}
				response, err := api.RunningStatus()
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(response)
			}
		}
	case connectCommand.name:
		{
			bobURl := args[2]
			api := commands_bob.API{
				Client:http.DefaultClient,
				BaseURL:bobURl,
			}
			_, err := api.RunningStatus()
			if err != nil {
				fmt.Println("Error")
				return
			}
			configfile, err := os.Create(".bob_config")
			n, err := configfile.Write([]byte(bobURl))
			if err != nil {
				fmt.Println("Error")
			}
			if n != 0 {
				fmt.Println("Connected")
			}
		}
	}
}
