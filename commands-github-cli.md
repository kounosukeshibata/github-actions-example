# 基本的なコマンド

## Github Cli のログイン

gh auth login

## ヘルプコマンド

gh --help

## リポジトリの一覧取得

gh repo list

## カレントブランチでプルリクエストを作成

gh pr create --fill-first --web

## リポジトリの作成

gh repo create my-app --public --clone --add-readme

## リモートリポジトリからクローン

gh repo clone tmknom/example-github-cicd

## ワークフロー実行ログの確認

- 実行中
  gh run watch
- 終了後
  gh run view

# ワークフローの実行

## 手動実行

gh workflow run manual.yml -f greeting=goodbye

# variables の登録

gh variables set USERNAME --body 'octocat'
gh secret set PASSWORD --body 'SuperSecret!'

# キャッシュ一覧の取得

gh cache list

# キャッシュを削除

gh cache delete "<キャッシュキー>"
gh cache delete --all

# 自動マージ

gh pr merge <NUMBER | URL | BRANCH> --merge
・status チェックが全完了してからマージする
gh pr merge <NUMBER | URL | BRANCH> --merge --auto
・自動承認する
gh pr review <NUMBER | URL | BRANCH> --approve

# リリースノートの作成

gh release create v0.1.0 --title "v0.1.0" --notes "Wonderful Text"
gh release create v0.1.0 --title "v0.1.0" --notes "Wonderful Text" example.txt
gh release create v0.2.0 --title "v0.2.0" --generate-notes
・特定のバージョンにアセットのアップロード
gh release upload v0.1.0 example.txt

# プルリクエスト

・プルリクの作成
gh pr create --fill-first
gh pr create --fill-first --label "enhancement"
・マージ
gh pr merge --merge
