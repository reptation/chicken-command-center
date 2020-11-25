# Chicken Command Center
Monitor and control a remote coop using Django running on a Raspberry Pi

## Creating the Development Environment
Install python, create a virtualenv, clone the repo, and install dependencies:
```
pyenv install 3.9.0b3
pyenv virtualenv 3.9.0b3 ccc
pyenv activate ccc
git clone https://github.com/reptation/chicken-command-center.git
python -m pip install -r requirements.txt
```

## Setting Up the Raspberry Pi
Enable the camera module:
```
sudo raspi-config
# Camera module is under "Interfacing Options"
sudo reboot
```
