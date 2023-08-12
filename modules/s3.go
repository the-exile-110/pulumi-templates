package modules

import (
	"fmt"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateS3Bucket(ctx *pulumi.Context, projectName, env, resourceName string) (*s3.Bucket, error) {
	bucketName := fmt.Sprintf("%s-%s-%s", projectName, env, resourceName)
	return s3.NewBucket(ctx, bucketName, nil)
}
