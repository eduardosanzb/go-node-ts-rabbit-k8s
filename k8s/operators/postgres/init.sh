#!/bin/bash

echo "Installing the kubegres operator..."
kubectl apply -f https://raw.githubusercontent.com/reactive-tech/kubegres/v1.15/kubegres.yaml

BASEDIR=$(dirname $0)
echo "Load secrets for postgres"
kubectl apply -f $BASEDIR/secrets.yaml

echo "Create the cluster for the postgres"
kubectl apply -f $BASEDIR/cluster.yaml

