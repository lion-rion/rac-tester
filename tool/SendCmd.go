package tool

import (
	"fmt"
	"net"
	"rac-tester/proto/pb_gen"
	pb "rac-tester/proto/pb_gen"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"
)

func SendCmd(cmd Commad) {

	var ipv4 string
	var port string
	var addr string
	if cmd.IsSim {
		ipv4 = "127.0.0.1"
		port = "20106"
		addr = ipv4 + ":" + port
	} else {
		ipv4 = "192.168.0." + strconv.Itoa(int(cmd.Id)+100)
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

	for i := 0; i < cmd.Loop; i++ {
		pe := &pb.GrSim_Robot_Command{
			Id:          &cmd.Id,
			Kickspeedx:  &cmd.Kickspeedx,
			Kickspeedz:  &cmd.Kickspeedz,
			Veltangent:  &cmd.Veltangent,
			Velnormal:   &cmd.Velnormal,
			Velangular:  &cmd.Velangular,
			Spinner:     &cmd.Spinner,
			Wheelsspeed: &cmd.Wheelsspeed,
		}

		var timestamp float64 = float64(time.Now().UnixNano() / 1e6)
		var isteamyellow bool = false

		command := &pb_gen.GrSim_Commands{
			Timestamp:     &timestamp,
			Isteamyellow:  &isteamyellow,
			RobotCommands: []*pb_gen.GrSim_Robot_Command{pe},
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

	Stop(cmd.Id, conn)
}
