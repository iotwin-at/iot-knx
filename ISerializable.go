package knx

type ISerializable interface {
	Pack() []byte
}
