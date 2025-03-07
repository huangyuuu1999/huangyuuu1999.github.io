package main

import (
	"bytes"
	"container/heap"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"strings"
)

const placeholder byte = '#'

type TreeNode struct {
	val         byte
	freq        int
	left, right *TreeNode
}

type TreeNodeHeap []TreeNode

func (t TreeNodeHeap) Len() int {
	return len(t)
}

func (t TreeNodeHeap) Less(i, j int) bool {
	if t[i].freq < t[j].freq {
		return true
	} else if t[i].freq == t[j].freq {
		return t[i].val < t[j].val
	}
	return false
}

func (t TreeNodeHeap) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t *TreeNodeHeap) Push(x any) {
	*t = append(*t, x.(TreeNode))
}

func (t *TreeNodeHeap) Pop() any {
	old := *t
	n := len(old)
	x := old[n-1]
	*t = old[:n-1]
	return x
}

func counter(s string, verbose bool) map[byte]int {
	n, m := len(s), map[byte]int{}
	for i := 0; i < n; i++ {
		m[s[i]]++
	}
	if verbose {
		fmt.Printf("{ ")
		for k, v := range m {
			fmt.Printf("%v:%v ", string(k), v)
		}
		fmt.Printf("}\n")
	}
	return m
}

// 构建哈夫曼 编码树
func buildTreeFromHeap(th *TreeNodeHeap) *TreeNode {
	if len(*th) == 1 {
		root := heap.Pop(th).(TreeNode)
		return &root
	}
	node1, node2 := heap.Pop(th).(TreeNode), heap.Pop(th).(TreeNode)
	if node1.freq > node2.freq {
		node1, node2 = node2, node1
	}
	mergeNode := TreeNode{placeholder, node1.freq + node2.freq, &node1, &node2}
	heap.Push(th, mergeNode)
	return buildTreeFromHeap(th)
}

func getEncodeDictionary(s string) (map[byte]string, map[string]byte) {
	m := counter(s, true)
	treeNodes := []TreeNode{}
	for b, freq := range m {
		treeNodes = append(treeNodes, TreeNode{b, freq, nil, nil})
	}
	t := TreeNodeHeap(treeNodes)
	heap.Init(&t)
	root := buildTreeFromHeap(&t)

	path := []string{} // 保存 "0" "1", 后面再处理为位

	byteToCode := map[byte]string{}
	codeToByte := map[string]byte{}

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.left == root.right {
			code := strings.Join(path, "")
			byteToCode[root.val] = code
			codeToByte[code] = root.val
			return
		}
		path = append(path, "0")
		dfs(root.left)
		path = path[:len(path)-1]

		path = append(path, "1")
		dfs(root.right)
		path = path[:len(path)-1]
	}
	dfs(root)
	printDictionary(byteToCode, codeToByte)
	return byteToCode, codeToByte
}

func printDictionary(a map[byte]string, b map[string]byte) {
	fmt.Println("==== encode dictionary ====")
	for k, v := range a {
		fmt.Printf("%v -> encode -> %v \n", string(k), v)
	}

	fmt.Println("==== decode dictionary ====")
	for k, v := range b {
		fmt.Printf("%v -> decode -> %v \n", k, string(v))
	}
}

// 原始字符压缩成字节编码
func compressString(s string, byteTocode map[byte]string) []byte {
	bytes := []byte{}
	for i := 0; i < len(s); i++ {
		bytes = append(bytes, []byte(byteTocode[s[i]])...)
	}
	return bytes
}

func bytesToBits(bytes []byte) ([]byte, int) {
	// 把[]byte{'1','1','0','0', '1','1','0','0'} 变成比特位 11001100
	l := len(bytes)
	whiteSpace := (8 - l%8) % 8
	byteCnt := (l + 7) / 8
	ans := make([]byte, byteCnt)
	for j, b := range bytes {
		if b == '1' {
			ans[j/8] |= 1 << (7 - j%8)
		}
	}
	fmt.Printf("bits: %08b\n", ans)
	fmt.Printf("whiteSpace: %d\n", whiteSpace)
	return ans, whiteSpace
}

// 反向操作：将 []byte 和 whiteSpace 转换回原始的比特数据
func bitsToBytes(bits []byte, whiteSpace int) []byte {
	// 计算原始比特数据的长度
	bitLength := len(bits)*8 - whiteSpace
	// 创建一个切片来存储结果
	result := make([]byte, 0, bitLength)
	// 遍历每个字节
	for i := 0; i < len(bits); i++ {
		// 遍历每个比特（从最高位到最低位）
		for j := 7; j >= 0; j-- {
			// 如果已经处理完所有有效比特，则退出
			if len(result) >= bitLength {
				break
			}
			// 提取当前比特的值
			bit := (bits[i] >> j) & 1
			if bit == 1 {
				result = append(result, '1')
			} else {
				result = append(result, '0')
			}
		}
	}
	return result
}

