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

	length := payloadLen //+ HeadSize
	size := (length-1)/BodySize + 1
	remainder := payloadLen % BodySize

	return size, remainder
}
