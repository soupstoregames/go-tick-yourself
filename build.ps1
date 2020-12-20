docker build --tag "go-tick-yourself:dev" --file "Dockerfile" --quiet .
docker build --tag "go-tick-yourself:specs" --file "Dockerfile.specs" --quiet .

docker network create "go-tick-yourself"
docker run  --rm --name "go-tick-yourself.db" --network "go-tick-yourself" --env "POSTGRES_USER=soupstoregames" --env "POSTGRES_PASSWORD=twitch2020" --env "POSTGRES_DB=gotickyourself" --detach "postgres:13"
Start-Sleep -Seconds 1
docker run  --rm --name "go-tick-yourself.app" --network "go-tick-yourself" --env "DB_USER=soupstoregames" --env "DB_PASSWORD=twitch2020" --env "DB_NAME=gotickyourself" --env "DB_HOST=go-tick-yourself.db"  --detach "go-tick-yourself:dev"
Start-Sleep -Seconds 1
docker run  --rm --name "go-tick-yourself.specs" --network "go-tick-yourself" "go-tick-yourself:specs"

docker stop "go-tick-yourself.app"
docker stop "go-tick-yourself.db"
docker network rm "go-tick-yourself"
