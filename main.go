package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	profileName := "xxx"
	session := session.Must(session.NewSessionWithOptions(session.Options{Profile: profileName}))
	svc := ec2.New(session, aws.NewConfig().WithRegion("ap-northeast-1"))
	output, err := svc.DescribeSecurityGroupRules(&ec2.DescribeSecurityGroupRulesInput{})
	if err != nil {
		log.Fatal(err)
	}

	var fullOpen []*ec2.SecurityGroupRule
	for _, rule := range output.SecurityGroupRules {
		if rule.CidrIpv4 != nil {
			if *rule.IsEgress == false && *rule.IpProtocol == "tcp" && *rule.CidrIpv4 == "0.0.0.0/0" {
				fullOpen = append(fullOpen, rule)
			}
		}
	}

	fmt.Println(fullOpen)
}
