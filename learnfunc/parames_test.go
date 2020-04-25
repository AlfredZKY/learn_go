package learnfunc


import "testing"

func TestFuncParams(t*testing.T){
	array1 := [3]string{"a","b","c"}
	t.Logf("The array:%v\n",array1)
	array2 := modifyArray(array1)
	t.Logf("The array:%v\n",array2)
	t.Logf("The array:%v\n",array1)

	slice1 := []string{"x","y","z"}
	t.Logf("The slice:%v\n",slice1)
	slice2 := modifySlice(slice1)
	t.Logf("The slice:%v\n",slice2)
	t.Logf("The slice:%v\n",slice1)

	complexArray1 := [3][]string{
		[]string{"d","e","f"},
		[]string{"g","h","i"},
		[]string{"j","k","l"},
	}

	t.Logf("The complex array:%v\n",complexArray1)
	complexArray2 := modifyComplexArray(complexArray1)
	t.Logf("The modify complex array:%v\n",complexArray2)
	t.Logf("The original complex array:%v\n",complexArray1)
}