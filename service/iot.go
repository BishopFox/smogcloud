package service

import (
    "fmt"
    "github.com/BishopFox/smogcloud-research/util"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/iot"
    "github.com/remeh/sizedwaitgroup"
)

func GetIoT(sess *session.Session, region string, group *sizedwaitgroup.SizedWaitGroup) {
    defer group.Done()
    results := util.CreateResultsObject()
    hostnames := []string{}
    svc := iot.New(sess, &aws.Config{Region: aws.String(region)})
    result, err := svc.ListDomainConfigurations(nil)
    if err != nil {
        fmt.Println(err)
    }

    for _, domainConfiguration := range result.DomainConfigurations {

        listParams := &iot.DescribeDomainConfigurationInput{
            DomainConfigurationName: domainConfiguration.DomainConfigurationName,
        }

        resultConfig, err := svc.DescribeDomainConfiguration(listParams)
        if err != nil {
            fmt.Println(err)
        }
        hostnames = append(hostnames, *resultConfig.DomainName)
    }

    results.Hostnames = util.UniqueStrings(hostnames)

    util.Save("iot", region, results)
}
