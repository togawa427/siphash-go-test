package main

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/dchest/siphash"
)

func siphashLoop(n int) time.Duration {
	start := time.Now()
	msg := []byte("44c241260f1c6faa09f782c9") // 固定メッセージ

	for i := 0; i < n; i++ {
		randomKey := make([]byte, 16)
		_, err := rand.Read(randomKey)
		if err != nil {
			panic(err)
		}

		key1 := binary.LittleEndian.Uint64(randomKey[:8])
		key2 := binary.LittleEndian.Uint64(randomKey[8:])
		siphash.Hash(key1, key2, msg)
	}

	return time.Since(start)
}

// --- 平均時間を計測してCSV出力 ---
func main() {
	outputFile := "out/" + time.Now().Format("2006-01-02-15:04:05") + ".csv"
	file, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// CSVヘッダ
	writer.Write([]string{"KeyCount", "Trial", "ExecutionTime(ms)"})

	// 100〜10000まで100ずつ増やす
	var testCases []int
	for i := 100; i <= 10000; i += 100 {
		testCases = append(testCases, i)
	}
	numTrials := 10

	for _, keyCount := range testCases {
		fmt.Printf("\n=== Testing with %d keys ===\n", keyCount)

		var totalTime float64
		for i := 1; i <= numTrials; i++ {
			elapsed := siphashLoop(keyCount)
			ms := float64(elapsed.Microseconds()) / 1000.0
			totalTime += ms
			fmt.Printf("[%d/%d] %d keys → %.3f ms\n", i, numTrials, keyCount, ms)

			writer.Write([]string{
				fmt.Sprintf("%d", keyCount),
				fmt.Sprintf("%d", i),
				fmt.Sprintf("%.3f", ms),
			})
		}

		avg := totalTime / float64(numTrials)
		fmt.Printf("Average for %d keys: %.3f ms\n", keyCount, avg)
		writer.Write([]string{
			fmt.Sprintf("%d", keyCount),
			"average",
			fmt.Sprintf("%.3f", avg),
		})
	}

	fmt.Printf("\n結果を %s に出力しました\n", outputFile)
}
