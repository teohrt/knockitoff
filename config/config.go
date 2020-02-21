package config

import (
	"net"

	"github.com/teohrt/knockitoff/packets"
)

type Config struct {
	SrcIP   string `yaml:"srcIP"`
	DstIP   string `yaml:"dstIP"`
	DstMAC  string `yaml:"dstMAC"`
	SrcMAC  string `yaml:"srcMAC"`
	SrcPort int    `yaml:"srcPort"`
	DstPort int    `yaml:"dstPort"`
}

// TODO: Additional field validation
func ConvertToPacketBase(c *Config) (*packets.PacketBase, error) {
	srcMAC, err := net.ParseMAC(c.SrcMAC)
	if err != nil {
		return nil, err
	}
	dstMAC, err := net.ParseMAC(c.DstMAC)
	if err != nil {
		return nil, err
	}
	return &packets.PacketBase{
		SrcIP:   net.ParseIP(c.SrcIP),
		SrcMAC:  srcMAC,
		SrcPort: uint16(c.SrcPort),
		DstIP:   net.ParseIP(c.DstIP),
		DstMAC:  dstMAC,
		DstPort: uint16(c.DstPort),
	}, nil
}
