package models

import ()

type Tasks struct {
	Identifier string `bson:"identifier" json:"identifier"`
	Tasks      []Task `bson:"tasks" json:"tasks"`
}

func NewTasksCollection(identifier string, tasks []Task) *Tasks {
	return &Tasks{Identifier: identifier, Tasks: tasks}
}

type Task struct {
	Day  string `bson:"day" json:"day"`
	Task string `bson:"task" json:"task"`
}

func NewTask(day string, task string) *Task {
	return &Task{Day: day, Task: task}
}

type Post struct {
	Post      string `bson:"post" json:"post"`
	Author    string `bson:"author" json:"author"`
	Responses []Post `json:"responses"`
	Likes     int    `json:"likes"`
	Shares    int    `json:"shares"`
	Career    string `json:"career"`
}

func NewPost(post string, author string, responses []Post, likes int, shares int, career string) *Post {
	return &Post{Post: post, Author: author, Responses: responses, Likes: likes, Shares: shares, Career: career}
}

type Profile struct {
	Identifier string `bson:"identifier" json:"identifier"`
	UserName   string `bson:"username" json:"username"`
	Career     string `bson:"career" json:"career"`
	Posts      []Post `bson:"posts" json:"posts"`
	Likes      []Post `bson:"likes" json:"likes"`
	Shares     []Post `bson:"shared_posts" json:"shared_posts"`
}

func NewProfile(identifier string, username string, career string) *Profile {
	return &Profile{Identifier: identifier, UserName: username, Career: career}
}

// type Curricular struct {
// 	pw.CurricularSubject
// 	Identifier string `bson:"identifier" json:"identifier"`
// }
//
// type Kardex struct {
// 	pw.Subject
// 	Identifier string `bson:"identifier" json:"identifier"`
// }
//
// type Moodle struct {
// 	pw.Assigment
// 	ControlNumber string `bson:"control_number" json:"control_number"`
// }
//
// type ClassRoom struct {
// 	pw.Assigment
// 	ControlNumber string `bson:"control_number" json:"control_number"`
// }
//
// type Activites struct {
// 	pw.Assigment
// 	ControlNumber string `bson:"control_number" json:"control_number"`
// }
//
// type Calendar struct {
// 	Date          Date   `bson:"date" json:"date"`
// 	Description   string `bson:"description" json:"description"`
// 	IsDone        bool   `bson:"is_done" json:"is_done"`
// 	ControlNumber string `bson:"control_number" json:"control_number"`
// }
//
// type Date struct {
// 	Day   string `bson:"day" json:"day"`
// 	Month string `bson:"month" json:"month"`
// 	Year  string `bson:"year" json:"year"`
// }
