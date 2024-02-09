NET=`docker network ls --format json| \
  jq 'select(.Driver | contains("bridge")) | {Name: .Name}' | \
  jq 'select(.Name != "bridge")' | jq .Name | tr -d \"`
  echo "mynet=${NET}" > .env
  echo "mynet=${NET}"


