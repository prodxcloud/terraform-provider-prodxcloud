provider "aws" {
  region = "us-east-1"
}

module "s3_bucket" {
  source = "../../"

  bucket_name      = "example-prodxcloud-bucket"
  environment     = "dev"
  enable_versioning = true
  enable_encryption = true
  force_destroy    = true

  tags = {
    Project     = "Example"
    Owner       = "Platform Team"
    Cost_Center = "12345"
  }

  lifecycle_rules = [
    {
      id      = "archive"
      enabled = true
      prefix  = "archive/"
      
      transition_days         = 30
      transition_storage_class = "STANDARD_IA"
      
      expiration_days = 90
    },
    {
      id      = "logs"
      enabled = true
      prefix  = "logs/"
      
      transition_days         = 60
      transition_storage_class = "GLACIER"
      
      expiration_days = 365
    }
  ]
}

output "bucket_name" {
  value = module.s3_bucket.bucket_id
}

output "bucket_arn" {
  value = module.s3_bucket.bucket_arn
}

output "bucket_versioning_status" {
  value = module.s3_bucket.bucket_versioning_status
}

output "bucket_encryption_status" {
  value = module.s3_bucket.bucket_encryption_status
} 