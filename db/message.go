package DB

// For unmarshalling Client.Message from JSON
// I regret everything
type Message struct {
	Source      Address
	Dest        Address
	Timestamp   int64
	ClientId    string
	MessageType string
}

type Address struct {
	IP   string
	Port int
}
