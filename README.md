#☁️ Smogcloud 

Find exposed AWS cloud assets that you did not know you had. A comprehensive asset inventory is step one to any capable security program. We made smogcloud to enable security engineers, penetration testers, and AWS administrators to monitor the collective changes that create dynamic and ephemeral internet-facing assets on a more frequent basis. May be useful to identify:

 - Internet-facing FQDNs and IPs across one or hundreds of AWS accounts
 - Misconfigurations or vulnerabilities
 - Assets that are no longer in use
 - Services not currently monitored 
 - Shadow IT

##🛠Getting Started
1. Install and setup go
2. Install smogcloud using the following command

    ``` 
        go get -u github.com/BishopFox/smogcloud
    ```
3. Set up aws environment variable for the account you wish to query. We suggest utilizing a read-only [Security Auditor](https://medium.com/@HorosAWSData/how-to-add-an-aws-user-with-security-audit-access-819f0aef7cee) role. The following commands can be used to set environment variables:

    ```
    export AWS_ACCOUNT_ID=''            # Describe account
    export AWS_ACCESS_KEY_ID=''         # Access key for aws account
    export AWS_SECRET_ACCESS_KEY=''     # Secret key for aws account
    ```

4. Run the application

    ```
    smogcloud
    ```
    or
    ```
    go run main.go
    ```

## Current Services
Supported services for extracting internet exposures:

    * API Gateway
    * CloudFront
    * EC2
    * Elastic Kubernetes Service
    * Elastic Beanstalk
    * Elastic Search
    * Elastic Load Balancing 
    * IoT
    * Lightsail
    * MediaStore
    * Relational Database Service
    * Redshift
    * Route53
    * S3

##🔎 AWS Patterns

From studying Open API documentation on RESTful AWS endpoints we determined these are the patterns of exposure URIs that you may find in AWS accounts. It is important to understand how to interact with these native services to test them for vulnerabilities and other misconfigurations. Security engineers may want to monitor Cloudtrail logs or build DNS monitoring for requests to these services. 

- s3
  - https://{user_provided}.s3.amazonaws.com
- cloudfront
  - https://{random_id}.cloudfront.net
- ec2
  - ec2-{ip-seperated}.compute-1.amazonaws.com
- es
  - https://{user_provided}-{random_id}.{region}.es.amazonaws.com
- elb
  - http://{user_provided}-{random_id}.{region}.elb.amazonaws.com:80
  - https://{user_provided}-{random_id}.{region}.elb.amazonaws.com:443
- elbv2
  - https://{user_provided}-{random_id}.{region}.elb.amazonaws.com
- rds
  - mysql://{user_provided}.{random_id}.{region}.rds.amazonaws.com:3306
  - postgres://{user_provided}.{random_id}.{region}.rds.amazonaws.com:5432
- route53
  - {user_provided}
- execute-api
  - https://{random_id}.execute-api.{region}.amazonaws.com/{user_provided}
- cloudsearch
  - https://doc-{user_provided}-{random_id}.{region}.cloudsearch.amazonaws.com
- transfer
  - sftp://s-{random_id}.server.transfer.{region}.amazonaws.com
- iot 
  - mqtt://{random_id}.iot.{region}.amazonaws.com:8883
  - https://{random_id}.iot.{region}.amazonaws.com:8443
  - https://{random_id}.iot.{region}.amazonaws.com:443
- mq
  - https://b-{random_id}-{1,2}.mq.{region}.amazonaws.com:8162
  - ssl://b-{random_id}-{1,2}.mq.{region}.amazonaws.com:61617
- kafka
  - b-{1,2,3,4}.{user_provided}.{random_id}.c{1,2}.kafka.{region}.amazonaws.com
  - {user_provided}.{random_id}.c{1,2}.kafka.{region}.amazonaws.com
- cloud9
  - https://{random_id}.vfs.cloud9.{region}.amazonaws.com
- mediastore
  - https://{random_id}.data.mediastore.{region}.amazonaws.com.
- kinesisvideo
  - https://{random_id}.kinesisvideo.{region}.amazonaws.com
- mediaconvert
  - https://{random_id}.mediaconvert.{region}.amazonaws.com
- mediapackage
  - https://{random_id}.mediapackage.{region}.amazonaws.com/in/v1/{random_id}/channel

##📌 References
* [AWS SDK Go](https://docs.aws.amazon.com/sdk-for-go/api/)
* [API-guru Open API for AWS](https://github.com/APIs-guru/openapi-directory/tree/master/APIs/amazonaws.com)
* [aws-cli](https://github.com/aws/aws-cli)

##👨‍💻  Authors

* **Oscar Salazar** - *Initial work* - [Bishop Fox](https://github.com/tracertea)
* **Rob Ragan** - *Initial work* - [Bishop Fox](https://github.com/basicScandal) [🐦@sweepthatleg](https://twitter.com/sweepthatleg)
* **Brandon Gaudet** - *Initial work* - [Bishop Fox](https://github.com/brandondgaudet)

##📣 Acknowledgments

Thank you for inspiration
* [Cloudmapper](https://github.com/duo-labs/cloudmapper)
* [AWS Public IPs](https://github.com/arkadiyt/aws_public_ips)
* [John Backes & Tiros](https://aws.amazon.com/blogs/security/aws-security-profile-john-backes-senior-software-development-engineer/)
* [IAM Access Analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/what-is-access-analyzer.html)
* [Cartography](https://github.com/lyft/cartography)

## License

Smogcloud is licensed under GPLv3, some subcomponents have seperate licenses. 
