package main

import (
	"fmt"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"pulumi-templates/modules"
)

const (
	projectName = "test-project"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// 获取当前堆栈名称（环境）
		stack := ctx.Stack()

		// 根据当前环境创建S3 bucket
		bucket, err := modules.CreateS3Bucket(ctx, projectName, stack, "bucket")
		if err != nil {
			return fmt.Errorf("failed to create S3 bucket: %v", err)
		}

		// 输出S3 bucket的ARN
		ctx.Export("bucketARN", bucket.Arn)
		return nil
	})
}
