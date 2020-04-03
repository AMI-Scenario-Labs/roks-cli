package main

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli"
)

func pullPolicy(name string) string {
	answer := "none"

	switch name {
	case "ace", "tracing", "mq", "datapower":
		answer = "ibm-anyuid-scc"
	case "apic":
		answer = "ibm-anyuid-hostpath-scc"
	case "assetrepo":
		answer = "ibm-privileged-scc"
	case "aspera":
		answer = "ibm-anyuid-hostaccess-scc"
	case "eventstreams", "es":
		answer = "ibm-restricted-scc"
	default:

	}

	return answer
}
func verifySCC(c *cli.Context) error {
	// 0 = namespace , 1 = type
	num := c.NumFlags()
	fmt.Printf("%d\n", int(num))

	if c.NumFlags() == 2 {
		tagName := c.FlagNames()
		fmt.Printf("\n%s\n", tagName[1])

		returnFunction := pullPolicy(namespaceName)
		fmt.Printf("this is the pullPolicy(): %s", returnFunction)

	} else {

	}
	// t := c.Args().First()
	// fmt.Printf("%s", string(t))

	// if c.NumFlags() > 0 {
	// 	tagName := c.FlagNames()
	// 	fmt.Printf("%s", tagName)

	// } else {
	// 	cli.ShowSubcommandHelp(c)
	// 	fmt.Printf("\n%s\n\n", string("Flag --service is required for this command"))
	// }

	// if c.Args().Present() {
	// 	t := c.Args().First()

	// 	// cmd1 := exec.Command("oc", "project", t)
	// 	cmd1 := exec.Command("echo", t)
	// 	out1, err1 := cmd1.CombinedOutput()

	// 	if err1 != nil {
	// 		//log.Fatal(err1)
	// 		fmt.Printf("\n%s\n\n", string(noConnectionString))
	// 	}
	// 	fmt.Printf("\ncombined out:\n%s\n", string(out1))
	// } else {
	// 	cli.ShowSubcommandHelp(c)
	// 	fmt.Printf("\n%s\n\n", string(noArgumentProvided))
	// }

	return nil
}

func addSCC(c *cli.Context) error {
	// 0 = namespace , 1 = type

	//	if c.NumFlags() == 2 {
	// tagName := c.FlagNames()
	// fmt.Printf("\n%s\n", tagName[1])

	returnFunction := pullPolicy(serviceName)
	if returnFunction == "none" {
		fmt.Printf("\n%s\n\n", string(wrongServiceProvided))
	} else {
		cmd1 := exec.Command("oc", "adm", "policy", "add-scc-to-group", returnFunction, "system:serviceaccounts:"+namespaceName)
		out1, err1 := cmd1.CombinedOutput()
		if err1 != nil {
			fmt.Printf("\n%s\n\n", string(noConnectionString))
		} else {
			fmt.Printf("\ncombined out:\n%s\n", string(out1))
		}
	}

	// } else {

	// }
	// cmd1 := exec.Command("oc", "projects")
	// out1, err1 := cmd1.CombinedOutput()

	// if err1 != nil {
	// 	//log.Fatal(err1)
	// 	fmt.Printf("\n%s\n\n", string(noConnectionString))
	// }
	// fmt.Printf("\ncombined out:\n%s\n", string(out1))

	return nil
}
