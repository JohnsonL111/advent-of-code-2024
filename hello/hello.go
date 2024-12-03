package main // package is a way to group functions. It is made up of all the files in the same directory

import "fmt" // package for printing to console (part of go std)

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
