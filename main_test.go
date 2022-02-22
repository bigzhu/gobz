package main

import (
	"log"
	"reflect"
	"testing"

	_ "gobz/apibz"
	_ "gobz/confbz"
	_ "gobz/functionbz"
	_ "gobz/ginbz"
	_ "gobz/httpbz"
	_ "gobz/jsonbz"
	_ "gobz/modelsbz"
	_ "gobz/rsabz"
	_ "gobz/services/callbackbz"
	_ "gobz/services/mailbz"
	_ "gobz/services/oauthbz"
	_ "gobz/services/smsbz"
	_ "gobz/telegram"
	_ "gobz/websocketbz"
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/postgres"
	_ "github.com/ugorji/go/codec"
)

func init() {
	// test log 时候显示文件全路径
	log.SetFlags(log.Llongfile | log.LstdFlags)
}
func TestSetupPlay(t *testing.T) {
	tests := []struct {
		name string
		want *gin.Engine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetupPlay(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetupPlay() = %v, want %v", got, tt.want)
			}
		})
	}
}
