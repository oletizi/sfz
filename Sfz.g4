grammar Sfz;

sfz
  : sfzObject* EOF
  ;

sfzObject
  :
  ( WHITESPACE
  | NEWLINE
  )*
  ( headerObject
    ( ( WHITESPACE
       | NEWLINE
       )*
       headerObject
    | ( WHITESPACE
      | NEWLINE
      )+
      opcodeStatement
    )*
  )
  ;

headerObject
  : '<' header '>'
  ;

header
  : 'global'
  | 'group'
  | 'region'
  ;

opcodeStatement
  : opcode '=' STRING
  ;

opcode
  : STRING?
  ;

//value
//  : STRING?
//  ;

STRING
  : ~[=<> \t\n\r]+
  ;

NEWLINE: ('\r\n' | '\n' | '\r');

WHITESPACE: [ \t];

BLOCK_COMMENT: '/*' .*? '*/' -> skip;

LINE_COMMENT: '//' ~[\r\n]* -> skip;

HASH_COMMENT: '#' ~[\r\n]* -> skip;