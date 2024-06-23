- [Context](#context)
  - [GitHub context](#github-context)
  - [Runner context](#runner-context)
- [環境変数](#環境変数)
  - [中間環境変数 (Intermediate environment variables)](#中間環境変数-intermediate-environment-variables)
- [Variables](#variables)
- [Secrets](#secrets)
- [Glob](#glob)
- [ログ](#ログ)
  - [ログのグループ化](#ログのグループ化)
- [ジョブの並列実行](#ジョブの並列実行)
- [マトリックス](#マトリックス)
- [バージョニング](#バージョニング)
  - [Gitタグ](#gitタグ)

## Context
GitHub Actions には実行時の情報などを保持する Context がある。

コンテキストの例
- github
- job
- runner
- inputs

コンテキストから値を取り出したい場合は `${{ github.actor }}` のように記述する。

### GitHub context
- github.run_id: ワークフローの実行 ID
- github.head_ref: プルリクエストのブランチ

### Runner context
- runner.name: ランナーの名前
- runner.os: ランナーの OS

## 環境変数
ワークフロー・ジョブ・ステップのレベルで環境変数を設定することが出来る。

環境変数設定の例
```yml
env:
  BRANCH: main
```

下記2つの方法で環境変数にアクセスできる
```yml
  steps:
    - run: echo ${{ env.BRANCH }}
    - run: echo $BRANCH
```

### 中間環境変数 (Intermediate environment variables)
上記コードでは シェルコマンド へ直接変数を埋め込んだ。
しかし、この実装はアンチパターンである。

理由は、コンテキストによっては特殊文字が含まれていてシェルの実行に意図しない影響を与える可能性があるためである。

この問題を解決するためには環境変数経由で context を渡して参照を行う。

```yml
jobs: 
  print:
    runs-on: ubuntu-latest
    env:
      ACTOR: ${{ github.actor }}
    steps:
      - run: echo "${ACTOR}" # 環境変数経由で context の参照を行う
```

このようにすることで、トークン分割やパス名展開が抑制される。

トークン分割
- コマンドをトークン(単語)に分割する処理
- ls -l /tmp は `ls`、`-l`、`/tmp` の3つのトークンに分割される

パス名展開
- トークン分割後の置換のことを指す
- ワイルドカード文字が含まれている場合はそれらをディレクトリ内のファイル名に展開を行う


## Variables
複数のワークフローで同じ値を使いたい場合は Variables を使うことが出来る。
これは Github のページから設定を行う。

参考: https://docs.github.com/ja/actions/learn-github-actions/variables

設定後、vars contextから値の参照が出来る。
```yml
jobs:
  print:
    runs-on: ubuntu-latest
    env:
      USERNAME: ${{ vars.USERNAME }}
    steps:
      - run: echo "${USERNAME}"
```

## Secrets
機密情報などは Secrets で扱う。
以下の特徴がある。

- Secrets へ登録した情報は暗号化される
- ログ出力時にマスクされる
- 登録後に値の確認ができなくなる

Secrets の登録方法は以下を参照する。
参考: https://docs.github.com/ja/actions/security-guides/using-secrets-in-github-actions

variables と同じように secrets contextから値の参照が出来る。

```yml
jobs:
  print:
    runs-on: ubuntu-latest
    env:
      PASSWORD: ${{ secrets.PASSWORD }}
    steps:
      - run: echo "${PASSWORD}" # ログ出力時にマスクされる
```

## Glob
Glob とは Unix 系環境において、ワイルドカードでファイル名のパターンを指定するための仕組みである。

- `*` : `/` を除く、ゼロ文字以上の文字列にマッチ
- `**`: `/` を含む、ゼロ文字以上の文字列にマッチ
- `?` : 直前に指定した文字に対して、ゼロ文字か1文字にマッチ
- `[]` : 指定した文字範囲のみ含まれる文字列にマッチ

## ログ
Secrets や Variables に `ACTIONS_STE_DEBUG` true を設定することで、ログの詳細を確認することが出来る。

Bash のトレーシングオプションを利用することで、ログの詳細を確認することが出来る。

```yml
jobs:
  log:
    runs-on: ubuntu-latest
    steps:
      - run: |
          set -x
          echo "Hello, World!"
          date
          hostname
```

### ログのグループ化
```yml
jobs:
  log:
    runs-on: ubuntu-latest
    steps:
      - run: |
        echo "::group::Show environment variables"
        printenv
        echo "::endgroup::" 
```

## ジョブの並列実行
以下のような形でジョブを並列実行することが出来る。
```yml
jobs:
  job1:
    runs-on: ubuntu-latest
    steps:
      - run: echo "Job 1"
  job2:
    runs-on: ubuntu-latest
    steps:
      - run: echo "Job 2"
  job3:
    runs-on: ubuntu-latest
    steps:
      - run: echo "Job 3"
```

## マトリックス
1つのジョブ定義で複数のジョブを実行することが出来る。
```yml
jobs:
  build:
    strategy:
      matrix:
        os: [ubuntu-latest, windows-latest, macos-latest]
  runs-on: ${{ matrix.os }}
    steps:
      - run: echo "Hello, World!"
```


## バージョニング
バージョニングには沢山の手法があるが、よく使われるのは セマンティックバージョニング(Semantic versioning) である。
「1.2.3」のように表記し、3つの要素で構成される。

1. メジャーバージョン: 後方互換性がない大きな変更
2. マイナーバージョン: 後方互換性がある機能追加
3. パッチバージョン: 後方互換性があるバグ修正

### Gitタグ
バージョンの付与方法にはいくつかの手法がある。
ソースコードへの記述、利用指定技術への依存など。

しかし実装技術に依存しない方法がある。`git tag` でタグを付与する方法である。

```sh
# タグの作成
git tag v1.0.0

# タグのpush
git push origin v1.0.0
```

参考: https://git-scm.com/book/ja/v2/Git-%E3%81%AE%E5%9F%BA%E6%9C%AC-%E3%82%BF%E3%82%B0
