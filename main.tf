terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"  # Change this to your desired region
}

module "s3_bucket" {
  source  = "prodxcloud/s3/aws"
  version = "1.0.0"

  bucket_name         = "my-prodxcloud-bucket"
  environment        = "production"
  enable_versioning  = true
  enable_encryption  = true
  
  tags = {
    Environment = "production"
    Project     = "my-project"
    Managed_by  = "terraform"
    Owner       = "platform-team"
  }
}

output "bucket_id" {
  value = module.s3_bucket.bucket_id
}

output "bucket_arn" {
  value = module.s3_bucket.bucket_arn
}

output "bucket_domain_name" {
  value = module.s3_bucket.bucket_domain_name
} 