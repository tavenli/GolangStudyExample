package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/*
bool，0 为 false，其它为 true
byte = uint8 的值范围（8bit * 1）：0~255
int8 的值范围（8bit * 1）：-128 ~ 127
uint16 的值范围（8bit * 2）：0 ~ 65535
int16 的值范围（8bit * 2）：-32768 ~ 32767
uint32 的值范围（8bit * 4）：0 ~ 4294967295
int = int32 的值范围（8bit * 4）：-2147483648 ~ 2147483647
uint64 的值范围（8bit * 8）：0 ~ 18446744073709551615
int64 的值范围（8bit * 8）：-9223372036854775808 ~ 9223372036854775807
float32 的值范围（8bit * 4）：3.4E-38 ~ 3.4E+38
float64 的值范围（8bit * 8）：1.7E-308 ~ 1.7E+308
*/

func ReadUint16(val []byte) uint16 {
	result := binary.LittleEndian.Uint16(val)
	return result
}

func WriteUint16(buf *bytes.Buffer, val uint16) error {
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, val)
	_, err := buf.Write(bs)
	return err
}

func ReadInt32(val []byte) int32 {
	result := binary.LittleEndian.Uint32(val)
	return int32(result)
}

func WriteInt32(buf *bytes.Buffer, val int32) error {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(val))
	_, err := buf.Write(bs)
	return err
}

func ReadUint32(val []byte) uint32 {
	result := binary.LittleEndian.Uint32(val)
	return result
}

func WriteUint32(buf *bytes.Buffer, val uint32) error {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, val)
	_, err := buf.Write(bs)
	return err
}

func ReadInt64(val []byte) int64 {
	result := binary.LittleEndian.Uint64(val)
	return int64(result)
}

func WriteInt64(buf *bytes.Buffer, val int64) error {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(val))
	_, err := buf.Write(bs)
	return err
}

func ReadUint64(val []byte) uint64 {
	result := binary.LittleEndian.Uint64(val)
	return result
}

func WriteUint64(buf *bytes.Buffer, val uint64) error {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, val)
	_, err := buf.Write(bs)
	return err
}

func BinaryReadInt32(val []byte) int32 {
	var result int32
	err := BinaryReadAny(val, &result)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	return result
}

func BinaryReadAny(val []byte, result interface{}) error {
	bufRead := bytes.NewReader(val)
	return binary.Read(bufRead, binary.LittleEndian, result)

}

func BinaryWriteInt32(buf *bytes.Buffer, val int32) error {
	return BinaryWriteAny(buf, val)
}

func BinaryWriteAny(buf *bytes.Buffer, val interface{}) error {
	return binary.Write(buf, binary.LittleEndian, val)
}

func _writeAny(val interface{}) []byte {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, val)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.Bytes())
	return buf.Bytes()
}

func WriteBYTE(data *bytes.Buffer, val int) {
	//BYTE 长度：1
	buf := make([]byte, 1)
	buf[0] = byte(val)

	data.Write(buf)
}

func WriteWORD(data *bytes.Buffer, val int) {
	//WORD 长度：2
	buf := make([]byte, 2)
	buf[0] = byte(val)
	buf[1] = byte(val >> 8)

	data.Write(buf)
}

func WriteDWORD(data *bytes.Buffer, val int) {
	//DWORD 长度：4
	buf := make([]byte, 4)
	buf[0] = byte(val)
	buf[1] = byte(val >> 8)
	buf[2] = byte(val >> 16)
	buf[3] = byte(val >> 24)

	data.Write(buf)
}

func WriteTCHAR(data *bytes.Buffer, size int, val string) {
	//TCHAR 长度：由size指定
	buf := []byte(val)
	data.Write(buf)
	//
	for i := 0; i < size-len(buf); i++ {
		//剩余的补0
		data.WriteByte(0)
	}

}

func WriteUnicodeTCHAR(data *bytes.Buffer, size int, val string) {
	//Unicode TCHAR 长度：size*2
	realSize := size * 2
	buf := []byte(val)
	dataBuffer := make([]byte, realSize)
	k := 0
	for j := 0; j < len(buf); j++ {
		dataBuffer[k] = buf[j]
		dataBuffer[k+1] = byte(0)
		k = k + 2
	}

	data.Write(dataBuffer)
	//

}

func WriteInt(data *bytes.Buffer, val int) {
	//Byte 长度：4
	buf := make([]byte, 4)
	buf[0] = byte(val)
	buf[1] = byte(val >> 8)
	buf[2] = byte(val >> 16)
	buf[3] = byte(val >> 24)

	data.Write(buf)
}

func _ReadInt_(data *bytes.Buffer, val []byte) (int, error) {

	return data.Read(val)
}

func ReadWord(val []byte) uint16 {
	//binary.LittleEndian.Uint16(rData[4:6])
	return binary.LittleEndian.Uint16(val)
}

func ReadDWord(val []byte) uint32 {
	return binary.LittleEndian.Uint32(val)
}

func ReadTCHAR(val []byte) string {
	return string(val)
}
