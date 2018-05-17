package ghost

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"cloud-deploy.io/go-st"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGhostAppBasic(t *testing.T) {
	resourceName := "ghost_app.test"
	envName := fmt.Sprintf("ghost_app_acc_env_basic_%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGhostAppDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGhostAppConfig(envName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckGhostAppExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", envName),
					resource.TestCheckResourceAttr(resourceName, "env", "dev"),
					resource.TestCheckResourceAttr(resourceName, "region", "eu-west-1"),
					resource.TestCheckResourceAttr(resourceName, "log_notifications.0", "ghost-devops@domain.com"),
					resource.TestCheckResourceAttr(resourceName, "autoscale.0.max", "3"),
					resource.TestCheckResourceAttr(resourceName, "environment_variables.0.key", "myvar"),
				),
			},
			{
				Config: testAccGhostAppConfigUpdated(envName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckGhostAppExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", envName),
					resource.TestCheckResourceAttr(resourceName, "env", "dev"),
					resource.TestCheckResourceAttr(resourceName, "region", "eu-west-2"),
					resource.TestCheckResourceAttr(resourceName, "log_notifications.0", "ghost-devops2@domain.com"),
					resource.TestCheckResourceAttr(resourceName, "autoscale.0.max", "2"),
					resource.TestCheckResourceAttr(resourceName, "environment_variables.0.key", "myvar2"),
				),
			},
			{
				Config: testAccGhostAppConfigOmitEmpty(envName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckGhostAppExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", envName),
					resource.TestCheckResourceAttr(resourceName, "env", "dev"),
					resource.TestCheckResourceAttr(resourceName, "autoscale.0.min", "0"),
					resource.TestCheckResourceAttr(resourceName, "autoscale.0.max", "0"),
					resource.TestCheckResourceAttr(resourceName, "instance_monitoring", "false"),
					resource.TestCheckResourceAttr(resourceName, "environment_infos.0.public_ip_address", "false"),
					resource.TestCheckResourceAttr(resourceName, "modules.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "features.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "environment_variables.#", "0"),
				),
			},
		},
	})
}

func testAccCheckGhostAppExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Ghost Application ID is set")
		}

		log.Printf("[INFO] Try to connect to Ghost and get all apps")
		client := testAccProvider.Meta().(*ghost.Client)
		_, err := client.GetApps()
		if err != nil {
			return fmt.Errorf("Ghost environment not reachable: %v", err)
		}

		return nil
	}
}

func testAccCheckGhostAppDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*ghost.Client)

	// Iterates through ghost apps
	for _, rs := range s.RootModule().Resources {
		// Skip resources that aren't ghost apps
		if rs.Type != "ghost_app" {
			continue
		}

		app_id := rs.Primary.ID

		// Try to get ghost app
		_, err := client.GetApp(app_id)
		if err == nil {
			return fmt.Errorf("[INFO] Ghost app still exists: " + app_id)
		}
	}

	return nil
}

