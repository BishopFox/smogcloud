package service

import (
	"fmt"
	"github.com/BishopFox/smogcloud/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/remeh/sizedwaitgroup"
)

func GetRedshift(sess *session.Session, region string, group *sizedwaitgroup.SizedWaitGroup) {
	defer group.Done()
	results := util.CreateResultsObject()

	hostnames := []string{}
	uris := []string{}
	dbnames := []string{}

	svc := redshift.New(sess, &aws.Config{Region: aws.String(region)})

	for {
		listParams := &redshift.DescribeClustersInput{}
		result, err := svc.DescribeClusters(listParams)
		if err != nil {
			fmt.Println("\n Cannot list functions for region " + region)
			fmt.Println(err)
		}

		for _, cluster := range result.Clusters {
			if *cluster.PubliclyAccessible {
				if cluster.Endpoint != nil {
					hostnames = append(hostnames, *cluster.Endpoint.Address)

					uri := fmt.Sprintf("%s://%s:%d", "jdbc:redshift", *cluster.Endpoint.Address, *cluster.Endpoint.Port)
					uris = append(uris, uri)

					dbnames = append(dbnames, *cluster.DBName)

				}
			}
		}

		if result.Marker != nil {
			listParams = &redshift.DescribeClustersInput{
				Marker: result.Marker,
			}
		} else {
			break
		}
	}

	results.Hostnames = util.UniqueStrings(hostnames)
	results.URIs = util.UniqueStrings(uris)

	util.Save("redshift", region, results)

}
