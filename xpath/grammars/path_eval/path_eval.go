// Code generated by goyacc -o path_eval.go -p pathEval path_eval.y. DO NOT EDIT.

//line path_eval.y:27

package path_eval

import __yyfmt__ "fmt"

//line path_eval.y:28

import (
	"encoding/xml"

	"github.com/steiler/yang-parser/xpath"
	"github.com/steiler/yang-parser/xpath/xutils"
)

//line path_eval.y:39
type pathEvalSymType struct {
	yys     int
	sym     *xpath.Symbol /* Symbol table entry */
	val     float64       /* Numeric value */
	name    string        /* NodeType or AxisName */
	xmlname xml.Name      /* For NameTest */
}

const NUM = 57346
const DOTDOT = 57347
const DBLSLASH = 57348
const DBLCOLON = 57349
const ERR = 57350
const FUNC = 57351
const NODETYPE = 57352
const AXISNAME = 57353
const LITERAL = 57354
const NAMETEST = 57355
const OR = 57356
const AND = 57357
const NE = 57358
const EQ = 57359
const GT = 57360
const GE = 57361
const LT = 57362
const LE = 57363
const DIV = 57364
const MOD = 57365
const UNARYMINUS = 57366

var pathEvalToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUM",
	"DOTDOT",
	"DBLSLASH",
	"DBLCOLON",
	"ERR",
	"FUNC",
	"NODETYPE",
	"AXISNAME",
	"LITERAL",
	"NAMETEST",
	"OR",
	"AND",
	"NE",
	"EQ",
	"GT",
	"GE",
	"LT",
	"LE",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"DIV",
	"MOD",
	"UNARYMINUS",
	"'|'",
	"'('",
	"')'",
	"','",
	"'['",
	"']'",
	"'.'",
	"'@'",
}

var pathEvalStatenames = [...]string{}

const pathEvalEofCode = 1
const pathEvalErrCode = 2
const pathEvalInitialStackSize = 16

//line path_eval.y:256

//line yacctab:1
var pathEvalExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 14,
	6, 30,
	25, 30,
	-2, 27,
}

const pathEvalPrivate = 57344

const pathEvalLast = 168

var pathEvalAct = [...]int8{
	2, 67, 19, 32, 66, 12, 8, 6, 4, 9,
	5, 96, 100, 101, 16, 57, 55, 97, 98, 59,
	61, 54, 103, 37, 63, 90, 64, 53, 41, 33,
	7, 35, 50, 35, 51, 52, 62, 40, 29, 38,
	95, 48, 49, 45, 47, 44, 46, 68, 69, 70,
	72, 73, 71, 36, 39, 78, 79, 85, 60, 83,
	80, 81, 82, 88, 89, 92, 61, 65, 94, 84,
	93, 56, 61, 86, 87, 74, 75, 76, 77, 25,
	37, 38, 38, 34, 26, 27, 33, 24, 35, 30,
	61, 61, 28, 43, 42, 94, 20, 22, 11, 99,
	31, 58, 102, 21, 17, 23, 91, 25, 37, 38,
	36, 39, 26, 27, 33, 24, 35, 18, 15, 14,
	13, 10, 3, 1, 0, 0, 11, 0, 31, 0,
	0, 0, 0, 23, 0, 25, 37, 38, 36, 39,
	26, 27, 33, 24, 35, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 31, 0, 0, 0,
	0, 23, 0, 0, 0, 0, 36, 39,
}

var pathEvalPact = [...]int16{
	103, -1000, -1000, 23, 13, 77, 25, 19, 8, -1000,
	-2, 103, -1000, -1000, -18, 76, 33, -1000, -1000, -1000,
	-1000, 18, -1000, 103, -1000, -1000, -4, -1000, 20, -18,
	-1000, -1000, 18, 41, -1000, -1000, -1000, -1000, -1000, -1000,
	103, 103, 103, 103, 103, 103, 103, 103, 103, 103,
	103, 103, 103, 131, -1000, -1000, 103, -1000, 18, 18,
	18, 18, 33, -6, 75, -18, -18, -1000, 33, -1000,
	13, 77, 25, 25, 19, 19, 19, 19, 8, 8,
	-1000, -1000, -1000, -1000, -23, -1000, 33, 33, -1000, -1000,
	-1000, -1000, -14, -18, -1000, -1000, -1000, -1000, 103, -19,
	-1000, 103, -9, -1000,
}

var pathEvalPgo = [...]int8{
	0, 123, 0, 122, 8, 10, 7, 30, 6, 9,
	121, 5, 120, 119, 118, 14, 3, 117, 1, 104,
	103, 97, 2, 96, 92, 38, 4, 89, 83, 71,
	69, 40,
}

