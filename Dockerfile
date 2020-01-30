FROM ruby:2.5
RUN curl -sL https://deb.nodesource.com/setup_10.x | bash -
RUN apt-get update -yqq \
    && apt-get install -yqq --no-install-recommends \
    apt-utils \
    build-essential \
    nodejs \
    tzdata \
    default-libmysqlclient-dev \
    zlib1g-dev \
    libyaml-dev \
    apt-transport-https \
    libicu-dev \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/app/web
RUN mkdir -p /usr/app/mysql
RUN mkdir -p /usr/app/bundle
ENV BUNDLE_PATH=/usr/app/bundle

RUN gem install bundler
RUN bundle config --global silence_root_warning 1
ADD Gemfile /usr/app/web/Gemfile
ADD Gemfile Gemfile.lock
RUN bundle install --deployment
RUN bundle lock --add-platform x86-mingw32 x86-mswin32 x64-mingw32 java
RUN gem install unicorn