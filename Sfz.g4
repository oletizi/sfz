grammar Sfz;

sfz
  : sfzObject* EOF
  ;

sfzObject
  :
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
  : LT headerName GT
  ;

headerName
  : GLOBAL
  | GROUP
  | REGION
  ;

LT
  : '<'
  ;

GT
  : '>'
  ;

GLOBAL
  : 'global'
  ;

GROUP
  : 'group'
  ;

REGION
  : 'region'
  ;


opcodeStatement
  : opcode WHITESPACE? '=' WHITESPACE? value
  ;

opcode
  :
  ( HIKEY
  | KEY
  | LOKEY
  | SAMPLE
  )
  ;

value
  :
  INT
  ;

HIKEY: 'hikey';
KEY: 'key';
LOKEY: 'lokey';
SAMPLE: 'sample';

INT: '0'..'9'+;

fragment FLOAT: (INT | INT? '.' INT);

fragment Text: ~[\n\r]+;

NEWLINE: ('\r\n' | '\n' | '\r');

WHITESPACE: [ \t];

BLOCK_COMMENT: '/*' .*? '*/' -> skip;

LINE_COMMENT: '//' ~[\r\n]* -> skip;

HASH_COMMENT: '#' ~[\r\n]* -> skip;

//header_phrase
//  : '<'? 'region' '>'?
//  ;

//header
//  :
//  ( 'region'
//  | 'group'
//  | 'control'
//  | 'global'
//  | 'curve'
//  | 'effect'
//  | 'master'
//  | 'midi'
//  | 'sampler'
//  )
//  ;

//opcode_phrase
//  :
//  ( int_opcode WHITESPACE? '=' WHITESPACE? int_value
//  | float_opcode WHITESPACE? '=' WHITESPACE? float_value
//  | text_opcode WHITESPACE? '=' WHITESPACE? text_value
//  )
//  ;
//
//int_opcode
//  :
//  ( 'hikey'
//  | 'key'
//  | 'lokey'
//  )
//  ;
//
//float_opcode
//  :
//  ( 'ampeg_release'  )
//  ;
//
//text_opcode
//  :
//  ('sample')
//  ;
//
//int_value: INT;
//
//float_value: FLOAT;
//
//text_value: TEXT;


