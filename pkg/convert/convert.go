//类型转换
package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	return strconv.Atoi(s.String())
}

func (s StrTo) MustInt() int {
	v,_:=strconv.Atoi(s.String())
	return v
}

func (s StrTo) Uint32() (uint32, error) {
	v,err:=strconv.Atoi(s.String())
	return uint32(v),err
}

func (s StrTo) MustUint32() uint32 {
	v,_:=strconv.Atoi(s.String())
	return uint32(v)
}
