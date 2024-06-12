package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathName := CASPathTransformFunc(key)
	expectedOriginalKey := "6804429f74181a63c50c3d81d733a12f14a353ff"
	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	if pathName.Pathname != expectedPathName {
		t.Errorf("HAVE %s want %s", pathName, expectedPathName)
	}

	if pathName.Filename != expectedOriginalKey {
		t.Errorf("HAVE %s want %s", pathName, expectedPathName)
	}

}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)
	key := "momsspecial"
	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, err := ioutil.ReadAll(r)
	fmt.Print(string(b))
	if string(b) != string(data) {
		t.Errorf("Want %s have %s", data, b)
	}

}