var pathEvalR1 = [...]int8{
	0, 1, 2, 3, 3, 4, 4, 5, 5, 5,
	6, 6, 6, 6, 6, 7, 7, 7, 8, 8,
	8, 8, 9, 9, 10, 10, 11, 11, 11, 11,
	14, 13, 13, 17, 17, 17, 17, 17, 17, 17,
	17, 12, 12, 19, 19, 19, 20, 15, 15, 15,
	22, 22, 22, 22, 22, 24, 24, 25, 26, 26,
	18, 29, 30, 31, 21, 23, 27, 27, 28, 16,
}

var pathEvalR2 = [...]int8{
	0, 1, 1, 1, 3, 1, 3, 1, 3, 3,
	1, 3, 3, 3, 3, 1, 3, 3, 1, 3,
	3, 3, 1, 2, 1, 3, 1, 1, 3, 3,
	1, 1, 2, 3, 1, 1, 3, 4, 6, 8,
	1, 1, 1, 1, 2, 1, 1, 1, 3, 1,
	3, 2, 2, 1, 1, 2, 1, 1, 1, 2,
	3, 1, 1, 1, 2, 3, 1, 1, 1, 1,
}

var pathEvalChk = [...]int16{
	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, 23, -11, -12, -13, -14, -15, -19, -17, -22,
	-23, -20, -21, 30, 12, 4, 9, 10, -24, -25,
	-27, 25, -16, 11, -28, 13, 35, 5, 6, 36,
	14, 15, 17, 16, 20, 18, 21, 19, 22, 23,
	24, 26, 27, 29, -9, -18, -29, 33, 25, -16,
	25, -16, -15, -2, 30, -25, -26, -18, -15, 7,
	-4, -5, -6, -6, -7, -7, -7, -7, -8, -8,
	-9, -9, -9, -11, -30, -2, -15, -15, -22, -22,
	31, 31, -2, -26, -18, -31, 34, 31, 32, -2,
	31, 32, -2, 31,
}

var pathEvalDef = [...]int8{
	0, -2, 1, 2, 3, 5, 7, 10, 15, 18,
	22, 0, 24, 26, -2, 0, 41, 42, 31, 47,
	49, 43, 45, 0, 34, 35, 0, 40, 0, 53,
	54, 46, 0, 0, 56, 57, 66, 67, 69, 68,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 23, 32, 0, 61, 0, 0,
	0, 0, 44, 0, 0, 51, 52, 58, 64, 55,
	4, 6, 8, 9, 11, 12, 13, 14, 16, 17,
	19, 20, 21, 25, 0, 62, 28, 29, 48, 65,
	33, 36, 0, 50, 59, 60, 63, 37, 0, 0,
	38, 0, 0, 39,
}

var pathEvalTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	30, 31, 24, 22, 32, 23, 35, 25, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 36, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 33, 3, 34, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 29,
}

var pathEvalTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	26, 27, 28,
}

var pathEvalTok3 = [...]int8{
	0,
}

var pathEvalErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	pathEvalDebug        = 0
	pathEvalErrorVerbose = false
)

type pathEvalLexer interface {
	Lex(lval *pathEvalSymType) int
	Error(s string)
}

type pathEvalParser interface {
	Parse(pathEvalLexer) int
	Lookahead() int
}

type pathEvalParserImpl struct {
	lval  pathEvalSymType
	stack [pathEvalInitialStackSize]pathEvalSymType
	char  int
}

func (p *pathEvalParserImpl) Lookahead() int {
	return p.char
}

func pathEvalNewParser() pathEvalParser {
	return &pathEvalParserImpl{}
}

const pathEvalFlag = -1000

