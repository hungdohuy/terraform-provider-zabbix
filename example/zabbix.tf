provider "zabbix" {
    server_url = "http://127.0.0.1/api_jsonrpc.php"
    username = "Admin"
    password = "zabbix"
}
resource "zabbix_info" "zabbix_api_version" {
}