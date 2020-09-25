// Define a grammar called Hello
grammar Sfz;
NEWLINE: ('\n'|'\r\n');
TEXT: ('a'..'z'|'A'..'Z'|'0'..'9')+;
OPCODE: ('a'..'z'|'A'..'Z'|'_')+;
CONTEXT: ('a'..'z'|'A'..'Z')+;
sfz : line+ EOF;
line : (context | statement) NEWLINE?;
context: '<' CONTEXT '>';
statement: OPCODE '=' value;
value: TEXT;
