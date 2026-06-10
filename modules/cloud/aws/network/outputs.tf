output "vpc_id" {
  description = "ID of the VPC."
  value       = aws_vpc.this.id
}

output "vpc_cidr_block" {
  description = "CIDR block of the VPC."
  value       = aws_vpc.this.cidr_block
}

output "public_subnet_ids" {
  description = "IDs of the public subnets, ordered by availability_zones."
  value       = [for key in sort(keys(aws_subnet.public)) : aws_subnet.public[key].id]
}

output "private_subnet_ids" {
  description = "IDs of the private subnets, ordered by availability_zones."
  value       = [for key in sort(keys(aws_subnet.private)) : aws_subnet.private[key].id]
}

output "public_route_table_ids" {
  description = "IDs of public route tables."
  value       = [aws_route_table.public.id]
}

output "private_route_table_ids" {
  description = "IDs of private route tables, ordered by availability_zones."
  value       = [for key in sort(keys(aws_route_table.private)) : aws_route_table.private[key].id]
}

output "nat_gateway_ids" {
  description = "IDs of NAT gateways created by this module."
  value       = [for key in sort(keys(aws_nat_gateway.this)) : aws_nat_gateway.this[key].id]
}

output "internet_gateway_id" {
  description = "ID of the internet gateway."
  value       = aws_internet_gateway.this.id
}
