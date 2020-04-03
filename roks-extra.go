package main

// oc adm policy add-scc-to-group ibm-mq-integration-scc system:serviceaccounts:mq;
// oc adm policy add-scc-to-group ibm-anyuid-scc system:serviceaccounts:mq;
// oc adm policy add-scc-to-group ibm-anyuid-scc system:serviceaccounts:tracing;
// oc adm policy add-scc-to-group ibm-anyuid-scc system:serviceaccounts:ace;
// oc adm policy add-scc-to-group ibm-anyuid-hostpath-scc system:serviceaccounts:apic;
// oc adm policy add-scc-to-group ibm-anyuid-hostaccess-scc system:serviceaccounts:aspera;
// oc adm policy add-scc-to-group  ibm-restricted-scc  system:servieaccounts:eventstreams;
// oc adm policy add-scc-to-group ibm-privileged-scc system:serviceaccounts:assetrepo;

// app.
// 	app.Action = func(c *cli.Context) error {
// 	fmt.Println("boom! I say!")
// 	return nil
// }

// app.Commands = []cli.Command{
// 	{
// 		Name:    "add",
// 		Aliases: []string{"a"},
// 		Usage:   "add a task to the list",
// 		Action: func(c *cli.Context) error {
// 			fmt.Println("added task: ", c.Args().First())
// 			return nil
// 		},
// 	},
// 	{
// 		Name:    "complete",
// 		Aliases: []string{"c"},
// 		Usage:   "complete a task on the list",
// 		Action: func(c *cli.Context) error {
// 			fmt.Println("completed task: ", c.Args().First())
// 			return nil
// 		},
// 	},
// 	{
// 		Name:    "template",
// 		Aliases: []string{"t"},
// 		Usage:   "options for task templates",
// 		Subcommands: []cli.Command{
// 			{
// 				Name:  "add",
// 				Usage: "add a new template",
// 				Action: func(c *cli.Context) error {
// 					fmt.Println("new task template: ", c.Args().First())
// 					return nil
// 				},
// 			},
// 			{
// 				Name:  "remove",
// 				Usage: "remove an existing template",
// 				Action: func(c *cli.Context) error {
// 					fmt.Println("removed task template: ", c.Args().First())
// 					return nil
// 				},
// 			},
// 		},
// 	},
// }

// Flags: []cli.Flag{
// 	// cli.BoolFlag{Name: "serve, s"},
// 	// cli.BoolFlag{Name: "option, o"},
// 	cli.StringFlag{Name: "message, m"},
//},

// userFlags = []cli.Flag{
// 	cli.StringFlag{
// 		Name:        "Full Name",
// 		Value:       "First Last",
// 		Destination: &fullname,
// 	},
// 	//cli.StringFlag{Name: "serve, s"},
// 	// cli.StringFlag{Name: "serve, s"},
// 	// cli.StringFlag{Name: "option, o"},
// 	// cli.StringFlag{Name: "message, m"},
// }

// {
// 	Name:     "users",
// 	Aliases:  []string{"haxorx", "pwn"},
// 	Usage:    "hack target",
// 	Action:   act,
// 	Category: "Offline Installer",
// },
// {
// 	Name:     "template",
// 	Aliases:  []string{"t"},
// 	Usage:    "options for task templates",
// 	Category: "Sample",
// 	Subcommands: []*cli.Command{
// 		{
// 			Name:  "add-template",
// 			Usage: "add a new template",
// 			Action: func(c *cli.Context) error {
// 				fmt.Println("new task template: ", c.Args().First())
// 				return nil
// 			},
// 		},
// 		{
// 			Name:  "remove",
// 			Usage: "remove an existing template",
// 			Action: func(c *cli.Context) error {
// 				fmt.Println("removed task template: ", c.Args().First())
// 				return nil
// 			},
// 		},
// 	},
// },

