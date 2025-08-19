grammar TSgen;

// =================== PARSER RULES ===================

program
  : statement* EOF
  ;

statement
  : stateDecl
  | transitionDecl
  | docComment
  | pkgDecl
  | importDecl
  | protocolDecl
  | fileDecl
  | rawDecl
  | s0Decl
  | fDecl
  ;

// e.g.  state S2 final
stateDecl
  : 'state' from=ID ('final')?
  ;

LBRACK   : '[' ;
RBRACK   : ']' ;
STAR     : '*' ;
ELLIPSIS : '...' ;
WHEN     : 'when' ;
TRUE     : 'true' ;
FALSE    : 'false' ;
transitionDecl
  : from=ID '->' label=ID ':' method=ID '(' tArgs? ')' '->' to=ID ( WHEN outcome=(TRUE|FALSE) )?
  ;

tArgs : tParam (',' tParam)* ;

tParam : ID typeExpr | typeExpr ;

docComment
  : COMMENT
  ;

// --- config ---

pkgDecl       : 'pkg' ID ;
importDecl    : 'import' STRING ;
protocolDecl  : 'protocol' ID ;
fileDecl      : 'file' FILENAME ;

// ----- raw method declaration -----
rawDecl
  : 'raw' from=ID label=ID 'name' mname=ID '(' rawInputs? ')' 'returns' rawOutputs
  ;

// One or more comma-separated parameters
rawInputs : rawParam (',' rawParam)* ;

// A parameter may be "name type" or just "type" (unnamed)
rawParam : (ID typeExpr)  | typeExpr ;

// One or more comma-separated result types
rawOutputs : typeExpr (',' typeExpr)* ;

// A simplified Go-ish type expression:
//   optional "..." for variadics
//   zero or more "*" pointers
//   zero or more "[]" slices
//   then a qualified identifier like "json.Decoder", "byte", "error"
typeExpr : ELLIPSIS? (STAR)* (LBRACK RBRACK)* qualID ;

// Qualified identifier: "pkg.Type" or just "Type"
qualID : ID ('.' ID)* ;

// explicit start / final state
s0Decl        : 's0' ID ;
fDecl         : 'fstate' ID ;

// =================== LEXER RULES ===================

COMMENT     : '//' ~[\r\n]* -> channel(HIDDEN);
FILENAME    : [a-zA-Z0-9_./-]+ '.go';
ID          : [a-zA-Z_][a-zA-Z_0-9]*;
STRING      : '"' (~["\r\n])* '"' ;
WS          : [ \t\r\n]+ -> channel(HIDDEN);
