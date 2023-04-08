package main

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/instance"
	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v10/provider"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
    stack := cdktf.NewTerraformStack(scope, &id)

    awsprovider.NewAwsProvider(stack, jsii.String("AWS"), &awsprovider.AwsProviderConfig{
        Region: jsii.String("us-east-1"),
    })

    instance := instance.NewInstance(stack, jsii.String("compute"), &instance.InstanceConfig{
        Ami:          jsii.String("ami-06e46074ae430fba6"),
        InstanceType: jsii.String("t2.micro"),
		SubnetId:     jsii.String("subnet-06003e038d6766cf4"),
    })

    cdktf.NewTerraformOutput(stack, jsii.String("public_ip"), &cdktf.TerraformOutputConfig{
        Value: instance.PublicIp(),
    })

    return stack
}

func main() {
    app := cdktf.NewApp(nil)
    stack := NewMyStack(app, "aws_instance")
	fmt.Println("stack:", stack)
    // cdktf.NewRemoteBackend(stack, &cdktf.RemoteBackendProps{
    //     Hostname:     jsii.String("app.terraform.io"),
    //     Organization: jsii.String("<YOUR_ORG>"),
    //     Workspaces:   cdktf.NewNamedRemoteWorkspace(jsii.String("learn-cdktf")),
    // })

    app.Synth()
}
