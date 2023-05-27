resource "aws_iam_role_policy" "tr_lambda_role_policy_policy" {
  name = "${var.prefix}_tr_lambda_policy"
  role = aws_iam_role.this.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "dynamodb:GetItem"
        ]
        Resource = [
          var.dynamodb_table_arn
        ]
      }
    ]
  })
}