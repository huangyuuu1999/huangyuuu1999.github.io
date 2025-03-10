package main

import (
	"bytes"
	"fmt"
)

type BitWriter struct {
	buffer    bytes.Buffer // 用于存储字节数据
	currByte  byte         // 当前正在写入的字节
	bitOffset uint         // 当前字节的位偏移（0-7）
	validBits uint         // 最后一个字节的有效位数
}

// NewBitWriter 创建一个新的 BitWriter
func NewBitWriter() *BitWriter {
	return &BitWriter{}
}

// WriteBit 写入一个位（0 或 1）
func (bw *BitWriter) WriteBit(bit byte) error {
	if bit != 0 && bit != 1 {
		return fmt.Errorf("invalid bit value: %d (must be 0 or 1)", bit)
	}

	// 将位写入当前字节
	bw.currByte |= bit << (7 - bw.bitOffset)
	bw.bitOffset++

	// 如果当前字节已满，将其写入缓冲区
	if bw.bitOffset == 8 {
		bw.buffer.WriteByte(bw.currByte)
		bw.currByte = 0
		bw.bitOffset = 0
		bw.validBits = 8 // 当前字节已满，有效位数为 8
	}

	return nil
}

// WriteBits 写入多个位
func (bw *BitWriter) WriteBits(bits uint64, n uint) error {
	if n > 64 {
		return fmt.Errorf("cannot write more than 64 bits at a time")
	}

	for i := uint(0); i < n; i++ {
		bit := byte((bits >> (n - 1 - i)) & 1) // 从最高位开始写入
		if err := bw.WriteBit(bit); err != nil {
			return err
		}
	}

	return nil
}

// Flush 将剩余的位写入缓冲区（如果当前字节未满）
func (bw *BitWriter) Flush() error {
	if bw.bitOffset > 0 {
		bw.buffer.WriteByte(bw.currByte)
		bw.validBits = bw.bitOffset // 记录最后一个字节的有效位数
		bw.currByte = 0
		bw.bitOffset = 0
	}
	return nil
}

// Bytes 返回写入的字节数据
func (bw *BitWriter) Bytes() []byte {
	return bw.buffer.Bytes()
}

// ValidBits 返回最后一个字节的有效位数
func (bw *BitWriter) ValidBits() uint {
	return bw.validBits
}

func testBitWritter() {
	bw := NewBitWriter()

	// 写入单个位
	bw.WriteBit(1)
	bw.WriteBit(0)
	bw.WriteBit(1)

	// 写入多个位
	bw.WriteBits(0b1100, 4) // 写入 4 位：1100

	// 刷新缓冲区
	bw.Flush()

	// 获取写入的字节数据
	data := bw.Bytes()
	fmt.Printf("Written data: %08b\n", data) // 输出：10101100

	// 获取最后一个字节的有效位数
	validBits := bw.ValidBits()
	fmt.Printf("Last byte has %d valid bits\n", validBits) // 输出：Last byte has 7 valid bits
}
