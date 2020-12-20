require "httpclient"
require "docker"
require "json"
require "pg"
require "pathname"

Docker.url = "tcp://127.0.0.1:2375"

RSpec.configure do |config|
  $docker = Pathname.new('/.dockerenv').exist?

  config.before(:suite) do
    if $docker
      $pg = PG.connect(
        :host => 'go-tick-yourself.db',
        :port => '5432',
        :dbname => 'gotickyourself',
        :user => 'soupstoregames',
        :password => 'twitch2020',
      )
      $endpoint = "http://go-tick-yourself.app:8080"
    else
      network = Docker::Network.create("rspec")
      pgContainer = create_postgres()
      pgContainer.start
      pgPort = pgContainer.json["NetworkSettings"]["Ports"]["5432/tcp"][0]["HostPort"]
      sleep 1

      serverContainer = create_server()
      serverContainer.start
      serverPort = serverContainer.json["NetworkSettings"]["Ports"]["8080/tcp"][0]["HostPort"]
      sleep 1

      $pg = PG.connect(
        :host => 'localhost',
        :port => pgPort,
        :dbname => 'gotickyourself',
        :user => 'soupstoregames',
        :password => 'twitch2020',
      )

      $endpoint = "http://127.0.0.1:#{serverPort}"
    end

    $client = HTTPClient.new    
  end

  config.after(:suite) do
    if !$docker
      serverContainer.stop
      serverContainer.delete
      pgContainer.stop
      pgContainer.delete
      network.delete
    end
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
      "POSTGRES_DB=gotickyourself",
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
