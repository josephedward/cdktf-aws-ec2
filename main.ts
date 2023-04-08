import { Construct } from "constructs";
import { App, TerraformStack, TerraformOutput } from "cdktf";
import { AwsProvider, } from "@cdktf/provider-aws/lib/provider";
import { Instance, } from "@cdktf/provider-aws/lib/instance";

class MyStack extends TerraformStack {
  constructor(scope: Construct, id: string) {
    super(scope, id);

    new AwsProvider(this, "AWS", {
      region: "us-east-1",
    });

// ec2 instance without a group name
    const ec2Instance = new Instance(this, "compute", {
      ami: "ami-06e46074ae430fba6",
      instanceType: "t2.micro",
      subnetId: "subnet-06003e038d6766cf4",
    });

    new TerraformOutput(this, "public_ip", {
      value: ec2Instance.publicIp,
    });
  }
}

const app = new App();
const stack = new MyStack(app, "aws_instance");
stack.prepareStack();

app.synth();
