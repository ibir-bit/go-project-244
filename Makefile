lint:
	golangci-lint run

# Сборка
build:
	mkdir -p bin
	go build -o bin/gendiff ./cmd/gendiff

# Запуск (пример)
run: build
	./bin/gendiff file1.json file2.json

# Очистка
clean:
	rm -rf bin/