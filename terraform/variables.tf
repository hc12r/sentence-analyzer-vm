variable "aws_region" {
  description = "AWS region to deploy resources"
  type        = string
  default     = "eu-north-1"
}

variable "ami_id" {
  description = "AMI ID for the EC2 instance (Ubuntu 20.04 LTS)"
  type        = string
  default     = "ami-0261755bbcb8c4a84" # Ubuntu 20.04 LTS in us-east-1
}

variable "instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t3.micro" # Free tier eligible
}

variable "key_name" {
  description = "Name of the SSH key pair to use for the EC2 instance"
  type        = string
}

variable "private_key_path" {
  description = "Path to the private SSH key file"
  type        = string
}

variable "app_name" {
  description = "Name of the application"
  type        = string
  default     = "sentence-analyzer"
}