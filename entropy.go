package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
)

const bufSize = 8388608 // 8Mb

func calcEntropy(path string) (float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var (
		freq  [256]uint64
		total float64
	)

	read := bufio.NewReaderSize(file, bufSize)
	buf := make([]byte, bufSize)

	for {
		done, err := read.Read(buf)
		if done > 0 {
			total += float64(done)
			for _, val := range buf[:done] {
				freq[val]++
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
	}

	if total == .0 {
		return 0, nil
	}

	var ent float64
	for _, chr := range freq {
		if chr > 0 {
			pnt := float64(chr) / total
			ent += pnt * math.Log2(pnt)
		}
	}

	return -ent, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <file>\n", filepath.Base(os.Args[0]))
		return
	}

	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	ent, err := calcEntropy(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%.3f %s\n", ent, path)
}
