package controllers

import (
	"github.com/aakashverma1124/go-jwt-authentication/database"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	helpers "github.com/aakashverma1124/go-jwt-authentication/helpers"
	"github.com/aakashverma1124/go-jwt-authentication/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user"
var validate = validator.New()

func hashPassword()  {

}

func VerifyPassword() {

}

func SignUp() {

}

func Login() {

}

func GetUsers() {

}

func GetUser() {

}