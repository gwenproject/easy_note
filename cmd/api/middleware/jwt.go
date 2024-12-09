package middleware

import (
    "context"
    "time"

    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/protocol/consts"
    "github.com/gwen0x4c3/easy_note/cmd/api/rpc"
    "github.com/gwen0x4c3/easy_note/kitex_gen/kuser"
    "github.com/gwen0x4c3/easy_note/pkg/constants"
    "github.com/gwen0x4c3/easy_note/pkg/errno"
    "github.com/hertz-contrib/jwt"
)

var AuthMiddleware *jwt.HertzJWTMiddleware

func init() {
    AuthMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
        Key:        []byte(constants.SecretKey),
        Timeout:    time.Hour,
        MaxRefresh: time.Hour * 24,
        PayloadFunc: func(data interface{}) jwt.MapClaims {
            if v, ok := data.(int64); ok {
                return jwt.MapClaims{
                    constants.IdentityKey: v,
                }
            }
            return jwt.MapClaims{}
        },
        HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
            switch e := e.(type) {
            case errno.ErrNo:
                return e.ErrMsg
            default:
                return e.Error()
            }
        },
        LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
            c.JSON(consts.StatusOK, map[string]interface{}{
                "code":   errno.SuccessCode,
                "token":  token,
                "expire": expire.Format(time.RFC3339),
            })
        },
        Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
            c.JSON(code, map[string]interface{}{
                "code":    errno.AuthorizationFailedErrCode,
                "message": message,
            })
        },
        Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
            var req kuser.CheckUserRequest
            if err := c.Bind(&req); err != nil {
                return "", jwt.ErrMissingLoginValues
            }

            if (len(req.UserName) == 0) || (len(req.Password) == 0) {
                return "", jwt.ErrMissingLoginValues
            }

            return rpc.CheckUser(context.Background(), &req)
        },
        IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
            claims := jwt.ExtractClaims(ctx, c)
            return int64(claims[constants.IdentityKey].(float64))
        },
        TokenLookup:   "header: Authorization, query: token, cookie: jwt",
        TokenHeadName: "Bearer",
        TimeFunc:      time.Now,
    })
}
