package org.letizi.sfz.parser;

import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CommonTokenStream;
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
        final String filename = getClass().getClassLoader().getResource("test.sfz").getFile();
        logger.info("Testing file: " + filename);
        final CharStream in = CharStreams.fromFileName(filename);
        final SfzLexer lexer = new SfzLexer(in);
        final SfzParser parser = new SfzParser(new CommonTokenStream(lexer));
        new TestSfzListener().visitSfz(parser.sfz());
    }

    static class TestSfzListener extends SfzBaseVisitor<SfzParser.LineContext> {
        private Logger logger = LogManager.getLogger(TestSfzListener.class.getName());

        @Override
        public SfzParser.LineContext visitLine(SfzParser.LineContext ctx) {
            logger.info("visitLine: ", ctx);
            return ctx;
        }
    }
}