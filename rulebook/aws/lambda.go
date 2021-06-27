package aws

//
//import (
//	logger "../../logger"
//	utilities "../../utilities"
//	"github.com/aws/aws-sdk-go/aws"
//	"github.com/aws/aws-sdk-go/aws/session"
//	"github.com/aws/aws-sdk-go/service/lambda"
//)
//
//type lambdaFunction struct {}
//
//func init(){
//	InstMap["lambdaFunction"] = new(lambdaFunction)
//}
//
//func (r *lambdaFunction) ExecuteRules(resources string) {
//	logger.Info("checking for lambda")
//	functionName := "oh-test-function"
//	region := "us-east-1"
//	lambdaComplianceRulesCheck(functionName, region)
//}
//
//func lambdaComplianceRulesCheck(functionName, region string) {
//	// Example sending a request using the GetBucketEncryptionRequest method.
//	sess := session.Must(session.NewSessionWithOptions(session.Options{
//		SharedConfigState: session.SharedConfigEnable,
//	}))
//	svc := lambda.New(sess, &aws.Config{Region: aws.String(region)})
//
//	getInput := &lambda.GetFunctionInput{
//		FunctionName: &functionName,
//	}
//	response, err := svc.GetFunction(getInput)
//	logger.InfoS(response.String())
//	if err != nil {
//		logger.ErrorS(err.Error())
//	}
//	if response != nil && response.Configuration != nil {
//		lambdaConfig := response.Configuration
//		utilities.ExecuteRule(
//			lambdaConfig.VpcConfig != nil,
//			"pass	- lambda public rule",
//			"fail	- lambda public rule",
//			"Lambda doesn't have vpc configuration - can be public",
//		)
//		utilities.ExecuteRule(
//			lambdaConfig.KMSKeyArn != nil,
//			"pass	- lambda kmsKey encryption rule",
//			"fail	- lambda kmsKey encryption rule",
//			"Lambda is not encrypted",
//		)
//		//check for security group rules
//		//check for environment variables encryption
//		//check for iam rules
//		//check for kms key role policy
//	}
//}
//
