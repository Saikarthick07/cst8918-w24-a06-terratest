output "resource_group_name" {
  value       = azurerm_resource_group.rg.name
  description = "Name of the Azure resource group"
}

output "vm_name" {
  value       = azurerm_linux_virtual_machine.webserver.name
  description = "Name of the deployed virtual machine"
}

output "nic_name" {
  value       = azurerm_network_interface.webserver.name
  description = "Name of the network interface associated with the VM"
}

output "public_ip" {
  value       = azurerm_public_ip.webserver.ip_address
  description = "Public IP address assigned to the virtual machine"
}
