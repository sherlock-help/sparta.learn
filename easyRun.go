// File: main.go
package main

import (
  "encoding/json"
  "fmt"
  "net/http"

  "github.com/Sirupsen/logrus"
  sparta "github.com/Sirupsen/Sparta"
)

func helloWorld(event *json.RawMessage, context *sparta.LambdaContext, w http.ResponseWriter, logger *logrus.Logger) {
  fmt.Fprintf(w, "Hello World!")
}

func main() {
  var lambdaFunctions []*sparta.LambdaAWSInfo

  helloWorldFn := sparta.NewLambda(sparta.IAMRoleDefinition{},
    helloWorld,
    nil)
  lambdaFunctions = append(lambdaFunctions, helloWorldFn)
  sparta.Main("MyHelloWorldStack",
    "Simple Sparta application that demonstrates core functionality",
    lambdaFunctions,
    nil,
    nil)
}