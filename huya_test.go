package tv_test

import (
	"testing"

	"github.com/go-olive/tv"
)

func TestHuya_Snap(t *testing.T) {
	u := "https://www.huya.com/520588"
	huya, err := tv.NewWithUrl(u)
	if err != nil {
		println(err.Error())
		return
	}
	huya.Snap()
	t.Log(huya)
}
