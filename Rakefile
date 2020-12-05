require 'colorize'
require "rake/clean"
require "rspec/core/rake_task"

serviceName = "go-tick-yourself"

task default: %w[all]

desc 'Create generated code'
task :codegen do
  puts "\nGenerating Go code ...".colorize(:cyan)
 sh("go generate ./...")
end

desc 'Run the unit tests'
task :unittest do
  puts "\nRake: Unit tests ...".colorize(:cyan)
  sh "go test ./..."
end

desc 'Build a local docker image'
task :build do
  puts "\nRake: Building Linux AMD64 ...".colorize(:cyan)
  ENV['GOOS'] = 'linux'
  ENV['GOARCH'] = 'amd64'
  ENV['CGO_ENABLED'] = '0'
  sh "go build -o bin/#{serviceName}-#{ENV['GOOS']}-#{ENV['GOARCH']} main.go"

  puts "\nRake: Building Docker image ...".colorize(:cyan)
  sh "docker build -t soupstoregames/#{serviceName}:dev ."
end

RSpec::Core::RakeTask.new(:spec) do |t|
  puts "\nRake: Verifying specifications ...".colorize(:cyan)
  t.pattern = Dir.glob("specs/**/*.rb")
  t.rspec_opts = "--format documentation"
end

task :all => [:codegen, :unittest, :build, :spec]

CLEAN << "bin"