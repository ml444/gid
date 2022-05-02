package core

type Converter interface {
	ConvertToGen(id Ider) uint64 // 合成一个长整型的ID
	ConvertToExp(id uint64, out Ider)     // 拆解长整型的ID
}
