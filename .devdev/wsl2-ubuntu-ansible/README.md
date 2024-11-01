# ansible playbook for my wsl2-ubuntu environment.

## Install Ansible & Apply
```bash
apt update
apt install -y ansible unzip
curl -L -o playbook.zip https://gist.github.com/enuesaa/5bf668ebdf61853a9eebad1d80c234b3/archive/main.zip
unzip playbook.zip
cd 5bf668ebdf61853a9eebad1d80c234b3-main
ansible-playbook playbook.yml
```
