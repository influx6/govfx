package reflection_test

import (
	"fmt"
	"testing"

	"github.com/ardanlabs/kit/tests"
	"github.com/influx6/faux/reflection"
)

func init() {
	tests.Init("")
}

// mosnter provides a basic struct test case type.
type monster struct {
	Name string
}

// TestGetArgumentsType validates reflection API GetArgumentsType functions
// results.
func TestGetArgumentsType(t *testing.T) {
	tests.ResetLog()
	defer tests.DisplayLog()

	f := func(m monster) string {
		return fmt.Sprintf("Monster[%s] is ready!", m.Name)
	}

	args, err := reflection.GetFuncArgumentsType(f)
	if err != nil {
		t.Fatalf("\t%s\tShould be able to retrieve function arguments lists: %s", tests.Failed, err)
	} else {
		t.Logf("\t%s\tShould be able to retrieve function arguments lists", tests.Success)
	}

	newVals := reflection.MakeArgumentsValues(args)
	t.Logf("%+s", newVals)

	if nlen, alen := len(newVals), len(args); nlen != alen {
		t.Fatalf("\t%s\tShould have matching new values lists for arguments: %d %d", tests.Failed, nlen, alen)
	} else {
		t.Logf("\t%s\tShould have matching new values lists for arguments: %d %d", tests.Success, nlen, alen)
	}

}
