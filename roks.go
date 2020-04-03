package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

type cpsvcversion struct {
	version string
	ppa     string
	chart   string
}

type cpservice struct {
	name       string
	download   string
	capability string
	scc        string
	version    string
	ppa        string
	chart      string
	install    string
}

var (
	cmds          []*cli.Command
	userFlags     []cli.Flag
	sccFlags      []cli.Flag
	serviceName   string
	namespaceName string
	cACE          = cpservice{
		name:       "ace",
		capability: "App & Data Integration",
		scc:        "ibm-anyuid-scc",
		ppa:        "IBM-App-Connect-Enterprise-for-IBM-Cloud-Pak-for-Integration-2.2.0.tgz",
		version:    "2.2.0",
		download:   "https://raw.githubusercontent.com/IBM/charts/master/repo/entitled/ibm-ace-dashboard-icp4i-prod-2.2.0.tgz",
		chart:      "ibm-ace-dashboard-icp4i-prod-2.2.0.tgz",
		install:    "http://charts.roks.ibm.cloud.cloudpakintegration.io/",
	}
	cAPIC = cpservice{
		name:       "apic",
		capability: "API Lifecycle",
		scc:        "ibm-anyuid-hostpath-scc",
		ppa:        "IBM-API-Connect-Enterprise-for-IBM-Cloud-Pak-for-Integration-1.0.3.tgz",
		version:    "1.0.3",
		download:   "https://raw.githubusercontent.com/IBM/charts/master/repo/entitled/ibm-apiconnect-icp4i-prod-1.0.3.tgz",
		chart:      "ibm-apiconnect-icp4i-prod-1.0.3.tgz",
		install:    "http://charts.roks.ibm.cloud.cloudpakintegration.io/",
	}
	cASPERA = cpservice{
		name:       "aspera",
		capability: "High Speed Transfer",
		scc:        "ibm-anyuid-hostaccess-scc",
		ppa:        "IBM-Aspera-High-Speed-Transfer-Server-for-IBM-Cloud-Pak-for-Integration-1.2.3.tgz",
		version:    "1.2.3",
		download:   "https://raw.githubusercontent.com/IBM/charts/master/repo/entitled/ibm-aspera-hsts-icp4i-1.2.3.tgz",
		chart:      "ibm-aspera-hsts-icp4i-1.2.3.tgz",
		install:    "http://charts.roks.ibm.cloud.cloudpakintegration.io/",
	}
	cDATAPOWER = cpservice{
		name:       "datapower",
		capability: "Gateway",
		scc:        "ibm-anyuid-scc",
		ppa:        "IBM-DataPower-Virtual-Edition-for-IBM-Cloud-Pak-for-Integration-1.0.4.tgz",
		version:    "1.0.4",
		download:   "https://raw.githubusercontent.com/IBM/charts/master/repo/entitled/ibm-datapower-icp4i-1.0.4.tgz",
		chart:      "ibm-datapower-icp4i-1.0.4.tgz",
		install:    "http://charts.roks.ibm.cloud.cloudpakintegration.io/",
	}
	cEVENTSTREAMS = cpservice{
		name:       "eventstreams",
		capability: "Events",
		scc:        "ibm-restricted-scc",
		ppa:        "IBM-Event-Streams-for-IBM-Cloud-Pak-for-Integration-1.3.2-for-OpenShift.tgz",
		version:    "1.3.2",
		download:   "https://raw.githubusercontent.com/IBM/charts/master/repo/entitled/ibm-eventstreams-rhel-icp4i-prod-1.3.2.tgz",
		chart:      "ibm-eventstreams-rhel-icp4i-prod-1.3.2.tgz",
		install:    "http://charts.roks.ibm.cloud.cloudpakintegration.io/",
	}
	cMQ = cpservice{
		name:       "mq",
		capability: "Messaging",
		scc:        "ibm-anyuid-scc",
		ppa:        "IBM-MQ-Advanced-for-IBM-Cloud-Pak-for-Integration-4.1.0.tgz",
		version:    "4.1.0",
		download:   "https://raw.githubusercontent.com/IBM/charts/master/repo/entitled/ibm-mqadvanced-server-integration-prod-4.1.0.tgz",
		chart:      "ibm-mqadvanced-server-integration-prod-4.1.0.tgz",
		install:    "http://charts.roks.ibm.cloud.cloudpakintegration.io/",
	}
	cTRACING = cpservice{
		name:       "tracing",
		capability: "Tracing",
		scc:        "ibm-anyuid-scc",
		ppa:        "IBM-Cloud-Pak-for-Integration-Operations-Dashboard-1.0.0.tgz",
		version:    "1.0.0",
		download:   "https://raw.githubusercontent.com/IBM/charts/master/repo/entitled/ibm-icp4i-tracing-prod-1.0.0.tgz",
		chart:      "ibm-icp4i-tracing-prod-1.0.0.tgz",
		install:    "http://charts.roks.ibm.cloud.cloudpakintegration.io/",
	}
	cASSETREPO = cpservice{
		name:       "assetrepo",
		capability: "Asset Repository",
		scc:        "ibm-privileged-scc",
		ppa:        "IBM-Cloud-Pak-for-Integration-Asset-Repository-2.2.0.tgz",
		version:    "2.2.0",
		download:   "https://raw.githubusercontent.com/IBM/charts/master/repo/entitled/ibm-icp4i-asset-repo-prod-2.2.0.tgz",
		chart:      "ibm-icp4i-asset-repo-prod-2.2.0.tgz",
		install:    "http://charts.roks.ibm.cloud.cloudpakintegration.io/",
	}
	cpiServices          [7]string
	fullname             string
	mynamespace          string
	subDomain            string
	userEmailId          string
	userCloudID          string
	wrongServiceProvided string = "Invalid Capability Type Provided.  Please view the help for valid names"
	noArgumentProvided   string = "Argument is required for this command"
	noConnectionString   string = "Please refresh your OpenShift Console Connection in your Terminal"
)

