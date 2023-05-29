resource "aws_lambda_function" "this" {
   depends_on = [
    aws_cloudwatch_log_group.this,
  ]

  filename         = "../build/zip/main.zip"
  # filename         = data.archive_file.this.output_path
  function_name    = "${var.prefix}_this"
  role             = var.lambda_role_for_dynamodb_arn
  handler          = "main"
  # source_code_hash = data.archive_file.this.output_base64sha256
  runtime          = "go1.x"
  timeout          = 29
  environment {
    variables = {
      TABLE_NAME = var.dynamodb_table.name
    }
  }
}

resource "aws_cloudwatch_log_group" "this" {
  name              = "example"
  retention_in_days = 3
}



