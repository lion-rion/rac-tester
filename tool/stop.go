package tool

import (
	"net"
	"rac-tester/proto/pb_gen"
	pb "rac-tester/proto/pb_gen"
	"time"

	"google.golang.org/protobuf/proto"
)

func Stop(robotid uint32, conn net.Conn) {

	var kickspeedx float32 = 0
	var kickspeedz float32 = 0
	var veltangent float32 = 0
	var velnormal float32 = 0
	var velangular float32 = 0
	var spinner bool = false
	var wheelsspeed bool = false

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

		_, err := conn.Write([]byte(marshalpacket))
		//:debug println("send : %v", marshalpacket)
		if err != nil {
			panic(err)
		}
		time.Sleep(3 * time.Millisecond)
	}
}
