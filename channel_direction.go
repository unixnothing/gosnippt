package main

import "fmt"

func main() {
    im_ping := make(chan string, 1)
    im_pang := make(chan string, 1)

    ping(im_ping, "Initial meesage")   // arg1 <- arg2
    pong(im_ping, im_pang); // arg1 -> arg2

    fmt.Println(<-im_pang)
}

func ping(im_receive_only chan <- string, msg string) {
    im_receive_only <- msg
}

func pong(im_send_only <-chan string, im_receive_only chan <- string) {
    //im_receive_only <- im_send_only   //can't use chanel send to chanel directly.
    tmp := <-im_send_only
    im_receive_only <- tmp
}


