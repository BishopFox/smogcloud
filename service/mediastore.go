package service

import (
    "fmt"
    "github.com/BishopFox/smogcloud/util"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/mediastore"
    "github.com/remeh/sizedwaitgroup"
    "strings"
)

func GetMediaStore(sess *session.Session, region string, group *sizedwaitgroup.SizedWaitGroup) {
    defer group.Done()
    uris := []string{}
    results := util.CreateResultsObject()

    svc := mediastore.New(sess, &aws.Config{Region: aws.String(region)})
    result, err := svc.ListContainers(nil)
    if err != nil {
        fmt.Println(err)
    }

    for _, container := range result.Containers {
        domain := strings.ToLower(util.GetHostnameFromUrl(*container.Endpoint))
        fmt.Println(domain)
        results.Hostnames = append(results.Hostnames, domain)
        uri := strings.ToLower(*container.Endpoint)

        uris = append(uris, uri)
    }

    results.URIs = util.UniqueStrings(uris)

    util.Save("mediastore", region, results)

}