// 频率表定长占用 256int 256 * 8 byte
func counterToFreqTable(counter map[byte]int) []int {
	freqTable := make([]int, 256) //频率表数组
	for b, f := range counter {
		freqTable[b] = f
	}
	return freqTable
}

func intSliceToByteSlice(intSlice []int) ([]byte, error) {
	buf := new(bytes.Buffer)

	// 遍历 int 切片，将每个 int 写入缓冲区
	for _, v := range intSlice {
		err := binary.Write(buf, binary.LittleEndian, int64(v)) // 使用 int64 保证兼容性
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func byteSliceToIntSlice(byteSlice []byte) ([]int, error) {
	// 检查字节切片的长度是否是 int64 的整数倍
	if len(byteSlice)%8 != 0 {
		return nil, errors.New("byte slice length is not a multiple of 8")
	}

	// 创建一个缓冲区来读取字节切片
	buf := bytes.NewReader(byteSlice)

	// 计算需要读取的 int64 数量
	numInts := len(byteSlice) / 8
	intSlice := make([]int, numInts)

	// 遍历字节切片，读取每个 int64 并转换为 int
	for i := 0; i < numInts; i++ {
		var int64Val int64
		err := binary.Read(buf, binary.LittleEndian, &int64Val)
		if err != nil {
			return nil, err
		}
		intSlice[i] = int(int64Val)
	}

	return intSlice, nil
}

func compressFile(srcPath, destPath string) {
	/*
		freqTable(256byte) + whiteSpace(1byte) + content(...)
	*/
	info, err := os.Stat(srcPath)
	if err != nil {
		fmt.Println("无法获取文件信息:", err)
		return
	} else if info.Size() >= 10*1024*1024 {
		panic("文件过大")
	}
	data, err := os.ReadFile(srcPath)
	if err != nil {
		fmt.Println("文件读取失败:", err)
		return
	}
	s := string(data)
	byteToCode, _ := getEncodeDictionary(s)

	bytes := compressString(s, byteToCode)
	bitContent, whiteSpace := bytesToBits(bytes)

	freqTable := counterToFreqTable(counter(s, false))
	fmt.Printf("freqTable: %v\n", freqTable)

	byteSlice, _ := intSliceToByteSlice(freqTable)
	whiteSpaceByte := byte(whiteSpace)

	_, err = os.Stat(destPath)
	if os.IsExist(err) {
		fmt.Println("目标文件已存在")
		return
	}

	byteSlice = append(byteSlice, whiteSpaceByte)
	byteSlice = append(byteSlice, bitContent...)

	os.WriteFile(destPath, byteSlice, 0644)
}

const FreqTableSize = 2048

func unCompressFile(srcPath, destPath string) {
	data, err := os.ReadFile(srcPath)
	if err != nil {
		fmt.Println("文件读取失败:", err)
		return
	}
	freqTable := data[:FreqTableSize]
	whiteSpace := data[FreqTableSize]
	content := data[FreqTableSize+1:]
	// fmt.Printf("freqTable: %v\n", freqTable)
	// fmt.Printf("whiteSpace: %v\n", whiteSpace)
	// fmt.Printf("content: %v\n", content)
	freq, _ := byteSliceToIntSlice(freqTable)

	// fmt.Printf("freq: %v\n", freq)

	contentBytes := bitsToBytes(content, int(whiteSpace))

	// fmt.Printf("contentBytes: %v\n", string(contentBytes))

	nodes := []TreeNode{}
	for i, v := range freq {
		if v != 0 {
			nodes = append(nodes, TreeNode{byte(i), v, nil, nil})
		}
	}
	t := TreeNodeHeap(nodes)
	heap.Init(&t)
	root := buildTreeFromHeap(&t)

	// 一个bit一个bit地读取content，遍历树，翻译成字节'a','b'... 最后写入解压的文件
	unCompressData := []byte{}
	cur := root
	for k := 0; k < len(contentBytes); k++ {
		if contentBytes[k] == '0' {
			cur = cur.left
		} else if contentBytes[k] == '1' {
			cur = cur.right
		}
		if cur.left == cur.right { //leaf
			unCompressData = append(unCompressData, cur.val)
			cur = root
		}
	}
	fmt.Printf("unCompressData: %v\n", unCompressData)
	os.WriteFile(destPath, unCompressData, 0644)
}

func main() {

	compressFile("a.txt", "a.bin")
	unCompressFile("a.bin", "a.copt.txt")
}
