package utils

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
)

// 开奖前需要的参数
type LotteryInput struct {
	PreviousSeedHash string // 上次开奖的种子哈希（开奖前已知）
	TimeStamp        int64  // 开奖时间戳（开奖前已知）
}

// 开奖结果
type LotteryResult struct {
	LuckyNumber      int    // 幸运号码（开奖后生成）
	SeedString       string // 种子字符串（开奖后生成）
	CurrentSeedHash  string // 当前种子哈希（开奖后生成，用于下一轮）
	TimeStamp        int64  // 开奖时间戳（开奖前已知）
	PreviousSeedHash string
	SessionId        int
	Gid              int
}

// 验证需要的参数
type VerifyInput struct {
	PreviousSeedHash string `json:"previous_seed_hash"` // 上次开奖的种子哈希（开奖前已知）
	TimeStamp        int64  `json:"time_stamp"`         // 开奖时间戳（开奖前已知）
	LuckyNumber      int    `json:"lucky_number"`       // 公布的幸运号码（开奖后生成）
	CurrentSeedHash  string `json:"current_seed_hash"`  // 公布的当前种子哈希（开奖后生成）
}

// GenerateZeroHash 计算字符串 "0" 的 SHA256 哈希
func GenerateZeroHash() string {
	hash := sha256.Sum256([]byte("0"))
	return hex.EncodeToString(hash[:])
}

// GenerateLuckyNumber 生成幸运号码（1-8）
func GenerateLuckyNumber(input LotteryInput) (*LotteryResult, error) {
	// 步骤 1：检查 PreviousSeedHash
	if input.PreviousSeedHash == "" {
		input.PreviousSeedHash = GenerateZeroHash() // 首次开奖用 SHA256("0")
	} else if len(input.PreviousSeedHash) != 64 {
		return nil, fmt.Errorf("invalid PreviousSeedHash format")
	}

	// 步骤 2：组合种子因子（只使用上一轮hash和时间戳）
	timeStampStr := fmt.Sprintf("%d", input.TimeStamp)
	fmt.Println("timeStampStr::", timeStampStr)
	seedStr := strings.Join([]string{
		input.PreviousSeedHash,
		timeStampStr,
	}, "_")

	// 步骤 3：计算种子字符串的 SHA256 哈希
	hash := sha256.Sum256([]byte(seedStr))
	seedHash := hex.EncodeToString(hash[:])

	// 步骤 4：生成幸运号码
	num := binary.BigEndian.Uint64(hash[:8])
	luckyNumber := int(num%8) + 1

	return &LotteryResult{
		LuckyNumber:      luckyNumber,
		SeedString:       seedStr,
		CurrentSeedHash:  seedHash,
		TimeStamp:        input.TimeStamp,
		PreviousSeedHash: input.PreviousSeedHash,
	}, nil
}

// VerifyLottery 验证幸运号码和当前种子哈希
func VerifyLottery(input VerifyInput) (bool, string, error) {
	// 步骤 1：检查 PreviousSeedHash
	if input.PreviousSeedHash == "" {
		input.PreviousSeedHash = GenerateZeroHash()
	} else if len(input.PreviousSeedHash) != 64 {
		return false, "", fmt.Errorf("invalid PreviousSeedHash format")
	}

	// 步骤 2：组合种子因子（只使用上一轮hash和时间戳）
	timeStampStr := fmt.Sprintf("%d", input.TimeStamp)
	seedStr := strings.Join([]string{
		input.PreviousSeedHash,
		timeStampStr,
	}, "_")

	// 步骤 3：计算种子字符串的 SHA256 哈希
	hash := sha256.Sum256([]byte(seedStr))
	seedHash := hex.EncodeToString(hash[:])

	// 步骤 4：生成幸运号码
	num := binary.BigEndian.Uint64(hash[:8])
	calculatedLuckyNumber := int(num%8) + 1

	// 步骤 5：验证幸运号码
	if calculatedLuckyNumber != input.LuckyNumber {
		return false, seedStr, fmt.Errorf("lucky number mismatch: calculated %d, expected %d", calculatedLuckyNumber, input.LuckyNumber)
	}

	// 步骤 6：验证当前种子哈希
	if seedHash != input.CurrentSeedHash {
		return false, seedStr, fmt.Errorf("current seed hash mismatch: calculated %s, expected %s", seedHash, input.CurrentSeedHash)
	}

	return true, seedStr, nil
}

// 打印开奖结果
func PrintLotteryResult(roundNumber int, input LotteryInput, result *LotteryResult) {
	fmt.Printf("\n=== 第 %d 轮开奖 ===\n", roundNumber)
	fmt.Println("\n开奖参数:")
	fmt.Printf("上一轮种子哈希: %s\n", input.PreviousSeedHash)
	fmt.Printf("开奖时间戳: %d\n", input.TimeStamp)

	fmt.Println("\n开奖结果:")
	fmt.Printf("幸运号码: %d\n", result.LuckyNumber)
	fmt.Printf("种子字符串: %s\n", result.SeedString)
	fmt.Printf("当前种子哈希: %s\n", result.CurrentSeedHash)

	// 打印计算过程
	fmt.Println("\n计算过程:")
	fmt.Printf("1. 组合种子字符串: %s_%d\n", input.PreviousSeedHash, input.TimeStamp)
	fmt.Printf("2. 计算SHA256哈希: %s\n", result.CurrentSeedHash)
	fmt.Printf("3. 取前8字节: %x\n", result.CurrentSeedHash[:16])
	fmt.Printf("4. 计算幸运号码: %d\n", result.LuckyNumber)
}

// 打印验证结果
func PrintVerifyResult(verifyInput VerifyInput, isValid bool, seedStr string) {
	fmt.Println("\n=== 验证详情 ===")
	fmt.Println("\n验证参数:")
	fmt.Printf("上一轮种子哈希: %s\n", verifyInput.PreviousSeedHash)
	fmt.Printf("开奖时间戳: %d\n", verifyInput.TimeStamp)
	fmt.Printf("幸运号码: %d\n", verifyInput.LuckyNumber)
	fmt.Printf("当前种子哈希: %s\n", verifyInput.CurrentSeedHash)

	fmt.Println("\n验证过程:")
	fmt.Printf("1. 组合种子字符串: %s\n", seedStr)

	// 重新计算哈希
	hash := sha256.Sum256([]byte(seedStr))
	seedHash := hex.EncodeToString(hash[:])
	fmt.Printf("2. 计算SHA256哈希: %s\n", seedHash)

	// 重新计算幸运号码
	num := binary.BigEndian.Uint64(hash[:8])
	calculatedLuckyNumber := int(num%8) + 1
	fmt.Printf("3. 取前8字节: %x\n", seedHash[:16])
	fmt.Printf("4. 计算幸运号码: %d\n", calculatedLuckyNumber)

	fmt.Println("\n验证结果:")
	fmt.Printf("幸运号码验证: %v\n", calculatedLuckyNumber == verifyInput.LuckyNumber)
	fmt.Printf("哈希验证: %v\n", seedHash == verifyInput.CurrentSeedHash)
	fmt.Printf("总体验证: %v\n", isValid)
}
