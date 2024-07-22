#!/bin/bash
sudo apt update
tar -C /tmp/ -xzf <(curl -L https://dl.pstmn.io/download/latest/linux64) && sudo mv /tmp/Postman /opt/
sudo apt install curl
curl -o- "https://dl-cli.pstmn.io/install/linux64.sh" | sh
sudo apt install nodejs npm -y
sudo npm install -g newman
sudo npm install -g newman-reporter-htmlextra

