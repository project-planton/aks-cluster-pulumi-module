package outputs

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/azure/akscluster"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

func PulumiOutputsToStackOutputsConverter(pulumiOutputs auto.OutputMap,
	input *akscluster.AksClusterStackInput) *akscluster.AksClusterStackOutputs {
	return &akscluster.AksClusterStackOutputs{}
}
