resource "aws_dynamodb_table_item" "this" {
  count = var.item_count

  table_name = aws_dynamodb_table.this.name
  hash_key   = aws_dynamodb_table.this.hash_key
  range_key  = aws_dynamodb_table.this.range_key

  item = <<ITEM
{
  "id": {"S": "${format("%03d", count.index + 1)}"},
  "name": {"S": "テストデータ${count.index + 1}"}
}
ITEM
}