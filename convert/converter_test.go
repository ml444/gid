package convert

import (
	"github.com/ml444/gid/server"
	"testing"
)

func TestIdConvert_ConvertToGen(t *testing.T) {
	//self := Convertor{*server.NewIdMeta(1)}

	var idArr = [2]map[string]uint64{
		{
			"machine": 1, "sequence": 0, "time": 1323400746, "method": 2,
			"type": 0, "version": 0, "result": 2307230695474331649,
		},
		{
			"machine": 1, "sequence": 0, "time": 1323475, "method": 2,
			"type": 1, "version": 0, "result": 6918950098101600257,
		},
	}
	for i := 0; i < len(idArr); i++ {
		dict := idArr[i]
		id := &server.Id{} // {0 1 0 1320370752 2 0 0}
		id.SetMachine(dict["machine"])
		id.SetSeq(dict["sequence"])
		id.SetTime(dict["time"])
		id.SetGenMethod(dict["method"])
		id.SetType(dict["type"])
		id.SetVersion(dict["version"])
		convertor := NewConvertor(server.NewIdMeta(dict["type"]))
		longId := convertor.ConvertToGen(id)
		if longId != dict["result"] {
			t.Errorf("%d != %d", longId, dict["result"])
		}
	}
}

func TestIdConvert_ConvertToExp(t *testing.T) {
	var idArr = [2]map[string]uint64{
		{
			"machine": 1, "sequence": 0, "time": 1323400746, "method": 2,
			"type": 0, "version": 0, "result": 2307230695474331649,
		},
		{
			"machine": 1, "sequence": 0, "time": 1323475, "method": 2,
			"type": 1, "version": 0, "result": 6918950098101600257,
		},
	}

	for i := 0; i < len(idArr); i++ {
		dict := idArr[i]
		convertor := NewConvertor(server.NewIdMeta(dict["type"]))
		idObj := convertor.ConvertToExp(dict["result"])
		if idObj.GetTime() != dict["time"] ||
			idObj.GetType() != dict["type"] ||
			idObj.GetSeq() != dict["sequence"] ||
			idObj.GetMachine() != dict["machine"] {
			t.Error("The explain longId is wrong.")
		}
	}
}

func BenchmarkIdConvert_ConvertToGen(b *testing.B) {
	var idArr = [2]map[string]uint64{
		{
			"machine": 1, "sequence": 0, "time": 1323400746, "method": 2,
			"type": 0, "version": 0, "result": 2307230695474331649,
		},
		{
			"machine": 1, "sequence": 0, "time": 1323475, "method": 2,
			"type": 1, "version": 0, "result": 6918950098101600257,
		},
	}
	for i := 0; i < len(idArr); i++ {
		dict := idArr[i]
		id := &server.Id{} // {0 1 0 1320370752 2 0 0}
		id.SetMachine(dict["machine"])
		id.SetSeq(dict["sequence"])
		id.SetTime(dict["time"])
		id.SetGenMethod(dict["method"])
		id.SetType(dict["type"])
		id.SetVersion(dict["version"])
		convertor := NewConvertor(server.NewIdMeta(dict["type"]))
		b.ResetTimer() // 重置定时器
		for i := 0; i < b.N; i++ {
			longId := convertor.ConvertToGen(id)
			if longId != dict["result"] {
				b.Errorf("%d != %d", longId, dict["result"])
			}
		}

	}
}

// go test -bench=. -run=none

//func BenchmarkIdConvert_ConvertToExp_Parallel(b *testing.B) {
//
//}

// 并发测试
//func BenchmarkTemplateParallel(b *testing.B) {
//	temp := template.Must(template.New("test").Parse("Hello, {{.}}!"))
//	b.RunParallel(func(pb *testing.PB) {
//		var buf bytes.Buffer
//		for pb.Next() {
//			buf.Reset()
//			temp.Execute(&buf, "World")
//		}
//	})
//}
