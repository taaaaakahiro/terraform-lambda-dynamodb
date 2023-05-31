resource "aws_cloudwatch_log_group" "this" {
  name              = "${var.prefix}_example_this"
  retention_in_days = 3
}