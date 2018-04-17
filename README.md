# apns
A Golang package for sending Apple Push 

This is push demo, use github.com/sideshow/apns2 library.

How to create certificate.pem and key.unencrypted.pem ?

##Use

* 1.Download cer in https://developer.apple.com/account/ios/certificate/
* 2.openssl x509 -inform der -in aps_development.cer -out certificate.pem
* 3.openssl pkcs12 -clcerts -nokeys -out cert.pem -in cert.p12 -passin pass:P12_PASS
* 4.openssl pkcs12 -nocerts  -out key.pem -in cert.p12 -passin pass:P12_PASS -passout pass:TMP_PASS
* 5.openssl rsa -in key.pem -out key.unencrypted.pem -passin pass:TMP_PASS
* tips:P12_PASS is your cert.p12 password, TMP_PASS is tmp password