package main

import (
	aksclusterv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/azure/akscluster/v1"
	"github.com/pkg/errors"
	"github.com/project-planton/aks-cluster-pulumi-module/pkg"
	"github.com/project-planton/pulumi-module-golang-commons/pkg/stackinput"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &aksclusterv1.AksClusterStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		return pkg.Resources(ctx, stackInput)
	})
}
