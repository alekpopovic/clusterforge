locals {
  topic_names = {
    for name, topic in var.topics : name => topic.fifo_topic && !endswith(name, ".fifo") ? "${name}.fifo" : name
  }
}

resource "aws_sns_topic" "this" {
  for_each = var.topics

  name                        = local.topic_names[each.key]
  fifo_topic                  = each.value.fifo_topic
  content_based_deduplication = each.value.content_based_deduplication
  tags                        = var.tags
}

resource "aws_sns_topic_subscription" "this" {
  for_each = var.subscriptions

  topic_arn            = aws_sns_topic.this[each.value.topic].arn
  protocol             = each.value.protocol
  endpoint             = each.value.endpoint
  raw_message_delivery = each.value.raw_message_delivery
}
