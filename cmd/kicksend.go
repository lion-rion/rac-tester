package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

var robotid uint32 = 0
var kickspeedx float32 = 0
var kickspeedz float32 = 0
var veltangent float32 = 0
var velnormal float32 = 0
var velangular float32 = 0
var spinner bool = false
var wheelsspeed bool = false

func KickTest() {
	prompt := promptui.Select{
		Label:     "which robot?",
		Items:     []string{"0", "1", "2"},
		CursorPos: 0,
	}
	id, result, err := prompt.Run() //入力を受け取る
	if err != nil {
		fmt.Println(result)
		fmt.Println(err)
		return
	}
	fmt.Println("You choose ", id)

}

func SendCommnad(robotid int) {
	// ipv4 := "192.168.0." + strconv.Itoa(robotid+100)
	// port := "20011"
	// addr := ipv4 + ":" + port
	// conn, err := net.Dial("udp4", addr)
	// if err != nil {
	// 	panic(err)
	// }
}
