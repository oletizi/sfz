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
  | 'bend_down'
  | 'bend_up'
  | 'hikey'
  | 'hivel'
  | 'key'
  | 'lokey'
  | 'lovel'
  | 'pitch_keycenter'
  | 'sample'
  | 'seq_length'
  | 'seq_position'
  | 'sw_default'
  | 'sw_hikey'
  | 'sw_last'
  | 'sw_lokey'
  )
  ;

value
  :
  ( INT
  | FLOAT
  | PATH
  | GENERATOR
  )
  ;

GENERATOR
  :
  '*' [a-z]+
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