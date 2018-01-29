#/bin/bash

DIR=$PWD

cd monitor/src

#windows quirk
export MSYS_NO_PATHCONV=1

docker build . -t cachet-monitor

#docker rm --force cachet-monitor
echo "starting monitor..."
docker run -it --rm \
--link cachet \
--link google-hello \
-e CACHET_SERVER_URL=http://cachet:8000/api/v1 \
-e CACHET_TOKEN=zlvNVV0VxKkuRL0WM5ww \
-v $PWD/config:/app/Configs cachet-monitor