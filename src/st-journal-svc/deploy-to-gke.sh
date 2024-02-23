#!/bin/bash

SERVICE_NAME="$1"
PORT="$2"
IMAGE_TAG="$3"
GIT_REF="$4"
PR_NUMBER="$5"

if [[ "$GIT_REF" == "refs/heads/develop" ]] && [[ "$PR_NUMBER" == "" ]]; then
    SUFFIX="-dev"
elif [[ "$PR_NUMBER" != "" ]]; then
    SUFFIX="-pr-$PR_NUMBER"
else
    SUFFIX=""
fi

DEPLOYMENT_FILE="./k8s/${SERVICE_NAME}-deployment.yml"

cp "./k8s/resources.yml" "$DEPLOYMENT_FILE"

sed -i "s|DEPLOYMENT_NAME_PLACEHOLDER|${SERVICE_NAME}${SUFFIX}|g" "$DEPLOYMENT_FILE"
sed -i "s|APP_LABEL_PLACEHOLDER|${SERVICE_NAME}${SUFFIX}|g" "$DEPLOYMENT_FILE"
sed -i "s|CONTAINER_NAME_PLACEHOLDER|${SERVICE_NAME}|g" "$DEPLOYMENT_FILE"
sed -i "s|IMAGE_PLACEHOLDER|${IMAGE_TAG}|g" "$DEPLOYMENT_FILE"
sed -i "s|CONTAINER_PORT_PLACEHOLDER|${PORT}|g" "$DEPLOYMENT_FILE"
sed -i "s|SERVICE_PORT_PLACEHOLDER|${PORT}|g" "$DEPLOYMENT_FILE"
sed -i "s|SERVICE_NAME_PLACEHOLDER|${SERVICE_NAME}${SUFFIX}|g" "$DEPLOYMENT_FILE"

kubectl apply -f "$DEPLOYMENT_FILE" --namespace=${NAMESPACE}

kubectl rollout status deployment/${SERVICE_NAME}${SUFFIX} --namespace=${NAMESPACE}


DEPLOYMENT_FILE="${GITHUB_WORKSPACE}/k8s/${{ matrix.service }}-resources.yml"

if [ "${{ github.ref }}" == "refs/heads/develop" ] && [ "${{ github.event_name }}" != "pull_request" ]; then
  SUFFIX="-dev"
elif [ "${{ github.event_name }}" == "pull_request" ] && [ "${{ github.base_ref }}" != "develop" ]; then
  SUFFIX="-pr-${{ github.event.pull_request.number }}"
else
  SUFFIX=""
fi

cp "${GITHUB_WORKSPACE}/k8s/resources.yml" "$DEPLOYMENT_FILE"
sed -i "s|DEPLOYMENT_NAME_PLACEHOLDER|${{ matrix.service }}$SUFFIX|g" "$DEPLOYMENT_FILE"
sed -i "s|APP_LABEL_PLACEHOLDER|${{ matrix.service }}$SUFFIX|g" "$DEPLOYMENT_FILE"
sed -i "s|CONTAINER_NAME_PLACEHOLDER|${{ matrix.service }}|g" "$DEPLOYMENT_FILE"
sed -i "s|IMAGE_PLACEHOLDER|${IMAGE_TAG}|g" "$DEPLOYMENT_FILE"
sed -i "s|CONTAINER_PORT_PLACEHOLDER|${{ matrix.port }}|g" "$DEPLOYMENT_FILE"
sed -i "s|SERVICE_NAME_PLACEHOLDER|${{ matrix.service }}$SUFFIX|g" "$DEPLOYMENT_FILE"
sed -i "s|SERVICE_PORT_PLACEHOLDER|${{ matrix.port }}|g" "$DEPLOYMENT_FILE"
kubectl apply -f "$DEPLOYMENT_FILE" --namespace=${{ env.NAMESPACE }}
kubectl rollout status deployment/${{ matrix.service }}$SUFFIX --namespace=${{ env.NAMESPACE }}