# ProdXCloud AWS S3 Bucket Module

This Terraform module creates an AWS S3 bucket with standardized configuration and best practices for the ProdXCloud platform.

## Features

- Standardized bucket creation with best practices
- Optional versioning
- Server-side encryption
- Access logging
- Lifecycle rules
- Tags management
- Public access controls

## Usage

```hcl
module "s3_bucket" {
  source  = "prodxcloud/s3/aws"
  version = "1.0.0"

  bucket_name         = "my-unique-bucket"
  environment        = "production"
  enable_versioning  = true
  enable_encryption  = true
  
  tags = {
    Environment = "production"
    Project     = "my-project"
    Managed_by  = "terraform"
  }
}
```

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 1.0.0 |
| aws | >= 5.0.0 |

## Providers

| Name | Version |
|------|---------|
| aws | >= 5.0.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| bucket_name | The name of the S3 bucket | `string` | n/a | yes |
| environment | Environment (e.g., prod, staging, dev) | `string` | n/a | yes |
| enable_versioning | Enable versioning on the bucket | `bool` | `true` | no |
| enable_encryption | Enable server-side encryption | `bool` | `true` | no |
| tags | A map of tags to assign to the bucket | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| bucket_id | The name of the bucket |
| bucket_arn | The ARN of the bucket |
| bucket_domain_name | The bucket domain name |

## Examples

### Basic Usage
```hcl
module "s3_bucket" {
  source = "prodxcloud/s3/aws"

  bucket_name  = "my-app-bucket"
  environment = "production"
}
```

### With All Features Enabled
```hcl
module "s3_bucket" {
  source = "prodxcloud/s3/aws"

  bucket_name         = "my-app-bucket"
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
```

## Development

1. Install dependencies:
   ```bash
   terraform init
   ```

2. Make your changes

3. Run tests:
   ```bash
   terraform test
   ```

## License

Apache 2.0 Licensed. See LICENSE for full details. 