package main

import (
	"log"
	"reflect"
	"testing"

	_ "github.com/bigzhu/gobz/apibz"
	_ "github.com/bigzhu/gobz/confbz"
	_ "github.com/bigzhu/gobz/functionbz"
	_ "github.com/bigzhu/gobz/ginbz"
	_ "github.com/bigzhu/gobz/httpbz"
	_ "github.com/bigzhu/gobz/jsonbz"
	_ "github.com/bigzhu/gobz/modelsbz"
	_ "github.com/bigzhu/gobz/rsabz"
	_ "github.com/bigzhu/gobz/services/callbackbz"
	_ "github.com/bigzhu/gobz/services/mailbz"
	_ "github.com/bigzhu/gobz/services/oauthbz"
	_ "github.com/bigzhu/gobz/services/smsbz"
	_ "github.com/bigzhu/gobz/telegram"
	_ "github.com/bigzhu/gobz/websocketbz"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/ugorji/go/codec"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
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
