package main

import (
    "fmt"
    // Repo = sifdata or sifdata-go
    // Sub-package = infrastructure/XXX (version - see standard go)/
    "github.com/nsip/sifgo-data/infrastructure/environment"
    // "github.com/nsip/sifgo-data/sifau/v3.4.5/SchoolInfos"
)

func main() {
    fmt.Println("End of line")

    var mye = environment.NewEnvironmentRequest();
    fmt.Println(mye.SolutionId)

}
