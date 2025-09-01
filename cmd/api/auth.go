package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"regexp"
)

var store = sessions.NewCookieStore([]byte("<placeholder-for-now-lol>"))

func validatePassword(password string) error {
	if ok, _ := regexp.MatchString(`^.{8,}$`, password); !ok {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	if ok, _ := regexp.MatchString(`[a-z]`, password); !ok {
		return fmt.Errorf("password must contain a lowercase letter")
	}
	if ok, _ := regexp.MatchString(`[A-Z]`, password); !ok {
		return fmt.Errorf("password must contain an uppercase letter")
	}
	if ok, _ := regexp.MatchString(`\d`, password); !ok {
		return fmt.Errorf("password must contain a digit")
	}

	if ok, _ := regexp.MatchString(`[!@#$%^&*()_\-+={}\[\]:;'"|<>,.?/~]`, password); !ok {
		return fmt.Errorf("password must contain a special character")
	}

	return nil
}

func createUser(username, description, email, password string) error {
	sql := `
	SELECT MAX(uid) FROM userdata
	`
	var uid_nxt int
	err := pool.QueryRow(ctx, sql).Scan(&uid_nxt)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Returned an error while querying : %s", err.Error())
	}
	_, err = mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("Returned error : %s", err.Error())
	}
	if validatePassword(password) != nil {
		return fmt.Errorf("Invalid password")
	}
	pass_bytes, err_2 := bcrypt.GenerateFromPassword([]byte(password), 7)
	if err_2 != nil {
		return fmt.Errorf("Hashing returns : %s", err_2.Error())
	}
	sql_insert := `
	INSERT INTO userdata VALUES ($1,$2,$3,0,0,$4,$5) 
	`
	var uid_scan int
	_, err_1 := pool.Exec(ctx, sql_insert, uid_nxt+1, username, string(pass_bytes), description, email)
	fmt.Println(ctx, sql_insert, uid_nxt+1, username, string(pass_bytes), description, email)
	if err_1 != nil {
		return fmt.Errorf("Error with %s", err_1.Error())
	}

	fmt.Printf("Inserted record with %d", uid_scan)
	return nil
}

func validateUser(email, password string) error {
	sql := `
	SELECT password FROM userdata WHERE email=$1 
	`
	var password_usr string
	err := pool.QueryRow(ctx, sql, email).Scan(&password_usr)
	if err != nil {
		return fmt.Errorf("Unable to fetch user: %s", err.Error())
	}
	if err := bcrypt.CompareHashAndPassword([]byte(password_usr), []byte(password)); err != nil {
		return fmt.Errorf("Wrong password")
	}
	fmt.Println(bcrypt.CompareHashAndPassword([]byte(password_usr), []byte(password)))
	return nil
}
