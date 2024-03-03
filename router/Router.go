package router

import (
	"fmt"
	"gomountain/docs"
	"gomountain/service"

	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 跨域問題
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.GetHeader("Origin"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // No content, but the headers are sent
			return
		}

		c.Next()
	}
}

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware()) // 跨域問題
	r.Static("/assets/images", "./assets/images")
	r.Static("/assets/userImages", "./assets/userImages")
	docs.SwaggerInfo.BasePath = ""
	// Public routes
	public := r.Group("/")
	public.GET("/ping", service.GetIndex)
	public.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	public.POST("/user/createUser", service.CreateUser)
	public.POST("/user/findUserByNameAndPwd", service.FindUserByNameAndPwd)
	public.POST("/user/googleSignIn", service.GoogleSignIn)
	public.POST("/user/appleSignIn", service.AppleSignIn)
	public.POST("/user/RefreshToken", service.RefreshToken)
	public.GET("/user/getUserList", service.GetUserList)
	public.GET("/mountainRoad/getMountainRoad", service.GetMountainRoad)
	public.POST("/connect/connectToUs", service.ConnectToUsHandler)
	public.GET("/version/getVersion", service.GetVersion)

	// Private (authenticated) routes
	private := r.Group("/")
	private.Use(JWTAuthMiddleware())
	private.GET("/user/deleteUser", service.DeleteUser)
	private.POST("/user/updateUser", service.UpdateUser)
	private.GET("/user/getUserHistory", service.GetUserHistory)

	private.GET("/post/getPostList", service.GetPostList)
	private.POST("/post/createPost", service.CreatePost)
	private.GET("/post/deletePost", service.DeletePost)

	private.POST("/like/addLike", service.AddLike)

	private.POST("/comment/addComment", service.AddCommentHandler)
	private.GET("/comment/getComment", service.GetComment)

	private.POST("/favorite/addFavoriteMountainRoad", service.AddFavoriteMountainRoad)
	private.GET("/favorite/getFavoriteMountainRoad", service.GetFavoriteMountainRoad)
	private.DELETE("/favorite/delFavoriteMountainRoad", service.DelFavoriteMountainRoad)

	return r
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, BearerSchema) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BearerSchema)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Always check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Return the secret key (this should be read from environment variable or config file)
			// This is just an example and should not be used in production
			return []byte(viper.GetString("tokenSign")), nil
		})

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1, // 0 成功 -1失敗
				"message": err.Error(),
			})
			// c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// You can access the token claims (payload) here
			c.Set("userID", claims["userID"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1, // 0 成功 -1失敗
				"message": "Invalid token",
			})
			// c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}
