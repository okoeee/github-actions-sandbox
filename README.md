
## GitHub CLI のセットアップ

```sh
# インストール
brew install gh

# ログイン
gh auth login

# 設定例
? What account do you want to log into? GitHub.com
? What is your preferred protocol for Git operations on this host? HTTPS
? Authenticate Git with your GitHub credentials? Yes
? How would you like to authenticate GitHub CLI? Login with a web browser
```

## Go のセットアップ

```sh
# goenv のインストール
brew install goenv

# path の追加
vim ~/.zshrc

# 以下の設定を .zshrc に追加
export GOENV_ROOT=$HOME/.goenv
export PATH=$GOENV_ROOT/bin:$PATH
eval "$(goenv init -)"

# .zshrc の再読み込み
source ~/.zshrc

# go のインストール
goenv install 1.22.4

# go のバージョンを設定
goenv global 1.22.4

# go のバージョンを確認
goenv versions
```
