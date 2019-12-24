package mystruct

import (
	"fmt"
	"testing"
	"unsafe"
)

// AnimalCategory 代表动物分类学中的基本分类法。
type AnimalCategory struct {
	Kingdom string // 界。
	Phylum  string // 门。
	Class   string // 纲。
	Order   string // 目。
	Family  string // 科。
	Genus   string // 属。
	Species string // 种。
}

// Animal 再次对动物进行分类
type Animal struct {
	ScientificName string //学名
	AnimalCategory        //动物基本分类 嵌入字段，也叫匿名字段
}

// Cat 实例化一个对象
type Cat struct {
	Name string
	Animal
}

// func String
func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s", ac.Kingdom, ac.Phylum,
		ac.Class, ac.Order, ac.Family, ac.Genus, ac.Species)
}

// func String
func (a Animal) String() string {
	return fmt.Sprintf("%s (category:%s)", a.ScientificName, a.AnimalCategory)
}

// func String
func (cat Cat) String() string {
	return fmt.Sprintf("%s (category:%s,name:%q)", cat.ScientificName, cat.AnimalCategory, cat.Name)
}

// CatAnimal define a struct for catanimal
type CatAnimal struct {
	Name           string // 名字
	ScientificName string //学名
	Category       string // 动物学基本分类
}

// New new a CatAnimal object
func New(name, scientificname, category string) CatAnimal {
	return CatAnimal{
		Name:           name,
		ScientificName: scientificname,
		Category:       category,
	}
}

//SetName  assignment variable of Name
func (cat *CatAnimal) SetName(name string) {
	cat.Name = name

}

// SetNameOfCopy copy variable
func (cat CatAnimal) SetNameOfCopy(name string) {
	cat.Name = name
}

// func (cat CatAnimal)String(){
// 	fmt.Printf("address:is %x\n",unsafe.Pointer(&cat.Name))
// }

func (cat *CatAnimal) String() {
	fmt.Printf("address:is %x\n", unsafe.Pointer(&cat.Name))
}

func TestStructOperations(t *testing.T) {
	cat := New("little pig", "dog", "testing")
	fmt.Printf("address:is %x\n", unsafe.Pointer(&cat.Name))

	// cat cat* 打印出地址的结果可见最好采用cat*方式封装
	cat.String()
}

// GetName get the value of the variable
func (cat CatAnimal) GetName() string {
	return cat.Name
}

// GetScientificName get the value of the variable
func (cat CatAnimal) GetScientificName() string {
	return cat.ScientificName
}

// GetCategory get the value of the variable
func (cat CatAnimal) GetCategory() string {
	return cat.Category
}

func TestStructNull(t *testing.T) {
	// 空结构体的特点：1.不占用内存 2。地址不变
	var s struct{}
	var s1 struct{}
	t.Logf("Memory occupied by empty structures is %d", unsafe.Sizeof(s))
	t.Logf("s address is %p s1 address is %p, s or s1 compare %v", &s, &s1, &s == &s1)
}
