package expr

type Env map[string]float64

// 浮点数运算，支持二元操作符+，-，*， 和/；一元操作符-x和+x；调用pow(x,y)，sin(x)，和sqrt(x)的函数
// 支持括号和优先级
type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) (float64, error)
}

type Var string

// A literal is a numeric constant, e.g., 3.141.
type literal float64

// A unary represents a unary operator expression, e.g., -x.
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

// A binary represents a binary operator expression, e.g., x+y.
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

// A call represents a function call expression, e.g., sin(x).
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

//func (v Var) Eval(env Env) (float64, error) {
//	if val, ok := env[string(v)]; ok {
//		return val, nil
//	} else {
//		return 0, fmt.Errorf("undefined variable %q", v)
//	}
//}
//
//func (l literal) Eval(_ Env) (float64, error) {
//	return float64(l), nil
//}
//
//func (u unary) Eval(env Env) (float64, error) {
//	val, err := u.x.Eval(env)
//	if err != nil {
//		return 0, err
//	}
//	switch u.op {
//	case '+':
//		return +val, nil
//	case '-':
//		return -val, nil
//	}
//	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
//}
//
//func (b binary) Eval(env Env) (float64, error) {
//	switch b.op {
//	case '+':
//		return b.x.Eval(env) + b.y.Eval(env)
//	case '-':
//		return b.x.Eval(env) - b.y.Eval(env)
//	case '*':
//		return b.x.Eval(env) * b.y.Eval(env)
//	case '/':
//		return b.x.Eval(env) / b.y.Eval(env)
//	}
//	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
//}
//
//func (c call) Eval(env Env) (float64, error) {
//	switch c.fn {
//	case "pow":
//		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
//	case "sin":
//		return math.Sin(c.args[0].Eval(env))
//	case "sqrt":
//		return math.Sqrt(c.args[0].Eval(env))
//	}
//	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
//}
