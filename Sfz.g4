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
  : LT header GT
  ;

header
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
  ( AMPEG_RELEASE
  | HIKEY
  | KEY
  | LOKEY
  | SAMPLE
  )
  ;

value
  :
  ( INT
  | FLOAT
  | PATH
  )
  ;

AMPEG_RELEASE: 'ampeg_release';
HIKEY: 'hikey';
KEY: 'key';
LOKEY: 'lokey';
SAMPLE: 'sample';

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

//FILE_PATH:
//  ( FILENAME
//  | DIRNAME * SEPARATOR
//  )
//  ;

//Text: ~[\n\r]+;

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


