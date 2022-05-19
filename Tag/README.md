# タグについて
タグはreflectで確認するのが一般的

```
t := reflect.TypeOf(Structure)
```
これでtypeを取得できる。
```
field := t.Field(1)
field.Name
field.Tag.Get("tag name")
```
これでタグのNameとタグ名を取得する事ができる。
