package main

import (
	"github.com/pulumi/pulumi-eks/sdk/go/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		cfg := config.New(ctx, "")
		version := cfg.Require("version")
		stackName := ctx.Stack()

		cluster, err := eks.NewCluster(ctx, stackName, &eks.ClusterArgs{
			CreateOidcProvider: pulumi.Bool(true),
			Version:            pulumi.String(version),
		})
		if err != nil {
			return err
		}

		ctx.Export("kubeconfig", cluster.Kubeconfig)
		return nil
	})
}
