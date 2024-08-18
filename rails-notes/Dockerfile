FROM ruby:3.3.0

RUN apt-get update && \
    apt-get install -y build-essential git libvips pkg-config

WORKDIR /workspace

COPY Gemfile .
COPY Gemfile.lock .

RUN bundle install && \
    bundle exec bootsnap precompile --gemfile

COPY . .

EXPOSE 3000

CMD ["./bin/rails", "server"]
