package models

type Protocol int

const (
	ProtocolHTTPS Protocol = iota + 1
	ProtocolDNS
)

func (p Protocol) String() string {
	switch p {
	case ProtocolHTTPS:
		return "HTTPS"
	case ProtocolDNS:
		return "DNS"
	default:
		return "Unknown"
	}
}
