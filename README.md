# ToyProject Recruiting Community With Golang
## 目標
- Go言語の勉強
- Ginの勉強
- 個人プロジェクトの開始
- Clean Architectureの理解
- Go言語でのTDDの勉強

## 技術スタック
- Gin
  - シェアが一番高いため採用
- Gorm
  - シェアが一番高いため採用
  - `.ent`は後で勉強してみる
- Postgresql
- AWS
  - GCPと悩むがシェアが高い方を選択
- Terraform
  - Infra as a Codeの勉強

## 要件
### v0.1スコープ
- 募集の投稿ができる
- 投稿した募集の編集ができる
- 投稿した募集の詳細が見える
- 投稿した募集の削除ができる
- 新規ユーザー作成ができる
- 全てのユーザーが作成した投稿が見える
### v0.1スコープ以後
- ログインができる
- ログアウトができる
- ユーザー脱退ができる
- ユーザーが作成した投稿のみが見える
- ...どんどん追加していく！
- MSA化

## Architecture
### v0.1スコープ
- Clean Architectureを採用
### v0.1スコープ以後
- MSA導入

## ERD
![erd](resources/erd.png)


