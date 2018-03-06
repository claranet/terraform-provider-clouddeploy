package ghost

import (
	"log"

	"bitbucket.org/morea/go-st"
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
			"detailled_monitoring": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"autoscale": {
				Type:     schema.TypeList,
				Required: true,
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
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"size": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"security_groups": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Set:      schema.HashString,
						},
						"instance_tags": {
							Type:     schema.TypeMap,
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
						"parameters": {
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
						"post_boostrap": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"modules": {
				Type:     schema.TypeList,
				Optional: true,
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
						},
						"gid": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"initialized": {
							Type:     schema.TypeBool,
							Optional: true,
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

	log.Printf("[INFO] Testing Ghost client get all apps")
	apps, err := client.GetApps()
	if err == nil {
		log.Printf("All apps retrieved: %s", apps)
	} else {
		log.Printf("error: %v", err)
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