func testAccGhostAppConfig(name string) string {
	return fmt.Sprintf(`
      resource "ghost_app" "test" {
        name        = "%s"
        env         = "dev"
        role        = "webfront"
        description = "This is a test"

        region        = "eu-west-1"
        instance_type = "t2.micro"
        vpc_id        = "vpc-3f1eb65a"

        log_notifications = [
          "ghost-devops@domain.com",
        ]

        build_infos = {
          subnet_id    = "subnet-a7e849fe"
          ssh_username = "admin"
          source_ami   = "ami-03ce4474"
        }

        environment_infos = {
          instance_profile  = "iam.ec2.demo"
          key_name          = "ghost-demo"
          root_block_device = {
            name = "testblockdevice"
            size = 20
          }
          optional_volumes  = [{
            device_name = "/dev/xvdd"
            volume_type = "gp2"
            volume_size = 20
          }]
          subnet_ids        = ["subnet-a7e849fe"]
          security_groups   = ["sg-6814f60c", "sg-2414f60c"]
          instance_tags     = [{
            tag_name  = "Name"
            tag_value = "wordpress"
          },
          {
            tag_name  = "Type"
            tag_value = "front"
          }]
        }

        autoscale = {
          name = "autoscale"
          min  = 1
          max  = 3
        }

        modules = [{
          name       = "wordpress"
          pre_deploy = ""
          path       = "/var/www"
          scope      = "code"
          git_repo   = "https://github.com/KnpLabs/KnpIpsum.git"
        },
        {
          name        = "wordpress2"
          pre_deploy  = "ZXhpdCAx"
          post_deploy = "ZXhpdCAx"
          path        = "/var/www-test.test"
          scope       = "code"
          git_repo    = "https://github.com/KnpLabs/KnpIpsum.git"
        }]

        features = [{
          version     = "5.4"
          name        = "php5"
          provisioner = "salt"
        },
        {
          version     = ""
          name        = "package"
          provisioner = "ansible"
          parameters  = <<JSON
            {
              "package_name" : [
                "test",
                "nano"
              ]
            }
            JSON
        }]

        lifecycle_hooks = {
          pre_buildimage  = "#!/usr/bin/env bash"
          post_buildimage = "#!/usr/bin/env bash"
        }

        environment_variables = [{
          key   = "myvar"
          value = "myvalue"
        }]
      }
      `, name)
}

func testAccGhostAppConfigUpdated(name string) string {
	return fmt.Sprintf(`
      resource "ghost_app" "test" {
        name = "%s"
        env  = "dev"
        role = "webfront"

        region        = "eu-west-2"
        instance_type = "t2.micro"
        vpc_id        = "vpc-3f1eb65a"

        log_notifications = [
          "ghost-devops2@domain.com",
        ]

        build_infos = {
          subnet_id    = "subnet-a7e849fe"
          ssh_username = "admin"
          source_ami   = "ami-03ce4474"
        }

        environment_infos = {
          instance_profile  = "iam.ec2.demo"
          key_name          = "ghost-demo"
          root_block_device = {
            name = "testblockdevice"
            size = 20
          }
          subnet_ids        = ["subnet-a7e849fe"]
          security_groups   = ["sg-6814f60c"]
          instance_tags     = [{
            tag_name  = "Name"
            tag_value = "wordpress"
          },
          {
            tag_name  = "Type"
            tag_value = "front"
          }]
        }

        autoscale = {
          name = "autoscale"
          min  = 1
          max  = 2
        }

        modules = [{
          name       = "wordpress"
          pre_deploy = ""
          path       = "/var/www"
          scope      = "code"
          git_repo   = "https://github.com/KnpLabs/KnpIpsum.git"
        },
        {
          name        = "wordpress2"
          pre_deploy  = "ZXhpdCAx"
          post_deploy = "ZXhpdCAx"
          path        = "/var/www"
          scope       = "code"
          git_repo    = "https://github.com/KnpLabs/KnpIpsum.git"
        }]

        features = [{
          version     = "5.4"
          name        = "php5"
          provisioner = "salt"
        },
        {
          version     = ""
          name        = "package"
          provisioner = "ansible"
          parameters  = <<JSON
            {
              "package_name" : [
                "test2",
                "nano"
              ]
            }
            JSON
        },
        {
          version     = "2.2"
          name        = "apache2"
          provisioner = "salt"
        }]

        lifecycle_hooks = {
          pre_buildimage  = "#!/usr/bin/env bash"
        }

        environment_variables = [{
          key   = "myvar2"
          value = "myvalue2"
        }]
      }
      `, name)
}

func testAccGhostAppConfigOmitEmpty(name string) string {
	return fmt.Sprintf(`
      resource "ghost_app" "test" {
        name = "%s"
        env  = "dev"
        role = "webfront"

        region        = "eu-west-2"
        instance_type = "t2.micro"
        vpc_id        = "vpc-3f1eb65a"

        instance_monitoring = false

        build_infos = {
          subnet_id    = "subnet-a7e849fe"
          ssh_username = "admin"
          source_ami   = "ami-03ce4474"
        }

        environment_infos = {
          instance_profile  = "iam.ec2.demo"
          key_name          = "ghost-demo"
          public_ip_address = false
          instance_tags     = [{
            tag_name  = "Name"
            tag_value = "wordpress"
          }]
        }

      modules = []

      }
      `, name)
}

