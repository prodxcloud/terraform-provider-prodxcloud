output "bucket_id" {
  description = "The name of the bucket"
  value       = aws_s3_bucket.this.id
}

output "bucket_arn" {
  description = "The ARN of the bucket"
  value       = aws_s3_bucket.this.arn
}

output "bucket_domain_name" {
  description = "The bucket domain name"
  value       = aws_s3_bucket.this.bucket_domain_name
}

output "bucket_regional_domain_name" {
  description = "The bucket region-specific domain name"
  value       = aws_s3_bucket.this.bucket_regional_domain_name
}

output "bucket_versioning_status" {
  description = "The versioning status of the bucket"
  value       = try(aws_s3_bucket_versioning.this[0].versioning_configuration[0].status, "Disabled")
}

output "bucket_encryption_status" {
  description = "The server-side encryption configuration"
  value       = try(aws_s3_bucket_server_side_encryption_configuration.this[0].rule[0].apply_server_side_encryption_by_default[0].sse_algorithm, "Disabled")
} 