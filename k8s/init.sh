#!/bin/bash

BASEDIR=$(dirname $0)

echo "About to create a cluster using kind"
sh $BASEDIR/cluster/kind.sh sleep 3

echo "About to init a postgres operator with custom values. (Our DB)"
sh $BASEDIR/operators/postgres/init.sh

echo "About to init a rabbitmq operator with custom values. (Our broker)"
sh $BASEDIR/operators/rabbitmq/init.sh

sleep 3
echo "Deploying our app"
kubectl apply -f $BASEDIR/secrets -f $BASEDIR/deployments -f $BASEDIR/services -f $BASEDIR/ingress

sleep 5
echo "Seeding the db"
kubectl apply -f $BASEDIR/jobs
