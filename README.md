# Smog Cloud

Find cloud assets that no one wants exposed

## AWS Patterns

These are the patterns of exposure URIs that you may find in your AWS accounts

s3
    https://{user_provided}.s3.amazonaws.com
cloudfront
    https://{random_id}.cloudfront.net
ec2
    ec2-{ip-seperated}.compute-1.amazonaws.com
es
    https://{user_provided}-{random_id}.{region}.es.amazonaws.com
elb
    http://{user_provided}-{random_id}.{region}.elb.amazonaws.com:80
    https://{user_provided}-{random_id}.{region}.elb.amazonaws.com:443
elbv2
    https://{user_provided}-{random_id}.{region}.elb.amazonaws.com
rds
    mysql://{user_provided}.{random_id}.{region}.rds.amazonaws.com:3306
    postgres://{user_provided}.{random_id}.{region}.rds.amazonaws.com:5432
route53
    {user_provided}
execute-api
    https://{random_id}.execute-api.{region}.amazonaws.com/{user_provided}
cloudsearch
    https://doc-{user_provided}-{random_id}.{region}.cloudsearch.amazonaws.com
transfer
    sftp://s-{random_id}.server.transfer.{region}.amazonaws.com
iot 
    mqtt://{random_id}.iot.{region}.amazonaws.com:8883
    https://{random_id}.iot.{region}.amazonaws.com:8443
    https://{random_id}.iot.{region}.amazonaws.com:443
mq
    https://b-{random_id}-{1,2}.mq.{region}.amazonaws.com:8162
    ssl://b-{random_id}-{1,2}.mq.{region}.amazonaws.com:61617
kafka
    b-{1,2,3,4}.{user_provided}.{random_id}.c{1,2}.kafka.{region}.amazonaws.com
    {user_provided}.{random_id}.c{1,2}.kafka.useast-1.amazonaws.com
cloud9
    https://{random_id}.vfs.cloud9.{region}.amazonaws.com
mediastore
    https://{random_id}.data.mediastore.{region}.amazonaws.com.
kinesisvideo
    https://{random_id}.kinesisvideo.{region}.amazonaws.com
mediaconvert
    https://{random_id}.mediaconvert.{region}.amazonaws.com
mediapackage
    https://{random_id}.mediapackage.{region}.amazonaws.com/in/v1/{random_id}/channel

## Authors

* **Oscar Salazar** - *Initial work* - [Bishop Fox](https://github.com/tracertea)
* **Rob Ragan** - *Initial work* - [Bishop Fox](https://github.com/basicScandal)

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc

