package user

//User data struct
type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Approved bool   `json:"approved"`
}

//Users is an array of User's
type Users struct {
	// slice is go's version of a dynamic array
	Users []User
}

// NewUser is the Constructor for user
func NewUser(name string) *User {
	u := User{Name: name}
	return &u
}

// SetAge age for User
func (u *User) SetAge(age int) *User {
	u.Age = age
	return u
}

// Approve sets if the user is approved
func (u *User) Approve(approved bool) *User {
	u.Approved = approved
	return u
}

// Append adds a user struct to the Users slice
func (us *Users) Append(u User) *Users {
	if us.Users == nil {
		us.Users = []User{}
	}
	us.Users = append(us.Users, u)
	return us
}
