printf "***Install num2word\n"
cd src/ru/mail/noknod/num2word
go install
#printf "\n***Test num2word\n"
#go test

printf "\n***Install main\n"
cd ../../../../../
go install main

printf "\n***Run main\n"
bin/main
