Hosted Image:

Required: ollama

docker pull henryarrovin/json-mock-generator:latest

docker run -d -p 8080:8080 \
  --name=json-mock-api \
  --add-host=host.docker.internal:host-gateway \
  -e OLLAMA_URL=http://host.docker.internal:11434/api/generate \
  -e OLLAMA_MODEL=phi \
  henryarrovin/json-mock-generator:latest
