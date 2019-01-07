// +build ignore

// TEMPORARY AUTOGENERATED FILE: easyjson bootstapping code to launch
// the actual generator.

package main

import (
  "fmt"
  "os"

  "github.com/mailru/easyjson/gen"

  pkg "github.com/subiz/header/client"
)

func main() {
  g := gen.NewGenerator("client.pb_easyjson.go")
  g.SetPkg("client", "github.com/subiz/header/client")
  g.Add(pkg.EasyJSON_exporter_AllType(nil))
  g.Add(pkg.EasyJSON_exporter_Authorization(nil))
  g.Add(pkg.EasyJSON_exporter_Client(nil))
  g.Add(pkg.EasyJSON_exporter_Clients(nil))
  if err := g.Run(os.Stdout); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
