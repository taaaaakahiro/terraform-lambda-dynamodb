resource "aws_api_gateway_rest_api" "tr_api" {
  name = "${var.prefix}_tr_api"
}

resource "aws_api_gateway_method" "tr_api_get" {
  authorization = "NONE"
  http_method   = "ANY"
  resource_id   = aws_api_gateway_rest_api.tr_api.root_resource_id
  rest_api_id   = aws_api_gateway_rest_api.tr_api.id
  api_key_required = false
}

resource "aws_api_gateway_integration" "tr_api_get" {
  http_method             = aws_api_gateway_method.tr_api_get.http_method
  resource_id             = aws_api_gateway_rest_api.tr_api.root_resource_id
  rest_api_id             = aws_api_gateway_rest_api.tr_api.id
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.this.invoke_arn
}

resource "aws_api_gateway_deployment" "tr_api" {
  depends_on = [
    aws_api_gateway_integration.tr_api_get
  ]
  rest_api_id = aws_api_gateway_rest_api.tr_api.id
  stage_name  = "test"
  triggers = {
    redeployment = filebase64("${path.module}/aws_api_gateway_rest_api.tf")
  }
}

resource "aws_api_gateway_method_settings" "example" {
  rest_api_id = aws_api_gateway_rest_api.tr_api.id
  stage_name  = aws_api_gateway_deployment.tr_api.stage_name
  method_path = "*/*"

  settings {
    data_trace_enabled = true
    logging_level      = "INFO"
  }
}

resource "aws_lambda_permission" "tr_lambda_permit" {               # 追記
  statement_id  = "AllowAPIGatewayGetTrApi"                         # 追記
  action        = "lambda:InvokeFunction"                           # 追記 
  function_name = aws_lambda_function.this.arn                 # 追記
  principal     = "apigateway.amazonaws.com"                        # 追記
  source_arn    = "${aws_api_gateway_rest_api.tr_api.execution_arn}/*/*"           # 追記
}   


resource "aws_api_gateway_rest_api_policy" "admission_restrict_michizane_ip" {
  rest_api_id = aws_api_gateway_rest_api.tr_api.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "*"
      },
      "Action": "execute-api:Invoke",
      "Resource": "${aws_api_gateway_rest_api.tr_api.execution_arn}/*"
    }
  ]
}
EOF
}