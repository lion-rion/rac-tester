package cmd

import (
	"fmt"
	"net"
	"rac-tester/proto/pb_gen"
	pb "rac-tester/proto/pb_gen"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
	"google.golang.org/protobuf/proto"
)

func KickTest() {
	prompt := promptui.Select{
		Label:     "which robot?",
		Items:     []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
		CursorPos: 0,
	}
	id, result, err := prompt.Run() //入力を受け取る
	if err != nil {
		fmt.Println(result)
		fmt.Println(err)
		return
	}

	SendCommnad(uint32(id))
	for {
		fmt.Printf("again? (y/n): ")
		var input string
		fmt.Scanln(&input)
		if input == "y" || input == "Y" {
			SendCommnad(uint32(id))
		} else {
			break
		}
	}

	return //ここで終了
}

func SendCommnad(robotid uint32) {

	var kickspeedx float32 = 1
	var kickspeedz float32 = 0
	var veltangent float32 = 0
	var velnormal float32 = 0
	var velangular float32 = 0
	var spinner bool = false
	var wheelsspeed bool = false

	ipv4 := "192.168.0." + strconv.Itoa(int(robotid)+100)
	port := "20011"
	addr := ipv4 + ":" + port
	//conn, err := net.Dial("udp4", "127.0.0.1:20106")
	conn, err := net.Dial("udp4", addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("countdown")
	fmt.Println(3)
	time.Sleep(1 * time.Second)
	fmt.Println(2)
	time.Sleep(1 * time.Second)
	fmt.Println(1)
	time.Sleep(1 * time.Second)

	for i := 0; i < 100; i++ {
		pe := &pb.GrSim_Robot_Command{
			Id:          &robotid,
			Kickspeedx:  &kickspeedx,
			Kickspeedz:  &kickspeedz,
			Veltangent:  &veltangent,
			Velnormal:   &velnormal,
			Velangular:  &velangular,
			Spinner:     &spinner,
			Wheelsspeed: &wheelsspeed,
		}

		var timestamp float64 = float64(time.Now().UnixNano() / 1e6)
		var isteamyellow bool = false

		command := &pb.GrSim_Commands{
			Timestamp:     &timestamp,
			Isteamyellow:  &isteamyellow,
			RobotCommands: []*pb.GrSim_Robot_Command{pe},
		}
		packet := &pb_gen.GrSim_Packet{
			Commands: command,
		}
		marshalpacket, _ := proto.Marshal(packet)
		//println(marshalpacket)

		_, err = conn.Write([]byte(marshalpacket))
		//:debug println("send : %v", marshalpacket)
		if err != nil {
			panic(err)
		}
		time.Sleep(3 * time.Millisecond)

	}

	Stop(robotid)

	return

}
