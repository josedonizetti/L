package l

import "testing"
import "reflect"

func TestParseSimpleAssignStm(t *testing.T) {
  ast := Parse("a := 1")
  checkType(ast, AssignStm{}, t)
  stm := ast.(AssignStm)
  checkValue(stm.id, "a", t)
  checkType(stm.exp, NumExp{}, t)
  numExp := stm.exp.(NumExp)
  checkValue(numExp.val, 1, t)
}

func TestParseAssignAnOpExp(t *testing.T) {
  ast := Parse("a := 1 + 1")
  checkType(ast, AssignStm{}, t)
  stm := ast.(AssignStm)
  checkValue(stm.id, "a", t)
  checkType(stm.exp, OpExp{}, t)
  opExp := stm.exp.(OpExp)
  rightExp := opExp.right.(NumExp)
  leftExp := opExp.right.(NumExp)

  checkValue(opExp.operation, "+", t)
  checkValue(rightExp.val, 1, t)
  checkValue(leftExp.val, 1, t)
}

func TestParseAssignAnIdExp(t *testing.T) {
  ast := Parse("a := b")
  checkType(ast, AssignStm{}, t)
  stm := ast.(AssignStm)
  checkValue(stm.id, "a", t)
  checkType(stm.exp, IdExp{}, t)
  idExp := stm.exp.(IdExp)
  checkValue(idExp.id, "b", t)
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
