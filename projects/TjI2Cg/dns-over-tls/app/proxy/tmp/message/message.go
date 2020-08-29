package message

type DNSHeader struct{}

type DNSQuestion struct{}

type DNSAnswer struct{}

type DNSAuthority struct{}

type DNSAdditional struct{}

type DNSMessage struct {
	header     DNSHeader
	question   DNSQuestion
	answer     DNSAnswer
	authority  DNSAuthority
	additional DNSAdditional
}

func NewDNSMessage() *DNSMessage {
	return &DNSMessage{}
}
