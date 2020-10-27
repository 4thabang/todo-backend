FROM golang:1.15-alpine as base
# Create app folder 
RUN mkdir /main
# Copy our file in the host contianer to our contianer
ADD . /main
# Set /main to the go folder as workdir
WORKDIR /main
# Generate binary file from our /main
RUN go build
# Expose the port 3000
EXPOSE 3000:3000
# Run the main binarry file 
CMD ["./main"]