package aws

import (
	logger "../../logger"
	utilities "../../utilities"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Bucket struct{}

func init() {
	InstMap["s3Bucket"] = new(s3Bucket)
}

func (s s3Bucket) ExecuteRules(resources string) {
	bucket := "test2-xyx-hell"
	s3DefaultEncryptionCheck(bucket)
	s3VersioningCheck(bucket)
	s3PublicAclCheck(bucket)
}

// Create an Amazon S3 Client
func setS3BucketSession() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Fatal(err.Error())
	}
	client := s3.NewFromConfig(cfg)
	return client
}

func s3DefaultEncryptionCheck(bucket string) {
	//Load the Shared Aws Configuration (~/.aws/config)
	client := setS3BucketSession()
	getInput := &s3.GetBucketEncryptionInput{
		Bucket: aws.String(bucket),
	}
	response, err := client.GetBucketEncryption(context.TODO(), getInput)
	if err != nil {
		logger.Fatal("Unable to get bucket Encryption due to - " + err.Error())
	}
	if response != nil && response.ServerSideEncryptionConfiguration != nil {
		algorithm := aws.ToString((*string)(&response.ServerSideEncryptionConfiguration.Rules[0].ApplyServerSideEncryptionByDefault.SSEAlgorithm))
		logger.DebugS(algorithm)
		utilities.ExecuteRule(
			algorithm == "AES256",
			"s3 server side encryption rule",
			"Encryption feature is not enabled for s3 Bucket "+bucket,
		)
	} else {
		logger.StreamRed("Unable to get bucket Encryption Information")
	}
}

func s3VersioningCheck(bucket string) {
	// Example sending a request using the GetBucketEncryptionRequest method.
	client := setS3BucketSession()
	getInput := &s3.GetBucketVersioningInput{
		Bucket: aws.String(bucket),
	}
	response, err := client.GetBucketVersioning(context.TODO(), getInput)
	if err != nil {
		logger.Fatal("Unable to get bucket versioning info due to - " + err.Error())
	}
	if response != nil {
		versioning := aws.ToString((*string)(&response.Status))
		utilities.ExecuteRule(
			versioning == "Enabled",
			"s3 version rule",
			"Versioning feature is not enabled for s3 Bucket "+bucket,
		)
	} else {
		logger.StreamRed("Response does not have versioning information -" + err.Error())
	}
}

func s3PublicAclCheck(bucket string) {
	client := setS3BucketSession()
	getInput := &s3.GetBucketAclInput{
		Bucket: aws.String(bucket),
	}
	response, err := client.GetBucketAcl(context.TODO(), getInput)
	fmt.Println(response.Grants)
	if err != nil {
		logger.Fatal("Unable to get bucket acl due to - " + err.Error())
	}
	//if response != nil && response.ServerSideEncryptionConfiguration != nil {
	//	algorithm := aws.ToString((*string)(&response.ServerSideEncryptionConfiguration.Rules[0].ApplyServerSideEncryptionByDefault.SSEAlgorithm))
	//	logger.DebugS(algorithm)
	//	utilities.ExecuteRule(
	//		algorithm == "AES256",
	//		"pass	- s3 server side encryption rule",
	//		"fail	- s3 server side encryption rule",
	//		"Encryption feature is not enabled for s3 Bucket " + bucket,
	//	)
	//} else {
	//	logger.StreamRed("Unable to get bucket Encryption Information")
	//}

}
