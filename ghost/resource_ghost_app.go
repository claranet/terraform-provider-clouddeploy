package ghost

import (
	"log"

	"cloud-deploy.io/go-st"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGhostApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceGhostAppCreate,
		Read:   resourceGhostAppRead,
		Update: resourceGhostAppUpdate,
		Delete: resourceGhostAppDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"env": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role": {
				Type:     schema.TypeString,
				Required: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_monitoring": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"autoscale": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enable_metrics": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"min": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"load_balancer": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "elb",
						},
						"wait_before_deploy": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  10,
						},
						"wait_after_deploy": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  10,
						},
					},
				},
			},

			"build_infos": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ssh_username": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "admin",
						},
						"source_ami": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ami_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"environment_infos": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_profile": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"key_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"public_ip_address": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"root_block_device": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"size": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  20,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"security_groups": {
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"instance_tags": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"tag_value": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"subnet_ids": {
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"optional_volumes": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"volume_type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"volume_size": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"iops": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"launch_block_device_mappings": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"environment_variables": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"log_notifications": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"blue_green": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_blue_green": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"color": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"is_online": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"hooks": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"post_swap": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"pre_swap": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"alter_ego_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"features": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"provisioner": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"lifecycle_hooks": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pre_buildimage": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"post_buildimage": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"pre_bootstrap": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"post_bootstrap": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"modules": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"git_repo": {
							Type:     schema.TypeString,
							Required: true,
						},
						"path": {
							Type:     schema.TypeString,
							Required: true,
						},
						"scope": {
							Type:     schema.TypeString,
							Required: true,
						},
						"uid": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"gid": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"build_pack": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"pre_deploy": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"post_deploy": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"after_all_deploy": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"last_deployment": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"safe_deployment": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ha_backend": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"load_balancer_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"app_tag_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"api_port": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"wait_before_deploy": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"wait_after_deploy": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceGhostAppCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ghost.Client)

	name := d.Get("name").(string)
	d.SetId(name)

	log.Printf("[INFO] Creating Ghost app %s", d.Get("name").(string))
	app := expandGhostApp(d)

	eveMetadata, err := client.CreateApp(app)
	if err == nil {
		log.Println("[INFO] App created: " + eveMetadata.ID)
	} else {
		log.Fatalf("[ERROR] error: %v", err)
	}

	return nil
}

func resourceGhostAppRead(d *schema.ResourceData, meta interface{}) error {
	//client := meta.(*ghost.Client)
	log.Printf("[INFO] Reading Ghost app %s", d.Get("name").(string))
	// TODO retrieve Ghost App with d.Id()

	return nil
}

func resourceGhostAppUpdate(d *schema.ResourceData, meta interface{}) error {
	//client := meta.(*ghost.Client)
	log.Printf("[INFO] Updating Ghost app %s", d.Get("name").(string))
	return nil
}

func resourceGhostAppDelete(d *schema.ResourceData, meta interface{}) error {
	//client := meta.(*ghost.Client)
	log.Printf("[INFO] Deleting Ghost app %s", d.Get("name").(string))
	d.SetId("")
	return nil
}

// Get app from TF configuration
func expandGhostApp(d *schema.ResourceData) ghost.App {
	app := ghost.App{
		Name:               d.Get("name").(string),
		Env:                d.Get("env").(string),
		Role:               d.Get("role").(string),
		Region:             d.Get("region").(string),
		InstanceType:       d.Get("instance_type").(string),
		VpcID:              d.Get("vpc_id").(string),
		InstanceMonitoring: d.Get("instance_monitoring").(bool),

		Modules:              expandGhostAppModules(d),
		Features:             expandGhostAppFeatures(d),
		Autoscale:            expandGhostAppAutoscale(d),
		BuildInfos:           expandGhostAppBuildInfos(d),
		EnvironmentInfos:     expandGhostAppEnvironmentInfos(d),
		LifecycleHooks:       expandGhostAppLifecycleHooks(d),
		LogNotifications:     expandGhostAppStringList(d.Get("log_notifications").([]interface{})),
		EnvironmentVariables: expandGhostAppEnvironmentVariables(d),
	}

	return app
}

// Get modules from TF configuration
func expandGhostAppModules(d *schema.ResourceData) *[]ghost.Module {
	configs := d.Get("modules").([]interface{})
	modules := &[]ghost.Module{}

	// Add each module to modules list
	for _, config := range configs {
		data := config.(map[string]interface{})
		module := ghost.Module{
			Name:           data["name"].(string),
			GitRepo:        data["git_repo"].(string),
			Scope:          data["scope"].(string),
			Path:           data["path"].(string),
			BuildPack:      StrToB64(data["build_pack"].(string)),
			PreDeploy:      StrToB64(data["pre_deploy"].(string)),
			PostDeploy:     StrToB64(data["post_deploy"].(string)),
			AfterAllDeploy: StrToB64(data["after_all_deploy"].(string)),
			LastDeployment: data["last_deployment"].(string),
			GID:            data["gid"].(int),
			UID:            data["uid"].(int),
		}

		*modules = append(*modules, module)
	}

	return modules
}

// Get environment variables from TF configuration
func expandGhostAppEnvironmentVariables(d *schema.ResourceData) *[]ghost.EnvironmentVariable {
	configs := d.Get("environment_variables").([]interface{})
	environmentVariables := &[]ghost.EnvironmentVariable{}

	for _, config := range configs {
		data := config.(map[string]interface{})
		environmentVariable := ghost.EnvironmentVariable{
			Key:   data["key"].(string),
			Value: data["value"].(string),
		}

		*environmentVariables = append(*environmentVariables, environmentVariable)
	}

	return environmentVariables
}

// Get autoscale from TF configuration
func expandGhostAppAutoscale(d *schema.ResourceData) *ghost.Autoscale {
	config := d.Get("autoscale").([]interface{})
	if len(config) == 0 {
		return nil
	}
	data := config[0].(map[string]interface{})

	autoscale := &ghost.Autoscale{
		Name:          data["name"].(string),
		EnableMetrics: data["enable_metrics"].(bool),
		Min:           data["min"].(int),
		Max:           data["max"].(int),
	}

	return autoscale
}

// Get lifecycle_hooks from TF configuration
func expandGhostAppLifecycleHooks(d *schema.ResourceData) *ghost.LifecycleHooks {
	config := d.Get("lifecycle_hooks").([]interface{})
	if len(config) == 0 {
		return nil
	}
	data := config[0].(map[string]interface{})

	lifecycleHooks := &ghost.LifecycleHooks{
		PreBuildimage:  StrToB64(data["pre_buildimage"].(string)),
		PostBuildimage: StrToB64(data["post_buildimage"].(string)),
		PreBootstrap:   StrToB64(data["pre_bootstrap"].(string)),
		PostBootstrap:  StrToB64(data["post_boostrap"].(string)),
	}

	return lifecycleHooks
}

// Get features from TF configuration
func expandGhostAppFeatures(d *schema.ResourceData) *[]ghost.Feature {
	configs := d.Get("features").([]interface{})
	features := &[]ghost.Feature{}

	for _, config := range configs {
		data := config.(map[string]interface{})
		feature := ghost.Feature{
			Name:        data["name"].(string),
			Version:     data["version"].(string),
			Provisioner: data["provisioner"].(string),
		}

		*features = append(*features, feature)
	}

	return features
}

// Get build_infos from TF configuration
func expandGhostAppBuildInfos(d *schema.ResourceData) *ghost.BuildInfos {
	config := d.Get("build_infos").([]interface{})
	data := config[0].(map[string]interface{})

	buildInfos := &ghost.BuildInfos{
		SshUsername: data["ssh_username"].(string),
		SourceAmi:   data["source_ami"].(string),
		AmiName:     data["ami_name"].(string),
		SubnetID:    data["subnet_id"].(string),
	}

	return buildInfos
}

// Get environment_infos from TF configuration
func expandGhostAppEnvironmentInfos(d *schema.ResourceData) *ghost.EnvironmentInfos {
	config := d.Get("environment_infos").([]interface{})
	data := config[0].(map[string]interface{})

	environmentInfos := &ghost.EnvironmentInfos{
		InstanceProfile: data["instance_profile"].(string),
		KeyName:         data["key_name"].(string),
		PublicIpAddress: data["public_ip_address"].(bool),
		SecurityGroups:  expandGhostAppStringList(data["security_groups"].([]interface{})),
		SubnetIDs:       expandGhostAppStringList(data["subnet_ids"].([]interface{})),
		InstanceTags:    expandGhostAppInstanceTags(data["instance_tags"].([]interface{})),
		OptionalVolumes: expandGhostAppOptionalVolumes(data["optional_volumes"].([]interface{})),
		RootBlockDevice: expandGhostAppRootBlockDevice(data["root_block_device"].([]interface{})),
	}

	return environmentInfos
}

func expandGhostAppRootBlockDevice(d []interface{}) *ghost.RootBlockDevice {
	data := d[0].(map[string]interface{})
	rootBlockDevice := &ghost.RootBlockDevice{
		Name: data["name"].(string),
		Size: data["size"].(int),
	}

	return rootBlockDevice
}

func expandGhostAppOptionalVolumes(d []interface{}) *[]ghost.OptionalVolume {
	optionalVolumes := &[]ghost.OptionalVolume{}

	for _, config := range d {
		data := config.(map[string]interface{})
		optionalVolume := ghost.OptionalVolume{
			DeviceName: data["device_name"].(string),
			VolumeType: data["volume_type"].(string),
			VolumeSize: data["volume_size"].(int),
			Iops:       data["iops"].(int),
			LaunchBlockDeviceMappings: data["launch_block_device_mappings"].(bool),
		}

		*optionalVolumes = append(*optionalVolumes, optionalVolume)
	}

	return optionalVolumes
}

func expandGhostAppInstanceTags(d []interface{}) *[]ghost.InstanceTag {
	instanceTags := &[]ghost.InstanceTag{}

	for _, config := range d {
		data := config.(map[string]interface{})
		instanceTag := ghost.InstanceTag{
			TagName:  data["tag_name"].(string),
			TagValue: data["tag_value"].(string),
		}

		*instanceTags = append(*instanceTags, instanceTag)
	}

	return instanceTags
}

func expandGhostAppStringList(d []interface{}) []string {
	var stringList []string

	for _, str := range d {
		stringList = append(stringList, str.(string))
	}

	return stringList
}