//
/*

 2020/02/14 20:03:39 Terraform apply | null_resource.script (local-exec): namespace/mq created
 2020/02/14 20:03:39 Terraform apply | null_resource.script (local-exec): secret/ibm-entitlement-key created
 2020/02/14 20:03:39 Terraform apply | null_resource.script (local-exec): --- creating the scc for mq named ibm-mq-integration-scc
 2020/02/14 20:03:41 Terraform apply | null_resource.script (local-exec): securitycontextconstraints.security.openshift.io/ibm-mq-integration-scc created
 2020/02/14 20:03:41 Terraform apply | null_resource.script (local-exec): --- changing the scc of the mq to ibm-mq-integration-scc
 2020/02/14 20:03:41 Terraform apply | null_resource.script (local-exec): scc "ibm-mq-integration-scc" added to groups: ["system:serviceaccounts:mq"]
 2020/02/14 20:03:42 Terraform apply | null_resource.script (local-exec): namespace/ace created
 2020/02/14 20:03:43 Terraform apply | null_resource.script (local-exec): secret/ibm-entitlement-key created
 2020/02/14 20:03:43 Terraform apply | null_resource.script (local-exec): --- changing the scc of the ace to ibm-anyuid-scc
 2020/02/14 20:03:44 Terraform apply | null_resource.script (local-exec): scc "ibm-anyuid-scc" added to groups: ["system:serviceaccounts:ace"]
 2020/02/14 20:03:45 Terraform apply | null_resource.script (local-exec): namespace/apic created
 2020/02/14 20:03:45 Terraform apply | null_resource.script (local-exec): secret/ibm-entitlement-key created
 2020/02/14 20:03:45 Terraform apply | null_resource.script (local-exec): --- changing the scc of the apic to ibm-anyuid-hostpath-scc
 2020/02/14 20:03:46 Terraform apply | null_resource.script (local-exec): scc "ibm-anyuid-hostpath-scc" added to groups: ["system:serviceaccounts:apic"]
 2020/02/14 20:03:47 Terraform apply | null_resource.script: Still creating... (10s elapsed)
 2020/02/14 20:03:47 Terraform apply | null_resource.script (local-exec): namespace/eventstreams created
 2020/02/14 20:03:48 Terraform apply | null_resource.script (local-exec): secret/ibm-entitlement-key created
 2020/02/14 20:03:48 Terraform apply | null_resource.script (local-exec): --- changing the scc of the eventstreams to ibm-restricted-scc
 2020/02/14 20:03:49 Terraform apply | null_resource.script (local-exec): scc "ibm-restricted-scc" added to groups: ["system:serviceaccounts:eventstreams"]
 2020/02/14 20:03:50 Terraform apply | null_resource.script (local-exec): namespace/aspera created
 2020/02/14 20:03:50 Terraform apply | null_resource.script (local-exec): secret/ibm-entitlement-key created
 2020/02/14 20:03:50 Terraform apply | null_resource.script (local-exec): --- changing the scc of the aspera to ibm-anyuid-scc
 2020/02/14 20:03:51 Terraform apply | null_resource.script (local-exec): scc "ibm-anyuid-hostaccess-scc" added to groups: ["system:serviceaccounts:aspera"]
 2020/02/14 20:03:52 Terraform apply | null_resource.script (local-exec): namespace/datapower created
 2020/02/14 20:03:53 Terraform apply | null_resource.script (local-exec): secret/ibm-entitlement-key created
 2020/02/14 20:03:53 Terraform apply | null_resource.script (local-exec): --- changing the scc of the datapower to ibm-anyuid-scc
 2020/02/14 20:03:54 Terraform apply | null_resource.script (local-exec): scc "ibm-anyuid-scc" added to groups: ["system:serviceaccounts:datapower"]
 2020/02/14 20:03:54 Terraform apply | null_resource.script (local-exec): namespace/assetrepo created
 2020/02/14 20:03:55 Terraform apply | null_resource.script (local-exec): secret/ibm-entitlement-key created
 2020/02/14 20:03:55 Terraform apply | null_resource.script (local-exec): --- changing the scc of the assetrepo to ibm-privileged-scc
 2020/02/14 20:03:56 Terraform apply | null_resource.script (local-exec): scc "ibm-privileged-scc" added to groups: ["system:serviceaccounts:assetrepo"]
 2020/02/14 20:03:57 Terraform apply | null_resource.script (local-exec): namespace/tracing created
  2020/02/14 20:03:57 Terraform apply | null_resource.script (local-exec): namespace/tracing created
 2020/02/14 20:03:57 Terraform apply | null_resource.script: Still creating... (20s elapsed)
 2020/02/14 20:03:58 Terraform apply | null_resource.script (local-exec): secret/ibm-entitlement-key created
 2020/02/14 20:03:58 Terraform apply | null_resource.script (local-exec): --- changing the scc of the tracing to anyuid
 2020/02/14 20:03:58 Terraform apply | null_resource.script (local-exec): scc "anyuid" added to groups: ["system:serviceaccounts:tracing"]
*/
