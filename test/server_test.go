package main

import (
	"bytes"
	"log"
	"noa/codec"
	"noa/io"
	"testing"
	"time"
)

type TestChanelHandler struct {
}

func (handler TestChanelHandler) OnMessage(chanel *io.Chanel, data *bytes.Buffer) {
	log.Println(string(data.Bytes()))
}

func (handler TestChanelHandler) OnError(chanel *io.Chanel, err error) {
	log.Println(err.Error())
}

func (handler TestChanelHandler) OnClose() {
	log.Println("connection is closed")
}

func Test(t *testing.T) {

	server := io.NewServer(TestChanelHandler{}, codec.LengthSplitCodec{})
	err := server.Listen(10086)
	if err != nil {
		return
	}
	log.Println("server started!")
	time.Sleep(time.Hour * 2)
}
