package pbsclient

import "testing"

func TestParseQmgrOutputMultipleResources(t *testing.T) {
	sourceText := `Queue workq
    queue_type = Execution
    total_jobs = 0
    state_count = Transit:0 Queued:0 Held:0 Waiting:0 Running:0 Exiting:0 Begun:0
    enabled = True
    started = True

Queue test
    queue_type = Execution
    total_jobs = 0
    state_count = Transit:0 Queued:0 Held:0 Waiting:0 Running:0 Exiting:0 Begun:0
    resources_default.ncpus = 1
    resources_default.nodect = 1
    resources_default.nodes = 1
    resources_default.walltime = 01:00:00
    enabled = True
    started = True`
	parsedOutput := parseGenericQmgrOutput(sourceText)

	if len(parsedOutput) != 2 {
		t.Errorf("expected 2 output from parsing result but got %d", len(parsedOutput))
		return
	}
	if parsedOutput[0].objType != "Queue" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].objType, "Queue")
	}
	if parsedOutput[0].name != "workq" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].name, "workq")
	}
	if len(parsedOutput[0].attributes) != 5 {
		t.Errorf("expected 5 attributes but got %d", len(parsedOutput[0].attributes))
	}
	if parsedOutput[0].attributes["queue_type"] != "Execution" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["queue_type"], "Execution")
	}
	if parsedOutput[0].attributes["total_jobs"] != "0" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["total_jobs"], "0")
	}
	if parsedOutput[0].attributes["state_count"] != "Transit:0 Queued:0 Held:0 Waiting:0 Running:0 Exiting:0 Begun:0" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["state_count"], "Transit:0 Queued:0 Held:0 Waiting:0 Running:0 Exiting:0 Begun:0")
	}
	if parsedOutput[0].attributes["enabled"] != "True" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["enabled"], "True")
	}
	if parsedOutput[0].attributes["started"] != "True" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["started"], "True")
	}

	if parsedOutput[1].objType != "Queue" {
		t.Errorf("got %q, wanted %q", parsedOutput[1].objType, "Queue")
	}
	if parsedOutput[1].name != "test" {
		t.Errorf("got %q, wanted %q", parsedOutput[1].name, "test")
	}
	if len(parsedOutput[1].attributes) != 6 {
		t.Errorf("expected 6 attributes but got %d", len(parsedOutput[1].attributes))
	}
	if parsedOutput[1].attributes["queue_type"] != "Execution" {
		t.Errorf("got %q, wanted %q", parsedOutput[1].attributes["queue_type"], "Execution")
	}
	if parsedOutput[1].attributes["total_jobs"] != "0" {
		t.Errorf("got %q, wanted %q", parsedOutput[1].attributes["total_jobs"], "0")
	}
	if parsedOutput[1].attributes["state_count"] != "Transit:0 Queued:0 Held:0 Waiting:0 Running:0 Exiting:0 Begun:0" {
		t.Errorf("got %q, wanted %q", parsedOutput[1].attributes, "Transit:0 Queued:0 Held:0 Waiting:0 Running:0 Exiting:0 Begun:0")
	}
	if parsedOutput[1].attributes["resources_default"].(map[string]string)["ncpus"] != "1" {
		t.Errorf("got %q, wanted %q", parsedOutput[1].attributes["resources_default"], "ncpus=1")
	}
	if parsedOutput[1].attributes["resources_default"].(map[string]string)["nodect"] != "1" {
		t.Errorf("got %q, wanted %q", parsedOutput[1].attributes["resources_default"], "nodect=1")
	}
	if parsedOutput[1].attributes["resources_default"].(map[string]string)["nodes"] != "1" {
		t.Errorf("got %q, wanted %q", parsedOutput[1].attributes["resources_default"], "nodes=1")
	}
	if parsedOutput[1].attributes["resources_default"].(map[string]string)["walltime"] != "01:00:00" {
		t.Errorf("got %q, wanted %q", parsedOutput[1].attributes["resources_default"], "walltime=01:00:00")
	}
	if parsedOutput[1].attributes["enabled"] != "True" {
		t.Errorf("got %q, wanted %q", parsedOutput[1].attributes["enabled"], "True")
	}
	if parsedOutput[1].attributes["started"] != "True" {
		t.Errorf("got %q, wanted %q", parsedOutput[1].attributes["started"], "True")
	}
}

