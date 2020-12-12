require "httpclient"
require "docker"
require "json"
require "pg"

Docker.url = "tcp://127.0.0.1:2375"

RSpec.configure do |config|
  pgContainer = nil
  $serverContainer = nil
  network = nil

  config.before(:suite) do
    network = Docker::Network.create("rspec")

    #start postgres
    pgContainer = create_postgres()
    pgContainer.start
    pgPort = pgContainer.json["NetworkSettings"]["Ports"]["5432/tcp"][0]["HostPort"]

    sleep 10

    $pg = PG.connect(
      :host => 'localhost',
      :port => pgPort,
      :dbname => 'postgres',
      :user => 'soupstoregames',
      :password => 'twitch2020',
    )
    $pg.exec("CREATE DATABASE gotickyourself")
    $pg.close()
    $pg = PG.connect(
      :host => 'localhost',
      :port => pgPort,
      :dbname => 'gotickyourself',
      :user => 'soupstoregames',
      :password => 'twitch2020',
    )

    # start server
    $serverContainer = create_server()
    $serverContainer.start
    port = $serverContainer.json["NetworkSettings"]["Ports"]["8080/tcp"][0]["HostPort"]
    $endpoint = "http://127.0.0.1:#{port}"

    # create http client
    $client = HTTPClient.new

    # wait for servers to start
    sleep 2
  end

  config.after(:suite) do
    $serverContainer.stop
    $serverContainer.delete
    pgContainer.stop
    pgContainer.delete
    network.delete
  end
end

def create_server()
  return Docker::Container.create(
    "Image" => "soupstoregames/go-tick-yourself:dev",
    "Env" => [
      "DB_HOST=postgres-test"
    ],
    "HostConfig" => {
      "NetworkMode" => "rspec",
      "PortBindings" => {
        "8080/tcp" => [{}]
      },
    },
  )
end

def create_postgres()
  Docker::Image.create('fromImage' => 'postgres:12')
  return Docker::Container.create(
    'name' => 'postgres-test',
    "Image" => "postgres:12",
    "Env" => [
      "POSTGRES_USER=soupstoregames",
      "POSTGRES_PASSWORD=twitch2020"
    ],
    "HostConfig" => {
      "NetworkMode" => "rspec",
      "PortBindings" => {
        "5432/tcp" => [{}]
      },
    }
  )
end
