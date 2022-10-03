package main

type User struct {
	Name, Lastname	string
	Age				int
	Email, Password	string
}

func (u *User) changeName(name *string) {
	u.Name = *name
}

func (u *User) changeLastname(lastname *string) {
	u.Lastname = *lastname
}

func (u *User) changeEmail(email *string) {
	u.Email = *email
}

func (u *User) changePassword(pass *string) {
	u.Password = *pass
}

func main() {

}