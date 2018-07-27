#!/bin/sh

#INTERVAL=${BILLING_INTERVAL:-60}

#LOG_LEVEL=${BILLING_LOG_LEVEL:-"debug"}

#DB=${BILLING_DB:-"mysql"}
#DB_ADDR=${BILLING_DB_ADDR:-"10.39.0.251:3306"}
#DB_USERNAME=${BILLING_DB_USERNAME:-"tcepaas"}
#DB_PASSWORD=${BILLING_DB_PASSWORD:-"xboU58vQbAbN"}
#DB_DATABASE=${BILLING_DB_DATABASE:-"tenxcloud_2_0"}

#QUEUE=${BILLING_QUEUE:-"kafka"}
#QUEUE_ADDR=${BILLING_QUEUE_ADDR:-"10.39.0.119:9092"}
#QUEUE_TOPIC=${BILLING_QUEUE_TOPIC:-"paas_usage"}
#QUEUE_RETRY=${BILLING_QUEUE_RETRY:-5}
#ORDER_URL=${BILLING_ORDER_URL:-"http://metering-d-multi-cloud-platform.ipaas.enncloud.cn/subscription/createbyproduct"}

#sed -ig 's/${INTERVAL}/'$INTERVAL'/g' config.yaml
#sed -ig 's/${LOG_LEVEL}/'$LOG_LEVEL'/g' config.yaml
#sed -ig 's/${DB}/'$DB'/g' config.yaml
#sed -ig 's/${DB_ADDR}/'$DB_ADDR'/g' config.yaml
#sed -ig 's/${DB_USERNAME}/'$DB_USERNAME'/g' config.yaml
#sed -ig 's/${DB_PASSWORD}/'$DB_PASSWORD'/g' config.yaml
#sed -ig 's/${DB_DATABASE}/'$DB_DATABASE'/g' config.yaml
#sed -ig 's/${QUEUE}/'$QUEUE'/g' config.yaml
#sed -ig 's/${QUEUE_ADDR}/'$QUEUE_ADDR'/g' config.yaml
#sed -ig 's/${QUEUE_TOPIC}/'$QUEUE_TOPIC'/g' config.yaml
#sed -ig 's/${QUEUE_RETRY}/'$QUEUE_RETRY'/g' config.yaml
#sed -ig 's/${ORDER_URL}/'$ORDER_URL'/g' config.yaml

./harbor_upgrade -config config.yaml