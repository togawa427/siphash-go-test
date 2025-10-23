package main

import (
	"fmt"

	"github.com/dchest/siphash"
)

func main() {
	fmt.Println("Hello world")
	// SipHash用の16バイトのキーを指定（任意の値で設定）

	// XIAO側：const uint8_t key[] PROGMEM = {0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x00}; // ビーコンに記載する番号
	// var key1 uint64 = 0x0807060504030201
	// var key2 uint64 = 0x000F0E0D0C0B0A09

	// XIAO側：const uint8_t key[] PROGMEM = {0x0d, 0xbe, 0xb7, 0xc9, 0x01, 0xac, 0xeb, 0x36, 0x35, 0x3c, 0x8a, 0x0d, 0x23, 0x47, 0x60, 0xe3}; // ビーコンに記載する番号
	// var key1 uint64 = 0x36ebac01c9b7be0d
	// var key2 uint64 = 0xe36047230d8a3c35

	// XIAO側：const uint8_t key[] PROGMEM = {0x40, 0x50, 0xb4, 0x30, 0x2d, 0xa0, 0xe0, 0x3b, 0x54, 0xda, 0x69, 0x89, 0xcc, 0x3b, 0xf3, 0x35}	// fuma(実験)
	// var key1 uint64 = 0x3be0a02d30b45040
	// var key2 uint64 = 0x35f33bcc8969da54

	// rui(ローカル): 2e60aa6ae7f8f6030e9ab0673e3e7510
	// XIAO側：const uint8_t key[] PROGMEM = {0x2e, 0x60, 0xaa, 0x6a, 0xe7, 0xf8, 0xf6, 0x03, 0x0e, 0x9a, 0xb0, 0x67, 0x3e, 0x3e, 0x75, 0x10}
	var key1 uint64 = 0x03f6f8e76aaa602e
	var key2 uint64 = 0x10753e3e67b09a0e

	// ハッシュ化したいデータ
	// msg := []byte("05F7CA921B954D65")	// ランダム値
	// msg := []byte("c8908f0824969b694b597620") // ランダム値
	msg := []byte("44c241260f1c6faa09f782c9")
	// msg := []byte("d4ed8976601f0431ed50f404")
	// msg := []byte("332222")	// ランダム値 4c00090834fa106581b50a0412b51449

	// SipHashを計算
	hash := siphash.Hash(key1, key2, msg) // ここで第1, 第2引数にカスタム値を指定可能
	fmt.Printf("SipHash result: %x\n", hash)
	fmt.Printf("MSD: ffff%x%s\n", hash, string(msg))
}
