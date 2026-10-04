package cmd

import (
	"strings"
	"testing"
)

func TestDonePolecatActorIdentity(t *testing.T) {
	t.Parallel()

	rig, name, err := donePolecatActorIdentity("gastown/polecats/nux")
	if err != nil || rig != "gastown" || name != "nux" {
		t.Fatalf("polecat actor: got (%q, %q, %v), want (gastown, nux, nil)", rig, name, err)
	}

	for _, actor := range []string{"deacon/dogs/alpha", "dog"} {
		_, _, err := donePolecatActorIdentity(actor)
		if err == nil || !strings.Contains(err.Error(), "gt dog done") {
			t.Errorf("dog actor %q: got %v, want error pointing to gt dog done", actor, err)
		}
	}

	_, _, err = donePolecatActorIdentity("mayor/")
	if err == nil || strings.Contains(err.Error(), "gt dog done") {
		t.Errorf("mayor actor: got %v, want polecats-only error without dog hint", err)
	}
}
