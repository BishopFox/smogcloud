package service

import (
    "fmt"
    "github.com/BishopFox/smogcloud/util"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/elb"
    "github.com/aws/aws-sdk-go/service/elbv2"
    "github.com/remeh/sizedwaitgroup"
    "strings"
)

func GetElb(sess *session.Session, region string, group *sizedwaitgroup.SizedWaitGroup) {
    defer group.Done()
    results := util.CreateResultsObject()

    hostnames := []string{}
    uris := []string{}

    svc := elb.New(sess, &aws.Config{Region: aws.String(region)})
    for {
        listParams := &elb.DescribeLoadBalancersInput{}
        result, err := svc.DescribeLoadBalancers(listParams)
        if err != nil {
            fmt.Println("\n Cannot list functions for region " + region)
            fmt.Println(err)
        }

        for _, loadBalancer := range result.LoadBalancerDescriptions {
            if loadBalancer.CanonicalHostedZoneName != nil {
                hostnames = append(hostnames, *loadBalancer.CanonicalHostedZoneName)
                for _, listener := range loadBalancer.ListenerDescriptions {
                    uri := strings.ToLower(fmt.Sprintf("%s://%s:%d", *listener.Listener.Protocol, *loadBalancer.CanonicalHostedZoneName, *listener.Listener.LoadBalancerPort))
                    uris = append(uris, uri)
                }
            }
        }
        if result.NextMarker != nil {
            listParams = &elb.DescribeLoadBalancersInput{
                Marker: result.NextMarker,
            }
        } else {
            break
        }
    }

    svcV2 := elbv2.New(sess, &aws.Config{Region: aws.String(region)})

    for {
        listParams := &elbv2.DescribeLoadBalancersInput{}
        result, err := svcV2.DescribeLoadBalancers(listParams)
        if err != nil {
            fmt.Println("\n Cannot list functions for region " + region)
            fmt.Println(err)
        }

        for _, loadBalancer := range result.LoadBalancers {
            if *loadBalancer.Scheme != "internal" {
                hostnames = append(hostnames, *loadBalancer.DNSName)
            }
        }
        if result.NextMarker != nil {
            listParams = &elbv2.DescribeLoadBalancersInput{
                Marker: result.NextMarker,
            }
        } else {
            break
        }
    }

    results.Hostnames = util.UniqueStrings(hostnames)
    results.URIs = util.UniqueStrings(uris)

    util.Save("elb", region, results)

}
