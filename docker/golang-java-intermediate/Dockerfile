FROM circleci/golang:1.9

ENV JAVA_HOME=/usr/local/jdk-15
ENV PATH=${PATH}:${JAVA_HOME}/bin
ENV CLASSPATH=${CLASSPATH}:/usr/local/lib/antlr-4.8-complete.jar


RUN cd /usr/local/ && curl https://download.java.net/java/GA/jdk15/779bf45e88a44cbd9ea6621d33e33db1/36/GPL/openjdk-15_linux-x64_bin.tar.gz  | sudo tar xfvz -



