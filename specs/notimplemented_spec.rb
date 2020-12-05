require 'rspec'
require 'httpclient'
require 'json'

describe "GET / :: Root is not yet implemented" do
  before :all do
    @response = $client.get "#{$endpoint}/"
  end

  it "returns 501 NOT IMPLEMENTED" do
    expect(@response.status).to eq(501)
  end
end