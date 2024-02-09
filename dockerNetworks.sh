docker network ls --format json| \
  jq 'select(.Driver | contains("bridge")) | {Name: .Name}' | jq 'select(.Name != "bridge")'


