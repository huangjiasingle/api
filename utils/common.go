package utils

import ()

var (
	cmd = []byte{0xf0, 0xf0, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff}
)

func GenerateCmd(ForwardAddr, ControboxAddr, RelayAddr byte, isopen bool) []byte {
	if isopen {
		cmd[6] = 0xff
	}
	cmd[2] = ForwardAddr
	cmd[4] = ControboxAddr
	cmd[5] = RelayAddr
	u16 := CheckSum(cmd[2:8])
	cmd[8] = byte(u16 << 8 >> 8)
	cmd[9] = byte(u16 >> 8)
	return cmd
}
