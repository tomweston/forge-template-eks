package main

import (
	"strconv"

	"github.com/pulumi/pulumi-eks/sdk/go/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		cfg := config.New(ctx, "")

		desiredCapacity, err := strconv.Atoi(cfg.Require("desiredCapacity"))
		if err != nil {
			return err
		}

		minSize, err := strconv.Atoi(cfg.Require("minSize"))
		if err != nil {
			return err
		}

		maxSize, err := strconv.Atoi(cfg.Require("maxSize"))
		if err != nil {
			return err
		}

		version := cfg.Require("version")

		cluster, err := eks.NewCluster(ctx, "cluster", &eks.ClusterArgs{
			DesiredCapacity:    pulumi.Int(desiredCapacity),
			MinSize:            pulumi.Int(minSize),
			MaxSize:            pulumi.Int(maxSize),
			CreateOidcProvider: pulumi.Bool(true),
			Version:            pulumi.String(version),
			Fargate:            pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		ctx.Export("kubeconfig", cluster.Kubeconfig)
		return nil
	})
}
