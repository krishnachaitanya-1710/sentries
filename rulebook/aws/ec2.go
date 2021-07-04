package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/krishnachaitanya-1710/sentries/logger"
)

type elasticComputeCloud struct{}

func init() {
	InstMap["elasticComputeCloud"] = new(elasticComputeCloud)
}

func (e elasticComputeCloud) ExecuteRules(resources, skipRules []string) {

}

func setEc2Session(region string) *ec2.EC2 {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})
	return svc
}

func getVpcConfiguration(vpcId string) *ec2.DescribeVpcsOutput {
	client := setEc2Session(awsRegion)
	describeVpcsInput := &ec2.DescribeVpcsInput{
		VpcIds: []*string{
			aws.String(vpcId),
		},
	}
	response, err := client.DescribeVpcs(describeVpcsInput)
	if err != nil {
		logger.Fatal("unable to describe vpc configurations " + err.Error())
	}
	return response
}

func describeSecurityGroup(securityGroupIds []string) *ec2.DescribeSecurityGroupsOutput {
	client := setEc2Session(awsRegion)
	describeVpcsInput := &ec2.DescribeSecurityGroupsInput{
		GroupIds: aws.StringSlice(securityGroupIds),
	}
	response, err := client.DescribeSecurityGroups(describeVpcsInput)
	if err != nil {
		logger.Fatal("unable to describe security group " + err.Error())
	}
	return response
}

func isVpcDefault(vpcId string) *bool {
	client := setEc2Session(awsRegion)
	describeVpcsInput := &ec2.DescribeVpcsInput{
		VpcIds: []*string{
			aws.String(vpcId),
		},
	}
	response, err := client.DescribeVpcs(describeVpcsInput)
	if err != nil {
		logger.Fatal("unable to describe vpc configurations " + err.Error())
	}
	return response.Vpcs[0].IsDefault
}

func getSubnetConfiguration(vpcId string) *ec2.DescribeSubnetsOutput {
	client := setEc2Session(awsRegion)
	describeSubnetInput := &ec2.DescribeSubnetsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("vpc-id"),
				Values: []*string{
					aws.String(vpcId),
				},
			},
		},
	}
	response, err := client.DescribeSubnets(describeSubnetInput)
	if err != nil {
		logger.Fatal("unable to describe subnet configurations " + err.Error())
	}
	return response
}

func describeRouteTables(vpcId string) *ec2.DescribeRouteTablesOutput {
	client := setEc2Session(awsRegion)
	input := &ec2.DescribeRouteTablesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("vpc-id"),
				Values: []*string{
					aws.String(vpcId),
				},
			},
		},
	}
	response, err := client.DescribeRouteTables(input)
	if err != nil {
		logger.Fatal("unable to describe subnet configurations " + err.Error())
	}
	return response
}

//func getSubnetRouteConfiguration(subnetIds []string) *ec2.DescribeSubnetsOutput {
//	client := setEc2Session()
//	describeSubnetInput := &ec2.DescribeRouteTablesInput{
//		SubnetIds: subnetIds,
//	}
//	response, err := client.DescribeRouteTables(context.TODO(), describeSubnetInput)
//	if err != nil {
//		logger.Fatal(err.Error())
//	}
//	if response != nil {
//		fmt.Println(response.NextToken)
//		fmt.Println(err)
//	}
//	return response
//}

// Create an Amazon S3 Client fo v2
/*
func setEc2Session() *ec2.Client {
 cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithDefaultRegion("us-east-1"))
 if err != nil {
   logger.Fatal(err.Error())
 }
 client := ec2.NewFromConfig(cfg)
 return client
}
*/
