// Code generated from antlr/TSgen.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type TSgenLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var TSgenLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tsgenlexerLexerInit() {
	staticData := &TSgenLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'state'", "'final'", "'->'", "':'", "'('", "')'", "','", "'pkg'",
		"'import'", "'protocol'", "'file'", "'raw'", "'name'", "'returns'",
		"'.'", "'s0'", "'fstate'", "'['", "']'", "'*'", "'...'", "'when'", "'true'",
		"'false'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "LBRACK", "RBRACK", "STAR", "ELLIPSIS", "WHEN", "TRUE", "FALSE",
		"COMMENT", "FILENAME", "ID", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"LBRACK", "RBRACK", "STAR", "ELLIPSIS", "WHEN", "TRUE", "FALSE", "COMMENT",
		"FILENAME", "ID", "STRING", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 205, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 1, 24, 5, 24, 167, 8, 24, 10, 24, 12, 24, 170, 9, 24, 1,
		24, 1, 24, 1, 25, 4, 25, 175, 8, 25, 11, 25, 12, 25, 176, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 26, 1, 26, 5, 26, 185, 8, 26, 10, 26, 12, 26, 188, 9,
		26, 1, 27, 1, 27, 5, 27, 192, 8, 27, 10, 27, 12, 27, 195, 9, 27, 1, 27,
		1, 27, 1, 28, 4, 28, 200, 8, 28, 11, 28, 12, 28, 201, 1, 28, 1, 28, 0,
		0, 29, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28,
		57, 29, 1, 0, 6, 2, 0, 10, 10, 13, 13, 4, 0, 45, 57, 65, 90, 95, 95, 97,
		122, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122,
		3, 0, 10, 10, 13, 13, 34, 34, 3, 0, 9, 10, 13, 13, 32, 32, 209, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0,
		0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0,
		0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1,
		0, 0, 0, 0, 57, 1, 0, 0, 0, 1, 59, 1, 0, 0, 0, 3, 65, 1, 0, 0, 0, 5, 71,
		1, 0, 0, 0, 7, 74, 1, 0, 0, 0, 9, 76, 1, 0, 0, 0, 11, 78, 1, 0, 0, 0, 13,
		80, 1, 0, 0, 0, 15, 82, 1, 0, 0, 0, 17, 86, 1, 0, 0, 0, 19, 93, 1, 0, 0,
		0, 21, 102, 1, 0, 0, 0, 23, 107, 1, 0, 0, 0, 25, 111, 1, 0, 0, 0, 27, 116,
		1, 0, 0, 0, 29, 124, 1, 0, 0, 0, 31, 126, 1, 0, 0, 0, 33, 129, 1, 0, 0,
		0, 35, 136, 1, 0, 0, 0, 37, 138, 1, 0, 0, 0, 39, 140, 1, 0, 0, 0, 41, 142,
		1, 0, 0, 0, 43, 146, 1, 0, 0, 0, 45, 151, 1, 0, 0, 0, 47, 156, 1, 0, 0,
		0, 49, 162, 1, 0, 0, 0, 51, 174, 1, 0, 0, 0, 53, 182, 1, 0, 0, 0, 55, 189,
		1, 0, 0, 0, 57, 199, 1, 0, 0, 0, 59, 60, 5, 115, 0, 0, 60, 61, 5, 116,
		0, 0, 61, 62, 5, 97, 0, 0, 62, 63, 5, 116, 0, 0, 63, 64, 5, 101, 0, 0,
		64, 2, 1, 0, 0, 0, 65, 66, 5, 102, 0, 0, 66, 67, 5, 105, 0, 0, 67, 68,
		5, 110, 0, 0, 68, 69, 5, 97, 0, 0, 69, 70, 5, 108, 0, 0, 70, 4, 1, 0, 0,
		0, 71, 72, 5, 45, 0, 0, 72, 73, 5, 62, 0, 0, 73, 6, 1, 0, 0, 0, 74, 75,
		5, 58, 0, 0, 75, 8, 1, 0, 0, 0, 76, 77, 5, 40, 0, 0, 77, 10, 1, 0, 0, 0,
		78, 79, 5, 41, 0, 0, 79, 12, 1, 0, 0, 0, 80, 81, 5, 44, 0, 0, 81, 14, 1,
		0, 0, 0, 82, 83, 5, 112, 0, 0, 83, 84, 5, 107, 0, 0, 84, 85, 5, 103, 0,
		0, 85, 16, 1, 0, 0, 0, 86, 87, 5, 105, 0, 0, 87, 88, 5, 109, 0, 0, 88,
		89, 5, 112, 0, 0, 89, 90, 5, 111, 0, 0, 90, 91, 5, 114, 0, 0, 91, 92, 5,
		116, 0, 0, 92, 18, 1, 0, 0, 0, 93, 94, 5, 112, 0, 0, 94, 95, 5, 114, 0,
		0, 95, 96, 5, 111, 0, 0, 96, 97, 5, 116, 0, 0, 97, 98, 5, 111, 0, 0, 98,
		99, 5, 99, 0, 0, 99, 100, 5, 111, 0, 0, 100, 101, 5, 108, 0, 0, 101, 20,
		1, 0, 0, 0, 102, 103, 5, 102, 0, 0, 103, 104, 5, 105, 0, 0, 104, 105, 5,
		108, 0, 0, 105, 106, 5, 101, 0, 0, 106, 22, 1, 0, 0, 0, 107, 108, 5, 114,
		0, 0, 108, 109, 5, 97, 0, 0, 109, 110, 5, 119, 0, 0, 110, 24, 1, 0, 0,
		0, 111, 112, 5, 110, 0, 0, 112, 113, 5, 97, 0, 0, 113, 114, 5, 109, 0,
		0, 114, 115, 5, 101, 0, 0, 115, 26, 1, 0, 0, 0, 116, 117, 5, 114, 0, 0,
		117, 118, 5, 101, 0, 0, 118, 119, 5, 116, 0, 0, 119, 120, 5, 117, 0, 0,
		120, 121, 5, 114, 0, 0, 121, 122, 5, 110, 0, 0, 122, 123, 5, 115, 0, 0,
		123, 28, 1, 0, 0, 0, 124, 125, 5, 46, 0, 0, 125, 30, 1, 0, 0, 0, 126, 127,
		5, 115, 0, 0, 127, 128, 5, 48, 0, 0, 128, 32, 1, 0, 0, 0, 129, 130, 5,
		102, 0, 0, 130, 131, 5, 115, 0, 0, 131, 132, 5, 116, 0, 0, 132, 133, 5,
		97, 0, 0, 133, 134, 5, 116, 0, 0, 134, 135, 5, 101, 0, 0, 135, 34, 1, 0,
		0, 0, 136, 137, 5, 91, 0, 0, 137, 36, 1, 0, 0, 0, 138, 139, 5, 93, 0, 0,
		139, 38, 1, 0, 0, 0, 140, 141, 5, 42, 0, 0, 141, 40, 1, 0, 0, 0, 142, 143,
		5, 46, 0, 0, 143, 144, 5, 46, 0, 0, 144, 145, 5, 46, 0, 0, 145, 42, 1,
		0, 0, 0, 146, 147, 5, 119, 0, 0, 147, 148, 5, 104, 0, 0, 148, 149, 5, 101,
		0, 0, 149, 150, 5, 110, 0, 0, 150, 44, 1, 0, 0, 0, 151, 152, 5, 116, 0,
		0, 152, 153, 5, 114, 0, 0, 153, 154, 5, 117, 0, 0, 154, 155, 5, 101, 0,
		0, 155, 46, 1, 0, 0, 0, 156, 157, 5, 102, 0, 0, 157, 158, 5, 97, 0, 0,
		158, 159, 5, 108, 0, 0, 159, 160, 5, 115, 0, 0, 160, 161, 5, 101, 0, 0,
		161, 48, 1, 0, 0, 0, 162, 163, 5, 47, 0, 0, 163, 164, 5, 47, 0, 0, 164,
		168, 1, 0, 0, 0, 165, 167, 8, 0, 0, 0, 166, 165, 1, 0, 0, 0, 167, 170,
		1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 171, 1, 0,
		0, 0, 170, 168, 1, 0, 0, 0, 171, 172, 6, 24, 0, 0, 172, 50, 1, 0, 0, 0,
		173, 175, 7, 1, 0, 0, 174, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176,
		174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179,
		5, 46, 0, 0, 179, 180, 5, 103, 0, 0, 180, 181, 5, 111, 0, 0, 181, 52, 1,
		0, 0, 0, 182, 186, 7, 2, 0, 0, 183, 185, 7, 3, 0, 0, 184, 183, 1, 0, 0,
		0, 185, 188, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187,
		54, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189, 193, 5, 34, 0, 0, 190, 192,
		8, 4, 0, 0, 191, 190, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191, 1, 0,
		0, 0, 193, 194, 1, 0, 0, 0, 194, 196, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0,
		196, 197, 5, 34, 0, 0, 197, 56, 1, 0, 0, 0, 198, 200, 7, 5, 0, 0, 199,
		198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201, 202,
		1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 204, 6, 28, 0, 0, 204, 58, 1, 0,
		0, 0, 6, 0, 168, 176, 186, 193, 201, 1, 0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// TSgenLexerInit initializes any static state used to implement TSgenLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewTSgenLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TSgenLexerInit() {
	staticData := &TSgenLexerLexerStaticData
	staticData.once.Do(tsgenlexerLexerInit)
}

// NewTSgenLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewTSgenLexer(input antlr.CharStream) *TSgenLexer {
	TSgenLexerInit()
	l := new(TSgenLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &TSgenLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "TSgen.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TSgenLexer tokens.
const (
	TSgenLexerT__0     = 1
	TSgenLexerT__1     = 2
	TSgenLexerT__2     = 3
	TSgenLexerT__3     = 4
	TSgenLexerT__4     = 5
	TSgenLexerT__5     = 6
	TSgenLexerT__6     = 7
	TSgenLexerT__7     = 8
	TSgenLexerT__8     = 9
	TSgenLexerT__9     = 10
	TSgenLexerT__10    = 11
	TSgenLexerT__11    = 12
	TSgenLexerT__12    = 13
	TSgenLexerT__13    = 14
	TSgenLexerT__14    = 15
	TSgenLexerT__15    = 16
	TSgenLexerT__16    = 17
	TSgenLexerLBRACK   = 18
	TSgenLexerRBRACK   = 19
	TSgenLexerSTAR     = 20
	TSgenLexerELLIPSIS = 21
	TSgenLexerWHEN     = 22
	TSgenLexerTRUE     = 23
	TSgenLexerFALSE    = 24
	TSgenLexerCOMMENT  = 25
	TSgenLexerFILENAME = 26
	TSgenLexerID       = 27
	TSgenLexerSTRING   = 28
	TSgenLexerWS       = 29
)
