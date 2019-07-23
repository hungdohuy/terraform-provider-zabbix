package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hungdohuy/terraform-provider-zabbix/zabbix"
)

func main() {
	p := plugin.ServeOpts{
		ProviderFunc: zabbix.Provider,
	}

	plugin.Serve(&p)
}
