package l

import "strings"
import "strconv"

func isAssignStm(prog string) bool {
  return strings.Contains(prog, ":=")
}

func createAssignStm(prog string) AssignStm {
  ops := strings.Split(prog, ":=")
  val, _ := strconv.Atoi(strings.Trim(ops[1], " "))
  return AssignStm{ id: strings.Trim(ops[0], " "), exp: NumExp{ val: val } }
}

func isPrintStm(prog string) bool {
  return strings.Contains(prog, "print")
}

func createPrintStm(prog string) PrintStm {
  return PrintStm{}
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
