// Package common provides common functions.
package common

import (
	"net"
	"strings"

	log "../../logrus"
	"../../dns"
)

var ReservedIPNetworkList []*net.IPNet

func init() {

	ReservedIPNetworkList = getReservedIPNetworkList()
}

func IsIPMatchList(ip net.IP, ipnl []*net.IPNet, isLog bool, name string) bool {

	for _, ip_net := range ipnl {
		if ip_net.Contains(ip) {
			if isLog {
				log.Debug("Matched: IP network " + name + " " + ip.String() + " " + ip_net.String())
			}
			return true
		}
	}

	return false
}

func HasAnswer(m *dns.Msg) bool { return len(m.Answer) != 0 }

func HasSubDomain(s string, sub string) bool {

	return strings.HasSuffix(sub, "."+s) || s == sub
}

func getReservedIPNetworkList() []*net.IPNet {

	ipnl := make([]*net.IPNet, 0)
	localCIDR := []string{"127.0.0.0/8", "10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16", "100.64.0.0/10"}
	for _, c := range localCIDR {
		_, ip_net, err := net.ParseCIDR(c)
		if err != nil {
			break
		}
		ipnl = append(ipnl, ip_net)
	}
	return ipnl
}

func FindRecordByType(msg *dns.Msg, t uint16) string {

	for _, rr := range msg.Answer {
		if rr.Header().Rrtype == t {
			items := strings.SplitN(rr.String(), "\t", 5)
			return items[4]
		}
	}

	return ""
}
