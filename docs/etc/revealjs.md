# reveal.jsについて

reveal.jsは、簡単にスライドを作成するためのフレームワークでこんな感じのスライドが作成できる  
<https://revealjs.com/#/>

Githubでも「GitPitch」というスライド作成ツールを使ってみたことがあるが、  
reveal.jsはどうなんだろうと思ったのでreveal.jsを使ってみた記録をまとめていく  
ちなみにGitPitchについてはURLを参照: <https://gitpitch.com/docs>

## 起動方法

基本的なことはgitに書いてくれている  
<https://github.com/hakimel/reveal.js​>

```zsh
// reveal.jsをローカル環境に落としてくる
$ git clone https://github.com/hakimel/reveal.js.git

$ cd reveal.js

// 依存解決を行う
$ npm install

// Appの起動
$ npm start
```

既存コードは、Appを起動したら<http://localhost:8000> でスライド状態を確認できる  
シンプルなスライドが2枚ある状態を確認  
<img src="https://user-images.githubusercontent.com/61812388/79940832-c51d8700-849d-11ea-95e6-80f3e27badb6.png" width="600px">

## index.htmlの中身

スライドの内容をいじるには、reveal.jsディレクトリ配下のindex.htmlを編集する  

- cssディレクトリ配下のスタイルシートを使用しデザイン調整
- PDFとしてデータをエクスポートする際のデザイン調整
- sectionタグによるスライドの作成
- プレゼンテーションの設定適用の処理実行
- Markdown記法に対応

ざっくりいうとこんな感じで行なっているみたい  

## 早速遊んでみる

### デザインを変更してみる

設定しているスタイルシートを以下のように変更

```html
-    <link rel="stylesheet" href="css/theme/black.css">
+    <link rel="stylesheet" href="css/theme/moon.css" />
```

<img src="https://user-images.githubusercontent.com/61812388/79940672-5f30ff80-849d-11ea-8916-5dfa99f71d1e.png" width="600px">

当たり前だがスタイルが変わった  
今回はすでに存在しているファイルに変更してみたのであんまり面白さはないが、自分でcssを準備すれば無限にデザインは広がりそう

### スライドを作成してみる

sectionタグによってスライドの作成できるようなので、スライドを3枚にしてみる

```html
 <div class="reveal">
   <div class="slides">
     <section>Slide 1</section>
     <section>Slide 2</section>
+    <section>Slide 3</section>
   </div>
 </div>
```

<img src="https://user-images.githubusercontent.com/61812388/79940675-61935980-849d-11ea-91f5-e2239b8128b0.png" width="600px">

縦にもページを追加する

```html
 <div class="reveal">
   <div class="slides">
-    <section>Slide 1</section>
+    <section>
+      <section>Slide 1</section>
+      <section>Slide 1.1</section>
+      <section>Slide 1.2</section>
+    </section>
     <section>Slide 2</section>
     <section>Slide 3</section>
   </div>
 </div>
 ```

<img src="https://user-images.githubusercontent.com/61812388/79940681-62c48680-849d-11ea-890f-e6b6eabffa0d.png" width="600px">

クリックだけジャなくvimみたいにh, j, k, lでの操作も出来たし、エスケープキーで構成確認や移動操作も出来た  

<img src="https://user-images.githubusercontent.com/61812388/79940685-63f5b380-849d-11ea-8da1-bd5dcea2222e.png" width="600px">

### Markdown記法で書いてみる

section要素に特定の属性を追加する、かつ、textareaで記載することでMarkdownに対応するらしい  

```html
 <div class="reveal">
   <div class="slides">
     <section>
       <section>Slide 1</section>
       <section>Slide 1.1</section>
       <section>Slide 1.2</section>
     </section>
-    <section>Slide 2</section>
+    <section data-markdown>
+      <textarea data-template>
+        ## Slide 2
+
+        Slide 2 !!
+      </textarea>
+      </section>
    <section>Slide 3</section>
  </div>
</div>
```

<img src="https://user-images.githubusercontent.com/61812388/79940687-6526e080-849d-11ea-8870-ee86e2ef91b6.png" width="600px">

何かフォントが変わっている。markdownのtemplayteは違う設定のスタイルになってるんだな  
別で.mdファイルを別に準備して読み込む方法もあるということなのでその方法もやってみる

```html
 <div class="reveal">
   <div class="slides">
     <section>
       <section>Slide 1</section>
       <section>Slide 1.1</section>
       <section>Slide 1.2</section>
     </section>
-    <section>Slide 2</section>
+    <section data-markdown="sample.md"></section>
    <section>Slide 3</section>
  </div>
</div>
```

同じ階層にさっきの内容を記載したsample.mdを準備

```zsh
$ cat sample.md

## Slide 2

Slide 2 !!
```

data-separatorなどの属性を追加することでページ区切りも可能なので、  
mdファイルのみでスライド情報を完結することもできるみたい

```html
 <div class="reveal">
   <div class="slides">
     <section>
       <section>Slide 1</section>
       <section>Slide 1.1</section>
       <section>Slide 1.2</section>
     </section>
     <section>Slide 2</section>
    <section>Slide 3</section>
  </div>
</div>
```

上記を.mdファイルにまとめた

```html
 <div class="reveal">
   <div class="slides">
    <section data-markdown="sample.md"
             data-separator="^\n---\n"
             data-separator-vertical="^\n>>>\n"></section>
  </div>
</div>
```

```zsh
$ cat sample.md

Slide 1

>>>

Slide 1.1

>>>

Slide 1.2

--

## Slide 2

Slide 2 !!

--

Slide 3
```

今回は3つしか使わなかったが参考までに属性は以下に羅列する

data-markdown: スライドに使用するファイルの指定  
data-separator: スライドのページを区切るための文字列を指定  
data-separator-vertical: ページを縦に区切るための文字列を指定  
data-separator-notes: スピーカーノートとして記載するための文字列を指定  
data-charset: 外部ファイルをロードするときに使用する文字セットの指定  

## まとめ

基本的な書き方は網羅できた  
楽さでいうとGitPitchかなぁとなったが、自動再生したりタイマーを設定したりといった設定が可能らしいので、  
スライド側で出来ることについても確認してみて使いどころを考えてみたいと思った  
