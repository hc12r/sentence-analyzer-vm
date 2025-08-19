variable "aws_region" {
  description = "AWS region to deploy resources"
  type        = string
  default     = "eu-north-1"
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
