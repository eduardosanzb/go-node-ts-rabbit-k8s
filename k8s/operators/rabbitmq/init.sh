#!/bin/bash
echo "Installing rabbitmq operator"
kubectl apply -f "https://github.com/rabbitmq/cluster-operator/releases/latest/download/cluster-operator.yml"
sleep 10

BASEDIR=$(dirname $0)
echo "\nCreate configmap for rabbitmq definitions"
kubectl create configmap rabbit-definitions --from-file=$BASEDIR/definitions.json
sleep 10

echo "Create the cluster for the rabbitmq"
kubectl apply -f $BASEDIR/cluster.yaml
