package dpt

import "fmt"

type Dpt8011 V16

// Value implements IDpt.
func (d Dpt8011) Value() V16 {
	return V16(d)
}

// Name implements IDpt.
func (d Dpt8011) Name() string {
	return "DPT_Rotation_Angle"
}

// String implements IDpt.
func (d Dpt8011) String() string {
	return fmt.Sprintf("%v%s", V16(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt8011) Pack() []byte {
	return packV16(d.Value())
}

// Unit implements IDpt.
func (d Dpt8011) Unit() string {
	return "°"
}

func UnpackDpt8011(data []byte) (IDpt[V16], error) {
	v, err := unpackV16(data)
	return Dpt8011(v), err
}
