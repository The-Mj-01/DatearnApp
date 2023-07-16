package IP

import (
	"github.com/labstack/echo/v4"
	"net"
)

const (
	iPV6LocalAddress = "::1"
	ipv4LocalAddress = "127.0.01"
)

// ExtractFromEcho user's ip address
func ExtractFromEcho(c echo.Context) string {
	ip := c.RealIP()

	if ip == iPV6LocalAddress {
		return ipv4LocalAddress
	}

	return tooIPV4(ip)
}

// tooIPV4 converts user ip to ipv4 address
func tooIPV4(ip string) string {
	netIP := net.ParseIP(ip)
	if netIP == nil {
		return ""
	}
	return netIP.To4().String()
}
