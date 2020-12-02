# 時間割くん(開発中)
時間割を管理できるwebアプリのバックエンド部分。

# 機能一覧
- ログイン・ログアウト・ユーザー登録・削除
- 時間割作成
- 時間割表表示
- 科目登録
- 科目詳細表示
- 科目毎のメモ
- zoomが自動で開く

- 科目毎のtodo
- sns機能・学校科目ごとに掲示板相談

# 使用技術
- Golang(gqlgen, golang-migrate, gorm, gin)
- GraphQL
- Docker
- MySQL


## ローカルのdockerでDBを立ててテーブル作成
rootディレクトリで以下のコマンドを実行
```
docker-compose up -d
migrate -database mysql://user:password@/zikanwarikun -path internal/db/migrations/mysql up
```