func TestParseQmgrResourcesAvailableOutput(t *testing.T) {
	sourceText := `Node pbs
    Mom = pbs
    Port = 15002
    pbs_version = unavailable
    ntype = PBS
    state = state-unknown,down
    resources_available.host = pbs
    resources_available.vnode = pbs
    resources_assigned.accelerator_memory = 0kb
    resources_assigned.hbmem = 0kb
    resources_assigned.mem = 0kb
    resources_assigned.naccelerators = 0
    resources_assigned.ncpus = 0
    resources_assigned.vmem = 0kb
    resv_enable = True
    sharing = default_shared`
	parsedOutput := parseGenericQmgrOutput(sourceText)

	if len(parsedOutput) != 1 {
		t.Errorf("expected 1 output from parsing result but got %d", len(parsedOutput))
		return
	}
	if parsedOutput[0].objType != "Node" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].objType, "Node")
	}
	if parsedOutput[0].name != "pbs" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].name, "pbs")
	}
	if len(parsedOutput[0].attributes) != 9 {
		t.Errorf("expected 9 attributes but got %d", len(parsedOutput[0].attributes))
	}
	if parsedOutput[0].attributes["Mom"] != "pbs" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["Mom"], "pbs")
	}
	if parsedOutput[0].attributes["Port"] != "15002" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["Port"], "15002")
	}
	if parsedOutput[0].attributes["pbs_version"] != "unavailable" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["pbs_version"], "unavailable")
	}
	if parsedOutput[0].attributes["ntype"] != "PBS" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["ntype"], "PBS")
	}
	if parsedOutput[0].attributes["state"] != "state-unknown,down" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["state"], "state-unknown,down")
	}
	if parsedOutput[0].attributes["resources_available"].(map[string]string)["host"] != "pbs" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["resources_available"], "host=pbs")
	}
	if parsedOutput[0].attributes["resources_available"].(map[string]string)["vnode"] != "pbs" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["resources_available"], "vnode=pbs")
	}
	if parsedOutput[0].attributes["resources_assigned"].(map[string]string)["accelerator_memory"] != "0kb" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["resources_assigned"], "accelerator_memory=0kb")
	}
	if parsedOutput[0].attributes["resources_assigned"].(map[string]string)["hbmem"] != "0kb" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["resources_assigned"], "hbmem=0kb")
	}
	if parsedOutput[0].attributes["resources_assigned"].(map[string]string)["mem"] != "0kb" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["resources_assigned"], "mem=0kb")
	}
	if parsedOutput[0].attributes["resources_assigned"].(map[string]string)["naccelerators"] != "0" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["resources_assigned"], "naccelerators=0")
	}
	if parsedOutput[0].attributes["resources_assigned"].(map[string]string)["ncpus"] != "0" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["resources_assigned"], "ncpus=0")
	}
	if parsedOutput[0].attributes["resources_assigned"].(map[string]string)["vmem"] != "0kb" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["resources_assigned"], "vmem=0kb")
	}
	if parsedOutput[0].attributes["resv_enable"] != "True" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["resv_enable"], "True")
	}
	if parsedOutput[0].attributes["sharing"] != "default_shared" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["sharing"], "default_shared")
	}
}

