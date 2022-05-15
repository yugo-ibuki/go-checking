## メモ

空の配列を作る時、以下のように数を指定する事がある。
```
arr := make([]Type, 0, len(t))
```

インスタンス化する方法
1. new()関数で作成(ゼロ値のみ)
```
p1 := new(Person)
```
2. var変数宣言(ゼロ値のみ)
```
var p2 Person
```
3. 複合リテラル(composite literal)で作成(初期値を設定できる)
```
p3 := &Person{
    FirstName: "苗字",
    LastName: "名前",
}
```
