package service

import (
    "fmt"
    "github.com/BishopFox/smogcloud-research/util"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/rds"
    "github.com/remeh/sizedwaitgroup"
)

func GetRds(sess *session.Session, region string, group *sizedwaitgroup.SizedWaitGroup) {
    defer group.Done()
    results := util.CreateResultsObject()

    hostnames := []string{}
    uris := []string{}

    svc := rds.New(sess, &aws.Config{Region: aws.String(region)})

    for {
        listParams := &rds.DescribeDBInstancesInput{}
        result, err := svc.DescribeDBInstances(listParams)
        if err != nil {
            fmt.Println("\n Cannot list functions for region " + region)
            fmt.Println(err)
        }

        for _, db := range result.DBInstances {
            if *db.PubliclyAccessible {
                if db.Endpoint != nil {
                    hostnames = append(hostnames, *db.Endpoint.Address)
                    scheme := ""
                    if *db.Engine == "aurora" {
                        scheme = "mysql"
                    }
                    if *db.Engine == "aurora-mysql" {
                        scheme = "mysql"
                    }
                    if *db.Engine == "aurora-postgresql" {
                        scheme = "postgres"
                    }
                    if *db.Engine == "mariadb" {
                        scheme = "mariadb"
                    }
                    if *db.Engine == "mysql" {
                        scheme = "mysql"
                    }
                    if *db.Engine == "oracle-ee" {
                        scheme = "oracle"
                    }
                    if *db.Engine == "oracle-se2" {
                        scheme = "oracle"
                    }
                    if *db.Engine == "oracle-se1" {
                        scheme = "oracle"
                    }
                    if *db.Engine == "oracle-se" {
                        scheme = "oracle"
                    }
                    if *db.Engine == "postgres" {
                        scheme = "postgres"
                    }
                    if *db.Engine == "sqlserver-ee" {
                        scheme = "mssql"
                    }
                    if *db.Engine == "sqlserver-se" {
                        scheme = "mssql"
                    }
                    if *db.Engine == "sqlserver-ex" {
                        scheme = "mssql"
                    }
                    if *db.Engine == "sqlserver-web" {
                        scheme = "mssql"
                    }
                    uri := fmt.Sprintf("%s://%s:%d", scheme, *db.Endpoint.Address, *db.Endpoint.Port)
                    uris = append(uris, uri)
                }
            }
        }
        if result.Marker != nil {
            listParams = &rds.DescribeDBInstancesInput{
                Marker: result.Marker,
            }
        } else {
            break
        }
    }

    results.Hostnames = util.UniqueStrings(hostnames)
    results.URIs = util.UniqueStrings(uris)

    util.Save("rds", region, results)

}
