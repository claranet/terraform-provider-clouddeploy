package ghost

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	"cloud-deploy.io/go-st"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceGhostApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceGhostAppCreate,
		Read:   resourceGhostAppRead,
		Update: resourceGhostAppUpdate,
		Delete: resourceGhostAppDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Read:   schema.DefaultTimeout(1 * time.Minute),
			Update: schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: MatchesRegexp(`^[a-zA-Z0-9_.+-]*$`),
			},
			"env": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: MatchesRegexp(`^[a-z0-9\-\_]*$`),
			},
			"role": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: MatchesRegexp(`^[a-z0-9\-\_]*$`),
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: MatchesRegexp(`^vpc-[a-z0-9]*$`),
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
				Type:             schema.TypeList,
				Optional:         true,
				MaxItems:         1,
				DiffSuppressFunc: SuppressDiffEmptyStruct(),
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
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validation.IntAtLeast(0),
						},
						"max": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validation.IntAtLeast(0),
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
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "admin",
							ValidateFunc: MatchesRegexp(`^[a-z\_][a-z0-9\_\-]{0,30}$`),
						},
						"source_ami": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: MatchesRegexp(`^ami-[a-z0-9]*$`),
						},
						"ami_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_id": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: MatchesRegexp(`^subnet-[a-z0-9]*$`),
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
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: MatchesRegexp(`^[a-zA-Z0-9\+\=\,\.\@\-\_]{1,128}$`),
						},
						"key_name": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: MatchesRegexp(`^[a-zA-Z0-9\.\-\_]{1,255}$`),
						},
						"public_ip_address": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"root_block_device": {
							Type:             schema.TypeList,
							Optional:         true,
							MaxItems:         1,
							DiffSuppressFunc: SuppressDiffEmptyStruct(),
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"size": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      20,
										ValidateFunc: validation.IntAtLeast(20),
									},
									"name": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: MatchesRegexp(`^$|^(/[a-z0-9]+/)?[a-z0-9]+$`),
									},
								},
							},
						},
						"security_groups": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: MatchesRegexp(`^sg-[a-z0-9]*$`),
							},
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
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: MatchesRegexp(`^subnet-[a-z0-9]*$`),
							},
							Optional: true,
						},
						"optional_volumes": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_name": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: MatchesRegexp(`^/dev/xvd[b-m]$`),
									},
									"volume_type": {
										Type:     schema.TypeString,
										Required: true,
										ValidateFunc: validation.StringInSlice([]string{
											"gp2", "io1", "standard", "st1", "sc1"}, false),
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
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: MatchesRegexp(`^[a-zA-Z_]+[a-zA-Z0-9_]*$`),
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"log_notifications": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: MatchesRegexp(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`),
				},
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
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"blue", "green"}, false),
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
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: MatchesRegexp(`^[a-zA-Z0-9\.\-\_]*$`),
						},
						"version": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: MatchesRegexp(`^[a-zA-Z0-9\.\-\_\/:~\+=\,]*$`),
						},
						"provisioner": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: MatchesRegexp(`^[a-zA-Z0-9]*$`),
						},
						"parameters": {
							Type:             schema.TypeString,
							Optional:         true,
							ValidateFunc:     validation.ValidateJsonString,
							DiffSuppressFunc: SuppressDiffFeatureParameters(),
						},
					},
				},
			},
			"lifecycle_hooks": {
				Type:             schema.TypeList,
				Optional:         true,
				MaxItems:         1,
				DiffSuppressFunc: SuppressDiffEmptyStruct(),
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
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: MatchesRegexp(`^[a-zA-Z0-9\.\-\_]*$`),
						},
						"git_repo": {
							Type:     schema.TypeString,
							Required: true,
						},
						"path": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: MatchesRegexp(`^(/[a-zA-Z0-9\.\-\_]+)+$`),
						},
						"scope": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"system", "code"}, false),
						},
						"uid": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      0,
							ValidateFunc: validation.IntAtLeast(0),
						},
						"gid": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      0,
							ValidateFunc: validation.IntAtLeast(0),
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
							Computed: true,
						},
					},
				},
			},
			"safe_deployment": {
				Type:             schema.TypeList,
				Optional:         true,
				MaxItems:         1,
				DiffSuppressFunc: SuppressDiffEmptyStruct(),
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ha_backend": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"load_balancer_type": {
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								"elb", "alb", "haproxy"}, false),
							Default: "elb",
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
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validation.IntAtLeast(0),
							Default:      10,
						},
						"wait_after_deploy": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validation.IntAtLeast(0),
							Default:      10,
						},
					},
				},
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGhostAppCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ghost.Client)

	log.Printf("[INFO] Creating Ghost app %s", d.Get("name").(string))
	app := expandGhostApp(d)

	eveMetadata, err := client.CreateApp(app)
	if err != nil {
		return fmt.Errorf("[ERROR] error creating Ghost app: %v", err)
	}

	d.Set("etag", *eveMetadata.Etag)
	d.SetId(eveMetadata.ID)

	return resourceGhostAppRead(d, meta)
}

func resourceGhostAppRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ghost.Client)

	log.Printf("[INFO] Reading Ghost app %s", d.Get("name").(string))

	app, err := client.GetApp(d.Id())
	if err != nil {
		// If app was not found, return nil to show that app is gone
		if err.Error()[len(err.Error())-3:] == "404" {
			d.SetId("")
			log.Printf("[WARN] Ghost app (%s) not found, removing from state", d.Id())
			return nil
		}
		return fmt.Errorf("[ERROR] error reading Ghost app: %v", err)
	}

	if err := flattenGhostApp(d, app); err != nil {
		return fmt.Errorf("[ERROR] error reading Ghost app: %v", err)
	}

	return nil
}

func resourceGhostAppUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ghost.Client)

	log.Printf("[INFO] Updating Ghost app %s", d.Get("name").(string))

	app_updated := expandGhostApp(d)

	eveMetadata, err := client.UpdateApp(&app_updated, d.Id(), d.Get("etag").(string))
	if err != nil {
		ec := err.Error()[len(err.Error())-3:]
		if ec == "412" {
			return fmt.Errorf(`[ERROR] error updating Ghost app: app has been updated since
				last plan, you should run plan again: %v`, err)
		}
		return fmt.Errorf("[ERROR] error updating Ghost app: %v", err)
	}

	d.Set("etag", *eveMetadata.Etag)

	return resourceGhostAppRead(d, meta)
}

func resourceGhostAppDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ghost.Client)

	log.Printf("[INFO] Deleting Ghost app %s", d.Get("name").(string))

	err := client.DeleteApp(d.Id(), d.Get("etag").(string))
	if err != nil {
		ec := err.Error()[len(err.Error())-3:]
		if ec == "412" {
			return fmt.Errorf(`[ERROR] error deleting Ghost app: app has been updated since
					last destroy plan, you should run destroy plan again: %v`, err)
		}
		return fmt.Errorf("[ERROR] error deleting Ghost app: %v", err)
	}

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

		Modules:              expandGhostAppModules(d.Get("modules").([]interface{})),
		Features:             expandGhostAppFeatures(d.Get("features").([]interface{})),
		Autoscale:            expandGhostAppAutoscale(d.Get("autoscale").([]interface{})),
		BuildInfos:           expandGhostAppBuildInfos(d.Get("build_infos").([]interface{})),
		EnvironmentInfos:     expandGhostAppEnvironmentInfos(d.Get("environment_infos").([]interface{})),
		LifecycleHooks:       expandGhostAppLifecycleHooks(d.Get("lifecycle_hooks").([]interface{})),
		LogNotifications:     expandGhostAppStringList(d.Get("log_notifications").([]interface{})),
		EnvironmentVariables: expandGhostAppEnvironmentVariables(d.Get("environment_variables").([]interface{})),
		SafeDeployment:       expandGhostAppSafeDeployment(d.Get("safe_deployment").([]interface{})),
	}

	return app
}

func flattenGhostApp(d *schema.ResourceData, app ghost.App) error {
	d.Set("name", app.Name)
	d.Set("env", app.Env)
	d.Set("role", app.Role)
	d.Set("region", app.Region)
	d.Set("instance_type", app.InstanceType)
	d.Set("vpc_id", app.VpcID)
	d.Set("instance_monitoring", app.InstanceMonitoring)
	d.Set("etag", app.Etag)

	d.Set("modules", flattenGhostAppModules(app.Modules))
	d.Set("build_infos", flattenGhostAppBuildInfos(app.BuildInfos))
	d.Set("environment_infos", flattenGhostAppEnvironmentInfos(app.EnvironmentInfos))
	d.Set("features", flattenGhostAppFeatures(app.Features))
	d.Set("autoscale", flattenGhostAppAutoscale(app.Autoscale))
	d.Set("lifecycle_hooks", flattenGhostAppLifecycleHooks(app.LifecycleHooks))
	d.Set("log_notifications", flattenGhostAppStringList(app.LogNotifications))
	d.Set("environment_variables", flattenGhostAppEnvironmentVariables(app.EnvironmentVariables))
	d.Set("safe_deployment", flattenGhostAppSafeDeployment(app.SafeDeployment))

	return nil
}

// Get modules from TF configuration
func expandGhostAppModules(d []interface{}) *[]ghost.Module {
	modules := &[]ghost.Module{}

	// Add each module to modules list
	for _, config := range d {
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
			GID:            data["gid"].(int),
			UID:            data["uid"].(int),
		}

		*modules = append(*modules, module)
	}

	return modules
}

func flattenGhostAppModules(modules *[]ghost.Module) []interface{} {
	moduleList := []interface{}{}

	for _, module := range *modules {
		values := map[string]interface{}{
			"name":             module.Name,
			"git_repo":         module.GitRepo,
			"path":             module.Path,
			"scope":            module.Scope,
			"uid":              module.UID,
			"gid":              module.GID,
			"build_pack":       B64ToStr(module.BuildPack),
			"pre_deploy":       B64ToStr(module.PreDeploy),
			"post_deploy":      B64ToStr(module.PostDeploy),
			"after_all_deploy": B64ToStr(module.AfterAllDeploy),
			"last_deployment":  module.LastDeployment,
		}

		moduleList = append(moduleList, values)
	}

	return moduleList
}

// Get environment variables from TF configuration
func expandGhostAppEnvironmentVariables(d []interface{}) *[]ghost.EnvironmentVariable {
	environmentVariables := &[]ghost.EnvironmentVariable{}

	for _, config := range d {
		data := config.(map[string]interface{})
		environmentVariable := ghost.EnvironmentVariable{
			Key:   data["key"].(string),
			Value: data["value"].(string),
		}

		*environmentVariables = append(*environmentVariables, environmentVariable)
	}

	return environmentVariables
}

func flattenGhostAppEnvironmentVariables(environmentVariables *[]ghost.EnvironmentVariable) []interface{} {
	environmentVariableList := []interface{}{}

	if environmentVariables == nil {
		return nil
	}

	for _, environmentVariable := range *environmentVariables {
		values := map[string]interface{}{
			"key":   environmentVariable.Key,
			"value": environmentVariable.Value,
		}

		environmentVariableList = append(environmentVariableList, values)
	}

	return environmentVariableList
}

// Get autoscale from TF configuration
func expandGhostAppAutoscale(d []interface{}) *ghost.Autoscale {
	if len(d) == 0 {
		return nil
	}
	data := d[0].(map[string]interface{})

	autoscale := &ghost.Autoscale{
		Name:          data["name"].(string),
		EnableMetrics: data["enable_metrics"].(bool),
		Min:           data["min"].(int),
		Max:           data["max"].(int),
	}

	return autoscale
}

func flattenGhostAppAutoscale(autoscale *ghost.Autoscale) []interface{} {
	values := []interface{}{}

	if autoscale == nil {
		return nil
	}

	values = append(values, map[string]interface{}{
		"name":           autoscale.Name,
		"enable_metrics": autoscale.EnableMetrics,
		"min":            autoscale.Min,
		"max":            autoscale.Max,
	})

	return values
}

// Get lifecycle_hooks from TF configuration
func expandGhostAppLifecycleHooks(d []interface{}) *ghost.LifecycleHooks {
	if len(d) == 0 {
		return nil
	}
	data := d[0].(map[string]interface{})

	lifecycleHooks := &ghost.LifecycleHooks{
		PreBuildimage:  StrToB64(data["pre_buildimage"].(string)),
		PostBuildimage: StrToB64(data["post_buildimage"].(string)),
		PreBootstrap:   StrToB64(data["pre_bootstrap"].(string)),
		PostBootstrap:  StrToB64(data["post_bootstrap"].(string)),
	}

	return lifecycleHooks
}

func flattenGhostAppLifecycleHooks(lifecycleHooks *ghost.LifecycleHooks) []interface{} {
	values := []interface{}{}

	if lifecycleHooks == nil {
		return nil
	}

	values = append(values, map[string]interface{}{
		"pre_buildimage":  B64ToStr(lifecycleHooks.PreBuildimage),
		"post_buildimage": B64ToStr(lifecycleHooks.PostBuildimage),
		"pre_bootstrap":   B64ToStr(lifecycleHooks.PreBootstrap),
		"post_bootstrap":  B64ToStr(lifecycleHooks.PostBootstrap),
	})

	return values
}

// Get features from TF configuration
func expandGhostAppFeatures(d []interface{}) *[]ghost.Feature {
	features := &[]ghost.Feature{}

	for _, config := range d {
		data := config.(map[string]interface{})

		feature := ghost.Feature{
			Name:        data["name"].(string),
			Version:     data["version"].(string),
			Provisioner: data["provisioner"].(string),
		}

		if param := data["parameters"]; param != nil {
			var jsonDoc interface{}
			if err := json.Unmarshal([]byte(param.(string)), &jsonDoc); err != nil {
				log.Printf("Error loading feature paramaters json: %v", err)
			}
			feature.Parameters = jsonDoc
		}

		*features = append(*features, feature)
	}

	return features
}

func flattenGhostAppFeatures(features *[]ghost.Feature) []interface{} {
	featureList := []interface{}{}

	if features == nil {
		return nil
	}

	for _, feature := range *features {
		values := map[string]interface{}{
			"name":        feature.Name,
			"version":     feature.Version,
			"provisioner": feature.Provisioner,
			"parameters":  feature.Parameters,
		}

		if feature.Parameters != nil {
			params_json, err := json.Marshal(feature.Parameters)
			if err == nil {
				values["parameters"] = string(params_json)
			}
		}

		featureList = append(featureList, values)
	}

	return featureList
}

func SuppressDiffFeatureParameters() schema.SchemaDiffSuppressFunc {
	return func(k, old, new string, d *schema.ResourceData) bool {
		var oldJson, newJson interface{}

		if err := json.Unmarshal([]byte(old), &oldJson); err != nil {
			log.Printf("Error loading feature parameters json: %v", err)
		}

		if err := json.Unmarshal([]byte(new), &newJson); err != nil {
			log.Printf("Error loading feature parameters json: %v", err)
		}

		// If the new parameters structure is equivalent to the old one,
		// ignores the diff during plan
		return reflect.DeepEqual(oldJson, newJson)
	}
}

// Get build_infos from TF configuration
func expandGhostAppBuildInfos(d []interface{}) *ghost.BuildInfos {
	data := d[0].(map[string]interface{})

	buildInfos := &ghost.BuildInfos{
		SshUsername: data["ssh_username"].(string),
		SourceAmi:   data["source_ami"].(string),
		SubnetID:    data["subnet_id"].(string),
	}

	return buildInfos
}

func flattenGhostAppBuildInfos(buildInfos *ghost.BuildInfos) []interface{} {
	values := []interface{}{}

	if buildInfos == nil {
		return nil
	}

	values = append(values, map[string]interface{}{
		"ssh_username": buildInfos.SshUsername,
		"source_ami":   buildInfos.SourceAmi,
		"ami_name":     buildInfos.AmiName,
		"subnet_id":    buildInfos.SubnetID,
	})

	return values
}

// Get environment_infos from TF configuration
func expandGhostAppEnvironmentInfos(d []interface{}) *ghost.EnvironmentInfos {
	data := d[0].(map[string]interface{})

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

func flattenGhostAppEnvironmentInfos(environmentInfos *ghost.EnvironmentInfos) []interface{} {
	values := []interface{}{}

	if environmentInfos == nil {
		return nil
	}

	values = append(values, map[string]interface{}{
		"instance_profile":  environmentInfos.InstanceProfile,
		"key_name":          environmentInfos.KeyName,
		"public_ip_address": environmentInfos.PublicIpAddress,
		"security_groups":   flattenGhostAppStringList(environmentInfos.SecurityGroups),
		"subnet_ids":        flattenGhostAppStringList(environmentInfos.SubnetIDs),
		"instance_tags":     flattenGhostAppInstanceTags(environmentInfos.InstanceTags),
		"optional_volumes":  flattenGhostAppOptionalVolume(environmentInfos.OptionalVolumes),
		"root_block_device": flattenGhostAppRootBlockDevice(environmentInfos.RootBlockDevice),
	})

	return values
}

func expandGhostAppRootBlockDevice(d []interface{}) *ghost.RootBlockDevice {
	if len(d) == 0 {
		return nil
	}

	data := d[0].(map[string]interface{})

	rootBlockDevice := &ghost.RootBlockDevice{
		Name: data["name"].(string),
		Size: data["size"].(int),
	}

	return rootBlockDevice
}

func flattenGhostAppRootBlockDevice(rootBlockDevice *ghost.RootBlockDevice) []interface{} {
	values := []interface{}{}

	if rootBlockDevice == nil {
		return nil
	}

	values = append(values, map[string]interface{}{
		"name": rootBlockDevice.Name,
		"size": rootBlockDevice.Size,
	})

	return values
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

func flattenGhostAppOptionalVolume(optionalVolumes *[]ghost.OptionalVolume) []interface{} {
	OptionalVolumeList := []interface{}{}

	if optionalVolumes == nil {
		return nil
	}

	for _, OptionalVolume := range *optionalVolumes {
		values := map[string]interface{}{
			"device_name": OptionalVolume.DeviceName,
			"volume_type": OptionalVolume.VolumeType,
			"volume_size": OptionalVolume.VolumeSize,
			"iops":        OptionalVolume.Iops,
			"launch_block_device_mappings": OptionalVolume.LaunchBlockDeviceMappings,
		}

		OptionalVolumeList = append(OptionalVolumeList, values)
	}

	return OptionalVolumeList
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

func flattenGhostAppInstanceTags(instanceTags *[]ghost.InstanceTag) []interface{} {
	InstanceTagList := []interface{}{}

	if instanceTags == nil {
		return nil
	}

	for _, instanceTag := range *instanceTags {
		values := map[string]interface{}{
			"tag_name":  instanceTag.TagName,
			"tag_value": instanceTag.TagValue,
		}

		InstanceTagList = append(InstanceTagList, values)
	}

	return InstanceTagList
}

func expandGhostAppStringList(d []interface{}) []string {
	var stringList []string

	for _, str := range d {
		stringList = append(stringList, str.(string))
	}

	return stringList
}

func flattenGhostAppStringList(strings []string) []interface{} {
	stringList := []interface{}{}

	if strings == nil {
		return nil
	}

	for _, str := range strings {
		stringList = append(stringList, str)
	}

	return stringList
}

func expandGhostAppSafeDeployment(d []interface{}) *ghost.SafeDeployment {
	if len(d) == 0 {
		return nil
	}

	data := d[0].(map[string]interface{})

	safeDeployment := &ghost.SafeDeployment{
		WaitBeforeDeploy: data["wait_before_deploy"].(int),
		WaitAfterDeploy:  data["wait_after_deploy"].(int),
		LoadBalancerType: data["load_balancer_type"].(string),
	}

	if data["app_tag_value"] != nil {
		safeDeployment.AppTagValue = data["app_tag_value"].(string)
	}
	if data["ha_backend"] != nil {
		safeDeployment.HaBackend = data["ha_backend"].(string)
	}
	if data["api_port"] != nil {
		safeDeployment.ApiPort = data["api_port"].(int)
	}

	return safeDeployment
}

func flattenGhostAppSafeDeployment(safeDeployment *ghost.SafeDeployment) []interface{} {
	values := []interface{}{}

	if safeDeployment == nil {
		return nil
	}

	value := map[string]interface{}{
		"wait_before_deploy": safeDeployment.WaitBeforeDeploy,
		"wait_after_deploy":  safeDeployment.WaitAfterDeploy,
		"load_balancer_type": safeDeployment.LoadBalancerType,
	}

	if safeDeployment.AppTagValue != "" {
		value["app_tag_value"] = safeDeployment.AppTagValue
	}

	if safeDeployment.HaBackend != "" {
		value["ha_backend"] = safeDeployment.HaBackend
	}

	if safeDeployment.ApiPort != 0 {
		value["api_port"] = safeDeployment.ApiPort
	}

	values = append(values, value)

	return values
}

// Check that the struct is empty meaning there's no change
func hasNoChange(k string, d *schema.ResourceData) bool {
	if d == nil {
		return true
	}

	switch k {
	case "autoscale.#":
		val, ok := d.GetOk("autoscale")
		if !ok {
			return true
		}
		autoscale := expandGhostAppAutoscale(val.([]interface{}))
		return autoscale == nil || (autoscale.EnableMetrics &&
			autoscale.Max == 0 && autoscale.Min == 0 && autoscale.Name == "")
	case "lifecycle_hooks.#":
		lifecycle_hooks := expandGhostAppLifecycleHooks(d.Get("lifecycle_hooks").([]interface{}))
		return lifecycle_hooks == nil || (lifecycle_hooks.PostBootstrap == "" &&
			lifecycle_hooks.PostBuildimage == "" && lifecycle_hooks.PreBootstrap == "" &&
			lifecycle_hooks.PreBuildimage == "")
	case "environment_infos.0.root_block_device.#":
		environment_infos := expandGhostAppEnvironmentInfos((d.Get("environment_infos").([]interface{})))
		return environment_infos.RootBlockDevice == nil ||
			(environment_infos.RootBlockDevice.Name == "" &&
				environment_infos.RootBlockDevice.Size == 0)
	default:
		return false
	}
}

// Remove plan diffs due to empty struct created by ghost
func SuppressDiffEmptyStruct() schema.SchemaDiffSuppressFunc {
	return func(k, old, new string, d *schema.ResourceData) bool {
		list := []string{"autoscale.#", "lifecycle_hooks.#", "safe_deployment.#",
			"environment_infos.0.root_block_device.#"}

		return IsInList(k, list) && hasNoChange(k, d) && old == "1" && new == "0"
	}
}
