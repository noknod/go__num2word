package main


import  "fmt"

import  "ru/mail/noknod/num2word"


func main() {
    fmt.Println("\n===MainTest")
    var number int

    number = -2467
    fmt.Println(number, ": <" + num2word.NumToWord(number) + ">")

    number = 123456789
    fmt.Println(number, ": <" + num2word.NumToWord(number) + ">")
}