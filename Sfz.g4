// Define a grammar called Hello
grammar Sfz;
NEWLINE: ('\n'|'\r\n');
TEXT: ('a'..'z'|'A'..'Z'|'0'..'9')+;
OPCODE: ('a'..'z'|'A'..'Z'|'_')+;
sfz : line+ EOF ;
line : (context | statement) NEWLINE;
context: '<' TEXT '>';
statement: OPCODE '=' value;
value: TEXT;
