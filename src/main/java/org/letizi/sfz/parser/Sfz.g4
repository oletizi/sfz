// Define a grammar called Hello
grammar Sfz;
sfz
  : line+ EOF
  ;
line
  : (header_stmt | opcode_stmt) Newline?
  ;
header_stmt
  : '<' header '>'
  ;
header
  :
  (
  'region'
  | 'group'
  | 'control'
  | 'global'
  | 'curve'
  | 'effect'
  | 'master'
  | 'midi'
  | 'sampler'
  )
  ;

opcode_stmt
  : opcode Whitespace? '=' Whitespace? value
  ;

opcode
  :
  ( 'sample'
   | 'sw_default'
   | 'sw_lokey' )
  ;

value
  : Text
  ;

Text
  : [0-9a-zA-Z_\\/.*]+
  ;

Newline
  :
  ( '\r' '\n'
  | '\n' )
  ;

Whitespace
  : [ \t]+
  -> skip
  ;

BlockComment
    :   '/*' .*? '*/'
        -> skip
    ;

LineComment
    :   '//' ~[\r\n]*
        -> skip
    ;