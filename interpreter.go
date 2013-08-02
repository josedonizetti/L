package l

import "fmt"

// Environment
type Env map[string]interface{}

// Statements
type Stm interface {
  EvaluateStm(env Env)
}

type CompoundStm struct {
  stm1 Stm
  stm2 Stm
}

func (stm CompoundStm) EvaluateStm(env Env) {
  stm.stm1.EvaluateStm(env)
  stm.stm2.EvaluateStm(env)
}

type AssignStm struct {
  id string
  exp Exp
}

func (stm AssignStm) EvaluateStm(env Env) {
  value := stm.exp.EvaluateExp(env)
  env[stm.id] = value
}

type PrintStm struct {
  expList ExpList
}

func (stm PrintStm) EvaluateStm(env Env) {
  fmt.Print(stm.expList.EvaluateExpList(env))
}

// Expressions

type Exp interface {
  EvaluateExp(env Env) interface{}
}

type NumExp struct {
  val int
}

func (exp NumExp) EvaluateExp(env Env) interface{} {
  return exp.val
}

type OpExp struct {
  right, left Exp
  operation string
}

func (exp OpExp) EvaluateExp(env Env) interface{} {
  switch exp.operation {
  case "+":
    return exp.plus(env)
  case "*":
    return exp.times(env)
  case "-":
    return exp.minus(env)
  case "/":
    return exp.div(env)
  }
  return nil
}

func (exp OpExp) plus(env Env) int {
  v1 := exp.right.EvaluateExp(env).(int)
  v2 := exp.left.EvaluateExp(env).(int)
  return v1 + v2
}

func (exp OpExp) minus(env Env) int {
  v1 := exp.right.EvaluateExp(env).(int)
  v2 := exp.left.EvaluateExp(env).(int)
  return v1 - v2
}

func (exp OpExp) div(env Env) int {
  v1 := exp.right.EvaluateExp(env).(int)
  v2 := exp.left.EvaluateExp(env).(int)
  return v1 / v2
}

func (exp OpExp) times(env Env) int {
  v1 := exp.right.EvaluateExp(env).(int)
  v2 := exp.left.EvaluateExp(env).(int)
  return v1 * v2
}

type IdExp struct {
  id string
}

func (exp IdExp) EvaluateExp(env Env) interface{} {
  return env[exp.id]
}

type EseqExp struct {
  stm Stm
  exp Exp
}

func (exp EseqExp) EvaluateExp(env Env) interface{} {
  exp.stm.EvaluateStm(env)
  return exp.exp.EvaluateExp(env)
}

// Expression Lists

type ExpList interface {
  EvaluateExpList(env Env) interface{}
}

type LastExpList struct {
  head Exp
}

func (expList LastExpList) EvaluateExpList(env Env) interface{} {
  return expList.head.EvaluateExp(env)
}

type PairExpList struct {
  head Exp
  tail ExpList
}

func (expList PairExpList) EvaluateExpList(env Env) interface{} {
  return []interface{}{expList.head.EvaluateExp(env), expList.tail.EvaluateExpList(env) }
}

func interpret(env Env, stm Stm) {
  stm.EvaluateStm(env)
}

//func main() {
//  env := Env{}
//  prog := CompoundStm{ stm1: AssignStm{ id: "a", exp: OpExp{ right: NumExp{ val: 5}, left: NumExp{ val: 3 }, operation: "+" }},
//             stm2: CompoundStm{
//                stm1: AssignStm{id: "b", exp: EseqExp{ stm: PrintStm{ expList: PairExpList{ head: IdExp{id: "a"}, tail: LastExpList{ head: OpExp{ right: IdExp{id: "a"}, left: NumExp{val: 1}, operation: "-" }}}}, exp: OpExp{ right: NumExp{val: 10}, left: IdExp{id:"a"}, operation: "*"}}},
//                stm2: PrintStm{ expList: LastExpList{ head: IdExp{ id: "b"}}}}}
//  interpret(env, prog)
//  fmt.Println()
//}


// a := 5 + 3 ; b := ( print ( a , a - 1 ) , 10 * a ) ; print ( b )
// a := 5 + 3 - prog := AssignStm{id: "a", exp: OpExp{right: NumExp{num: 5}, left: NumExp{num: 5}, operation: "+"}}
// b := a + 3 - prog := AssignStm{id: "b", exp: OpExp{right: IdExp{id: "a"}, left: NumExp{num: 3}, operation: "+"}}
// print(b) - prog2 := PrintStm{expList: LastExpList{head: IdExp{id: "b"}}}
// print(a, a + 3) - prog2 := PrintStm{expList: PairExpList{head: IdExp{id: "a"}, tail: LastExpList{ head: OpExp{ right: IdExp{id: "a"}, left: NumExp{num: 3}, operation: "+" }}}}
// prog2 := AssignStm{id: "b", exp: EseqExp{ stm: PrintStm{ expList: PairExpList{ head: IdExp{id: "a"}, tail: LastExpList{ head: OpExp{ right: IdExp{id: "a"}, left: NumExp{num: 1}, operation: "+" }}}}, exp: OpExp{ right: NumExp{num: 10}, left: IdExp{id:"a"}, operation: "+"}}}
