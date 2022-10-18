package block

const (
	BlockSize int = 4096
	AlignSize     = 4096
	HeadSize      = 8
	FootSize      = 8
	BodySize      = BlockSize - HeadSize - FootSize
	//BodySize = 4080
)

func GetSize(payloadLen int) (int, int) {

	length := payloadLen + HeadSize
	bodySize := BlockSize - FootSize
	size := (length-1)/bodySize + 1
	remainder := payloadLen % bodySize

	return size, remainder
}
