#!/bin/bash

DOCKER_COMPOSE_ADDR="http://localhost:8080/query"
HELM_ADDR="http://localhost:8081/query"

echo "Running performance tests..."
echo -e "\tDocker Compose"
start=$(date -u +%s)
./bin/guacone collect files ../guac-data/docs/ --gql-addr $DOCKER_COMPOSE_ADDR >/dev/null 2>&1
end=$(date -u +%s)
elapsed_docker=$((end - start))

echo -e "\tHelm"
start=$(date -u +%s)
./bin/guacone collect files ../guac-data/docs/ --gql-addr $HELM_ADDR >/dev/null 2>&1
end=$(date -u +%s)
elapsed_helm=$((end - start))

echo -e "\n\n\n"

printf "%-15s%-20s%s\n" "Test" "Deployment Env" "Time"
printf "%-15s%-20s%s\n" "----" "--------------" "----"
printf "%-15s%-20s%s seconds\n" "Performance" "Docker Compose" "$elapsed_docker"
printf "%-15s%-20s%s seconds\n" "Performance" "Helm" "$elapsed_helm"
