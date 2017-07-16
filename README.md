Golang Text to Speech with Alpine's Flite

https://pkgs.alpinelinux.org/package/v3.4/main/x86/flite

Docker:
 - move to backend directory 'cd /be'
 - docker build -t wury/say .
 - (list of flite commands) docker run --rm wury/say flite -h
 - (write output using docker command) docker run --rm -v $(pwd)/data:/data -w /data wury/say flite -t hello wury output.wav

 Using makefile:
  - make build
  - docker run --rm -v $(pwd)/data:/data -w /data wury/say "golang"
-----------------------------------------------------------------------------------------------

# How to run Server and Client

  Run Server:
    - build first,
    - $ cd /be
    - $ make build
    - ./run.sh Or docker run -p 8080:8080 wury/say
  Run Client
    - $ cd client
    - $ go run main.go -b "localhost:8080" "Hello Go"
    - you'll see the result at the same directory (output.wav)
