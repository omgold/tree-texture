EXECUTABLE=tree-texture

default: $(EXECUTABLE)

all: examples

$(EXECUTABLE): tree-texture.go
	go build $<

examples: e1.png e2.png e3.png e4.png e5.png e6.png e7.png e8.png

e1.png: $(EXECUTABLE)
	./$(EXECUTABLE) -f $@ -sb 0.9 -is 1.2 -rs 0.318309886 -re 0.318309886 -di 0.75 -dor 0.25  -dm 0.8 -sl 0.75 -rml 0.2 -rmb 0.2 -bd 5 -nl 3.1 -ng 0.4 -na 0.3 -bj 0 -seed 2 -cx 0.8 -cy 0.2

e2.png: $(EXECUTABLE)
	./$(EXECUTABLE) -f $@ -is 0.9 -rs 0.42 -re 0.42 -di 0.5 -dor 0.25 -dm 0.8 -rml 1e-4 -rmb 0.3 -bd 3 -na 0.28 -bj 0.15 -seed 4 -cx 0.8 -cy 0.2

e3.png: $(EXECUTABLE)
	./$(EXECUTABLE) -f $@ -is 0.9 -rs 0.45 -re 0 -dor 0.6 -dol 0.3 -dm 0.8 -rml 1e-4 -rmb 0.3 -bd 2 -na 0.28 -seed 1 -cx 0.8 -cy 0.2

e4.png: $(EXECUTABLE)
	./$(EXECUTABLE) -f $@ -sb 0.95 -rml 1e-4 -rmb 1e-4 -shapel 2 -shapeb 2 -bd 2 -na 0.28 -seed 1 -cx 0.8 -cy 0.2

e5.png: $(EXECUTABLE)
	./$(EXECUTABLE) -f $@ -sb 0.95 -rml 1e-4 -rmb 1e-4 -shapel 2 -shapeb 1.5 -bd 2 -nl 1.5 -ng 1.1 -na 0.4 -seed 1 -cx 0.8 -cy 0.2

e6.png: $(EXECUTABLE)
	./$(EXECUTABLE) -f $@ -sb 0.95 -dol 0.2 -shapel 0.5 -shapeb 1.5 -bd 2 -nl 3 -ng 0.8 -na 0.2 -seed 1 -cx 0.8 -cy 0.2

e7.png: $(EXECUTABLE)
	./$(EXECUTABLE) -f $@ -sb 0.95 -di 0.7 -dor 0.2 -dol 0.5 -shapel 0.5 -shapeb 1.5 -bd 2 -nl 2.5 -ng 0.8 -seed 1 -cx 0.8 -cy 0.2

e8.png: $(EXECUTABLE)
	./$(EXECUTABLE) -f $@ -sb 0.7 -re 0.9 -di 0.7 -dor 0.2 -dol 0.5 -shapel -1 -bd 3 -nl 2.5 -ng 0.8 -seed 1 -cx 0.8 -cy 0.2
