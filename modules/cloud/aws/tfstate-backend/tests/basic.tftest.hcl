mock_provider "aws" {
  mock_resource "aws_s3_bucket" {
    defaults = {
      arn = "arn:aws:s3:::clusterforge-dev-tfstate-example"
    }
  }

  mock_resource "aws_dynamodb_table" {
    defaults = {
      arn = "arn:aws:dynamodb:eu-central-1:123456789012:table/clusterforge-dev-tf-locks"
    }
  }
}

run "default_backend_resources_are_planned" {
  command = plan

  variables {
    name                = "clusterforge-tfstate"
    environment         = "dev"
    bucket_name         = "clusterforge-dev-tfstate-example"
    dynamodb_table_name = "clusterforge-dev-tf-locks"
  }

  assert {
    condition     = aws_s3_bucket.state.bucket == "clusterforge-dev-tfstate-example"
    error_message = "Expected S3 state bucket to be planned."
  }

  assert {
    condition     = length(aws_s3_bucket_versioning.state) == 1
    error_message = "Expected versioning to be enabled by default."
  }

  assert {
    condition     = aws_s3_bucket_versioning.state[0].versioning_configuration[0].status == "Enabled"
    error_message = "Expected S3 versioning status to be Enabled."
  }

  assert {
    condition     = length(aws_s3_bucket_server_side_encryption_configuration.state) == 1
    error_message = "Expected encryption configuration to be enabled by default."
  }

  assert {
    condition     = aws_s3_bucket_server_side_encryption_configuration.state[0].rule[0].apply_server_side_encryption_by_default[0].sse_algorithm == "AES256"
    error_message = "Expected default server-side encryption to use AES256."
  }

  assert {
    condition = (
      aws_s3_bucket_public_access_block.state.block_public_acls &&
      aws_s3_bucket_public_access_block.state.block_public_policy &&
      aws_s3_bucket_public_access_block.state.ignore_public_acls &&
      aws_s3_bucket_public_access_block.state.restrict_public_buckets
    )
    error_message = "Expected all S3 public access block settings to be enabled."
  }

  assert {
    condition     = aws_dynamodb_table.locks.name == "clusterforge-dev-tf-locks" && aws_dynamodb_table.locks.hash_key == "LockID"
    error_message = "Expected DynamoDB lock table with LockID hash key."
  }
}

run "versioning_and_encryption_can_be_disabled" {
  command = plan

  variables {
    name                = "clusterforge-tfstate"
    environment         = "dev"
    bucket_name         = "clusterforge-dev-tfstate-example"
    dynamodb_table_name = "clusterforge-dev-tf-locks"
    enable_versioning   = false
    enable_encryption   = false
  }

  assert {
    condition     = length(aws_s3_bucket_versioning.state) == 0
    error_message = "Expected no versioning resource when versioning is disabled."
  }

  assert {
    condition     = length(aws_s3_bucket_server_side_encryption_configuration.state) == 0
    error_message = "Expected no encryption resource when encryption is disabled."
  }
}
