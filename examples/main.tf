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

resource "prodxcloud_s3_bucket" "example" {
  bucket     = "my-prodxcloud-bucket"
  acl        = "private"
  versioning = true

  tags = {
    Environment = "production"
    ManagedBy   = "terraform"
  }
}

resource "prodxcloud_ec2_instance" "example" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"
  subnet_id     = "subnet-12345678"
  key_name      = "my-key-pair"

  tags = {
    Name        = "example-instance"
    Environment = "production"
  }
} 