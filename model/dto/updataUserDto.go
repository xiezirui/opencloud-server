package dto

type UpNameDto struct {
	ID   uint
	Name string
}

type UpPasswordDto struct {
	ID       uint
	Password string
}

type CheckOldPassowrdDto struct {
	ID          uint
	OldPassword string
}
