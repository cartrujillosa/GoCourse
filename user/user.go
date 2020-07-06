package user

type User struct {
	ID      string   `json:"id,omitempty"`
	Lessons []string `json:"lessons,omitempty"`
}

var usersInMemory []User

func GetUsers() *[]User {
	return &usersInMemory
}

func AddUser(id string, lessons []string) {
	newUser := User{
		ID:      id,
		Lessons: lessons,
	}
	usersInMemory = append(usersInMemory, newUser)
	return
}

func UpdateUser(i int, lessons []string) {
	usersInMemory[i].Lessons = append(usersInMemory[i].Lessons, lessons...)
	return
}

func UserExists(id string) *int {
	for i, user := range usersInMemory {
		if user.ID == id {
			return &i
		}
	}
	return nil
}
