FROM metabase/metabase

RUN curl -L -o /plugins/csv-metabase-driver.jar \
    https://github.com/Markenson/csv-metabase-driver/releases/download/v1.3.1/csv.metabase-driver.jar

RUN chmod 644 /plugins/csv-metabase-driver.jar

RUN apk upgrade
