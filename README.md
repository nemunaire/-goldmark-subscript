# goldmark-subscript

[GoldMark](https://github.com/yuin/goldmark/) subscript extension.

This implements the [`subscript`](https://pandoc.org/MANUAL.html#extension-superscript-subscript) of pandoc.

```markdown
H~2~O is a liquid.  2^10^ is 1024.
```

```html
<p>H<sub>2</sub>O is a liquid.  2^10^ is 1024.</p>
```

```go
var md = goldmark.New(subscript.Enable)
var source = []byte("H~2~O is a liquid.  2^10^ is 1024.")
err := md.Convert(source, os.Stdout)
if err != nil {
    log.Fatal(err)
}
```