// Variables used for unit tests
var (
	app = ghost.App{
		Name:               "app_name",
		Env:                "test",
		Role:               "web",
		Description:        "My app",
		Region:             "us-west-1",
		InstanceType:       "t2.micro",
		VpcID:              "vpc-123456",
		InstanceMonitoring: false,

		Modules: &[]ghost.Module{{
			Name:      "my_module",
			GitRepo:   "https://github.com/test/test.git",
			Scope:     "system",
			Path:      "/",
			BuildPack: StrToB64("#!/usr/bin/env bash"),
			PreDeploy: StrToB64("#!/usr/bin/env bash"),
		}},
		Features: &[]ghost.Feature{{
			Name:        "feature",
			Version:     "1.0",
			Provisioner: "ansible",
		}},
		Autoscale: &ghost.Autoscale{
			Name:          "autoscale",
			EnableMetrics: false,
			Min:           0,
			Max:           3,
		},
		BuildInfos: &ghost.BuildInfos{
			SshUsername: "admin",
			SourceAmi:   "ami-1",
			SubnetID:    "subnet-1",
		},
		EnvironmentInfos: &ghost.EnvironmentInfos{
			InstanceProfile: "profile",
			KeyName:         "key",
			PublicIpAddress: false,
			SecurityGroups:  []string{"sg-1", "sg-2"},
			SubnetIDs:       []string{"subnet-1", "subnet-2"},
			InstanceTags: &[]ghost.InstanceTag{{
				TagName:  "name",
				TagValue: "val",
			}},
			OptionalVolumes: &[]ghost.OptionalVolume{{
				DeviceName: "my_device",
				VolumeType: "gp2",
				VolumeSize: 20,
				Iops:       3000,
			}},
			RootBlockDevice: &ghost.RootBlockDevice{
				Name: "rootblock",
				Size: 20,
			},
		},
		LifecycleHooks: &ghost.LifecycleHooks{
			PreBuildimage:  StrToB64("#!/usr/bin/env bash"),
			PostBuildimage: StrToB64("#!/usr/bin/env bash"),
		},
		LogNotifications: []string{"log_not@email.com"},
		EnvironmentVariables: &[]ghost.EnvironmentVariable{{
			Key:   "env_var_key",
			Value: "env_var_value",
		}},
		SafeDeployment: &ghost.SafeDeployment{
			LoadBalancerType: "elb",
			WaitBeforeDeploy: 10,
			WaitAfterDeploy:  10,
		},
	}
)

