package cmd

import (
	"fmt"
	//"io/ioutil"
	"os"
	//"path/filepath"

	"clairvoyance/log"
	"github.com/spf13/cobra"

	//"clairvoyance/app/reporting"
	"clairvoyance/app/terraform"
)

/*
In order for a report to be done, a tfexec config should be populated and we need to ensure the following
values have been captured.

The following options for additional reporting functionality.
	clairvoyance report:
		--command <show/plan/apply> (Performs limited Terraform CLI logic, a more comprehensive report behaviour is used)
		--path <working_directory>
		--output [<discord>, <stdout>]

		TODO: *what does a config file look like, where is this loaded from? (based off tfexc cfg?)
		--config <clairvoyance_config>

	clairvoyance report --path ~/noobshack --output discord
	clairvoyance report --command show --path ~/noobshack --output stdout
*/

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Reports terraform drift to Discord",
	Long: `Reports terraform drift to Discord
		Usage:
		clairvoyance report`,

	Run: func(cmd *cobra.Command, args []string) {
		// Figure out where to send data
		optOutput, _ := cmd.Flags().GetString("output")

		// Get version of Terraform binary to use
		//var binaryDir = os.Getenv("GOPATH") + "/src/clairvoyance/tfinstall/terraform_0.13.2"
		var tfBinary = "/usr/bin/terraform"
		//tfBinary := terraform.DetectBinary(binaryDir, terraformVersion)

		// Setup Terraform Version to use
		var _, tfVersionSet = os.LookupEnv("CLAIRVOYANCE_TERRAFORM_VERSION")
		var terraformVersion string
		_ = terraformVersion

		if tfVersionSet {
			terraformVersion = os.Getenv("CLAIRVOYANCE_TERRAFORM_VERSION")
		} else {
			// should be "" or "latest" - will hardcode to latest version for now
			terraformVersion = "0.14.3"
		}

		// Setup projects to plan
		// Point to directory container directories to plan
		//var workingDir = os.Getenv("CLAIRVOYANCE_WORKING_DIR")

		// for each dir in CLAIRVOYANCE_WORKING_DIR with *.tf, add dir to list
		var projects []string
		//projects := terraform.PopulateProjectList(workingDir)

		/*
			//terraformDir = os.Getenv("CLAIRVOYANCE_WORKING_DIR")
			terraformDir = "/home/reulan/noobshack/infrastructure/deploy"

			projectFileInfo, err := ioutil.ReadDir(terraformDir)
			if err != nil {
				log.Fatal(err)
			}

				for _, project := range projectFileInfo {
					var absServicePath = fmt.Sprintf("%s/%s", terraformDir, project.Name())
					projects = append(projects, absServicePath)
				}
		*/

		projects = []string{
			"/home/reulan/noobshack/gameservers/rust",
			"/home/reulan/noobshack/gameservers/csgo",
			"/home/reulan/noobshack/gameservers/minecraft",
			"/home/reulan/noobshack/infrastructure/bootstrap/cluster/noobshack/ingress-controller",
			"/home/reulan/noobshack/infrastructure/deploy/atlantis",
		}

		var terraformServices []*terraform.TerraformService

		for _, absProjectPath := range projects {
			//absProjectPath, _ := filepath.Abs(absProjectPath)

			// terraform init
			service := terraform.ConfigureTerraform(absProjectPath, tfBinary)
			terraform.Init(service)

			// terraform show
			state := terraform.Show(service)

			// terraform plan
			// TODO: tf plan (with -out=out.tfplan)
			planOptions := fmt.Sprintf("-out=%s/out.tfplan", absProjectPath)
			var po []string = []string{planOptions}
			fmt.Println(po)
			var isPlanned bool = terraform.Plan(service)
			planPath := fmt.Sprintf("%s/out.tfplan", absProjectPath)
			var rawPlan = terraform.ShowPlanFileRaw(service, planPath)
			//log.Printf("rawPlan: %s", rawPlan)
			planString := terraform.ResourceModificationCount(rawPlan)
			modifiedResourceCount := terraform.ParseModificationCount(planString)
			summary := terraform.DriftDetection(isPlanned, state)

			_, projectName := terraform.GetProjectName(absProjectPath)
			tfService := terraform.ExtractDriftReportData(state, projectName, modifiedResourceCount, summary)

			terraformServices = append(terraformServices, tfService)
		}

		terraform.CreateTableStdout(terraformServices)
		//terraform.ResourceAddressList(state)

		// Where is the message going?
		if optOutput == "discord" {
			//log.Println("Outputting to Discord.")
			//reporting.SendMessageDiscord(message)
		} else if optOutput == "stdout" {
			//log.Println("Outputting to Stdout.")
		} else {
			log.Errorf("cmd/report - optOutput: [%s] not supported (discord, stdout)", optOutput)
		}
	},
}

func init() {
	fmt.Println("cmd/report/go running.")
	rootCmd.AddCommand(reportCmd)
	//reportCmd.Flags().StringP("command", "c", "show", "Performs a specific Terraform command against the given project. (defaults to Show)")
	//reportCmd.Flags().StringP("path", "p", "/path/to/terraform/project", "Specify the path of the Terraform project you'd like to report on")
	reportCmd.Flags().StringP("output", "o", "discord", "Choose the target medium to report to. (discord, stdout)")
}
