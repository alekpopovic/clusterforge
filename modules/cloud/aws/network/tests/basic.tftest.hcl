mock_provider "aws" {
  mock_resource "aws_vpc" {
    defaults = {
      id = "vpc-1234567890"
    }
  }

  mock_resource "aws_internet_gateway" {
    defaults = {
      id = "igw-1234567890"
    }
  }

  mock_resource "aws_subnet" {
    defaults = {
      id = "subnet-1234567890"
    }
  }

  mock_resource "aws_route_table" {
    defaults = {
      id = "rtb-1234567890"
    }
  }

  mock_resource "aws_eip" {
    defaults = {
      id = "eipalloc-1234567890"
    }
  }

  mock_resource "aws_nat_gateway" {
    defaults = {
      id = "nat-1234567890"
    }
  }
}

run "valid_vpc_with_two_availability_zones_nat_disabled" {
  command = plan

  variables {
    name                 = "clusterforge-dev"
    environment          = "dev"
    cidr_block           = "10.0.0.0/16"
    availability_zones   = ["eu-central-1a", "eu-central-1b"]
    public_subnet_cidrs  = ["10.0.0.0/24", "10.0.1.0/24"]
    private_subnet_cidrs = ["10.0.10.0/24", "10.0.11.0/24"]
    enable_nat_gateway   = false
  }

  assert {
    condition     = aws_vpc.this.cidr_block == "10.0.0.0/16"
    error_message = "Expected VPC CIDR block to match input."
  }

  assert {
    condition     = length(aws_subnet.public) == 2 && length(aws_subnet.private) == 2
    error_message = "Expected two public and two private subnets."
  }

  assert {
    condition     = length(aws_nat_gateway.this) == 0 && length(aws_eip.nat) == 0
    error_message = "Expected no NAT resources when NAT is disabled."
  }

  assert {
    condition     = length(output.public_subnet_ids) == 2 && length(output.private_subnet_ids) == 2
    error_message = "Expected subnet outputs for both availability zones."
  }
}

run "nat_enabled_single_gateway_plan" {
  command = plan

  variables {
    name                 = "clusterforge-dev"
    environment          = "dev"
    cidr_block           = "10.0.0.0/16"
    availability_zones   = ["eu-central-1a", "eu-central-1b"]
    public_subnet_cidrs  = ["10.0.0.0/24", "10.0.1.0/24"]
    private_subnet_cidrs = ["10.0.10.0/24", "10.0.11.0/24"]
    enable_nat_gateway   = true
    single_nat_gateway   = true
  }

  assert {
    condition     = length(aws_nat_gateway.this) == 1 && length(aws_eip.nat) == 1
    error_message = "Expected one NAT gateway and EIP for single NAT mode."
  }

  assert {
    condition     = length(aws_route.private_default) == 2
    error_message = "Expected private default routes for both private route tables."
  }
}

run "nat_enabled_per_az_plan" {
  command = plan

  variables {
    name                 = "clusterforge-dev"
    environment          = "dev"
    cidr_block           = "10.0.0.0/16"
    availability_zones   = ["eu-central-1a", "eu-central-1b"]
    public_subnet_cidrs  = ["10.0.0.0/24", "10.0.1.0/24"]
    private_subnet_cidrs = ["10.0.10.0/24", "10.0.11.0/24"]
    enable_nat_gateway   = true
    single_nat_gateway   = false
  }

  assert {
    condition     = length(aws_nat_gateway.this) == 2 && length(aws_eip.nat) == 2
    error_message = "Expected one NAT gateway and EIP per availability zone."
  }
}

run "invalid_subnet_length_mismatch_fails" {
  command = plan

  variables {
    name                 = "clusterforge-dev"
    environment          = "dev"
    cidr_block           = "10.0.0.0/16"
    availability_zones   = ["eu-central-1a", "eu-central-1b"]
    public_subnet_cidrs  = ["10.0.0.0/24"]
    private_subnet_cidrs = ["10.0.10.0/24", "10.0.11.0/24"]
  }

  expect_failures = [aws_vpc.this]
}
