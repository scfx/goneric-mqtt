FROM alpine:latest  

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY ./bin/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"] 
