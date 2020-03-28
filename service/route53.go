package service

import (
    "fmt"
    "github.com/BishopFox/smogcloud-research/util"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/route53"
    "github.com/remeh/sizedwaitgroup"
    "strings"
)

func GetRoute53(sess *session.Session, region string, group *sizedwaitgroup.SizedWaitGroup) {
    defer group.Done()
    results := util.CreateResultsObject()

    hostnames := []string{}
    aliases := []string{}

    svc := route53.New(sess, &aws.Config{})

    result, err := svc.ListHostedZones(nil)
    if err != nil {
        fmt.Println(err)
    }

    for _, zone := range result.HostedZones {
        if *zone.Config.PrivateZone == false {

            zoneId := zone.Id
            listParams := &route53.ListResourceRecordSetsInput{
                HostedZoneId: aws.String(*zoneId),
            }

            for {
                listOutput, _ := svc.ListResourceRecordSets(listParams)
                fmt.Printf("%+v", listOutput)
                for _, record := range listOutput.ResourceRecordSets {
                    hostname := *record.Name
                    if strings.HasSuffix(hostname, ".") {
                        hostname = hostname[:len(hostname)-1]
                    }
                    hostnames = append(hostnames, hostname)
                    if record.AliasTarget != nil {
                        aliases = append(aliases, *record.AliasTarget.DNSName)
                    }
                }
                if *listOutput.IsTruncated == true {
                    listParams = &route53.ListResourceRecordSetsInput{
                        HostedZoneId:          aws.String(*zoneId),
                        StartRecordName:       listOutput.NextRecordName,
                        StartRecordType:       listOutput.NextRecordType,
                        StartRecordIdentifier: listOutput.NextRecordIdentifier,
                    }
                } else {
                    break
                }
            }
        }
    }

    hostnames = append(hostnames, aliases...)
    results.Hostnames = util.UniqueStrings(hostnames)

    util.Save("route53", region, results)

}
