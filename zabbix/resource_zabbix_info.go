package zabbix

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hungdohuy/go-zabbix"
)

func resourceZabbixInfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceZabbixInfoCreate,
		Read:   resourceZabbixInfoRead,
		Update: resourceZabbixInfoUpdate,
		Delete: resourceZabbixInfoDelete,

		Schema: map[string]*schema.Schema{
			"zbx_api_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceZabbixInfoCreate(d *schema.ResourceData, m interface{}) error {
	return resourceZabbixInfoRead(d, m)
}

func resourceZabbixInfoRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*zabbix.Session)
	log.Printf("API version: %s", api.GetVersion)
	// d.Set("zbx_api_version", api.GetVersion)
	return nil
}

func resourceZabbixInfoUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceZabbixInfoDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
