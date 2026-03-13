package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const MAX_PACKETS = 2000
packet_ch := make(chan gopacket.Packet, MAX_PACKETS/4)

func main() {
	
}