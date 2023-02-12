package network

type (
	NetAddr string

	RPC struct {
		From    NetAddr
		Payload []byte
	}
)

type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(NetAddr, []byte) error
	Addr() NetAddr
}
