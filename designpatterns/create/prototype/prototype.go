package prototype

//原型模式

/**
如果你希望生成一个对象，其与另一个对象完全相同，该如何实现呢？
如果遍历对象的所有成员，将其依次复制到新对象中，会稍显麻烦，而且有些对象可能会有私有成员变量遗漏。
原型模式将这个克隆的过程委派给了被克隆的实际对象，被克隆的对象就叫做“原型”。
*/

type Type1 struct {
	Name string
}

func (t *Type1) Clone() *Type1 {
	tc := *t
	return &tc
}
