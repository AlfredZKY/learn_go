package readconfig

import "testing"

func TestReadConfJSON(t *testing.T) {
	var path string ="../conf.json"
	res := ReadConfJSON(path)
	t.Log(res.Enabled,res.Path)
}
