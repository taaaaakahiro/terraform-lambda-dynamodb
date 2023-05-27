resource "aws_dynamodb_table" "this" {
  name           = "${var.prefix}_lambda_dynamodb_table"
  billing_mode   = "PROVISIONED"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "id"
  range_key      = "name"

  attribute {
    name = "id"
    type = "S"
  }

   attribute {
    name = "name"
    type = "S"
  }

  attribute {
    name = "TopScore"
    type = "N"
  }

  # ttl {
  #   attribute_name = "TimeToExist"
  #   enabled        = false
  # }

  global_secondary_index {
    name               = "GameTitleIndex"
    hash_key           = "name"
    range_key          = "TopScore"
    write_capacity     = 10
    read_capacity      = 10
    projection_type    = "INCLUDE"
    non_key_attributes = ["id"]
  }

  tags = {
    Name        = "${var.prefix}-dynamodb-table"
    Environment = "production"
  }
}