package service

import (
    "fmt"
    "github.com/BishopFox/smogcloud-research/util"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/apigateway"
    "github.com/aws/aws-sdk-go/service/apigatewayv2"
    "github.com/remeh/sizedwaitgroup"
    "strings"
)

func GetAPIGateway(sess *session.Session, region string, group *sizedwaitgroup.SizedWaitGroup) {
    defer group.Done()
    results := util.CreateResultsObject()
    hostnames := []string{}
    uris := []string{}

    svc := apigateway.New(sess, &aws.Config{Region: aws.String(region)})
    result, err := svc.GetRestApis(nil)
    if err != nil {
        fmt.Println(err)
    }

    for _, api := range result.Items {

        hostname := fmt.Sprintf("%s.execute-api.%s.amazonaws.com", strings.ToLower(*api.Id), region)
        hostnames = append(hostnames, hostname)
        params := &apigateway.GetStagesInput{
            RestApiId: api.Id,
        }
        stageData, err := svc.GetStages(params)
        if err != nil {
            fmt.Println(err)
        }
        for _, stage := range stageData.Item {
            uri := fmt.Sprintf("%s/%s/%s", "https://", hostname, *stage.StageName)
            uris = append(uris, uri)
        }
    }

    svcV2 := apigatewayv2.New(sess, &aws.Config{Region: aws.String(region)})
    result2, err := svcV2.GetApis(nil)
    if err != nil {
        fmt.Println(err)
    }
    for _, api := range result2.Items {
        hostname := fmt.Sprintf("%s.execute-api.%s.amazonaws.com", strings.ToLower(*api.ApiId), region)
        hostnames = append(hostnames, hostname)

        uris = append(uris, *api.ApiEndpoint)
    }

    results.Hostnames = util.UniqueStrings(hostnames)
    results.URIs = util.UniqueStrings(uris)

    util.Save("apigateway", region, results)

}
