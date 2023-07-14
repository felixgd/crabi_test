# Build the Docker image
build:
	docker build --no-cache -t crabi-test .

# Build the Docker image with debug enabled
build-debug:
	docker run -p 8080:8080 -p 2345:2345 -d crabi-test-debug

# Run the Docker image with debug enabled
run-debug:
	docker run -p 8080:8080 -p 2345:2345 -d crabi-test-debug

# Run the Docker container
run:
	docker-compose up

# Clean up Docker artifacts
clean:
	docker-compose down --volumes --rmi all

# Stop the containers
stop:
	docker-compose stop

restart:
	docker-compose run
