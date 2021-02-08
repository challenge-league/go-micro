set -euxo pipefail
docker build "$PWD" -t localhost:32000/go-micro:dkozlov
#docker push localhost:32000/go-micro:dkozlov
