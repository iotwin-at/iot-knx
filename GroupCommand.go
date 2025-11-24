package knx

type GroupCommand struct {
	Command     Apci
	Destination DestAddr
	Data        []byte
}
