package monkey

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func Test_armv8ImmediateMovz(t *testing.T) {
	tests := []struct {
		name         string
		movx         int
		regN         int
		shift        int
		value        uint16
		assembleFunc func(int, int, uint16) []byte
		want         []byte
	}{
		{
			name:         "MOVZ x26, 0xffff",
			movx:         0b10,
			regN:         26,
			shift:        0,
			value:        0xffff,
			assembleFunc: armv8ImmediateMovz,
			want:         []byte{0xfa, 0xff, 0x9f, 0xd2},
		},
		{
			name:         "MOVK x0, 0x1234, LSL #32",
			movx:         0b11,
			regN:         0,
			shift:        2,
			value:        0x1234,
			assembleFunc: arm8ImmediateMovk,
			want:         []byte{0x80, 0x46, 0xc2, 0xf2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instruction := movx(tt.movx, tt.regN, tt.shift, tt.value)
			if got := instruction; !reflect.DeepEqual(got, tt.want) {
				t.Log(strings.Join(details("want", tt.want), " "))
				t.Log(strings.Join(details("got ", got), " "))
				t.Error("unexpected result")
				return
			}

			if !reflect.DeepEqual(instruction, tt.assembleFunc(tt.regN, tt.shift, tt.value)) {
				t.Error("mismatch against an assemble function")
			}
		})
	}
}

func details(head string, want []byte) []string {
	res := []string{head}
	for _, b := range want {
		res = append(res, strconv.FormatUint(uint64(b), 2))
	}
	lwant := make([]string, len(want))
	for i := range want {
		lwant[i] = strconv.FormatUint(uint64(want[i]), 16)
	}
	res = append(res, " : ", strings.Join(lwant, " "))
	return res
}
