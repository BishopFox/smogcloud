package service

import (
    "fmt"
    "github.com/BishopFox/smogcloud-research/util"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/elasticsearchservice"
    "github.com/remeh/sizedwaitgroup"
)

func GetElasticSearch(sess *session.Session, region string, group *sizedwaitgroup.SizedWaitGroup) {
    defer group.Done()
    results := util.CreateResultsObject()

    hostnames := []string{}
    uris := []string{}

    svc := elasticsearchservice.New(sess, &aws.Config{Region: aws.String(region)})

    result, err := svc.ListDomainNames(nil)
    if err != nil {
        fmt.Println("\n Cannot list functions for region " + region)
        fmt.Println(err)
    }
    for _, f := range result.DomainNames {
        param := &elasticsearchservice.DescribeElasticsearchDomainInput{
            DomainName: f.DomainName,
        }
        domainDescription, _ := svc.DescribeElasticsearchDomain(param)
        if domainDescription.DomainStatus != nil {
            if domainDescription.DomainStatus.Endpoint != nil {
                hostnames = append(hostnames, *domainDescription.DomainStatus.Endpoint)
                uris = append(uris, fmt.Sprintf("https://%s", *domainDescription.DomainStatus.Endpoint))
                uris = append(uris, fmt.Sprintf("https://%s/_plugin/kibana/", *domainDescription.DomainStatus.Endpoint))
            } else if domainDescription.DomainStatus.Endpoints != nil {
                for _, domain := range domainDescription.DomainStatus.Endpoints {
                    hostnames = append(hostnames, *domain)
                    uris = append(uris, fmt.Sprintf("https://%s", *domain))
                    uris = append(uris, fmt.Sprintf("https://%s/_plugin/kibana/", *domain))
                }
            } else {
                fmt.Println("the domain description does not contain DomainStatus endpoints")
                fmt.Printf("%+v", domainDescription)
            }
        } else {
            fmt.Println("the domain description does not contain DomainStatus endpoints")
            fmt.Printf("%+v", domainDescription)
        }

    }

    results.Hostnames = util.UniqueStrings(hostnames)
    results.URIs = util.UniqueStrings(uris)

    util.Save("elasticsearch", region, results)
}
