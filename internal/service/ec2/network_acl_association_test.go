package ec2_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func init() {
	resource.AddTestSweepers("aws_network_acl_association", &resource.Sweeper{
		Name: "aws_network_acl_association",
		F:    sweepNetworkAcls,
	})
}

func TestAccNetworkAclAssociation_basic(t *testing.T) {

	var networkAcl ec2.NetworkAcl
	resourceName := "aws_network_acl.acl_a"

	//lintignore:XAT001
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { acctest.PreCheck(t) },
		IDRefreshName: resourceName,
		Providers:     acctest.Providers,
		CheckDestroy:  testAccCheckNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkAclAssoc,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckNetworkACLExists(resourceName, &networkAcl),
					testAccCheckSubnetIsAssociatedWithAcl(resourceName, "aws_subnet.subnet_a"),
				),
			},
		},
	})
}

const testAccNetworkAclAssoc = `
resource "aws_vpc" "test_vpc" {
  cidr_block = "10.1.0.0/16"
  tags = {
    Name = "testAccAWSNetworkAclEsp"
  }
}
resource "aws_network_acl" "acl_a" {
  vpc_id = aws_vpc.test_vpc.id
  tags = {
    Name = "terraform test"
  }
}
resource "aws_subnet" "subnet_a" {
  vpc_id     = aws_vpc.test_vpc.id
  cidr_block = "10.1.33.0/24"
  tags = {
    Name = "terraform test"
  }
}
resource "aws_network_acl_association" "test" {
  network_acl_id = aws_network_acl.acl_a.id
  subnet_id      = aws_subnet.subnet_a.id
}
`
