package l

import "strings"
import "regexp"
import "strconv"
//import "fmt"

var numbersRegex = regexp.MustCompile("([0-9])*")
var operationRegex = regexp.MustCompile("[+ \\- * /]")
var variableRegex = regexp.MustCompile("([a-z])")

func createExp(prog string) (exp Exp) {
  prog =  strings.Replace(prog, " ", "", -1)
  switch {
  case variableRegex.MatchString(prog):
    exp = IdExp{id: prog}
  case operationRegex.MatchString(prog):
    args := operationRegex.Split(prog, -1)
    right := createExp(args[0])
    left := createExp(args[1])
    exp = OpExp{right: right, left: left, operation: operationRegex.FindString(prog) }
  case numbersRegex.MatchString(prog):
    val, _ := strconv.Atoi(prog)
    exp = NumExp{val: val}
  }

  return
}

func isAssignStm(prog string) bool {
  return strings.Contains(prog, ":=")
}

func createAssignStm(prog string) AssignStm {
  ops := strings.Split(prog, ":=")
  exp := createExp(strings.Trim(ops[1], " "))
  return AssignStm{ id: strings.Trim(ops[0], " "), exp: exp }
}

func isPrintStm(prog string) bool {
  return strings.Contains(prog, "print")
}

func createPrintStm(prog string) PrintStm {
  return PrintStm{}
}

func isOpExp(prog string) bool {
  return false
}

func createOpExp(prog string) OpExp {
  return OpExp{}
}

func isIdExp(prog string) bool {
  return false
}

func createIdExp(prog string) IdExp {
  return IdExp{}
}

func createStm(prog string) (stm Stm) {
  switch {
  case isAssignStm(prog):
    stm = createAssignStm(prog)
  case isPrintStm(prog):
    stm = createPrintStm(prog)
  default:
    panic("createStm error! :P")
  }
  return
}

func Parse(prog string) (stm Stm) {
  elements := strings.Split(prog, ";")
  size := len(elements)
  switch {
  case size == 1:
    stm = createStm(elements[0])
  case size > 1:
  default:
    panic("Parse error! :P")
  }
  return
}
