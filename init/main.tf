module "s3" {
  source = "../modules/s3"
}

module "dynamodb" {
  source     = "../modules/dynamodb"
  prefix     = local.prefix
  item_count = local.item_count
}

module "iam" {
  source             = "../modules/iam"
  prefix             = local.prefix
  dynamodb_table_arn = module.dynamodb.dynamodb_table.arn
}