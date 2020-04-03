package main

import (
	"fmt"
	"log"
	"os/exec"

	//docker "github.com/fsouza/go-dockerclient"
	"github.com/urfave/cli"
)

func dockerLogin(c *cli.Context) error {
	// docker login -u $(oc whoami) -p $(oc whoami -t) docker-registry.default.svc:5000
	// oc project
	// cmd := exec.Command("oc", "project", "default")
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	//log.Fatalf("cmd.Run() failed with %s\n", err)
	// 	//cli.ShowAppHelp(c)
	// 	fmt.Printf("\n%s\n\n", string(noConnectionString))
	// } else {
	// 	fmt.Printf("\nCommand Results:\n%s\n", string(out))
	// cmdUser := exec.Command("oc", "whoami")
	// cmdUserOut, err0 := cmdUser.CombinedOutput()
	// if err0 != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err0)
	// } else {
	// 	fmt.Printf("\nCommand Results:\n%s\n", string(cmdUserOut))
	// }

	// cmdPass := exec.Command("oc", "whoami", "-t")
	// cmdPassOut, err2 := cmdPass.CombinedOutput()
	// if err2 != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err2)
	// } else {
	// 	fmt.Printf("\nCommand Results:\n%s\n", string(cmdPassOut))
	// }
	// client, err := docker.NewClientFromEnv()
	// if err != nil {
	// 	panic(err)
	// }

	// imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s", imgs)

	//cmdStr := "sudo docker run -v ~/exp/a.out:/a.out ubuntu:14.04 /a.out -m 10m"
	cmdMaster := ""
	//out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()
	out2, err2 := exec.Command(cmdMaster).Output()
	fmt.Printf("%s", out2)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("%s", out2)

	cmdFullList := "docker login -u $(oc whoami) -p $(oc whoami -t) docker-registry.default.svc:5000"
	cmd1 := exec.Command("/bin/sh", "-c", cmdFullList)
	out1, err1 := cmd1.CombinedOutput()

	//+string(cmdUserOut)+" -p "+string(cmdPassOut)+" ")
	//, "login", "-u", , "-p", , "")

	if err1 != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err1)
	} else {
		fmt.Printf("%s\r", string(out1))

	}

	// }
	return nil
}

func datapowerConnect(c *cli.Context) error {
	// kubectl port-forward svc/docker-registry 9090:9090
	// oc project datapower
	// cmd := exec.Command("oc", "project", "datapower")
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	//log.Fatalf("cmd.Run() failed with %s\n", err)
	// 	//cli.ShowAppHelp(c)
	// 	fmt.Printf("\n%s\n\n", string(noConnectionString))
	// } else {
	// 	fmt.Printf("\nCommand Results:\n%s\n", string(out))
	// }

	// Run Additional Connections for DataPower
	cmd1 := exec.Command("kubectl", "port-forward", "pod/", "9090:9090")
	out1, err1 := cmd1.CombinedOutput()
	if err1 != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err1)
	} else {
		fmt.Printf("\nCommand Results:\n%s\n", string(out1))

	}
	return nil
}

func registryConnect(c *cli.Context) error {
	// kubectl port-forward svc/docker-registry 5000:5000
	// oc project default
	// cmd := exec.Command("oc", "project", "default")
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	//log.Fatalf("cmd.Run() failed with %s\n", err)
	// 	//cli.ShowAppHelp(c)
	// 	fmt.Printf("\n%s\n\n", string(noConnectionString))
	// } else {
	// 	fmt.Printf("\nCommand Results:\n%s\n", string(out))

	// }

	// cmd1 := exec.Command("kubectl", "port-forward", "svc/docker-registry", "5000:5000")
	// //, "port-forward", "svc/docker-registry", "5000:5000")
	// stdout, err := cmd1.StdoutPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := cmd1.Start(); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := cmd1.Wait(); err != nil {
	// 	log.Fatal(err)
	// }
	//out1, err1 := cmd1.CombinedOutput()
	// if err1 != nil {
	// 	fmt.Printf("\n%s\n\n", string(noConnectionString))
	// 	log.Fatalf("cmd.Run() failed with %s\n", err1)
	// } else {
	//fmt.Printf("\nCommand Results:\n%s\n", io.WriteString(stdout))
	fmt.Printf("\nCommand Results:\n%s\n", "test")

	// }

	return nil
}
