
TARGET = iko
TARGETDIR = build
SRCS = $(shell fd -e .go)

all: $(TARGETDIR)/$(TARGET)

.PHONY: run clean

run:
	./$(TARGETDIR)/$(TARGET)

$(TARGETDIR)/$(TARGET): $(SRCS)
	go build -o $@ .

clean:
	rm -r $(TARGETDIR)

