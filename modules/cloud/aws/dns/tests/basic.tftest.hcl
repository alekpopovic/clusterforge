mock_provider "aws" {
  mock_resource "aws_route53_zone" {
    defaults = {
      zone_id      = "Z1234567890EXAMPLE"
      name_servers = ["ns-1.example.net", "ns-2.example.net"]
    }
  }

  mock_resource "aws_route53_record" {
    defaults = {
      fqdn = "www.example.com"
    }
  }
}

run "created_zone_with_standard_record_plans" {
  command = plan

  variables {
    create_zone = true
    zone_name   = "example.com"
    records = {
      www = {
        name    = "www"
        type    = "A"
        ttl     = 300
        records = ["192.0.2.10"]
      }
    }
  }

  assert {
    condition     = aws_route53_zone.this[0].name == "example.com"
    error_message = "Expected hosted zone to be created for example.com."
  }

  assert {
    condition     = aws_route53_record.this["www"].type == "A"
    error_message = "Expected A record to be planned."
  }

  assert {
    condition     = output.zone_name == "example.com"
    error_message = "Expected zone_name output to match created zone."
  }
}

run "alias_record_without_ttl_plans" {
  command = plan

  variables {
    create_zone = true
    zone_name   = "example.com"
    records = {
      app = {
        name = "app"
        type = "A"
        alias = {
          name    = "dualstack.example-alb-123.eu-central-1.elb.amazonaws.com"
          zone_id = "Z215JYRZR1TBD5"
        }
      }
    }
  }

  assert {
    condition     = aws_route53_record.this["app"].alias[0].evaluate_target_health == true
    error_message = "Expected alias evaluate_target_health default to be true."
  }
}

run "record_must_set_exactly_one_target" {
  command = plan

  variables {
    create_zone = true
    zone_name   = "example.com"
    records = {
      invalid = {
        name    = "invalid"
        type    = "A"
        ttl     = 300
        records = ["192.0.2.10"]
        alias = {
          name    = "dualstack.example-alb-123.eu-central-1.elb.amazonaws.com"
          zone_id = "Z215JYRZR1TBD5"
        }
      }
    }
  }

  expect_failures = [var.records]
}

run "alias_records_must_not_set_ttl" {
  command = plan

  variables {
    create_zone = true
    zone_name   = "example.com"
    records = {
      invalid = {
        name = "invalid"
        type = "A"
        ttl  = 300
        alias = {
          name    = "dualstack.example-alb-123.eu-central-1.elb.amazonaws.com"
          zone_id = "Z215JYRZR1TBD5"
        }
      }
    }
  }

  expect_failures = [var.records]
}

# Existing-zone lookup is intentionally not covered here because the
# data.aws_route53_zone lookup requires real AWS API access. Use create_zone=true
# or pass zone_id in local tests to keep tests credential-free.
