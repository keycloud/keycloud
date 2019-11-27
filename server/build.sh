path=`pwd`
sudo -i -- sh -c "cd ${path} && export GOPATH=${path} && go get ."
sudo -i -- sh -c "cd ${path} && export GOPATH=${path} && go build -o server ."
echo "GO BUILD END"
