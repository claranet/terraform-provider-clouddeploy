package ghost

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"ghost": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProviderImpl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("GHOST_USER"); v == "" {
		t.Fatal("GHOST_USER must be set for acceptance tests")
	}

	if v := os.Getenv("GHOST_PASSWORD"); v == "" {
		t.Fatal("GHOST_PASSWORD must be set for acceptance tests")
	}

	if v := os.Getenv("GHOST_ENDPOINT"); v == "" {
		t.Fatal("GHOST_ENDPOINT must be set for acceptance tests")
	}
}
