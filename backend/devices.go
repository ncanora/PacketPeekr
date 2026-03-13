import (
	"os"
	"fmt"
	"net"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// Takes device name and returns handle to read packets from
func openHandle(iface string) (*pcap.Handle, error) {
	return pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
}

// Takes gopacket and MAC + IP addr
func parseARP(packet gopacket.Packet) (mac, ip string, ok bool) {
	layer := packet.Layer(layers.LayerTypeARP)
	if layer == nil {
		return "", "", false
	}
	arp := layer.(*layers.ARP)
	mac = net.HardwareAddr(arp.SourceHwAddress).String()
	ip = net.IP(arp.SourceProtAddress).String()
	return mac, ip, true
} 