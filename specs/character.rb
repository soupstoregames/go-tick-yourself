require 'rspec'
require 'httpclient'
require 'json'

describe "GET /character" do
  context "with a fresh character" do 
    before :all do
      $pg.exec_params("INSERT INTO characters (id, balance, reputation) VALUES ($1, $2, $3)", [1, 0, 0])
      @response = $client.get "#{$endpoint}/character"
    end

    after :all do
      $pg.exec("TRUNCATE TABLE characters")
    end

    it "returns 200 OK" do
      expect(@response.status).to eq(200)
    end

    it "returns a the character ID" do
      jsonResponse = JSON.parse(@response.body)
      expect(jsonResponse["id"]).to eq 1
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
      $pg.exec_params("INSERT INTO characters (id, balance, reputation) VALUES ($1, $2, $3)", [1, 10000, -100])
      @response = $client.get "#{$endpoint}/character"
    end

    after :all do
      $pg.exec("TRUNCATE TABLE characters")
    end

    it "returns 200 OK" do
      expect(@response.status).to eq(200)
    end

    it "returns a the character ID" do
      jsonResponse = JSON.parse(@response.body)
      expect(jsonResponse["id"]).to eq 1
    end

    it "returns a zero balance" do
      jsonResponse = JSON.parse(@response.body)
      expect(jsonResponse["balance"]).to eq 10000
    end

    it "returns a neutral reputation" do
      jsonResponse = JSON.parse(@response.body)
      expect(jsonResponse["reputation"]).to eq -100
    end
  end
end