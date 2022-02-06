package persons
//import 	"fmt"
/*
   subject: class test. use newFunc for superclass.
*/

type Object struct {
	classname string
}
func (o *Object) SetClassname(name string) {
	o.classname = name
}
func (o *Object) GetClassname() string {
	return o.classname
}
func NewObject() Object {
	//var obj = Object{"Object"}
	var obj Object
	obj.classname = "Object"
	return obj
}

type Human struct {
	Temper string
	Object
}
func (hu *Human) Ban() string {
	hu.Temper = "angry"
	return "nandayou"
}
func NewHuman() Human {
	//var hu Human = Human{ "calm", NewObject() }
	var hu Human
	hu.Temper = "easy"
	hu.Object = NewObject()
	hu.SetClassname("Human")
	return hu
}
		
type Person struct {
	Name string
	Human
}
func NewPerson() Person {
	//var ps Person = Person{ "noname", NewHuman() }
	var ps Person // = Person{}
	ps.Name = "Anon"
	ps.Human = NewHuman()
	ps.SetClassname("Person")
	return ps
}

func PConcat(p1, p2 Person) Person {
	np := NewPerson()
	np.Name = p1.Name + p2.Name
	np.Temper = p1.Temper + p2.Temper
	np.SetClassname(p1.classname)
	return np
}
func NConcat(p1, p2 *Person) *Person {
	np := new(Person)
	np.Name = p1.Name + p2.Name
	np.Temper = p1.Temper + p2.Temper
	np.SetClassname(p1.classname)
	return np
	//return &np
}
