package service

import (
    "fmt"
    "github.com/BishopFox/smogcloud-research/util"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/eks"
    "github.com/remeh/sizedwaitgroup"
)

func GetEks(sess *session.Session, region string, group *sizedwaitgroup.SizedWaitGroup) {
    defer group.Done()
    results := util.CreateResultsObject()

    hostnames := []string{}
    uris := []string{}

    svc := eks.New(sess, &aws.Config{Region: aws.String(region)})

    listParams := &eks.ListClustersInput{}

    for {
        result, err := svc.ListClusters(listParams)
        if err != nil {
            fmt.Println(err)
        }

        for _, item := range result.Clusters {
            listParams := &eks.DescribeClusterInput{
                Name: item,
            }
            cluster, err := svc.DescribeCluster(listParams)
            if err != nil {
                fmt.Println(err)
            }
            if cluster.Cluster != nil {
                uris = append(uris, *cluster.Cluster.Endpoint)
                hostnames = append(hostnames, util.GetHostnameFromUrl(*cluster.Cluster.Endpoint))
            }
        }

        if result.NextToken != nil {
            listParams = &eks.ListClustersInput{
                NextToken: result.NextToken,
            }
        } else {
            break
        }
    }

    results.Hostnames = util.UniqueStrings(hostnames)
    results.URIs = util.UniqueStrings(uris)

    util.Save("eks", region, results)
}
