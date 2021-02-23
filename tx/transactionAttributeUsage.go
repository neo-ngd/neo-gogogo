package tx

import "strconv"

// Transaction attribute usages
type TransactionAttributeUsage uint8

const (
	ContractHash TransactionAttributeUsage = 0x00 // 0

	ECDH02 TransactionAttributeUsage = 0x02 // 2
	ECDH03 TransactionAttributeUsage = 0x03 // 3

	Script TransactionAttributeUsage = 0x20 // 32

	Vote TransactionAttributeUsage = 0x30 // 48

	DescriptionUrl TransactionAttributeUsage = 0x81 // 129
	Description    TransactionAttributeUsage = 0x90 // 144

	Hash1  TransactionAttributeUsage = 0xa1 // 161
	Hash2  TransactionAttributeUsage = 0xa2
	Hash3  TransactionAttributeUsage = 0xa3
	Hash4  TransactionAttributeUsage = 0xa4
	Hash5  TransactionAttributeUsage = 0xa5
	Hash6  TransactionAttributeUsage = 0xa6
	Hash7  TransactionAttributeUsage = 0xa7
	Hash8  TransactionAttributeUsage = 0xa8
	Hash9  TransactionAttributeUsage = 0xa9
	Hash10 TransactionAttributeUsage = 0xaa
	Hash11 TransactionAttributeUsage = 0xab
	Hash12 TransactionAttributeUsage = 0xac
	Hash13 TransactionAttributeUsage = 0xad
	Hash14 TransactionAttributeUsage = 0xae
	Hash15 TransactionAttributeUsage = 0xaf // 175

	Remark   TransactionAttributeUsage = 0xf0 // 240
	Remark1  TransactionAttributeUsage = 0xf1
	Remark2  TransactionAttributeUsage = 0xf2
	Remark3  TransactionAttributeUsage = 0xf3
	Remark4  TransactionAttributeUsage = 0xf4
	Remark5  TransactionAttributeUsage = 0xf5
	Remark6  TransactionAttributeUsage = 0xf6
	Remark7  TransactionAttributeUsage = 0xf7
	Remark8  TransactionAttributeUsage = 0xf8
	Remark9  TransactionAttributeUsage = 0xf9
	Remark10 TransactionAttributeUsage = 0xfa
	Remark11 TransactionAttributeUsage = 0xfb
	Remark12 TransactionAttributeUsage = 0xfc
	Remark13 TransactionAttributeUsage = 0xfd
	Remark14 TransactionAttributeUsage = 0xfe
	Remark15 TransactionAttributeUsage = 0xff // 255
)

func (u TransactionAttributeUsage) String() string {
	switch {
	case u == 0:
		return "ContractHash"
	case u == 2:
		return "ECDH02"
	case u == 3:
		return "ECDH03"
	case u == 32:
		return "Script"
	case u == 48:
		return "Vote"
	case u == 129:
		return "DescriptionUrl"
	case u == 144:
		return "Description"
	case 161 <= u && u <= 175:
		return "Hash" + strconv.FormatUint(uint64(u - 160), 10)
	case u == 240:
		return "Remark"
	case 241 <= u && u <= 255:
		return "Remark" + strconv.FormatUint(uint64(u - 240), 10)
	default:
		return "TransactionAttributeUsage=" + strconv.FormatUint(uint64(u), 10)
	}
}

func NewTransactionAttributeUsageFromString(s string) TransactionAttributeUsage {
	switch s {
	case "ContractHash":
		return ContractHash
	case "ECDH02":
		return ECDH02
	case "ECDH03":
		return ECDH03
	case "Script":
		return Script
	case "Vote":
		return Vote
	case "DescriptionUrl":
		return DescriptionUrl
	case "Description":
		return Description
	case "Hash1":
	case "Hash2":
	case "Hash3":
	case "Hash4":
	case "Hash5":
	case "Hash6":
	case "Hash7":
	case "Hash8":
	case "Hash9":
	case "Hash10":
	case "Hash11":
	case "Hash12":
	case "Hash13":
	case "Hash14":
	case "Hash15":
		sub := s[4:]
		n, _ := strconv.Atoi(sub)
		return TransactionAttributeUsage(byte(n + 160))
	case "Remark":
		return Remark
	case "Remark1":
	case "Remark2":
	case "Remark3":
	case "Remark4":
	case "Remark5":
	case "Remark6":
	case "Remark7":
	case "Remark8":
	case "Remark9":
	case "Remark10":
	case "Remark11":
	case "Remark12":
	case "Remark13":
	case "Remark14":
	case "Remark15":
		sub := s[6:]
		n, _ := strconv.Atoi(sub)
		return TransactionAttributeUsage(byte(n + 240))
	default:
		return Remark
	}
	return Remark
}
