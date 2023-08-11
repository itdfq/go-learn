package builder

import "fmt"

//建造者模式， 假设业务需要建造一个房子对象，需要先打地基、建墙、建屋顶、建花园、放置家具……。我们需要非常多的步骤，并且这些步骤之间是有联系的，即使将各个步骤从一个大的构造函数抽出到其他小函数中，整个程序的层次结构看起来依然很复杂。
//如何解决呢？像这种复杂的有许多步骤的构造函数，就可以用建造者模式来设计。
//建造者模式的用处就在于能够分步骤创建复杂对象。

//建造者接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

//管理类
type Director struct {
	builder Builder
}

//构造函数
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

//建造
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

type BuilderA struct{}

func (a *BuilderA) Part1() {
	fmt.Println("part1")
}

func (a *BuilderA) Part2() {
	fmt.Println("part2")
}

func (a *BuilderA) Part3() {
	fmt.Println("part3")
}
