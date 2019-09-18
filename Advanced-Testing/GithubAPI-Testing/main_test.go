package main

import "testing"

type FakeReleaseInfoer struct {
	Tag string
	Err error
}

func (f FakeReleaseInfoer) GetLatestReleaseTag(repo string) (string, error) {

	if f.Err != nil {
		return "", f.Err
	}

	return f.Tag, nil
}

func TestGetReleaseTagMessage(t *testing.T) {
	f := FakeReleaseInfoer{
		Tag: "v1.0.1",
		Err: nil,
	}
	expectedMsg := "The latest release tag is v1.0.1"
	msg, err := getReleaseTagMessage(f, "dev/null")
	if err != nil {
		t.Fatalf("Expected Error is nil but it was : %s", err)
	}
	if expectedMsg != msg {
		t.Fatalf("Expected message  %s received %s ", expectedMsg, msg)
	}
}
