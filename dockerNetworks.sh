NET=`docker network ls --format json| \
  jq 'select(.Driver | contains("bridge")) | {Name: .Name}' | \
  jq 'select(.Name != "bridge")' | jq .Name | tr -d \"`
  echo "network=${NET}" > .env
  echo "network=${NET}"


