package modelsbz

import (
	"testing"
	"time"
)

// TxLog 账单安全审计记录
type TxLog struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	TransferID  int    `gorm:"not null;index"` // 账单id
	AdminRole   string `gorm:"not null"`       // 操作员角色
	AdminID     int    `gorm:"not null"`       // 操作员ID 0为系统操作
	Action      string `gorm:"not null"`       // 行为 延迟 DELAYED 放行 RELEASE 驳回 RETRIEVED
	Description string // 说明
}

func TestSetUnlogged(t *testing.T) {
	Connect()
	SetUnlogged(TxLog{}, TxLog{})
}

func Test_getTableName(t *testing.T) {
	Connect()
	type args struct {
		model interface{}
	}
	var v interface{}
	v = TxLog{}
	arg := args{v}
	tests := []struct {
		name          string
		args          args
		wantTableName string
	}{
		{"base", arg, "tx_logs"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTableName := getTableName(tt.args.model); gotTableName != tt.wantTableName {
				t.Errorf("getTableName() = %v, want %v", gotTableName, tt.wantTableName)
			}
		})
	}
}
