package cityhash

import (
	"testing"
)

type param struct {
	data []byte
	hash int64
}

var params = []param{
	param{
		data: []byte("test-key"),
		hash: 4534655211177281079,
	},
	param{
		data: []byte("hello world"),
		hash: 6617098184377254238,
	},
	param{
		data: []byte("测试"),
		hash: -18718088446212349,
	},
	param{
		data: []byte("测试1hello world"),
		hash: 6682620565324616919,
	},
	param{
		data: []byte("hell测试编码o world"),
		hash: -434759769599644576,
	},
	param{
		data: []byte("he我的测试编码llo world"),
		hash: 8070307409891745327,
	},
	param{
		data: []byte("hello world哈哈，测试"),
		hash: 6663914289503119216,
	},
}

func TestCityHash64(t *testing.T) {
	for i, p := range params {
		nhash, err := CityHash64(p.data, int64(len(p.data)))
		if err != nil {
			t.Fatalf("err:%v", nhash)
		}

		if p.hash != nhash {
			t.Fatalf("index:%v expect:%v nhash:%v data:%v", i, p.hash, nhash, string(p.data))
		}
	}

	t.Logf("success")
}
