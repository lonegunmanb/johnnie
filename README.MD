usage:
```golang
johnnie.Visit(v Walker, node ast.Node)
```
Where Walker is an interface and johnnie.DefaultWalker is a blank default walker you can use as embedded type.

```golang
type MyWalker struct {
    johnnie.DefaultWalker
}
//implement your own walk method here
```