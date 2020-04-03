package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/urfave/cli"
)

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func openClients(c *cli.Context) error {
	// icp-console.<domain>/common-nav/cli?useNav=icp4i
	// mac curl -kLo cloudctl-darwin-amd64-v3.2.2-1532 https://icp-console.divyasoni-bff24f8c28cd4601e6dda0244988633f-0001.sjc03.containers.appdomain.cloud:443/api/cli/cloudctl-darwin-amd64
	// linux 64bit  curl -kLo cloudctl-linux-amd64-v3.2.2-1532 https://icp-console.divyasoni-bff24f8c28cd4601e6dda0244988633f-0001.sjc03.containers.appdomain.cloud:443/api/cli/cloudctl-linux-amd64
	// windows curl -kLo cloudctl-win-amd64-v3.2.2-1532.exe https://icp-console.divyasoni-bff24f8c28cd4601e6dda0244988633f-0001.sjc03.containers.appdomain.cloud:443/api/cli/cloudctl-win-amd64.exe

	browserURL := "https://icp-console." + subDomain + "/common-nav/cli?useNav=icp4i"
	openbrowser(browserURL)

	return nil
}

func downloadCharts(c *cli.Context) error {

	returnValid := validCapability(serviceName)
	if !returnValid {
		fmt.Printf("\n%s\n\n", string(wrongServiceProvided))
	} else {
		returnData := pullChartData(serviceName)
		fmt.Printf("\n%s\n", returnData.chart)

		cmd1 := exec.Command("curl", "-O", returnData.download)
		out1, err1 := cmd1.CombinedOutput()
		if err1 != nil {
			log.Fatalf("\nProblem Running Download with WGET %s\n", err1)
		} else {

			fmt.Printf("\nStarting %s Chart Download:\n%s\n", string(returnData.name), string(out1))
		}
	}

	return nil
}

func pullChartData(name string) *cpservice {
	var answer cpservice

	switch name {
	case "ace":
		answer = cACE
	case "tracing":
		answer = cTRACING
	case "mq":
		answer = cMQ
	case "datapower":
		answer = cDATAPOWER
	case "apic":
		answer = cAPIC
	case "assetrepo":
		answer = cASSETREPO
	case "aspera":
		answer = cASPERA
	case "eventstreams", "es":
		answer = cEVENTSTREAMS
	default:

	}

	return &answer
}
