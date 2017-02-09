package des

import (
	"encoding/base64"
	"testing"
)

func TestDes(t *testing.T) {
	key := []byte("test_des")
	result, err := DesEncrypt([]byte("davygeek/tc/crypt/des"), key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(base64.StdEncoding.EncodeToString(result))

	origData, err := DesDecrypt(result, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(origData))
}

func Test3Des(t *testing.T) {
	key := []byte("2fe023f_sefiel#fi32lf3e!")
	result, err := TripleDesEncrypt([]byte("davygeek/tc/crypt/des"), key)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(base64.StdEncoding.EncodeToString(result))
	origData, err := TripleDesDecrypt(result, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(origData))
}
