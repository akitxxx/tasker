FROM ruby:2.6.3
RUN apt-get update
RUN apt-get -y install curl gnupg

RUN curl -sL https://deb.nodesource.com/setup_11.x  | bash -
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
RUN apt-get update -qq && apt-get -y install nodejs yarn

RUN mkdir /tasker
WORKDIR /tasker
COPY Gemfile /tasker/Gemfile
COPY Gemfile.lock /tasker/Gemfile.lock
RUN bundle install
COPY . /tasker

COPY entrypoint.sh /usr/bin/
RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]

CMD ["rails", "server", "-b", "0.0.0.0"]
