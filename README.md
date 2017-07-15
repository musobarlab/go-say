Golang Text to Speech with Flite's Alpine

Docker:
 - move to backend directory 'cd /be'
 - docker build -t wury/say .
 - (list of flite commands) docker run --rm wury/say flite -h
 - (write output using docker command) docker run --rm -v $(pwd)/data:/data -w /data wury/say flite -t hello wury output.wav

 Using makefile:
  - make build
  - docker run --rm -v $(pwd)/data:/data -w /data wury/say "golang"
