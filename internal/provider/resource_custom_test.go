package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceCustom(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCustom,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"custom_resource.foo", "state", regexp.MustCompile("^qwe")),
				),
			},
		},
	})
}

const testAccResourceCustom = `
resource "custom_resource" "foo" {
  input = "qwe"
  program_create = ["/usr/bin/sh", "-x", "-c", "echo $EXT_DIR && /usr/bin/cat $EXT_FILE_input > $EXT_FILE_state && /usr/bin/sha256sum $EXT_FILE_input > $EXT_FILE_id"]
  program_read = ["/usr/bin/sh", "-c", "cat $EXT_FILE_state"]
  program_update = ["/usr/bin/sh", "-c", "cat $EXT_FILE_input > $EXT_FILE_state && /usr/bin/sha256sum $EXT_FILE_input > $EXT_FILE_id"]
  program_delete = ["/usr/bin/sh", "-c", "echo -n > $EXT_FILE_state"]
}
`