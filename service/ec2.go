package service

import (
    "fmt"
    "github.com/BishopFox/smogcloud-research/util"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/remeh/sizedwaitgroup"
)

func GetEc2(sess *session.Session, region string, group *sizedwaitgroup.SizedWaitGroup) {
    defer group.Done()
    results := util.CreateResultsObject()

    hostnames := []string{}
    ips := []string{}
    natIps := []string{}

    svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})

    for {
        listParams := &ec2.DescribeInstancesInput{}
        result, err := svc.DescribeInstances(listParams)
        if err != nil {
            fmt.Println("\n Cannot list functions for region " + region)
            fmt.Println(err)
        }

        for _, reservation := range result.Reservations {
            for _, instance := range reservation.Instances {
                if instance.PublicIpAddress != nil {
                    ips = append(ips, *instance.PublicIpAddress)
                }
                if instance.PublicDnsName != nil {
                    if *instance.PublicDnsName != "" {
                        hostnames = append(hostnames, *instance.PublicDnsName)
                    }
                }
            }
        }

        if result.NextToken != nil {
            listParams = &ec2.DescribeInstancesInput{
                NextToken: result.NextToken,
            }
        } else {
            break
        }
    }

    for {
        listParams := &ec2.DescribeNatGatewaysInput{}
        result, err := svc.DescribeNatGateways(listParams)
        if err != nil {
            fmt.Println("\n Cannot list functions for region " + region)
            fmt.Println(err)
        }

        for _, natGateways := range result.NatGateways {
            for _, address := range natGateways.NatGatewayAddresses {
                natIps = append(natIps, *address.PublicIp)
            }
        }
        if result.NextToken != nil {
            listParams = &ec2.DescribeNatGatewaysInput{
                NextToken: result.NextToken,
            }
        } else {
            break
        }
    }

    listParams := &ec2.DescribeVpcEndpointServiceConfigurationsInput{}
    result, err := svc.DescribeVpcEndpointServiceConfigurations(listParams)
    if err != nil {
        fmt.Println("\n Cannot list functions for region " + region)
        fmt.Println(err)
    }

    for _, reservation := range result.ServiceConfigurations {
        for _, instance := range reservation.BaseEndpointDnsNames {
            hostnames = append(hostnames, *instance)
        }
    }

    ips = append(ips, natIps...)

    results.Hostnames = util.UniqueStrings(hostnames)
    results.IPs = util.UniqueStrings(ips)

    util.Save("ec2", region, results)
}
