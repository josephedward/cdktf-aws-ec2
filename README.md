# AWS EC2 via CDKTF 

This is a project to generate AWS EC2 resources via CDKTF.

**From :**
https://developer.hashicorp.com/terraform/tutorials/cdktf/cdktf-build

**Notes**
- removed remote backend for demo purposes, will use in future for state management in more complex examples 
- had to find an AMI specific to the region I was using
- had to add a subnet id from the default VPC (was using a sandboxed account)
- repeated same steps on golang branch
