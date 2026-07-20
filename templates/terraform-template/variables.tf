variable "environment" {
  type        = string
  description = "Target deployment environment (dev, stage, prod)"
}

variable "region_short" {
  type        = string
  description = "Region shortcode (e.g. ue1 for us-east-1)"
}

variable "bucket_name" {
  type        = string
  description = "Unique descriptive name for the S3 bucket"
}

variable "kms_key_arn" {
  type        = string
  description = "KMS Customer Managed Key ARN used for encryption-at-rest"
}

variable "force_destroy" {
  type        = bool
  default     = false
  description = "Allow bucket deletion even if it contains objects"
}

variable "tags" {
  type        = map(string)
  default     = {}
  description = "Map of custom tags to append to the resource"
}
