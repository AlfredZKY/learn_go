package readconfig

import "testing"

func TestReadConfIni(t *testing.T) {
	path := "../conf.ini"
	res := ReadConfIni(path)
	t.Log(res.Enabled,res.Path,res.Section.Enabled,res.Section.Path)
}
