package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"http"
)

func TestPush(t *testing.T) {
	s := &Service{driver: mgo};
	assert.Nil(t, service.add("uid123", 987));
	assert.Equal(t, service.findUid(987, "uid123"));
	assert.Equal(t, service.findA2Id("uid123", "uid123"));
}

func TestGet(t *testing.T) {
	resp := httptest.NewRecorder()
}
