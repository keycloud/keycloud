path=`pwd`
echo ${path}
sudo -i -- sh -c "cd ${path} && export GOPATH=${path} && go get . && go build ."
sudo -i -- sh -c "cd ${path} && export GOPATH=${path} && go build -o server ."
echo "END"
