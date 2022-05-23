package main_test

import (
	"strings"
	"testing"

	"github.com/asstart/hello-world-publish-go-mod"
)

func TestHello(t *testing.T) {
	res, err := main.HelloWorld()
	if err != nil {
		t.Fatalf("error while greeting: %v", err.Error())
	}
	if !strings.HasPrefix(res, "Hello") {
		t.Fatalf("greeting doesn't start from <Hello>: %v", res)
	}
}
