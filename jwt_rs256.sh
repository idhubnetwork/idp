mkdir -p keystore
openssl genrsa -out keystore/id_rsa 4096
openssl rsa -in keystore/id_rsa -pubout -outform PEM -out keystore/id_rsa.pub
cat keystore/id_rsa.pub
cat keystore/id_rsa
