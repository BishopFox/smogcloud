package service

import (
    "fmt"
    "github.com/BishopFox/smogcloud-research/util"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/cloudfront"
    "github.com/remeh/sizedwaitgroup"
)

func GetCloudFront(sess *session.Session, region string, group *sizedwaitgroup.SizedWaitGroup) {
    defer group.Done()
    results := util.CreateResultsObject()

    hostnames := []string{}
    aliases := []string{}
    origins := []string{}

    svc := cloudfront.New(sess, &aws.Config{})

    listParams := &cloudfront.ListDistributionsInput{}

    for {
        result, err := svc.ListDistributions(listParams)
        if err != nil {
            fmt.Println(err)
        }
        if result.DistributionList != nil {
            for _, item := range result.DistributionList.Items {
                hostnames = append(hostnames, *item.DomainName)
                for _, alias := range item.Aliases.Items {
                    aliases = append(aliases, *alias)
                }
                for _, origin := range item.Origins.Items {
                    origins = append(origins, *origin.DomainName)
                }

            }

            if *result.DistributionList.IsTruncated == true {
                listParams = &cloudfront.ListDistributionsInput{
                    Marker: result.DistributionList.NextMarker,
                }
            } else {
                break
            }
        } else {
            break
        }
    }
    hostnames = append(hostnames, aliases...)
    hostnames = append(hostnames, origins...)
    results.Hostnames = util.UniqueStrings(hostnames)

    util.Save("cloudfront", region, results)
}
