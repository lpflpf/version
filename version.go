package version

import (
    "fmt"
    "os"
)

var (
    BuildVersion string
    BuildTime    string
    BuildName    string
    CommitID     string
    GoVersion    string
)

func Show() {
    fmt.Printf("Commit ID  : %s\n", CommitID)
    fmt.Printf("Build  Name: %s\n", BuildName)
    fmt.Printf("Build  Time: %s\n", BuildTime)
    fmt.Printf("Build  Vers: %s\n", BuildVersion)
    fmt.Printf("Golang Vers: %s\n", GoVersion)
    os.Exit(0)
}
