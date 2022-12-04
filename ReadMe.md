# Fib

## 変更履歴
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
package internal

func fib(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegative
	} else if n == 0 {
		return 0, nil
	} else if n == 1 {
		return 1, nil
	}
	n1, _ := fib(n - 1)
	n2, _ := fib(n - 2)
	return n1 + n2, nil
}
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

