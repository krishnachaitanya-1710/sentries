package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/krishnachaitanya-1710/sentries/globalvar"
	"github.com/krishnachaitanya-1710/sentries/logger"
	"github.com/krishnachaitanya-1710/sentries/utilities"
	"os"
)

type s3Bucket struct{}

func init() {
	InstMap["s3Bucket"] = new(s3Bucket)
}

func (s s3Bucket) ExecuteRules(resources, skipRules []string) {
	buckets := s3ListBuckets()
	for _, bucket := range buckets {
		logger.DebugS("Applying compliance rules against :: " + bucket)
		region := gets3BucketRegion(bucket)
		bucketArn := "arn:aws:s3:::" + bucket
		s3DefaultEncryptionCheck(bucket, region)
		s3VersioningCheck(bucket, region)
		globalvar.ResourceCount++
		globalvar.ResourceArns = append(globalvar.ResourceArns, bucketArn)
	}
}

// Create an Amazon S3 Client
func setS3BucketSession(region string) *s3.S3 {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		logger.Fatal(err.Error())
	}
	client := s3.New(sess)
	return client
}

func gets3BucketRegion(bucket string) string {
	sess := session.Must(session.NewSession())
	ctx := context.Background()
	region, err := s3manager.GetBucketRegion(ctx, sess, bucket, awsRegion)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "NotFound" {
			fmt.Fprintf(os.Stderr, "unable to find bucket %s's region not found\n", bucket)
		}
	}
	return region
}

func s3ListBuckets() []string {
	client := setS3BucketSession(awsRegion)
	getInput := &s3.ListBucketsInput{}
	response, _ := client.ListBuckets(getInput)
	var bucketList []string
	for _, b := range response.Buckets {
		bucketName := aws.StringValue(b.Name)
		bucketList = append(bucketList, bucketName)
	}
	return bucketList
}

func s3DefaultEncryptionCheck(bucket, region string) {
	//Load the Shared Aws Configuration (~/.aws/config)
	client := setS3BucketSession(region)
	getInput := &s3.GetBucketEncryptionInput{
		Bucket: aws.String(bucket),
	}
	response, _ := client.GetBucketEncryption(getInput)
	if response != nil && response.ServerSideEncryptionConfiguration != nil {
		algorithm := **&response.ServerSideEncryptionConfiguration.Rules[0].ApplyServerSideEncryptionByDefault.SSEAlgorithm
		logger.DebugS(algorithm)
		utilities.ExecuteRule(
			algorithm == "AES256",
			"s3 server side encryption rule",
			"Encryption feature is not enabled for s3 Bucket "+bucket,
		)
	} else {
		logger.StreamRed("fail   - s3 server side encryption rule")
	}
}

func s3VersioningCheck(bucket, region string) {
	// Example sending a request using the GetBucketEncryptionRequest method.
	client := setS3BucketSession(region)
	getInput := &s3.GetBucketVersioningInput{
		Bucket: aws.String(bucket),
	}
	response, err := client.GetBucketVersioning(getInput)
	if err != nil {
		logger.Fatal("Unable to get bucket versioning info due to - " + err.Error())
	}
	if response != nil && response.Status != nil {
		versioning := **&response.Status
		utilities.ExecuteRule(
			versioning == "Enabled",
			"s3 version rule",
			"Versioning feature is not enabled for s3 Bucket "+bucket,
		)
	} else {
		logger.StreamRed("fail   - s3 version rule")
	}
}
