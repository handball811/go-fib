# Fib

## 変更履歴
### FibGen9の作成
- ChatGPTによる精製
- 出力内容は`fib_gen8.go`
- 3分程度完了
- 最初から三回の指示で完成
- 出力に問題はない
```
Go1.19で指定されたフィボナッチ数列の値を取得できるような関数を以下のコードの続きとして出力してください。
入力が負の値の場合ErrNegativeというエラーを返すようにしてください
ErrNegativeの定義を削除してください。すでに定義されています。
関数は行列を利用して高速化してください。
出力はコードブロックの形式にしてください
package internal

func fib(n int) (int, error) {
	return 0, nil
}


ErrNegativeの定義を削除してください。すでに定義されています。


出力されている値が一つずれています。
```
### FibGen8の作成
- ChatGPTによる精製
- 出力内容は`fib_gen8.go`
- 3分程度完了
- FibGen7の続きとして入力開始
- 出力が正しくない
```
返り値の型が正しくありません。修正してください
```

### FibGen7の作成
- ChatGPTによる精製
- 出力内容は`fib_gen7.go`
- 3分程度完了
- FibGen6の続きとして入力開始
- たしかにエラーは出なくなったがそういうことではない
```
返り値の型が正しくありません。修正してください
```

### FibGen6の作成
- ChatGPTによる精製
- 出力内容は`fib_gen6.go`
- 10分程度完了
- 一番最初の入力として開始
- 返り値の型が正しくない
```
Go1.19で指定されたフィボナッチ数列の値を取得できるような関数を以下のコードの続きとして出力してください。
入力が負の値の場合ErrNegativeというエラーを返すようにしてください
ErrNegativeの定義を削除してください。すでに定義されています。
関数は行列を利用して高速化してください。
出力はコードブロックの形式にしてください
package internal

func fib(n int) (int, error) {
	return 0, nil
}
```

### FibGen5の作成
- ChatGPTによる精製
- 出力内容は`fib_gen5.go`
- 3分程度完了
- FibGen4の続き
- エラーの解消を要求
- テストの修正は完了した
```
memoizedFib関数は入力が負の値の時にStackOverflowとなります。入力の値が負の値の時はエラーを返すように修正してください
```

### FibGen4の作成
- ChatGPTによる精製
- 出力内容は`fib_gen4.go`
- 3分程度完了
- FibGenTestの続き
- 速度改善を要求
- 関数名は変更
- テストは負の値の時に失敗する
```
fib関数の処理効率が悪いです。改善してください
```

### FibGenTestの作成
- ChatGPTによる精製
- 出力内容は`fib_gen_test.go`
- ３分程度で完了
- FibGen3の生成の続きとして生成する
- ただし関数名は変更しています
```
この関数のテストを出力してください
```

### FibGen3の作成
- ChatGPTによる精製
- 出力内容は`fib_gen3.go`
- エラー定義をなくしたい
- ３分程度で完了
- FibGen2の生成の続きとして生成する
- テストは問題なくパスしています
```
ErrNegativeの定義を削除してください。すでに定義されています。
```

### FibGen2の作成
- ChatGPTによる精製
- 出力内容は`fib_gen2.go`
- 問題点は///で削除したエラーを排除できなった点
- ５分程度で完了
- FibGen1の生成の続きとして生成する
- テストは問題なくパスしています
```
入力に負の値がある場合にErrNegativeという定義済みのグローバル変数をエラーとして出力するように変更してください
```

### FibGen1の作成
- ChatGPTによる精製
- 出力内容は`fib_gen1.go`
- ５分程度完了
以下の文章を入力
```
Go1.19で指定されたフィボナッチ数列の値を取得できるような関数を以下のコードの続きとして出力してください。
出力はコードブロックの形式にしてください
package internal

func fib(n int) (int, error) {
	return 0, nil
}
```
負の値のテスト時に無限ループに突入してしまう

### FibFastの作成
- 実装が少々面倒なFibonatti関数の実装
- 60分程度で完了
- 効率は指数関数的ではやい
- 検証速度をあげるためにテストも追加
### FibNormalの作成
- 実装が楽なFibonatti関数の実装
- 3分程度で完了
- 効率は通常程度

### FibLazyの作成
- 実装が楽なFibonatti関数の実装
- 3分程度で完了
- 効率が相当に悪い

### ひな形の作成

