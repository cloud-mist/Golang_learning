package main

import "github.com/headfirstgo/gadget"

// 定义一个接口类型,TapePlayer和 TapeRecorder都满足
type Player interface {
	Play(string)	
	Stop()
}

func playList(device Player, songs []string) {
	 for _,song := range songs {
	 	device.Play(song)
	 }
	 device.Stop()
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}

	// 修改变量类型来保存任何player
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)

	player = gadget.TapeRecorder{}
	playList(player, mixtape)

	
}

/*
Playing 9 to 5
Stopped!
Playing Jessie's Girl
Playing Whip It
Playing 9 to 5
Stopped!
*/ 
