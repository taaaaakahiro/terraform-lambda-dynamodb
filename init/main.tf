module "s3" {
    source = "../modules/s3"
}

module "dynamodb" {
    source = "../modules/dynamodb"
    item_count = 50
}