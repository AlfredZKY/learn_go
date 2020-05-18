package readconfig

import "testing"

func TestReadConfYaml(t *testing.T) {
	var c ConfigFileYaml
	path := "../conf.yaml"
	res := c.ReadConfYaml(path)
	t.Log(res.Enabled,res.Path)
}
