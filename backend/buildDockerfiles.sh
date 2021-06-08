#!/bin/bash

docker start
docker build -t user_service -f UserServiceDockerfile .
docker build -t content_service -f ContentServiceDockerfile .
docker build -t recommendation_service -f RecommendationServiceDockerfile .
docker-compose up