// Expanders Unit Tests
func TestExpandGhostAppStringList(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput []string
	}{
		{
			[]interface{}{
				"1", "2", "3",
			},
			[]string{
				"1", "2", "3",
			},
		},
		{
			nil,
			[]string{},
		},
	}

	for _, tc := range cases {
		output := expandGhostAppStringList(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandGhostAppInstanceTags(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *[]ghost.InstanceTag
	}{
		{
			[]interface{}{
				map[string]interface{}{
					"tag_name":  "name",
					"tag_value": "val",
				},
			},
			app.EnvironmentInfos.InstanceTags,
		},
		{
			nil,
			&[]ghost.InstanceTag{},
		},
	}

	for _, tc := range cases {
		output := expandGhostAppInstanceTags(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandGhostAppOptionalVolume(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *[]ghost.OptionalVolume
	}{
		{
			[]interface{}{
				map[string]interface{}{
					"device_name": "my_device",
					"volume_type": "gp2",
					"volume_size": 20,
					"iops":        3000,
					"launch_block_device_mappings": false,
				},
			},
			app.EnvironmentInfos.OptionalVolumes,
		},
		{
			nil,
			&[]ghost.OptionalVolume{},
		},
	}

	for _, tc := range cases {
		output := expandGhostAppOptionalVolumes(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandGhostAppRootBlockDevice(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *ghost.RootBlockDevice
	}{
		{
			[]interface{}{
				map[string]interface{}{
					"name": "rootblock",
					"size": 20,
				},
			},
			app.EnvironmentInfos.RootBlockDevice,
		},
		{
			nil,
			nil,
		},
	}

	for _, tc := range cases {
		output := expandGhostAppRootBlockDevice(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandGhostAppEnvironmentInfos(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *ghost.EnvironmentInfos
	}{
		{
			[]interface{}{
				map[string]interface{}{
					"instance_profile":  "profile",
					"key_name":          "key",
					"public_ip_address": false,
					"security_groups":   []interface{}{"sg-1", "sg-2"},
					"subnet_ids":        []interface{}{"subnet-1", "subnet-2"},
					"instance_tags": []interface{}{
						map[string]interface{}{
							"tag_name":  "name",
							"tag_value": "val",
						},
					},
					"optional_volumes": []interface{}{
						map[string]interface{}{
							"device_name": "my_device",
							"volume_type": "gp2",
							"volume_size": 20,
							"iops":        3000,
							"launch_block_device_mappings": false,
						},
					},
					"root_block_device": []interface{}{
						map[string]interface{}{
							"name": "rootblock",
							"size": 20,
						},
					},
				},
			},
			app.EnvironmentInfos,
		},
	}

	for _, tc := range cases {
		output := expandGhostAppEnvironmentInfos(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandGhostAppBuildInfos(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *ghost.BuildInfos
	}{
		{
			[]interface{}{
				map[string]interface{}{
					"ssh_username": "admin",
					"source_ami":   "ami-1",
					"subnet_id":    "subnet-1",
					"ami_name":     "",
				},
			},
			app.BuildInfos,
		},
	}

	for _, tc := range cases {
		output := expandGhostAppBuildInfos(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandGhostAppFeatures(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *[]ghost.Feature
	}{
		// Parameters nil
		{
			[]interface{}{
				map[string]interface{}{
					"name":        "feature",
					"version":     "1.0",
					"provisioner": "ansible",
					"parameters":  nil,
				},
			},
			&[]ghost.Feature{{
				Name:        "feature",
				Version:     "1.0",
				Provisioner: "ansible",
				Parameters:  map[string]interface{}{},
			},
			},
		},
		// Valid parameters json
		{
			[]interface{}{
				map[string]interface{}{
					"name":        "feature",
					"version":     "1",
					"provisioner": "ansible",
					"parameters": `{
            "package_name" : [
              "test",
              "nano"
            ]
          }`,
				},
			},
			&[]ghost.Feature{{
				Name:        "feature",
				Version:     "1",
				Provisioner: "ansible",
				Parameters: map[string]interface{}{
					"package_name": []interface{}{"test", "nano"},
				},
			}},
		},
		// Wrong parameters json
		{
			[]interface{}{
				map[string]interface{}{
					"name":        "feature",
					"version":     "1",
					"provisioner": "ansible",
					"parameters": `{
            "package_name" : [
              "test",
              "nano"
          }`,
				},
			},
			&[]ghost.Feature{{
				Name:        "feature",
				Version:     "1",
				Provisioner: "ansible",
				Parameters:  nil,
			}},
		},
		{
			nil,
			&[]ghost.Feature{},
		},
	}

	for _, tc := range cases {
		output := expandGhostAppFeatures(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandGhostAppLifecycleHooks(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *ghost.LifecycleHooks
	}{
		{
			[]interface{}{
				map[string]interface{}{
					"pre_buildimage":  "#!/usr/bin/env bash",
					"post_buildimage": "#!/usr/bin/env bash",
					"pre_bootstrap":   "",
					"post_bootstrap":  "",
				},
			},
			app.LifecycleHooks,
		},
		{
			nil,
			&ghost.LifecycleHooks{
				PreBuildimage:  "",
				PreBootstrap:   "",
				PostBuildimage: "",
				PostBootstrap:  "",
			},
		},
	}

	for _, tc := range cases {
		output := expandGhostAppLifecycleHooks(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandGhostAppAutoscale(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *ghost.Autoscale
	}{
		{
			[]interface{}{
				map[string]interface{}{
					"name":           "autoscale",
					"enable_metrics": false,
					"min":            0,
					"max":            3,
				},
			},
			app.Autoscale,
		},
		{
			nil,
			&ghost.Autoscale{
				Min:           0,
				Max:           0,
				Name:          "",
				EnableMetrics: false,
			},
		},
	}

	for _, tc := range cases {
		output := expandGhostAppAutoscale(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandGhostAppEnvironmentVariables(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *[]ghost.EnvironmentVariable
	}{
		{
			[]interface{}{
				map[string]interface{}{
					"key":   "env_var_key",
					"value": "env_var_value",
				},
			},
			app.EnvironmentVariables,
		},
		{
			nil,
			&[]ghost.EnvironmentVariable{},
		},
	}

	for _, tc := range cases {
		output := expandGhostAppEnvironmentVariables(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandGhostAppModules(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *[]ghost.Module
	}{
		{
			[]interface{}{
				map[string]interface{}{
					"name":             "my_module",
					"git_repo":         "https://github.com/test/test.git",
					"path":             "/",
					"scope":            "system",
					"build_pack":       "#!/usr/bin/env bash",
					"pre_deploy":       "#!/usr/bin/env bash",
					"post_deploy":      "",
					"after_all_deploy": "",
					"uid":              0,
					"gid":              0,
					"last_deployment":  "",
				},
			},
			app.Modules,
		},
	}

	for _, tc := range cases {
		output := expandGhostAppModules(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandGhostSafeDeployment(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *ghost.SafeDeployment
	}{
		{
			[]interface{}{
				map[string]interface{}{
					"wait_before_deploy": 10,
					"wait_after_deploy":  10,
					"load_balancer_type": "elb",
				},
			},
			app.SafeDeployment,
		},
		{
			[]interface{}{
				map[string]interface{}{
					"wait_before_deploy": 10,
					"wait_after_deploy":  10,
					"load_balancer_type": "elb",
					"api_port":           5001,
					"ha_backend":         "test",
					"app_tag_value":      "test",
				},
			},
			&ghost.SafeDeployment{
				ApiPort:          5001,
				AppTagValue:      "test",
				HaBackend:        "test",
				WaitBeforeDeploy: 10,
				WaitAfterDeploy:  10,
				LoadBalancerType: "elb",
			},
		},
		{
			nil,
			&ghost.SafeDeployment{
				ApiPort:          0,
				AppTagValue:      "",
				HaBackend:        "",
				WaitBeforeDeploy: 10,
				WaitAfterDeploy:  10,
				LoadBalancerType: "elb",
			},
		},
	}

	for _, tc := range cases {
		output := expandGhostAppSafeDeployment(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

// Flatteners Unit Tests
func TestFlattenGhostAppStringList(t *testing.T) {
	cases := []struct {
		Input          []string
		ExpectedOutput []interface{}
	}{
		{
			[]string{
				"1", "2", "3",
			},
			[]interface{}{
				"1", "2", "3",
			},
		},
		{
			nil,
			nil,
		},
	}

	for _, tc := range cases {
		output := flattenGhostAppStringList(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenGhostAppInstanceTags(t *testing.T) {
	cases := []struct {
		Input          *[]ghost.InstanceTag
		ExpectedOutput []interface{}
	}{
		{
			app.EnvironmentInfos.InstanceTags,
			[]interface{}{
				map[string]interface{}{
					"tag_name":  "name",
					"tag_value": "val",
				},
			},
		},
		{
			nil,
			nil,
		},
	}

	for _, tc := range cases {
		output := flattenGhostAppInstanceTags(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenGhostAppOptionalVolume(t *testing.T) {
	cases := []struct {
		Input          *[]ghost.OptionalVolume
		ExpectedOutput []interface{}
	}{
		{
			app.EnvironmentInfos.OptionalVolumes,
			[]interface{}{
				map[string]interface{}{
					"device_name": "my_device",
					"volume_type": "gp2",
					"volume_size": 20,
					"iops":        3000,
					"launch_block_device_mappings": false,
				},
			},
		},
		{
			nil,
			nil,
		},
	}

	for _, tc := range cases {
		output := flattenGhostAppOptionalVolume(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenGhostAppRootBlockDevice(t *testing.T) {
	cases := []struct {
		Input          *ghost.RootBlockDevice
		ExpectedOutput []interface{}
	}{
		{
			app.EnvironmentInfos.RootBlockDevice,
			[]interface{}{
				map[string]interface{}{
					"name": "rootblock",
					"size": 20,
				},
			},
		},
		{
			nil,
			nil,
		},
	}

	for _, tc := range cases {
		output := flattenGhostAppRootBlockDevice(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenGhostAppEnvironmentInfos(t *testing.T) {
	cases := []struct {
		Input          *ghost.EnvironmentInfos
		ExpectedOutput []interface{}
	}{
		{
			app.EnvironmentInfos,
			[]interface{}{
				map[string]interface{}{
					"instance_profile":  "profile",
					"key_name":          "key",
					"public_ip_address": false,
					"security_groups":   []interface{}{"sg-1", "sg-2"},
					"subnet_ids":        []interface{}{"subnet-1", "subnet-2"},
					"instance_tags": []interface{}{
						map[string]interface{}{
							"tag_name":  "name",
							"tag_value": "val",
						},
					},
					"optional_volumes": []interface{}{
						map[string]interface{}{
							"device_name": "my_device",
							"volume_type": "gp2",
							"volume_size": 20,
							"iops":        3000,
							"launch_block_device_mappings": false,
						},
					},
					"root_block_device": []interface{}{
						map[string]interface{}{
							"name": "rootblock",
							"size": 20,
						},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		output := flattenGhostAppEnvironmentInfos(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenGhostAppBuildInfos(t *testing.T) {
	cases := []struct {
		Input          *ghost.BuildInfos
		ExpectedOutput []interface{}
	}{
		{
			app.BuildInfos,
			[]interface{}{
				map[string]interface{}{
					"ssh_username": "admin",
					"source_ami":   "ami-1",
					"subnet_id":    "subnet-1",
					"ami_name":     "",
				},
			},
		},
	}

	for _, tc := range cases {
		output := flattenGhostAppBuildInfos(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenGhostAppFeatures(t *testing.T) {
	cases := []struct {
		Input          *[]ghost.Feature
		ExpectedOutput []interface{}
	}{
		{
			app.Features,
			[]interface{}{
				map[string]interface{}{
					"name":        "feature",
					"version":     "1.0",
					"provisioner": "ansible",
					"parameters":  nil,
				},
			},
		},
		{
			&[]ghost.Feature{{
				Name:        "feature",
				Version:     "1",
				Provisioner: "ansible",
				Parameters:  `{"package_name":["test","nano"]}`,
			}},
			[]interface{}{
				map[string]interface{}{
					"name":        "feature",
					"version":     "1",
					"provisioner": "ansible",
					"parameters":  `"{\"package_name\":[\"test\",\"nano\"]}"`,
				},
			},
		},
		{
			nil,
			nil,
		},
	}

	for _, tc := range cases {
		output := flattenGhostAppFeatures(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenGhostAppLifecycleHooks(t *testing.T) {
	cases := []struct {
		Input          *ghost.LifecycleHooks
		ExpectedOutput []interface{}
	}{
		{
			app.LifecycleHooks,
			[]interface{}{
				map[string]interface{}{
					"pre_buildimage":  "#!/usr/bin/env bash",
					"post_buildimage": "#!/usr/bin/env bash",
					"pre_bootstrap":   "",
					"post_bootstrap":  "",
				},
			},
		},
		{
			nil,
			nil,
		},
	}

	for _, tc := range cases {
		output := flattenGhostAppLifecycleHooks(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenGhostAppAutoscale(t *testing.T) {
	cases := []struct {
		Input          *ghost.Autoscale
		ExpectedOutput []interface{}
	}{
		{
			app.Autoscale,
			[]interface{}{
				map[string]interface{}{
					"name":           "autoscale",
					"enable_metrics": false,
					"min":            0,
					"max":            3,
				},
			},
		},
		{
			nil,
			nil,
		},
	}

	for _, tc := range cases {
		output := flattenGhostAppAutoscale(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenGhostAppEnvironmentVariables(t *testing.T) {
	cases := []struct {
		Input          *[]ghost.EnvironmentVariable
		ExpectedOutput []interface{}
	}{
		{
			app.EnvironmentVariables,
			[]interface{}{
				map[string]interface{}{
					"key":   "env_var_key",
					"value": "env_var_value",
				},
			},
		},
		{
			nil,
			nil,
		},
	}

	for _, tc := range cases {
		output := flattenGhostAppEnvironmentVariables(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenGhostAppModules(t *testing.T) {
	cases := []struct {
		Input          *[]ghost.Module
		ExpectedOutput []interface{}
	}{
		{
			app.Modules,
			[]interface{}{
				map[string]interface{}{
					"name":             "my_module",
					"git_repo":         "https://github.com/test/test.git",
					"path":             "/",
					"scope":            "system",
					"build_pack":       "#!/usr/bin/env bash",
					"pre_deploy":       "#!/usr/bin/env bash",
					"post_deploy":      "",
					"after_all_deploy": "",
					"uid":              0,
					"gid":              0,
					"last_deployment":  "",
				},
			},
		},
	}

	for _, tc := range cases {
		output := flattenGhostAppModules(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenGhostSafeDeployment(t *testing.T) {
	cases := []struct {
		Input          *ghost.SafeDeployment
		ExpectedOutput []interface{}
	}{
		{
			app.SafeDeployment,
			[]interface{}{
				map[string]interface{}{
					"wait_before_deploy": 10,
					"wait_after_deploy":  10,
					"load_balancer_type": "elb",
				},
			},
		},
		{
			&ghost.SafeDeployment{
				ApiPort:          5001,
				AppTagValue:      "test",
				HaBackend:        "test",
				WaitBeforeDeploy: 10,
				WaitAfterDeploy:  10,
				LoadBalancerType: "elb",
			},
			[]interface{}{
				map[string]interface{}{
					"wait_before_deploy": 10,
					"wait_after_deploy":  10,
					"load_balancer_type": "elb",
					"api_port":           5001,
					"ha_backend":         "test",
					"app_tag_value":      "test",
				},
			},
		},
		{
			nil,
			nil,
		},
	}

	for _, tc := range cases {
		output := flattenGhostAppSafeDeployment(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestSuppressDiffFeatures(t *testing.T) {
	suppressFunc := suppressDiffFeaturesParameters()

	cases := []struct {
		ParameterName  string
		OldValue       string
		NewValue       string
		ExpectedOutput bool
		ResourceData   *schema.ResourceData
	}{
		{"features.0.parameters", `{ "name" : "positive", "id" : 1 }`, `{ "name" : "positive", "id" : 1 }`, true, nil},
		{"features.0.parameters", `{ "name" : "positive", "id" : 1 }`, `{ "id" : 1, "name" : "positive" }`, true, nil},
		{"features.0.parameters", `{ "name" : "negative", "id" : 1 }`, `{ "id" : 1, "name" : "positive" }`, false, nil},
		{"features.0.parameters", `{ "name" : "negative", "id" : 1 }`, `{ "name" : "positive", "id" : 1 }`, false, nil},
		{"features.0.parameters", `{}`, "", true, &schema.ResourceData{}},
	}

	for _, tc := range cases {
		output := suppressFunc(tc.ParameterName, tc.OldValue, tc.NewValue, tc.ResourceData)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from SuppressDiffFeatureParameters.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestSuppressDiffEnvironmentInfos(t *testing.T) {
	suppressFunc := suppressDiffEnvironmentInfos()

	resource := resourceGhostApp()
	nonEmptyResourceData := resource.Data(&terraform.InstanceState{
		ID: "ghost_app.test.id",
	})
	flattenGhostApp(nonEmptyResourceData, app)

	cases := []struct {
		ParameterName  string
		OldValue       string
		NewValue       string
		ExpectedOutput bool
		ResourceData   *schema.ResourceData
	}{
		{"environment_infos.0.root_block_device.#", "1", "0", true, &schema.ResourceData{}},
		{"environment_infos.0.root_block_device.#", "1", "0", false, nonEmptyResourceData},
		{"environment_infos.0.root_block_device.0.name", "1", "0", false, &schema.ResourceData{}},
	}

	for _, tc := range cases {
		output := suppressFunc(tc.ParameterName, tc.OldValue, tc.NewValue, tc.ResourceData)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from SuppressDiffEnvironmentInfos.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestSuppressDiffAutoscale(t *testing.T) {
	suppressFunc := suppressDiffAutoscale()

	resource := resourceGhostApp()
	nonEmptyResourceData := resource.Data(&terraform.InstanceState{
		ID: "ghost_app.test.id",
	})
	flattenGhostApp(nonEmptyResourceData, app)

	cases := []struct {
		ParameterName  string
		OldValue       string
		NewValue       string
		ExpectedOutput bool
		ResourceData   *schema.ResourceData
	}{
		{"autoscale.#", "1", "0", true, &schema.ResourceData{}},
		{"autoscale.#", "1", "0", false, nonEmptyResourceData},
		{"autoscale.#", "0", "1", false, &schema.ResourceData{}},
		{"autoscale.0.min", "1", "0", false, &schema.ResourceData{}},
	}

	for _, tc := range cases {
		output := suppressFunc(tc.ParameterName, tc.OldValue, tc.NewValue, tc.ResourceData)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from SuppressDiffAutoscale.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestSuppressDiffLifecycleHooks(t *testing.T) {
	suppressFunc := suppressDiffLifecycleHooks()

	resource := resourceGhostApp()
	nonEmptyResourceData := resource.Data(&terraform.InstanceState{
		ID: "ghost_app.test.id",
	})
	flattenGhostApp(nonEmptyResourceData, app)

	cases := []struct {
		ParameterName  string
		OldValue       string
		NewValue       string
		ExpectedOutput bool
		ResourceData   *schema.ResourceData
	}{
		{"lifecycle_hooks.#", "1", "0", true, &schema.ResourceData{}},
		{"lifecycle_hooks.#", "1", "0", false, nonEmptyResourceData},
		{"lifecycle_hooks.0.pre_buildimage", "1", "0", false, &schema.ResourceData{}},
	}

	for _, tc := range cases {
		output := suppressFunc(tc.ParameterName, tc.OldValue, tc.NewValue, tc.ResourceData)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from SuppressDiffLifecycleHooks.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestSuppressDiffSafeDeployment(t *testing.T) {
	suppressFunc := suppressDiffSafeDeployment()

	resource := resourceGhostApp()
	nonEmptyResourceDataWithDefaults := resource.Data(&terraform.InstanceState{
		ID: "ghost_app.test.id",
	})
	flattenGhostApp(nonEmptyResourceDataWithDefaults, app)

	cases := []struct {
		ParameterName  string
		OldValue       string
		NewValue       string
		ExpectedOutput bool
		ResourceData   *schema.ResourceData
	}{
		{"safe_deployment.#", "1", "0", true, &schema.ResourceData{}},
		{"safe_deployment.#", "1", "0", true, nonEmptyResourceDataWithDefaults},
		{"safe_deployment.0.wait_before_deploy", "1", "0", false, &schema.ResourceData{}},
	}

	for _, tc := range cases {
		output := suppressFunc(tc.ParameterName, tc.OldValue, tc.NewValue, tc.ResourceData)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from SuppressDiffSafeDeployment.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
