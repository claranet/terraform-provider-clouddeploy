package ghost

import (
	"log"

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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Managed by Terraform",
			},
		},
	}
}

func resourceGhostAppCreate(d *schema.ResourceData, meta interface{}) error {
	//client := meta.(*ghost.Client)
	log.Printf("[INFO] Creating Ghost app %s", d.Get("name").(string))

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
