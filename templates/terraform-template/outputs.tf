output "bucket_id" {
  value       = aws_s3_bucket.this.id
  description = "Unique S3 bucket ID name identifier."
}

output "bucket_arn" {
  value       = aws_s3_bucket.this.arn
  description = "Amazon Resource Name (ARN) of the bucket."
}