func init() {

	sccFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "capability",
			Usage:       "--capability <ace,apic,mq,eventstreams,datapower,assetrepo,tracing,aspera>",
			Required:    true,
			Destination: &serviceName,
		},
		&cli.StringFlag{
			Name:        "namespace",
			Usage:       "--namespace <name>",
			Required:    true,
			Destination: &namespaceName,
		},
	}

	domainFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "subdomain",
			Usage:       "--subdomain <roks cluster subdomain>",
			Required:    true,
			Destination: &subDomain,
		},
	}

	downloadFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "capability",
			Usage:       "--capability <ace,apic,mq,eventstreams,datapower,assetrepo,tracing,aspera>",
			Required:    true,
			Destination: &serviceName,
		},
	}

	cmds = []*cli.Command{
		{
			Name:     "scc",
			Usage:    "managing scc policies for namespaces ",
			Category: "Installers",
			Subcommands: []*cli.Command{
				{
					Name:   "verify",
					Usage:  "verify namespace",
					Action: verifySCC,
					Flags:  sccFlags,
				},
				{
					Name:   "apply",
					Usage:  "apply namespace --service <name>",
					Action: addSCC,
					Flags:  sccFlags,
				},
			},
		},
		{
			Name:     "console",
			Category: "Console",
			Usage:    "managing console user acess",
			Subcommands: []*cli.Command{
				{
					Name:   "get-users",
					Usage:  "get-users",
					Action: usersGet,
				},
				{
					Name:   "add-user",
					Usage:  "adduser --email <emailaddress> --iamid <cloudid>",
					Action: addUser,
					Flags:  userFlags,
				},
			},
		},
		{
			Name:     "charts",
			Category: "Installers",
			Usage:    "managing charts for CP4i Capabilities",
			Subcommands: []*cli.Command{
				{
					Name:   "download",
					Usage:  "download --capability <ace,apic,mq,eventstreams,datapower,assetrepo,tracing,aspera>",
					Action: downloadCharts,
					Flags:  downloadFlags,
				},
				// {
				// 	Name:   "add-user",
				// 	Usage:  "adduser --email <emailaddress> --iamid <cloudid>",
				// 	Action: addUser,
				// 	Flags:  userFlags,
				// },
			},
		},
		{
			Name:     "clients",
			Category: "Installers",
			Usage:    "accessing CLI utilites to connect to cluster",
			Subcommands: []*cli.Command{
				{
					Name:   "open",
					Usage:  "open --subdomain <roks subdomain>",
					Action: openClients,
					Flags:  domainFlags,
				},
				// {
				// 	Name:   "add-user",
				// 	Usage:  "adduser --email <emailaddress> --iamid <cloudid>",
				// 	Action: addUser,
				// 	Flags:  userFlags,
				// },
			},
		},
		{
			Name:     "namespace",
			Aliases:  []string{"name"},
			Category: "Console",
			Usage:    "managing openshift namespaces/projects",
			Subcommands: []*cli.Command{
				{
					Name:    "list",
					Aliases: []string{"ls"},
					Usage:   "list",
					Action:  listNamespace,
				},
				{
					Name:    "switch",
					Aliases: []string{"sw"},
					Usage:   "switch namespace",
					Action:  switchNamespace,
					//Flags:   userFlags,
				},
				{
					Name:    "add",
					Aliases: []string{"a"},
					Usage:   "add namespace",
					Action:  addNamespace,
					//Flags:   userFlags,
				},
				{
					Name:    "delete",
					Aliases: []string{"del"},
					Usage:   "delete namespace",
					Action:  deleteNamespace,
					//Flags:   userFlags,
				},
			},
		},
		{
			Name: "conn",
			//Aliases:  []string{"c"},
			Category: "Connections",
			Usage:    "connecting to your ROKS CP4i Cluster",
			Subcommands: []*cli.Command{
				{
					Name:        "datapower-admin",
					Usage:       "datapower-admin pod --localport=port",
					Description: "Connect to DataPower WebGUI",
					Action:      datapowerConnect,
					//Category: "Console",
				},
				{
					Name:   "docker-login",
					Usage:  "Perform Docker Login to Local Registry",
					Action: dockerLogin,
					//Category: "Console",
				},
				{
					Name:   "registry-svc",
					Usage:  "Port Forward to Local Registry",
					Action: registryConnect,
					Flags:  userFlags,
				},
			},
		},
	}
}

func main() {

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.UseShortOptionHandling = true
	app.HideHelp = true
	app.HideVersion = false
	app.Name = "roks-cli"
	app.Version = "v0.1"
	app.Compiled = time.Now()
	app.Authors = []*cli.Author{
		{Name: "Jamil Spain", Email: "jamil.spain@us.ibm.com"},
	}

	app.Commands = cmds

	app.Usage = "Command Line Utility RedHat OpenShift with Cloud Pak for Integration"
	app.Description = "Working with RedHat OpenShift and Cloud Pak for Integration? This cli program assists with the configuration and management of both products."

	app.Action = noArgs
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func act(c *cli.Context) error {
	return cli.NewExitError("no arguments for subcommand", 3)
}

func validCapability(name string) bool {
	var isValid bool = false

	switch name {
	case "ace", "tracing", "mq", "datapower", "apic", "assetrepo", "aspera", "eventstreams", "es":
		isValid = true
	}

	return isValid
}

func noArgs(c *cli.Context) error {
	cli.ShowAppHelp(c)
	fmt.Println(" ")
	return nil
	//cli.NewExitError("no commands provided", 2)
}

func noConnection() error {
	//fmt.printf("")
	return nil
}
