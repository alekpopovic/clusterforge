output "bucket_name" {
  description = "S3 bucket name for Velero backups."
  value       = aws_s3_bucket.this.bucket
}

output "bucket_arn" {
  description = "S3 bucket ARN for Velero backups."
  value       = aws_s3_bucket.this.arn
}

output "kms_key_arn" {
  description = "KMS key ARN used for bucket encryption, when configured."
  value       = var.kms_key_arn == "" ? null : var.kms_key_arn
}
