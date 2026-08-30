
TARGET = ui
TARGETDIR = build
SRCS = $(fd -e .go)

all: $(TARGETDIR)/$(TARGET)

.PHONY: run clean

run:
	./$(TARGETDIR)/$(TARGET)

$(TARGETDIR)/$(TARGET): $(SRCS)
	go build -o $@ cmd/main.go

clean:
	rm -r $(TARGETDIR)

