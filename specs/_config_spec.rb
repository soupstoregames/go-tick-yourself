require "httpclient"
require "docker"
require 'json'
require 'pg'

Docker.url='tcp://127.0.0.1:2375'

RSpec.configure do |config|
  $serverContainer = nil

  config.before(:suite) do
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
#     $serverContainer.stop
  end
end

def create_server()
  return Docker::Container.create(
    "Image" => "soupstoregames/go-tick-yourself:dev",
    "HostConfig" => {
      "PortBindings" => {
        "8080/tcp" => [{}]
      },
    }
  )
end
