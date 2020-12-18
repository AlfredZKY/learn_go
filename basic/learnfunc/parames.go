package learnfunc


func modifyArray(a[3]string)[3]string{
	a[1] = "x"
	return a
}

func modifySlice(a[]string)[]string{
	a[1] = "i"
	return a 
}

func modifyComplexArray(a[3][]string)[3][]string{
	a[1][1] = "s"
	a[2] = []string{"o","p","q"}
	return a
}

