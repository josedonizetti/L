package l

import "testing"
import "reflect"

func TestParseSimpleAssignStm(t *testing.T) {
  ast := Parse("a := 1")
  checkType(ast, AssignStm{}, t)
  stm := ast.(AssignStm)
  checkType(stm.exp, NumExp{}, t)
  checkValue(stm.id, "a", t)
}

func TestParseAssignAnOpExp(t *testing.T) {
  ast := Parse("a := 1 + 1")
  checkType(ast, AssignStm{}, t)
  stm := ast.(AssignStm)
  checkType(stm.exp, OpExp{}, t)
  checkValue(stm.id, "a", t)
}

func TestParseAssignAnIdExp(t *testing.T) {
  ast := Parse("a := b")
  checkType(ast, AssignStm{}, t)
  stm := ast.(AssignStm)
  checkType(stm.exp, IdExp{}, t)
  checkValue(stm.id, "a", t)
}

func checkType(a, b interface{}, t *testing.T) {
 t1 := reflect.TypeOf(a)
 t2 := reflect.TypeOf(b)
 if t1 != t2 {
  t.Error("CheckType failed with types", t1, t2 )
 }
}

func checkValue(a, b interface{}, t *testing.T) {
  if a != b {
    t.Error("CheckValue failed with values \"" + a.(string) + "\" and \"" + b.(string) +"\"")
  }
}
