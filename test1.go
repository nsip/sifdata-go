package main

import (
    "fmt"
    "github.com/nsip/sifgo-data/infrastructure/environment"
)

func main() {
    fmt.Println("End of line")

    var mye = environment.NewEnvironmentRequest();
    fmt.Println(mye.SolutionId)

}
