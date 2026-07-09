locals {
  queue_names = {
    for name, queue in var.queues : name => queue.fifo_queue && !endswith(name, ".fifo") ? "${name}.fifo" : name
  }
}

resource "aws_sqs_queue" "dlq" {
  for_each = {
    for name, queue in var.queues : name => queue
    if queue.dead_letter_queue
  }

  name                        = "${trimsuffix(local.queue_names[each.key], ".fifo")}-dlq${each.value.fifo_queue ? ".fifo" : ""}"
  fifo_queue                  = each.value.fifo_queue
  content_based_deduplication = each.value.fifo_queue
  message_retention_seconds   = each.value.message_retention_seconds
  tags                        = var.tags
}

resource "aws_sqs_queue" "this" {
  for_each = var.queues

  name                       = local.queue_names[each.key]
  fifo_queue                 = each.value.fifo_queue
  visibility_timeout_seconds = each.value.visibility_timeout_seconds
  message_retention_seconds  = each.value.message_retention_seconds
  delay_seconds              = each.value.delay_seconds
  max_message_size           = each.value.max_message_size
  receive_wait_time_seconds  = each.value.receive_wait_time_seconds

  content_based_deduplication = each.value.fifo_queue

  redrive_policy = each.value.dead_letter_queue ? jsonencode({
    deadLetterTargetArn = aws_sqs_queue.dlq[each.key].arn
    maxReceiveCount     = each.value.max_receive_count
  }) : null

  tags = var.tags
}
