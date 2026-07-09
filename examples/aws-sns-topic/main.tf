module "topics" {
  source = "../../modules/cloud/aws/sns"

  topics = {
    events = {}
  }

  tags = {
    Project = "clusterforge"
  }
}
