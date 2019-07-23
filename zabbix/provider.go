package zabbix

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/hungdohuy/go-zabbix"
)

// Provider declaration for Zabbix
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ZABBIX_USER", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ZABBIX_PASSWORD", nil),
			},
			"server_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ZABBIX_SERVER_URL", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"zabbix_info": resourceZabbixInfo(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	api, err := zabbix.NewSession(d.Get("server_url").(string), d.Get("username").(string), d.Get("password").(string))

	if err != nil {
		panic(err)
	}

	return api, nil
}
