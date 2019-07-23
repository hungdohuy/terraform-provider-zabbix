# Terraform zabbix provider
Managing Zabbix with Terraform

Current support focus on Zabbix 4.0 LTS

## Build provider
```
go build -o build/terraform-provider-zabbix
```
## Setup local plugins
```
cp build/terraform-provider-zabbix ~/.terraform.d/plugins/
```

## Setup provider zabbix
```
provider "zabbix" {
    server_url = "http://zabbix/api_jsonrpc.php"
    username = "xxxxx"
    password = "xxxxx"
}
```