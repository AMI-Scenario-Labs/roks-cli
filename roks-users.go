package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/urfave/cli"
)

func usersGet(c *cli.Context) error {
	cmd := exec.Command("oc", "get", "users")
	out, err := cmd.CombinedOutput()
	if err != nil {
		//log.Fatalf("cmd.Run() failed with %s\n", err)
		//cli.ShowAppHelp(c)
		fmt.Printf("\n%s\n\n", string(noConnectionString))
	} else {
		fmt.Printf("\nCommand Results:\n%s\n", string(out))
	}

	return nil
}

func addUser(c *cli.Context) error {
	//var out1,err1,out2, err2;

	cmd1 := exec.Command("oc", "get", "users")
	out1, err1 := cmd1.CombinedOutput()

	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("combined out:\n%s\n", string(out1))

	cmd2 := exec.Command("oc", "projects")
	out2, err2 := cmd2.CombinedOutput()
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("combined out:\n%s\n", string(out2))

	return nil
}

func listNamespace(c *cli.Context) error {
	// oc projects
	cmd1 := exec.Command("oc", "projects")
	out1, err1 := cmd1.CombinedOutput()

	if err1 != nil {
		//log.Fatal(err1)
		fmt.Printf("\n%s\n\n", string(noConnectionString))
	}
	fmt.Printf("\ncombined out:\n%s\n", string(out1))

	return nil
}

func addNamespace(c *cli.Context) error {
	// oc project #namespace

	if c.Args().Present() {
		t := c.Args().First()

		cmd1 := exec.Command("oc", "new-project", t, "--display-name="+t)
		out1, err1 := cmd1.CombinedOutput()

		if err1 != nil {
			//log.Fatal(err1)
			fmt.Printf("\n%s\n\n", string(noConnectionString))
		}
		fmt.Printf("\ncombined out:\n%s\n", string(out1))
	} else {
		cli.ShowSubcommandHelp(c)
		fmt.Printf("\n%s\n\n", string(noArgumentProvided))
	}

	return nil
}

func switchNamespace(c *cli.Context) error {
	// oc project #namespace
	if c.Args().Present() {
		t := c.Args().First()
		cmd1 := exec.Command("oc", "project", t)
		out1, err1 := cmd1.CombinedOutput()

		if err1 != nil {
			//log.Fatal(err1)
			fmt.Printf("\n%s\n\n", string(noConnectionString))
		}
		fmt.Printf("\ncombined out:\n%s\n", string(out1))
	} else {
		cli.ShowSubcommandHelp(c)
		fmt.Printf("\n%s\n\n", string(noArgumentProvided))
	}

	return nil
}

func deleteNamespace(c *cli.Context) error {
	// oc project #namespace
	if c.Args().Present() {
		t := c.Args().First()
		cmd1 := exec.Command("oc", "delete", "project", t)
		out1, err1 := cmd1.CombinedOutput()

		if err1 != nil {
			//log.Fatal(err1)
			fmt.Printf("\n%s\n\n", string(noConnectionString))
		}
		fmt.Printf("\ncombined out:\n%s\n", string(out1))
	} else {
		cli.ShowSubcommandHelp(c)
		fmt.Printf("\n%s\n\n", string(noArgumentProvided))
	}

	return nil
}
