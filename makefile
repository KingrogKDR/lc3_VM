SRC = $(wildcard src/*.go)
OUT = lc3

all: $(OUT)

$(OUT):
	go build -o $(OUT) $(SRC)

clean:
	rm -f $(OUT)

run: $(OUT)
	./$(OUT)

rebuild: clean all
