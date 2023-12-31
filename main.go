package main

import (
    "fmt"
    "os"
    "os/user"
    "deadairx/monkey-lang/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Hello %s! This is the Monkey programming language!\n",
        user.Username)
    fmt.Printf("Type some commands here, go on, try it!\n")
    repl.Start(os.Stdin, os.Stdout)
}
