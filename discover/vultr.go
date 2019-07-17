package discover

import (
	"net"
)

const (
	vulHostnameURL = "http://169.254.169.254/v1/hostname"
	vulPublicIPv4URL = "http://169.254.169.254/v1/interfaces/0/ipv4/address"
	vulPublicIPv6URL = "http://169.254.169.254/v1/interfaces/0/ipv6/address"
)

// NewVultrDiscoverer returns a new Vultr network discoverer
func NewVultrDiscoverer() Discoverer {
	return NewDiscoverer(
		PublicHostnameDiscovererOption(vulHostname),
		PublicIPv4DiscovererOption(vulPublicIPv4),
		PublicIPv6DiscovererOption(vulPublicIPv6),
	)

}

func vulHostname() (string, error) {
	return StandardHostnameFromHTTP(vulHostnameURL, nil)
}

func vulPublicIPv4() (net.IP, error) {
	return StandardIPFromHTTP(vulPublicIPv4URL, nil)
}

func vulPublicIPv6() (net.IP, error) {
	return StandardIPFromHTTP(vulPublicIPv6URL, nil)
}
