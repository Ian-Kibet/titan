#!/usr/bin/sudo bash

set -eu

[ "$UID" -eq 0 ] || { echo "Root permission is required"; exit 1;}


# TODO: Check if firecracker is installed before proceeding 

export PATH=$PATH:/usr/bin
if ! [ -x "$(command -v firecracker)" ]; then
    echo "firecracker is already installed in PATH. To check installation run 'which firecracker'"
    exit
fi


# me="$(whoami)"

# # Checking that the script is running as root.
# if [ "$(id -nu)" != "root" ]; then
#     sudo --reset-timestamp

#     # creating a dialog box to enter user password 
#     pswd=$(whiptail --title "Primelab CPU Authentication" \
#     --passwordbox "To run this script, administrative privilege is \
#     required. \n[sudo] Password for user $me:" 14 52 3>&2 2>&1 1>&3-)
    
#     # executing the script with the password entered by user
#     exec sudo --stdin --prompt '' "$0" "$@" <<< "$pswd"

#     # if password is wrong it will return the status code 1
#     exit 1

# fi

release_url="https://github.com/firecracker-microvm/firecracker/releases"
latest=$(basename $(curl -fsSLI -o /dev/null -w  %{url_effective} ${release_url}/latest))
arch=`uname -m`
version=${latest}-$(uname -m)


curl -L ${release_url}/download/${latest}/firecracker-${latest}-${arch}.tgz \
| tar -xz

mv release-$version/firecracker-$version /usr/bin/firecracker
rm -rf release-$version

echo "Successfully installed firecracker in PATH :)"