package dto

type UpNameDto struct {
	ID   string
	Name string
}

type UpPasswordDto struct {
	ID       string
	Password string
}

type CheckOldPassowrdDto struct {
	ID          string
	OldPassword string
}
