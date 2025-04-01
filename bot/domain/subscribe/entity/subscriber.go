package entity

type Subscriber struct {
	id    int
	email string
	name  string
}

func NewSubscribe(
	id int,
	email string,
	name string,
) Subscriber {
	return Subscriber{
		id:    id,
		email: email,
		name:  name,
	}
}

func (a Subscriber) ID() int {
	return a.id
}

func (a Subscriber) Email() string {
	return a.email
}

func (a Subscriber) Name() string {
	return a.name
}
