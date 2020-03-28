package service

import (
    "fmt"
    "github.com/BishopFox/smogcloud-research/util"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/lightsail"
    "github.com/remeh/sizedwaitgroup"
)

func GetLightsail(sess *session.Session, region string, group *sizedwaitgroup.SizedWaitGroup) {
    defer group.Done()
    results := util.CreateResultsObject()

    hostnames := []string{}
    ips := []string{}

    svc := lightsail.New(sess, &aws.Config{Region: aws.String(region)})

    for {
        listParams := &lightsail.GetInstancesInput{}

        result, err := svc.GetInstances(listParams)
        if err != nil {
            fmt.Println(err)
        }
        for _, instance := range result.Instances {
            ips = append(ips, *instance.PublicIpAddress)

        }

        if result.NextPageToken != nil {
            listParams = &lightsail.GetInstancesInput{
                PageToken: result.NextPageToken,
            }
        } else {
            break
        }
    }

    for {
        listParams := &lightsail.GetLoadBalancersInput{}

        result, err := svc.GetLoadBalancers(listParams)
        if err != nil {
            fmt.Println(err)
        }
        for _, loadBalancer := range result.LoadBalancers {
            hostnames = append(hostnames, *loadBalancer.DnsName)
        }

        if result.NextPageToken != nil {
            listParams = &lightsail.GetLoadBalancersInput{
                PageToken: result.NextPageToken,
            }
        } else {
            break
        }
    }

    results.Hostnames = util.UniqueStrings(hostnames)
    results.IPs = util.UniqueStrings(ips)

    util.Save("lightsail", region, results)

}
