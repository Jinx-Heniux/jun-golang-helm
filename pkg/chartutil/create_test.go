package chartutil

import "testing"

func TestCreateFromTest2(t *testing.T) {
	srcdir := "/home/zhs2si/projects/go/jun-golang-helm/testdata/sourcecharts/boilerplate-based-on-utility-toolkit"

	destdir := "/home/zhs2si/projects/go/jun-golang-helm/testdata/outputs"

	if err := CreateFromTest2(nil, destdir, srcdir); err != nil {
		t.Fatal(err)
	}
}
