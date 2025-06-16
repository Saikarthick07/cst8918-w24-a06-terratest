variable "label_prefix" {
  type        = string
  default     = "saikarthick"
  description = "Prefix used for naming all Azure resources"
}

variable "region" {
  type        = string
  default     = "westus3"
  description = "Azure region where the resources will be deployed"
}

variable "admin_username" {
  type        = string
  default     = "azureadmin"
  description = "Admin username for the virtual machine"
}
