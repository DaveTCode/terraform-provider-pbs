package pbsclient

import (
	"testing"
)

func TestPbsServerParsing(t *testing.T) {
	sourceText := `Server abcd00
    server_state = Active
    server_host = abcd00.hsn.cm.abcd.sc.test.com
    scheduling = True
    total_jobs = 0
    state_count = Transit:0 Queued:0 Held:0 Waiting:0 Running:0 Exiting:0 Begun:0
    acl_roots = root
    managers = a.person@*,another.chap@*,
        root@scheduler02.hsn.cm.abcd.sc.test.com,
        root@scheduler01.hsn.cm.abcd.sc.test.com,
        root@*.test.com,
        root@*
    operators = admin@*,one.another@*,another.chap@*`

	parsedOutput, err := parseServerOutput([]byte(sourceText))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	if len(parsedOutput) != 1 {
		t.Errorf("expected 1 output from parsing result but got %d", len(parsedOutput))
		return
	}
	if parsedOutput[0].Name != "abcd00" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].Name, "abcd00")
	}
	if *parsedOutput[0].AclRoots != "root" {
		t.Errorf("got %q, wanted %q", *parsedOutput[0].AclRoots, "root")
	}
	if *parsedOutput[0].Managers != "a.person@*,another.chap@*,root@scheduler02.hsn.cm.abcd.sc.test.com,root@scheduler01.hsn.cm.abcd.sc.test.com,root@*.test.com,root@*" {
		t.Errorf("got %q, wanted %q", *parsedOutput[0].Managers, "a.person@*,another.chap@*,root@scheduler02.hsn.cm.abcd.sc.test.com,root@scheduler01.hsn.cm.abcd.sc.test.com,root@*.test.com,root@*")
	}
	if *parsedOutput[0].Operators != "admin@*,one.another@*,another.chap@*" {
		t.Errorf("got %q, wanted %q", *parsedOutput[0].Operators, "admin@*,one.another@*,another.chap@*")
	}
}
