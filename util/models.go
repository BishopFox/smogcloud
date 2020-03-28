package util

type Results struct {
	Hostnames []string `json:"hostnames"`
	URIs      []string `json:"uris"`
	IPs       []string `json:"ips"`
}

func (s Results) isEmpty() bool {
	return len(s.Hostnames)+len(s.URIs)+len(s.IPs) == 0
}

func CreateResultsObject() Results {
	return Results{
		Hostnames: []string{},
		URIs:      []string{},
		IPs:       []string{},
	}
}
