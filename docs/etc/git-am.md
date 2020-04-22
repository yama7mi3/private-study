# git でのパッチファイルの使い方

## パッチファイルについて

git format-patch でパッチファイルを作成すると、  
そのパッチファイルを受け取った人は、git am を使ってコミット履歴をキープしつつ自分のリポジトリにコミットできる

### 作成

```zsh
$ git format-patch -o {出力ディレクトリ} {コミット番号}
```

パッチファイルはコミット 1 つごとに 1 つずつ作成される

### 使用方法

```zsh
$ git am {出力ディレクトリ}/{出力ファイル}.patch
```

git am を取り消したい場合は

```zsh
$ git am --abort
```

コンフリクトが起きた場合
git statusで対象のファイルを確認し、git add で修正したファイルをステージングにあげる  
amでのコンフリクトは以下のオプションで解消できる
```zsh
$ git am --resolved
```

## git amについて

amとは、apply mailboxのこと  
以下のコマンドにより、現在チェックアウトしているブランチ上に、パッチメッセージがcommitとして記録される

## 参考

<https://git-scm.com/docs/git-am>
<https://git-scm.com/docs/git-format-patch>
<https://git-scm.com/docs/git-apply>
