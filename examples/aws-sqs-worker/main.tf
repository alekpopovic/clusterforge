module "queues" {
  source = "../../modules/cloud/aws/sqs"

  queues = {
    jobs = {
      dead_letter_queue = true
      max_receive_count = 5
    }
  }

  tags = {
    Project = "clusterforge"
  }
}
