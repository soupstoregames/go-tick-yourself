require 'rspec'
require 'httpclient'
require 'json'

describe "GET /character" do
  context "with a fresh character" do 
    before :all do
      @response = $client.get "#{$endpoint}/character"
    end

    it "returns 200 OK" do
      expect(@response.status).to eq(200)
    end

    it "returns a zero balance" do
      jsonResponse = JSON.parse(@response.body)
      expect(jsonResponse["balance"]).to eq 0
    end

    it "returns a neutral reputation" do
      jsonResponse = JSON.parse(@response.body)
      expect(jsonResponse["reputation"]).to eq 0
    end
  end

  context "with a non-default character" do
    before :all do
      # TODO: set up database with character info
      @response = $client.get "#{$endpoint}/character"
    end

    it "returns 200 OK" do
      expect(@response.status).to eq(200)
    end

    # TODO: check result matches what we put in database
  end
end