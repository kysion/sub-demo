package main

import (
  "fmt"
  _ "github.com/kysion/sub-demo/internal/packed"
  "reflect"

  //_ "github.com/kysion/sub-demo/internal/logic/hook"

  _ "github.com/kysion/sub-demo/internal/logic"

  "github.com/gogf/gf/v2/os/gctx"

  "github.com/kysion/sub-demo/internal/cmd"
)

func ccc[T any]() {
  var a T
  aT := reflect.TypeOf(a)
  as := aT.String()
  fmt.Println(as)
}

func main() {
  ccc[func(a1 int, a2 bool) bool]()
  cmd.Main.Run(gctx.GetInitCtx())
}
