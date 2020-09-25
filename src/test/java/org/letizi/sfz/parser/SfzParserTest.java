package org.letizi.sfz.parser;

import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CommonTokenStream;
import org.antlr.v4.runtime.ParserRuleContext;
import org.apache.logging.log4j.Level;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.apache.logging.log4j.core.config.Configurator;
import org.junit.jupiter.api.Test;


public class SfzParserTest {
    static {
        Configurator.setRootLevel(Level.INFO);
    }

    private static final Logger logger = LogManager.getLogger(SfzParserTest.class.getName());


    @Test
    public void testBasics() throws Exception {
        final String filename = getClass().getClassLoader().getResource("test-big.sfz").getFile();
        logger.info("Testing file: " + filename);
        final CharStream in = CharStreams.fromFileName(filename);
        final SfzLexer lexer = new SfzLexer(in);
        final SfzParser parser = new SfzParser(new CommonTokenStream(lexer));
        parser.addParseListener(new TestSfzListener());
        parser.sfz();
    }

    static class TestSfzListener extends SfzBaseListener {
        private final Logger logger = LogManager.getLogger(TestSfzListener.class.getName());

        public void exitHeader(SfzParser.HeaderContext ctx) {
            logger.info("exitHeader: " + ctx.getText());
        }

        public void exitOpcode(SfzParser.OpcodeContext ctx) {
            logger.info("exitOpcode: " + ctx.getText());
        }

        public void exitValue(SfzParser.ValueContext ctx) {
            logger.info("exitValue: " + ctx.getText());
        }
    }
}