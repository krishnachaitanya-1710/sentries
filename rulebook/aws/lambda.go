package aws

import (
	logger "../../logger"
	utilities "../../utilities"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type lambdaFunction struct{}

func init() {
	InstMap["lambdaFunction"] = new(lambdaFunction)
}

func (r *lambdaFunction) ExecuteRules(resources string) {
	logger.Info("checking for lambda")
	//functionName := "oh-test-function"
	//region := "us-east-1"
	//allowedRegions := "us-east-1"
	lambdas := listLambdaFunctions([]string{"us-east-1"})

	for _, lambdas := range lambdas {
		//fmt.Println(aws.StringValue(lambdas.FunctionName))
		lambdaComplianceRulesCheck(aws.StringValue(lambdas.FunctionName), "us-east-1")
	}

}

func listLambdaFunctions(allowedRegions []string) []*lambda.FunctionConfiguration {
	svc := setNewSession("us-east-1")
	input := &lambda.ListFunctionsInput{}
	response, err := svc.ListFunctions(input)
	fmt.Println(err)

	return response.Functions
}

func lambdaComplianceRulesCheck(functionName, region string) {
	// Example sending a request using the GetBucketEncryptionRequest method.
	svc := setNewSession(region)

	getInput := &lambda.GetFunctionInput{
		FunctionName: &functionName,
	}
	response, err := svc.GetFunction(getInput)
	//logger.InfoS(response.String())
	if err != nil {
		logger.ErrorS(err.Error())
	}
	if response != nil && response.Configuration != nil {
		lambdaConfig := response.Configuration
		utilities.ExecuteRule(
			lambdaConfig.VpcConfig != nil,
			"lambda function with out vpc rule",
			"Lambda Function is not using VPC configuration",
		)
		if lambdaConfig.VpcConfig != nil {
			//need to get the info about public subnets
			utilities.ExecuteRule(
				lambdaConfig.VpcConfig != nil,
				"lambda function public subnet rule",
				"Lambda Function is in a public subnet",
			)
		}
		//check for security group rules
		utilities.ExecuteRule(
			lambdaConfig.KMSKeyArn != nil,
			"lambda kmsKey encryption rule",
			"Lambda is not encrypted",
		)
		utilities.ExecuteRule(
			lambdaConfig.Environment.Variables != nil && lambdaConfig.KMSKeyArn != nil,
			"lambda environment variable encryption check rule",
			"Lambda environment variables are not encrypted",
		)
		//input := &lambda.ListTagsInput{
		//	Resource: aws.String(*lambdaConfig.FunctionArn),
		//}
		//res, err1 := svc.ListTags(input)
		//fmt.Println(utilities.GetKeysofMap(res.Tags))
		//fmt.Println(err1)
		//validateLambdaDependencies(lambdaConfig)
	}
}

func setNewSession(region string) *lambda.Lambda {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := lambda.New(sess, &aws.Config{Region: aws.String(region)})
	return svc
}

func validateLambdaDependencies(lambdaConfig *lambda.FunctionConfiguration) {
	vpcId := *lambdaConfig.VpcConfig.VpcId
	securityGroups := []string(aws.StringValueSlice(lambdaConfig.VpcConfig.SecurityGroupIds))
	securityGroupsConfig := describeSecurityGroup(securityGroups).SecurityGroups
	utilities.ExecuteRule(
		!*isVpcDefault(vpcId),
		"lambda function default vpc rule",
		"lambda function created in default vpc",
	)
	for _, value := range securityGroupsConfig {
		utilities.ExecuteRule(
			len(value.IpPermissions) == 0,
			"lambda security group ingress rule check",
			"Lambda security group is having ingress rules",
		)
		if value.IpPermissionsEgress != nil {
			egressRules := value.IpPermissionsEgress
			for _, egressRule := range egressRules {
				utilities.ExecuteRule(
					*egressRule.IpRanges[0].CidrIp != "0.0.0.0/0",
					"lambda security group egress rule check",
					"Lambda security group is having 0.0.0.0/0 egress rules",
				)
			}
		}
	}
}
