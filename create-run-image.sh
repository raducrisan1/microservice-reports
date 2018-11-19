#run this locally so that you do not need to restore all the time the external dependencies (go get)
#of course, in a CI/CD environment, you need to change this approach 
docker rm $(docker ps -aqf "name=microservice-reports")
docker build -t local/microservice-reports .
docker tag local/microservice-reports gcr.io/itdays-201118/microservice-reports
docker run \
    --name microservice-reports \
    -e REPORTS_LISTEN_ADDR=':3070' \
    -e MONGO_ADDR='mongodb://172.17.0.1:27017' \
    -p 3070:3070 \
    local/microservice-reports