# terraform-lambda-dynamodb

## golang

## terraform
 - v1.4.6
```sh
$ terraform -version ##バージョン確認
```

## tfenv
```sh
$ tfenv list-remote #インストール可能なバージョン確認
$ tfenv install x.x.x #バージョンをインストール
$ tfenv use x.x.x #バージョンを切り替え
```

### AWS
```sh
$ export AWS_ACCESS_KEY_ID= ＜アクセスキー入力＞
$ export AWS_SECRET_ACCESS_KEY=＜シークレットアクセスキー入力＞
```

## docker-compose
```sh
$ docker-compose config
```

## modules
- [x] S3
- [x] IAM
- [x] DynamoDB