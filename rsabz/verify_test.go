package rsabz

import (
	"crypto/rsa"
	"math/big"
	"reflect"
	"testing"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzn2U8G3XId6zpbUvZ7zv
yeJwF73mq7CKgVEyltaKcHpV5WAvtmBvUt2SnSkuSm2JKlOY6kq4riXvIOb2dL9Q
M4Jkder4ZWtNWI8D9+Z+ibvdE99N9blVq9aRHrM1DZcOM3GxLxbjtdntyF79g9ag
ttKAZt6zE8ZKKBr1Zf44bsrxRAtxLj2DHz3MXC6cbYo34TpvPch3jTTFMnMU/oyj
XC4KX3YHvsyz7klnjTBhrO3JmAkBRW7qMNw76jBkoi/MWiWJV4tawlLPltKNO/2C
ZEG3fNK7vYVEo+cHRVNBhWDOh2I8fKmu3NxL0SoKRzeOoPTs7ZLmmIbPOXDuQ2et
jwIDAQAB
-----END PUBLIC KEY-----`

const publicKeyNoTag = `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA11coKTbUugNt9ljSQHTItLryyvAH1EI4Ww+mJyrbwAf8mnLl7kPAWBGYeS5hl5SZTDB8D1NhnK8JjU5iScagwcFbn+LKT1Z/8BdLiezwbPpSEtYWQJtb1tpgxXNYwWeYyJGhnp7d3f+/rPhBG2Y36mWjNYCAESXP+KtuXdhvsjAVccp6yRJlyiTewa6nrWBT1A5O5Lfs17XOfPqKmLWxWA30rtyJT3oMdENNY/aex0n+P5xI6eUGCHBlOukXPObBfLgHujUiDgatsKZKSNPeNRUgMsL8TLt+9p8rlZpDNfb3KCWFxVe6zm4KuLcyPpvTuy3ybZ7Gk1+doFF0keSuIQIDAQAB`

func TestVerify(t *testing.T) {
	type args struct {
		plainText string
		signature string
		publicKey string
	}
	ok := args{plainText: "name=bigzhu", signature: "BVDNSKAswoOvfwuh03eAi66B2T3mqcz6/4l04dyyU8wtMGsKXQPQi8awN+d34XhQR5G8gDPgXH9RLdBK6PJ0NAWjJmxszRD39ClQ0gSAjH+GUYcTdm+jYEL3GKA2wubsEAxqQvP0fAFtUYzszD0l/hcovZUZL/GNxpqaJrs/IuMlj7ZwlN7dXQ0XjWTxVIMqa3NTIg4cYZheklYvdPkXW6RizetvvcbdFV6gJKXKHADIqLxiRqPJb8pFAR7WtK8fJQdSJl3xQKHvk8wj5BoPviJFugQkOAjpEaq9hV2iETxLF46rZKgfi1iLJIbAfEwfMLGRbUOPHn4mDi3jxJjvxw==", publicKey: publicKey}
	errArg := args{plainText: "dfjdifjid=4&sign_type=RSA2$fuck=you", signature: "rHariRxHXaDZ1Pmvy6TKXVvvCP91xpR6VBIofPnSMSl/hr2ZejCVGGW7PNOhCWgspcpzLi19JOVoAZtCHk2XKiU7OP5/FShRCWqlEI0gGa7JBLF7IcrsTMcVOq0LCUlQjBhAFZ3LIEtOVCzXTQ61A7bjrn1mWrM9uwIx63AtMZj3/pb0oUPzvFZKxi5r0R5XHOObjDwcBhI3kv7460cb9a/mMgT4sQwOJISogRNDbuvuFhTsbNyHsKICFVndikvhki9f9xtnuSwpvf5k0uj1/n2gSiRt//hiLMJGK0HIgX7uc8EAi9RfsevxMRGISqeQaGvtCn5y7a6nvMT4OGmQpg==", publicKey: publicKey}
	noTagArg := args{plainText: "dfjdifjid=4&sign_type=RSA2", signature: "rHariRxHXaDZ1Pmvy6TKXVvvCP91xpR6VBIofPnSMSl/hr2ZejCVGGW7PNOhCWgspcpzLi19JOVoAZtCHk2XKiU7OP5/FShRCWqlEI0gGa7JBLF7IcrsTMcVOq0LCUlQjBhAFZ3LIEtOVCzXTQ61A7bjrn1mWrM9uwIx63AtMZj3/pb0oUPzvFZKxi5r0R5XHOObjDwcBhI3kv7460cb9a/mMgT4sQwOJISogRNDbuvuFhTsbNyHsKICFVndikvhki9f9xtnuSwpvf5k0uj1/n2gSiRt//hiLMJGK0HIgX7uc8EAi9RfsevxMRGISqeQaGvtCn5y7a6nvMT4OGmQpg==", publicKey: publicKeyNoTag}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", ok, false},
		{"error", errArg, true},
		{"error", noTagArg, true}, // 没有 tag 不会正确干活
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Verify(tt.args.plainText, tt.args.signature, tt.args.publicKey); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getPublicKey(t *testing.T) {
	type args struct {
		publicKey string
	}
	arg := args{publicKey: publicKey}
	// 期望的公钥
	wantP := new(rsa.PublicKey)
	wantP.N = big.NewInt(0)
	wantP.N.UnmarshalText([]byte("27184214226514158133300405299262353320292689684496368351344204098139332725880164777937651648325899107719408786796723369854875267353567309901948699160814049702206646065668152118198503452408544361562495115576674205119500961500401332875489305487337697931238137393355451032229158978665510808420232724088506155489843106213322494867991334200807786660769718152922304261316066609887484718846661796668623876670281516113590497056901119195916278334180533028261230601929397812932279788664594049350506128797319605281983312818031636647569117770724434665193888841497146780558439732128192283654221133610248799040021183184084526214689"))
	wantP.E = 65537
	tests := []struct {
		name    string
		args    args
		wantP   *rsa.PublicKey
		wantErr bool
	}{
		{"base", arg, wantP, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, err := getPublicKey(tt.args.publicKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPublicKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("getPublicKey() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}
