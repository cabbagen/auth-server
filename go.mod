module go-gateway

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/go-redis/redis/v7 v7.0.0-beta.5
	github.com/jinzhu/gorm v1.9.12
	github.com/mojocn/base64Captcha v1.3.0
	github.com/sirupsen/logrus v1.4.2
)

replace (
	github.com/go-redis/redis/v7 => /users/xia/documents/project/go/pkg/mod/github.com/go-redis/redis/v7@v7.0.0-beta.5
	github.com/mojocn/base64Captcha => /users/xia/documents/project/go/pkg/mod/github.com/mojocn/base64!captcha@v1.3.0
	github.com/sirupsen/logrus => /users/xia/documents/project/go/pkg/mod/github.com/sirupsen/logrus@v1.4.2
	golang.org/x/crypto => /users/xia/documents/project/go/pkg/mod/github.com/golang/crypto
	golang.org/x/image => /users/xia/documents/project/go/pkg/mod/github.com/golang/image
	golang.org/x/mod => /users/xia/documents/project/go/pkg/mod/github.com/golang/mod
	golang.org/x/net => /users/xia/documents/project/go/pkg/mod/github.com/golang/net
	golang.org/x/sync => /users/xia/documents/project/go/pkg/mod/github.com/golang/sync
	golang.org/x/sys => /users/xia/documents/project/go/pkg/mod/github.com/golang/sys
	golang.org/x/text => /users/xia/documents/project/go/pkg/mod/github.com/golang/text
	golang.org/x/tools => /users/xia/documents/project/go/pkg/mod/github.com/golang/tools
	golang.org/x/xerrors => /users/xia/documents/project/go/pkg/mod/github.com/golang/xerrors
)
