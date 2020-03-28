package service

import (
    "fmt"
    "github.com/BishopFox/smogcloud/util"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/elasticbeanstalk"
    "github.com/remeh/sizedwaitgroup"
    "strings"
)

func GetElasticBeanstalk(sess *session.Session, region string, group *sizedwaitgroup.SizedWaitGroup) {
    defer group.Done()
    results := util.CreateResultsObject()
    hostnames := []string{}
    uris := []string{}

    svc := elasticbeanstalk.New(sess, &aws.Config{Region: aws.String(region)})
    result, err := svc.DescribeEnvironments(nil)
    if err != nil {
        fmt.Println(err)
    }

    for _, environment := range result.Environments {

        hostname := fmt.Sprintf("%s", strings.ToLower(*environment.CNAME))
        hostnames = append(hostnames, hostname)
        uri := fmt.Sprintf("%s/%s", "http://", hostname)
        uris = append(uris, uri)
    }

    results.Hostnames = util.UniqueStrings(hostnames)
    results.URIs = util.UniqueStrings(uris)

    util.Save("elasticbeanstalk", region, results)

}
