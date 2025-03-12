package main

import "fmt"

func main() {
    var masukan string
    fmt.Print("Masukkan kata: ")
    fmt.Scan(&masukan)

    isPalindrom := true
    panjang := len(masukan)

    for i := 0; i < panjang/2; i++ {
        if masukan[i] != masukan[panjang-1-i] {
            isPalindrom = false
            break
        }
    }

    if isPalindrom {
        fmt.Println("Kata ini adalah palindrom!")
    } else {
        fmt.Println("Kata ini bukan palindrom.")
    }
}