# ProdXCloud Terraform Provider

The ProdXCloud Terraform Provider is used to manage ProdXCloud resources using Terraform.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0.0
- [Go](https://golang.org/doc/install) >= 1.21

## Using the Provider

```hcl
terraform {
  required_providers {
    prodxcloud = {
      source = "prodxcloud/prodxcloud"
    }
  }
}

provider "prodxcloud" {
  region     = "us-west-2"
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
}
```

## Resources

### S3 Bucket

The `prodxcloud_s3_bucket` resource creates and manages S3 buckets.

```hcl
resource "prodxcloud_s3_bucket" "example" {
  bucket     = "my-bucket"
  acl        = "private"
  versioning = true

  tags = {
    Environment = "production"
  }
}
```

### EC2 Instance

The `prodxcloud_ec2_instance` resource creates and manages EC2 instances.

```hcl
resource "prodxcloud_ec2_instance" "example" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"
  subnet_id     = "subnet-12345678"
  key_name      = "my-key-pair"

  tags = {
    Name = "example-instance"
  }
}
```

## Development

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Build the provider:
   ```bash
   go build -o terraform-provider-prodxcloud
   ```
4. Run tests:
   ```bash
   go test ./...
   ```

## License

Apache 2.0 Licensed. See LICENSE for full details. 