func TestParseQmgrMultiLineResource(t *testing.T) {
	sourceTest := `Hook pbs_cgroups
    type = site
    enabled = false
    event = execjob_begin,execjob_epilogue,execjob_end,execjob_launch,
        execjob_attach,
        execjob_resize,
        execjob_abort,
        execjob_postsuspend,
        execjob_preresume,
        exechost_periodic,
        exechost_startup
    user = pbsadmin
    alarm = 90
    freq = 120
    order = 100
    debug = false
    fail_action = offline_vnodes`
	parsedOutput := parseGenericQmgrOutput(sourceTest)

	if len(parsedOutput) != 1 {
		t.Errorf("expected 1 output from parsing result but got %d", len(parsedOutput))
		return
	}
	if parsedOutput[0].objType != "Hook" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].objType, "Hook")
	}
	if parsedOutput[0].name != "pbs_cgroups" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].name, "pbs_cgroups")
	}
	if len(parsedOutput[0].attributes) != 9 {
		t.Errorf("expected 9 attributes but got %d", len(parsedOutput[0].attributes))
	}
	if parsedOutput[0].attributes["type"] != "site" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["type"], "site")
	}
	if parsedOutput[0].attributes["enabled"] != "false" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["enabled"], "false")
	}
	if parsedOutput[0].attributes["event"] != "execjob_begin,execjob_epilogue,execjob_end,execjob_launch,execjob_attach,execjob_resize,execjob_abort,execjob_postsuspend,execjob_preresume,exechost_periodic,exechost_startup" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["event"], "execjob_begin,execjob_epilogue,execjob_end,execjob_launch,execjob_attach,execjob_resize,execjob_abort,execjob_postsuspend,execjob_preresume,exechost_periodic,exechost_startup")
	}
	if parsedOutput[0].attributes["user"] != "pbsadmin" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["user"], "pbsadmin")
	}
	if parsedOutput[0].attributes["alarm"] != "90" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["alarm"], "90")
	}
	if parsedOutput[0].attributes["freq"] != "120" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["freq"], "120")
	}
	if parsedOutput[0].attributes["order"] != "100" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["order"], "100")
	}
	if parsedOutput[0].attributes["debug"] != "false" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["debug"], "false")
	}
	if parsedOutput[0].attributes["fail_action"] != "offline_vnodes" {
		t.Errorf("got %q, wanted %q", parsedOutput[0].attributes["fail_action"], "offline_vnodes")
	}

}

var pinnedi32 = int32(1)
var pinnedstr = "a"
var pinnedbool = true
var updateTests = []struct {
	old      any
	new      any
	expected []string
}{
	{true, true, []string{}},
	{int32(1), int32(1), []string{}},
	{"a", "a", []string{}},
	{int32(1), int32(2), []string{"/opt/pbs/bin/qmgr -c 'set queue workq test=2'"}},
	{&pinnedi32, (*int32)(nil), []string{"/opt/pbs/bin/qmgr -c 'unset queue workq test'"}},
	{(*int32)(nil), &pinnedi32, []string{"/opt/pbs/bin/qmgr -c 'set queue workq test=1'"}},
	{"a", "b", []string{"/opt/pbs/bin/qmgr -c 'set queue workq test=\"b\"'"}},
	{&pinnedstr, (*string)(nil), []string{"/opt/pbs/bin/qmgr -c 'unset queue workq test'"}},
	{(*string)(nil), &pinnedstr, []string{"/opt/pbs/bin/qmgr -c 'set queue workq test=\"a\"'"}},
	{true, false, []string{"/opt/pbs/bin/qmgr -c 'set queue workq test=false'"}},
	{&pinnedbool, (*bool)(nil), []string{"/opt/pbs/bin/qmgr -c 'unset queue workq test'"}},
	{(*bool)(nil), &pinnedbool, []string{"/opt/pbs/bin/qmgr -c 'set queue workq test=true'"}},
}

func TestGenerateUpdateAttributeCommand(t *testing.T) {
	for _, tt := range updateTests {
		commands, err := generateUpdateAttributeCommand(tt.old, tt.new, "queue", "workq", "test")

		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
		if len(commands) != len(tt.expected) {
			t.Errorf("expected %d command but got %d", len(tt.expected), len(commands))
			return
		}

		for i, command := range commands {
			if command != tt.expected[i] {
				t.Errorf("got %q, wanted %q", command, tt.expected[i])
			}
		}
	}
}
