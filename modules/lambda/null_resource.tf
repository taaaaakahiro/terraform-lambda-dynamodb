resource "null_resource" "go_build" {
  triggers = {
    always_run = "${timestamp()}"
  }

  provisioner "local-exec" {
    command = "cd ../pkg && GOOS=linux GOARCH=amd64 go build -o ../build/bin/main ../cmd/main.go"
  }
}

data "archive_file" "this" {
  depends_on  = [null_resource.go_build]
  type        = "zip"
  source_file = "../build/bin/main"
  output_path = "../build/zip/main.zip"
}
