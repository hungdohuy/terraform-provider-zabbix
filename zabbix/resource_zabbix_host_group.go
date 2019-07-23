package zabbix

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceZabbixHostGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceZabbixHostGroupCreate,
		Read:   resourceZabbixHostGroupRead,
		Update: resourceZabbixHostGroupUpdate,
		Delete: resourceZabbixHostGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceZabbixHostGroupCreate(d *schema.ResourceData, m interface{}) error {
	return resourceZabbixHostGroupRead(d, m)
}

func resourceZabbixHostGroupRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceZabbixHostGroupUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceZabbixHostGroupRead(d, m)
}

func resourceZabbixHostGroupDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