func pathEvalTokname(c int) string {
	if c >= 1 && c-1 < len(pathEvalToknames) {
		if pathEvalToknames[c-1] != "" {
			return pathEvalToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func pathEvalStatname(s int) string {
	if s >= 0 && s < len(pathEvalStatenames) {
		if pathEvalStatenames[s] != "" {
			return pathEvalStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func pathEvalErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !pathEvalErrorVerbose {
		return "syntax error"
	}

	for _, e := range pathEvalErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + pathEvalTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(pathEvalPact[state])
	for tok := TOKSTART; tok-1 < len(pathEvalToknames); tok++ {
		if n := base + tok; n >= 0 && n < pathEvalLast && int(pathEvalChk[int(pathEvalAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if pathEvalDef[state] == -2 {
		i := 0
		for pathEvalExca[i] != -1 || int(pathEvalExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; pathEvalExca[i] >= 0; i += 2 {
			tok := int(pathEvalExca[i])
			if tok < TOKSTART || pathEvalExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if pathEvalExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += pathEvalTokname(tok)
	}
	return res
}

func pathEvallex1(lex pathEvalLexer, lval *pathEvalSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(pathEvalTok1[0])
		goto out
	}
	if char < len(pathEvalTok1) {
		token = int(pathEvalTok1[char])
		goto out
	}
	if char >= pathEvalPrivate {
		if char < pathEvalPrivate+len(pathEvalTok2) {
			token = int(pathEvalTok2[char-pathEvalPrivate])
			goto out
		}
	}
	for i := 0; i < len(pathEvalTok3); i += 2 {
		token = int(pathEvalTok3[i+0])
		if token == char {
			token = int(pathEvalTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(pathEvalTok2[1]) /* unknown char */
	}
	if pathEvalDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", pathEvalTokname(token), uint(char))
	}
	return char, token
}

func pathEvalParse(pathEvallex pathEvalLexer) int {
	return pathEvalNewParser().Parse(pathEvallex)
}

func (pathEvalrcvr *pathEvalParserImpl) Parse(pathEvallex pathEvalLexer) int {
	var pathEvaln int
	var pathEvalVAL pathEvalSymType
	var pathEvalDollar []pathEvalSymType
	_ = pathEvalDollar // silence set and not used
	pathEvalS := pathEvalrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	pathEvalstate := 0
	pathEvalrcvr.char = -1
	pathEvaltoken := -1 // pathEvalrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		pathEvalstate = -1
		pathEvalrcvr.char = -1
		pathEvaltoken = -1
	}()
	pathEvalp := -1
	goto pathEvalstack

ret0:
	return 0

ret1:
	return 1

pathEvalstack:
	/* put a state and value onto the stack */
	if pathEvalDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", pathEvalTokname(pathEvaltoken), pathEvalStatname(pathEvalstate))
	}

	pathEvalp++
	if pathEvalp >= len(pathEvalS) {
		nyys := make([]pathEvalSymType, len(pathEvalS)*2)
		copy(nyys, pathEvalS)
		pathEvalS = nyys
	}
	pathEvalS[pathEvalp] = pathEvalVAL
	pathEvalS[pathEvalp].yys = pathEvalstate

pathEvalnewstate:
	pathEvaln = int(pathEvalPact[pathEvalstate])
	if pathEvaln <= pathEvalFlag {
		goto pathEvaldefault /* simple state */
	}
	if pathEvalrcvr.char < 0 {
		pathEvalrcvr.char, pathEvaltoken = pathEvallex1(pathEvallex, &pathEvalrcvr.lval)
	}
	pathEvaln += pathEvaltoken
	if pathEvaln < 0 || pathEvaln >= pathEvalLast {
		goto pathEvaldefault
	}
	pathEvaln = int(pathEvalAct[pathEvaln])
	if int(pathEvalChk[pathEvaln]) == pathEvaltoken { /* valid shift */
		pathEvalrcvr.char = -1
		pathEvaltoken = -1
		pathEvalVAL = pathEvalrcvr.lval
		pathEvalstate = pathEvaln
		if Errflag > 0 {
			Errflag--
		}
		goto pathEvalstack
	}

pathEvaldefault:
	/* default state action */
	pathEvaln = int(pathEvalDef[pathEvalstate])
	if pathEvaln == -2 {
		if pathEvalrcvr.char < 0 {
			pathEvalrcvr.char, pathEvaltoken = pathEvallex1(pathEvallex, &pathEvalrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if pathEvalExca[xi+0] == -1 && int(pathEvalExca[xi+1]) == pathEvalstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			pathEvaln = int(pathEvalExca[xi+0])
			if pathEvaln < 0 || pathEvaln == pathEvaltoken {
				break
			}
		}
		pathEvaln = int(pathEvalExca[xi+1])
		if pathEvaln < 0 {
			goto ret0
		}
	}
	if pathEvaln == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			pathEvallex.Error(pathEvalErrorMessage(pathEvalstate, pathEvaltoken))
			Nerrs++
			if pathEvalDebug >= 1 {
				__yyfmt__.Printf("%s", pathEvalStatname(pathEvalstate))
				__yyfmt__.Printf(" saw %s\n", pathEvalTokname(pathEvaltoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for pathEvalp >= 0 {
				pathEvaln = int(pathEvalPact[pathEvalS[pathEvalp].yys]) + pathEvalErrCode
				if pathEvaln >= 0 && pathEvaln < pathEvalLast {
					pathEvalstate = int(pathEvalAct[pathEvaln]) /* simulate a shift of "error" */
					if int(pathEvalChk[pathEvalstate]) == pathEvalErrCode {
						goto pathEvalstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if pathEvalDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", pathEvalS[pathEvalp].yys)
				}
				pathEvalp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if pathEvalDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", pathEvalTokname(pathEvaltoken))
			}
			if pathEvaltoken == pathEvalEofCode {
				goto ret1
			}
			pathEvalrcvr.char = -1
			pathEvaltoken = -1
			goto pathEvalnewstate /* try again in the same state */
		}
	}

	/* reduction by production pathEvaln */
	if pathEvalDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", pathEvaln, pathEvalStatname(pathEvalstate))
	}

	pathEvalnt := pathEvaln
	pathEvalpt := pathEvalp
	_ = pathEvalpt // guard against "declared and not used"

	pathEvalp -= int(pathEvalR2[pathEvaln])
	// pathEvalp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if pathEvalp+1 >= len(pathEvalS) {
		nyys := make([]pathEvalSymType, len(pathEvalS)*2)
		copy(nyys, pathEvalS)
		pathEvalS = nyys
	}
	pathEvalVAL = pathEvalS[pathEvalp+1]

	/* consult goto table to find next state */
	pathEvaln = int(pathEvalR1[pathEvaln])
	pathEvalg := int(pathEvalPgo[pathEvaln])
	pathEvalj := pathEvalg + pathEvalS[pathEvalp].yys + 1

	if pathEvalj >= pathEvalLast {
		pathEvalstate = int(pathEvalAct[pathEvalg])
	} else {
		pathEvalstate = int(pathEvalAct[pathEvalj])
		if int(pathEvalChk[pathEvalstate]) != -pathEvaln {
			pathEvalstate = int(pathEvalAct[pathEvalg])
		}
	}
	// dummy call; replaced with literal code
	switch pathEvalnt {

	case 1:
		pathEvalDollar = pathEvalS[pathEvalpt-1 : pathEvalpt+1]
//line path_eval.y:68
		{
			getProgBldr(pathEvallex).CodeFn(
				getProgBldr(pathEvallex).StorePathEval,
				"storePathEval")
		}
	case 26:
		pathEvalDollar = pathEvalS[pathEvalpt-1 : pathEvalpt+1]
//line path_eval.y:118
		{
			getProgBldr(pathEvallex).CodeEvalLocPathExists()
		}
	case 28:
		pathEvalDollar = pathEvalS[pathEvalpt-3 : pathEvalpt+1]
//line path_eval.y:123
		{
			getProgBldr(pathEvallex).CodeEvalLocPathExists()
		}
	case 29:
		pathEvalDollar = pathEvalS[pathEvalpt-3 : pathEvalpt+1]
//line path_eval.y:127
		{
			getProgBldr(pathEvallex).CodeEvalLocPathExists()
		}
	case 40:
		pathEvalDollar = pathEvalS[pathEvalpt-1 : pathEvalpt+1]
//line path_eval.y:150
		{
			getProgBldr(pathEvallex).UnsupportedName(xutils.NODETYPE, pathEvalDollar[1].name)
		}
	case 46:
		pathEvalDollar = pathEvalS[pathEvalpt-1 : pathEvalpt+1]
//line path_eval.y:171
		{
			getProgBldr(pathEvallex).CodePathOper('/')
		}
	case 55:
		pathEvalDollar = pathEvalS[pathEvalpt-2 : pathEvalpt+1]
//line path_eval.y:196
		{
			getProgBldr(pathEvallex).UnsupportedName(xutils.AXISNAME, pathEvalDollar[1].name)
		}
	case 57:
		pathEvalDollar = pathEvalS[pathEvalpt-1 : pathEvalpt+1]
//line path_eval.y:202
		{
			getProgBldr(pathEvallex).CodeNameTest(pathEvalDollar[1].xmlname)
		}
	case 61:
		pathEvalDollar = pathEvalS[pathEvalpt-1 : pathEvalpt+1]
//line path_eval.y:214
		{
			getProgBldr(pathEvallex).CodePredStartIgnore()
		}
	case 63:
		pathEvalDollar = pathEvalS[pathEvalpt-1 : pathEvalpt+1]
//line path_eval.y:222
		{
			getProgBldr(pathEvallex).CodePredEndIgnore()
		}
	case 66:
		pathEvalDollar = pathEvalS[pathEvalpt-1 : pathEvalpt+1]
//line path_eval.y:234
		{
			getProgBldr(pathEvallex).CodePathOper('.')
		}
	case 67:
		pathEvalDollar = pathEvalS[pathEvalpt-1 : pathEvalpt+1]
//line path_eval.y:238
		{
			getProgBldr(pathEvallex).CodePathOper(xutils.DOTDOT)
		}
	case 68:
		pathEvalDollar = pathEvalS[pathEvalpt-1 : pathEvalpt+1]
//line path_eval.y:244
		{
			getProgBldr(pathEvallex).UnsupportedName(
				'@', "not yet implemented")
		}
	case 69:
		pathEvalDollar = pathEvalS[pathEvalpt-1 : pathEvalpt+1]
//line path_eval.y:251
		{
			getProgBldr(pathEvallex).UnsupportedName(
				xutils.DBLSLASH, "not yet implemented")
		}
	}
	goto pathEvalstack /* stack new state and value */
}
