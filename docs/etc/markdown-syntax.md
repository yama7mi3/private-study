# Markdownについて

Markdown とは、文章を記述するための記法の一つである  
Markdown で記述した文章は、githubやQiitaなどのWebサービスにおいてパーサを介してHTML文書へ変換される  
文章の構造を明示化できることや、簡単に記入できプラットフォームに依存しないこと等のメリットがある  

## Markdown 記法 一覧

### 文字装飾

*もしくは_で対象の文字を囲う  
_の場合は空白スペースが必要となる

#### 斜体

\*斜体*にしたい: *斜体*にしたい  
\_斜体_ にしたい: _斜体_ にしたい

#### 太字

\**太字**にしたい: **太字**にしたい  
\__太字__ にしたい: __太字__ にしたい

#### 斜体 & 太字

\***強調***したい: ***強調***したい  
\___強調___ したい: ___強調___ したい

#### 取り消し線

\~~取り消し~~たい: ~~取り消し~~たい

#### 引用 & 二重引用

\> 吾輩は猫である。  
\> 名前はまだ無い。  
\>> どこで生れたかとんと見当けんとうがつかぬ。
> 吾輩は猫である。  
> 名前は
まだ無い。
>> どこで生れたかとんと見当けんとうがつかぬ。

#### インラインコード

\` で対象の文字を囲う

\`インラインコード\`で表示: `インラインコード`で表示

#### コードブロック

\``` で対象の文字を囲う  
サービスによっては言語の指定が可能: [Rouge](http://rouge.jneen.net/)に準じているサービスが多そう  
基本的には指定した方が良さそう

\```  
public class SampleJava {  
  public static void main(String[] args) {  
    System.out.println("Hello World!");  
  }  
}  
\```

```
public class SampleJava {
  public static void main(String[] args) {
    System.out.println("Hello World!");
  }
}
```

\```java  
public class SampleJava {  
  public static void main(String[] args) {  
    System.out.println("Hello World!");  
  }  
}  
\```

```java
public class SampleJava {
  public static void main(String[] args) {
    System.out.println("Hello World!");
  }
}
```

### 文章構造

#### 見出し

\# を文字の先頭につける  
記号の後ろにはスペースを入れる  
\# の数で見出しレベルが変わる(見出しレベルは6まで対応)

\# 見出し1  
\## 見出し2  
\### 見出し3  
\#### 見出し4  
\##### 見出し5  
\###### 見出し6  

# 見出し1
## 見出し2
### 見出し3
#### 見出し4
##### 見出し5
###### 見出し6

#### 段落

行末に空白を入れる

**空白なし**  
吾輩は猫である。
名前はまだ無い。

**空白あり**  
吾輩は猫である。  
名前はまだ無い。

#### 箇条書き

-もしくは+もしくは*を文字の先頭につける  
記号の後ろにはスペースを入れる  
リスト対象から外れる場合は空行を入れる

\- リスト1  
\- リスト2  
リスト内に書かれる内容

リスト外に書かれる内容

- リスト1
- リスト2  
リスト内に書かれる内容

リスト外に書かれる内容

##### 箇条書きをネストする

ネストしたい文字の先頭に空白行もしくはタブを挿入する

- リスト1
  - リスト1.1  
- リスト2

#### 番号付きリスト

数値と半角ドットを文字の先頭につける  
番号の内容は何でもいいがすべて 1. 内容 で記載すると変更に強い  
記号の後ろにはスペースを入れる

1\. リスト1  
1\. リスト2

1. リスト1
1. リスト2

##### 番号付きリストをネストする

ネストしたい文字の先頭に空白行もしくはタブを挿入する

1. リスト1
    1. リスト1.1
    1. リスト1.2
1. リスト2

#### リンク

\[表示文字](リンクURL)の形式で記載する

\[Googleで検索](https://www.google.com/?hl=ja)  
[Googleで検索](https://www.google.com/?hl=ja)

#### 外部参照リンク

何度も使い回す際などに役立つリンク方法

\[テキスト][参照名]  
\[参照名]:URL

[某検索エンジン][Google]と[Google][Google]は何が違うか、否同じである

[Google]:https://www.google.com/?hl=ja

#### 画像

\!\[alt属性](画像URL)で記載する(alt属性とは代替テキストのこと)  
リンクと同じく外部参照やリンク付き画像なども可能

##### 基本的な使い方

\!\[alt属性](画像URL)  
![ヒツジ](https://user-images.githubusercontent.com/61812388/79941118-85a36a80-849e-11ea-9d42-536e1759f214.png)

タイトルを付けたい場合は\!\[alt属性](画像URL "タイトル")  
![ヒツジ](https://user-images.githubusercontent.com/61812388/79941118-85a36a80-849e-11ea-9d42-536e1759f214.png "もふもふ")

##### 外部参照の場合

\![alt属性][参照]  
\[参照]: 画像URL

![ヒツジ][1]  

[1]: https://user-images.githubusercontent.com/61812388/79941118-85a36a80-849e-11ea-9d42-536e1759f214.png


##### リンク付き画像

\[\!\[alt属性](画像URL)](リンクURL)  
[![ヒツジ](https://user-images.githubusercontent.com/61812388/79941118-85a36a80-849e-11ea-9d42-536e1759f214.png)](https://www.google.com/?hl=ja)


##### 外部リンク付き画像

\[\!\[alt属性](画像URL)][参照]  
\[参照]: 参照URL

[![ヒツジ](https://user-images.githubusercontent.com/61812388/79941118-85a36a80-849e-11ea-9d42-536e1759f214.png)][1]

[1]: https://www.google.com/?hl=ja

##### リサイズ

imgタグにして、width で調整する  
\<img src="画像URL" width=数値>  
<img src="https://user-images.githubusercontent.com/61812388/79941118-85a36a80-849e-11ea-9d42-536e1759f214.png" width=50%>

#### テーブル

コロンの位置で文字の配置を決める
その他の文字装飾にも対応
要素の前後にはスペースを入れること

| Left align | Right align | Center align |
|------------|------------:|:------------:|
| This       |        This |     This     |
| column     |      column |    column    |
| will       |        will |     will     |
| be         |          be |      be      |
| left       |       right |    center    |
| aligned    |     aligned |   aligned    |

#### 水平線

3つ以上のハイフン、アスタリスク、アンダースコアをならべる  
空白を入れても問題はないが、入れない方が好ましい

\---  
\- - -

---
- - -
