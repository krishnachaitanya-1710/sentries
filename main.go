package main

import (
	"encoding/json"
	"fmt"
	"github.com/krishnachaitanya-1710/sentries/cli"
	"github.com/krishnachaitanya-1710/sentries/globalvar"
	"github.com/krishnachaitanya-1710/sentries/logger"
	"github.com/krishnachaitanya-1710/sentries/rulebook/aws"
	"github.com/krishnachaitanya-1710/sentries/utilities"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	c := cli.NewCommand("inspect", "Inspects Cloud Infrastructure against pre-built rules")

	app := cli.New(globalvar.AppName + " - security and compliance scanner tool, built to scan cloud infrastructure and to run compliance and security rules to mark violations").
		WithOption(cli.NewOption("outFile", "Provide outFile name to store sentries response, defaulted to 'sentries_response.json'")).
		WithOption(cli.NewOption("resources", "Specify to apply sentries rules against specific resources instead all resources")).
		WithOption(cli.NewOption("skip-rules", "Provides violation context for failed compliance rules")).
		WithOption(cli.NewOption("violation-context", "Provides violation context for failed compliance rules")).
		WithCommand(c)

	args := os.Args[1:]
	processVersion(args)
	inspect()
	if len(args) < 1 {
		os.Exit(app.Run(os.Args, os.Stdout))
	}
}

func inspect() {
	if utilities.IsFlagPassed("inspect") {
		for _, service := range aws.InstMap {
			service.ExecuteRules(globalvar.Resources, globalvar.SkipRules)
		}
		generateResponseFile()
		logger.InfoS("Sentries applied " + strconv.Itoa(globalvar.RuleCount) + " compliance rules against " + strconv.Itoa(globalvar.ResourceCount) + " resources")
	} else if utilities.IsFlagPassed("inspect") {
		logger.ErrorS("Inspect flag is missing in the arguments")
	}
}

func processVersion(args []string) {
	for _, arg := range args {
		if arg == "-v" || arg == "-version" || arg == "version" || arg == "--version" {
			fmt.Println("Sentries v" + globalvar.Version)
			break
		}
	}
}

func generateResponseFile() {
	data := globalvar.ComplianceSummary{
		ComplianceStatus:     globalvar.ComplianceStatus,
		TotalRulesApplied:    globalvar.RuleCount,
		TotalResourceScanned: globalvar.ResourceCount,
		RuleNames:            utilities.RemoveDuplicatesFromSlice(globalvar.RuleNames),
		ResourceArns:         globalvar.ResourceArns,
		ViolationContext:     globalvar.ViolationContext,
	}
	file, _ := json.MarshalIndent(data, "", "")
	var responseFileName = globalvar.ResponseFile
	for _, arg := range os.Args {
		if strings.Contains(arg, "outFile=") {
			responseFileName = strings.Split(arg, "=")[1]
		}
	}
	_ = ioutil.WriteFile(responseFileName, file, 0644)
}
