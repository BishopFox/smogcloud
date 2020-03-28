package main

import (
    "github.com/BishopFox/smogcloud-research/service"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/remeh/sizedwaitgroup"
)

func main() {
    swg := sizedwaitgroup.New(5)
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // call global services
    swg.Add()
    go service.GetCloudFront(sess, "global", &swg)
    swg.Add()
    go service.GetRoute53(sess, "global", &swg)
    swg.Add()
    go service.GetS3(sess, "global", &swg)

    // loop through regions
    regions := []string{"us-east-1", "us-east-2", "us-west-1", "us-west-2", "ap-east-1", "ap-south-1", "ap-northeast-3", "ap-northeast-2", "ap-southeast-1", "ap-southeast-2", "ap-northeast-1", "ca-central-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-west-1", "eu-west-2", "eu-west-3", "eu-north-1", "me-south-1", "sa-east-1"}
    for _, region := range regions {
        swg.Add()
        go service.GetAPIGateway(sess, region, &swg)
        swg.Add()
        go service.GetEc2(sess, region, &swg)
        swg.Add()
        go service.GetEks(sess, region, &swg)
        swg.Add()
        go service.GetElasticBeanstalk(sess, region, &swg)
        swg.Add()
        go service.GetElasticSearch(sess, region, &swg)
        swg.Add()
        go service.GetElb(sess, region, &swg)
        swg.Add()
        go service.GetLightsail(sess, region, &swg)
        swg.Add()
        go service.GetMediaStore(sess, region, &swg)
        swg.Add()
        go service.GetRds(sess, region, &swg)
        swg.Add()
        go service.GetRedshift(sess, region, &swg)
        swg.Add()
        go service.GetIoT(sess, region, &swg)
    }
    swg.Wait()
}
