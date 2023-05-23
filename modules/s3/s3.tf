resource "aws_s3_bucket" "this" {
  bucket = "log-s3-backet"
  //タグの設定
  tags = {
    Name = "log-s3-backet"
  }
}