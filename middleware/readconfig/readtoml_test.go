package readconfig

import "testing"

func TestReadConf(t *testing.T) {
	// /root/.lotusstorage
	p, err := ReadConfToml("../worker_task_config.toml")
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("Person %s", p.MonitorUnit.LocalIP)
}
