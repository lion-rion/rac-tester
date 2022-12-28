package cmd

import (
	"errors"
	"fmt"
	"net"
	"rac-tester/proto/pb_gen"
	pb "rac-tester/proto/pb_gen"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
	"google.golang.org/protobuf/proto"
)

func KickTest(isSim bool) {

	var id int

	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("RobotID must be 0-11")
		}
		id, _ = strconv.Atoi(input)
		if id < 12 && id >= 0 {
		} else {
			return errors.New("RobotID must be 0-11")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . | bold }} ",
		Valid:   "{{ . | bold | green }} ",
		Invalid: "{{ . | bold |red }} ",
		Success: "{{ . | bold | green }} ",
	}

	prompt := promptui.Prompt{
		Label:     "RobotID",
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Println(result)
		fmt.Println(err)
		return
	}
	if isSim {
		SendCommnad(uint32(id), isSim)
	} else {
		SendCommnad(uint32(id), isSim)
	}

	for {
		prompt := promptui.Select{
			Label:     "Again?",
			Items:     []string{"Yes", "No"},
			CursorPos: 0,
		}
		idx, _, err := prompt.Run() //入力を受け取る

		if idx == 0 {
			SendCommnad(uint32(id), isSim)
		} else {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	return //ここで終了
}

func SendCommnad(robotid uint32, isSim bool) {

	var kickspeedx float32 = 1
	var kickspeedz float32 = 0
	var veltangent float32 = 0
	var velnormal float32 = 0
	var velangular float32 = 0
	var spinner bool = false
	var wheelsspeed bool = false

	var ipv4 string
	var port string
	var addr string
	if isSim {
		ipv4 = "127.0.0.1"
		port = "20106"
		addr = ipv4 + ":" + port
	} else {
		ipv4 = "192.168.0." + strconv.Itoa(int(robotid)+100)
		port = "20011"
		addr = ipv4 + ":" + port
	}

	conn, err := net.Dial("udp4", addr)
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("CountDown...")
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

	Stop(robotid, conn)
}
