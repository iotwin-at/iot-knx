package knx

type ISerializable interface {
	ToBytes() []byte
}
