echo creating docker bridge network
docker network create go-confluent-network
echo starting Confluent Kafka
docker-compose up -d
echo starting Producer API
cd ./producer && docker-compose up -d
echo starting Consumer API
cd ../consumer && docker-compose up -d