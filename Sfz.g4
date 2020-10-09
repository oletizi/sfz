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
  : opcode WHITESPACE? '=' WHITESPACE? value
  ;

opcode
  :
  ( 'ampeg_release'
  | 'hikey'
  | 'key'
  | 'lokey'
  | 'sample'
  )
  ;

value
  :
  ( INT
  | FLOAT
  | PATH
  )
  ;


INT: '0'..'9'+;

FLOAT: (INT | INT? '.' INT);

PATH
  :
  SEPARATOR* PATH_SEGMENT (SEPARATOR+ PATH_SEGMENT)*
  ;

fragment SEPARATOR
  : [/\\]
  ;

fragment PATH_SEGMENT
  : [a-zA-Z0-9._\-]+
  ;

NEWLINE: ('\r\n' | '\n' | '\r');

WHITESPACE: [ \t];

BLOCK_COMMENT: '/*' .*? '*/' -> skip;

LINE_COMMENT: '//' ~[\r\n]* -> skip;

HASH_COMMENT: '#' ~[\r\n]* -> skip;