package service

import (
    "fmt"
    "github.com/BishopFox/smogcloud/util"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/remeh/sizedwaitgroup"
)

func GetS3(sess *session.Session, region string, group *sizedwaitgroup.SizedWaitGroup) {
    defer group.Done()
    results := util.CreateResultsObject()

    hostnames := []string{}

    svc := s3.New(sess, &aws.Config{})
    result, err := svc.ListBuckets(nil)
    if err != nil {
        fmt.Println(err)
    }

    for _, bucket := range result.Buckets {
        name := *bucket.Name
        hostnames = append(hostnames, fmt.Sprintf("%s.s3.amazonaws.com", name))
    }

    results.Hostnames = util.UniqueStrings(hostnames)

    util.Save("s3", region, results)
}
