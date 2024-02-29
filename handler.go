package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	keyIncrement = "transactionIncrement"
)

func addIncrement(val string) (valInt int, err error) {
	expIncrement := 10 * time.Second

	// str to int
	valInt, err = strconv.Atoi(val)
	if err != nil {
		return
	}
	valInt++

	valStr := fmt.Sprintf("%d", valInt)

	// set redis
	err = set(ctx, keyIncrement, valStr, expIncrement)
	if err != nil {
		return
	}

	return
}

func lockHandler(c *gin.Context) {
	id := c.Param("id")

	// // set expiration time
	// expiration := 10 * time.Second
	// key := "transactionLock"

	// // lock with redis only
	// err := lockTrxRedisOnly(key, expiration)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// // lock transaction with redsync
	// mutex, err := lockTrx(key, expiration)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "There are other processes in progress, please try again later"})
	// 	return
	// }

	// get redis
	valExist, err := get(ctx, keyIncrement)
	if err != nil {
		fmt.Println("Error getting value from Redis:", err)
	}
	if valExist == "" {
		valExist = "0"
	}

	fmt.Println(id+". valExist ==>>", valExist)

	// time sleep
	time.Sleep(2 * time.Second)

	// increment
	val, err := addIncrement(valExist)
	if err != nil {
		fmt.Println("Error incrementing value:", err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction locked", "value": val})
	// defer func() {
	// 	// Pastikan untuk membuka kunci saat selesai
	// 	err := unlockTrx(*mutex)
	// 	if err != nil {
	// 		fmt.Println("Error unlocking transaction:", err)
	// 	}
	// }()
}
