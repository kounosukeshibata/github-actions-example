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

# コンテナレジストリにコンテナイメージを push する

## 個人アカウント名を環境変数へセット

export GHCR_USER=$(gh config get -h github.com user)

## コンテナイメージをビルド

docker build -t ghcr.io/${GHCR_USER}/example:latest docker/example/
docker build -t ghcr.io/${GHCR_USER}/auto-link:latest --label "org.opencontainers.image.source=https://github.com/${GHCR_USER}/my-repo" docker/example/

## GitHub CLI から GitHub Packages へのアクセス許可

gh auth refresh --scopes write:packages

## GitHub CLI から GitHub Packages へのログイン

gh auth token | docker login ghcr.io -u ${GHCR_USER} --password-stdin

## コンテナレジストリにコンテナイメージを push

docker push ghcr.io/${GHCR_USER}/example:latest
docker push ghcr.io/${GHCR_USER}/auto-link:latest

## コンテナイメージを pull する

docker pull ghcr.io/${GHCR_USER}/example:latest

# ワークフローの実行

gh workflow run 10_6_publish.yml -f version=0.1.0
