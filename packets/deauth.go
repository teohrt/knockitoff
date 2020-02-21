package packets

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// TODO: Initialize legitimate deauth
func NewDeauthPacket(p *PacketPreReq) ([]byte, error) {
	ethernetLayer := &layers.Ethernet{
		SrcMAC: p.SrcMAC,
		DstMAC: p.DstMAC,
	}
	ipLayer := &layers.IPv4{
		SrcIP: p.SrcIP,
		DstIP: p.DstIP,
	}
	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(p.SrcPort),
		DstPort: layers.TCPPort(p.DstPort),
	}
	tcpLayer.SetNetworkLayerForChecksum(ipLayer)
	return Serialize(
		ethernetLayer,
		ipLayer,
		tcpLayer,
		gopacket.Payload(p.Payload),
	)
}
