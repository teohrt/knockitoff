package packets

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
)

type PacketPreReq struct {
	SrcIP   net.IP
	DstIP   net.IP
	DstMAC  net.HardwareAddr
	SrcMAC  net.HardwareAddr
	SrcPort uint16
	DstPort uint16
	Seq     uint16
	Payload []byte
}

var SerializationOptions = gopacket.SerializeOptions{
	FixLengths:       true,
	ComputeChecksums: true,
}

func Serialize(layers ...gopacket.SerializableLayer) ([]byte, error) {
	buf := gopacket.NewSerializeBuffer()
	if err := gopacket.SerializeLayers(buf, SerializationOptions, layers...); err != nil {
		fmt.Println("Serialization error!")
		return nil, err
	}
	return buf.Bytes(), nil
}
