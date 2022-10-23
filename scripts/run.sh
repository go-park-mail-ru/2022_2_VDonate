# Clearing all stopped containers
docker container prune -f
# UP backend docker compose
docker-compose -f ../deployments/docker-compose.yaml up -d
