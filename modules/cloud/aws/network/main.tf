locals {
  name        = trimspace(var.name)
  environment = trimspace(var.environment)

  common_tags = merge(var.tags, {
    Name        = local.name
    Environment = local.environment
  })

  subnet_indexes = {
    for index, zone in var.availability_zones : tostring(index) => {
      availability_zone   = zone
      public_subnet_cidr  = var.public_subnet_cidrs[index]
      private_subnet_cidr = var.private_subnet_cidrs[index]
    }
  }

  nat_gateway_indexes = var.enable_nat_gateway ? (
    var.single_nat_gateway ? { "0" = local.subnet_indexes["0"] } : local.subnet_indexes
  ) : {}
}

resource "aws_vpc" "this" {
  cidr_block           = var.cidr_block
  enable_dns_hostnames = var.enable_dns_hostnames
  enable_dns_support   = var.enable_dns_support

  lifecycle {
    precondition {
      condition     = length(var.public_subnet_cidrs) == length(var.availability_zones)
      error_message = "Public subnet CIDR count must equal availability zone count."
    }

    precondition {
      condition     = length(var.private_subnet_cidrs) == length(var.availability_zones)
      error_message = "Private subnet CIDR count must equal availability zone count."
    }
  }

  tags = merge(local.common_tags, {
    Name = local.name
  })
}

resource "aws_internet_gateway" "this" {
  vpc_id = aws_vpc.this.id

  tags = merge(local.common_tags, {
    Name = "${local.name}-igw"
  })
}

resource "aws_subnet" "public" {
  for_each = local.subnet_indexes

  vpc_id                  = aws_vpc.this.id
  availability_zone       = each.value.availability_zone
  cidr_block              = each.value.public_subnet_cidr
  map_public_ip_on_launch = true

  tags = merge(
    local.common_tags,
    {
      Name                     = "${local.name}-public-${each.value.availability_zone}"
      Tier                     = "public"
      "kubernetes.io/role/elb" = "1"
    },
    var.public_subnet_tags
  )
}

resource "aws_subnet" "private" {
  for_each = local.subnet_indexes

  vpc_id            = aws_vpc.this.id
  availability_zone = each.value.availability_zone
  cidr_block        = each.value.private_subnet_cidr

  tags = merge(
    local.common_tags,
    {
      Name                              = "${local.name}-private-${each.value.availability_zone}"
      Tier                              = "private"
      "kubernetes.io/role/internal-elb" = "1"
    },
    var.private_subnet_tags
  )
}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.this.id

  tags = merge(local.common_tags, {
    Name = "${local.name}-public"
    Tier = "public"
  })
}

resource "aws_route" "public_default" {
  route_table_id         = aws_route_table.public.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.this.id
}

resource "aws_route_table_association" "public" {
  for_each = aws_subnet.public

  subnet_id      = each.value.id
  route_table_id = aws_route_table.public.id
}

resource "aws_eip" "nat" {
  for_each = local.nat_gateway_indexes

  domain = "vpc"

  tags = merge(local.common_tags, {
    Name = "${local.name}-nat-${each.value.availability_zone}"
  })

  depends_on = [aws_internet_gateway.this]
}

resource "aws_nat_gateway" "this" {
  for_each = local.nat_gateway_indexes

  allocation_id = aws_eip.nat[each.key].id
  subnet_id     = aws_subnet.public[each.key].id

  tags = merge(local.common_tags, {
    Name = "${local.name}-nat-${each.value.availability_zone}"
  })

  depends_on = [aws_internet_gateway.this]
}

resource "aws_route_table" "private" {
  for_each = local.subnet_indexes

  vpc_id = aws_vpc.this.id

  tags = merge(local.common_tags, {
    Name = "${local.name}-private-${each.value.availability_zone}"
    Tier = "private"
  })
}

resource "aws_route" "private_default" {
  for_each = var.enable_nat_gateway ? local.subnet_indexes : {}

  route_table_id         = aws_route_table.private[each.key].id
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = var.single_nat_gateway ? aws_nat_gateway.this["0"].id : aws_nat_gateway.this[each.key].id
}

resource "aws_route_table_association" "private" {
  for_each = aws_subnet.private

  subnet_id      = each.value.id
  route_table_id = aws_route_table.private[each.key].id
}
