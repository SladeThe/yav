package tests

import (
	"github.com/asaskevich/govalidator"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.uber.org/multierr"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/vbytes"
	"github.com/SladeThe/yav/vmap"
	"github.com/SladeThe/yav/vnumeric"
	"github.com/SladeThe/yav/vstring"
	"github.com/SladeThe/yav/vstruct"
)

var (
	ozzoContainsLowerAlpha = ozzo.NewStringRule(govalidator.HasLowerCase, "must contain lower alpha")
	ozzoContainsUpperAlpha = ozzo.NewStringRule(govalidator.HasUpperCase, "must contain upper alpha")
)

type Size struct {
	Width  uint16 `json:"width" validate:"required,min=32,max=512"`
	Height uint16 `json:"height" validate:"required,gt=31,lte=512"`
}

func (s Size) Validate() error {
	return multierr.Combine(
		yav.Chain(
			"width", s.Width,
			vnumeric.Required[uint16],
			vnumeric.Min[uint16](32),
			vnumeric.Max[uint16](512),
		),
		yav.Chain(
			"height", s.Height,
			vnumeric.Required[uint16],
			vnumeric.GreaterThan[uint16](31),
			vnumeric.LessThanOrEqual[uint16](512),
		),
	)
}

func (s Size) OzzoValidate() error {
	return ozzo.ValidateStruct(
		&s,
		ozzo.Field(
			&s.Width,
			ozzo.Required,
			ozzo.Min(32),
			ozzo.Max(512),
		),
		ozzo.Field(
			&s.Height,
			ozzo.Required,
			ozzo.Min(32),
			ozzo.Max(512),
		),
	)
}

type Account struct {
	ID string `json:"id" validate:"required,uuid"`

	Login    string `json:"login" validate:"required,min=4,max=20,alphanum,lowercase"`
	Password string `json:"password" validate:"required_with=Login,omitempty,min=8,max=32,contains_lower_alpha,contains_upper_alpha,contains_digit,contains_special_character,excludes_whitespace,text"`

	Email string `json:"email" validate:"required,min=6,max=100,email,lowercase"`
	Phone string `json:"phone" validate:"required,min=8,max=16,e164"`

	Age     uint8           `json:"age" validate:"omitempty,gte=18,lt=120"`
	Avatars map[Size][]byte `json:"avatars" validate:"omitempty,min=3,max=10,dive,keys,endkeys,required"`

	Secret    string `json:"secret" validate:"required,eq=secure"`
	PromoCode string `json:"promoCode" validate:"omitempty,oneof=BlackFriday2022 BlackFriday2023"`

	FirstName string `json:"firstName" validate:"omitempty,min=2,max=30,alpha,starts_with_upper_alpha,ends_with_lower_alpha"`
	LastName  string `json:"lastName" validate:"omitempty,min=2,max=30,alpha,starts_with_upper_alpha,ends_with_lower_alpha"`

	DisplayName string `json:"displayName" validate:"required_without_all=FirstName LastName,omitempty,min=2,max=50,title,alpha,uppercase"`
}

func (a Account) Validate() error {
	return multierr.Combine(
		yav.Chain(
			"id", a.ID,
			vstring.Required,
			vstring.UUID,
		),
		yav.Chain(
			"login", a.Login,
			vstring.Required,
			vstring.Between(4, 20),
			vstring.Alphanumeric,
			vstring.Lowercase,
		),
		yav.Chain(
			"password", a.Password,
			vstring.RequiredWithAny("Login", yav.RequiredWithAny().String(a.Login)),
			vstring.OmitEmpty,
			vstring.Between(8, 32),
			vstring.ContainsLowerAlpha,
			vstring.ContainsUpperAlpha,
			vstring.ContainsDigit,
			vstring.ContainsSpecialCharacter,
			vstring.ExcludesWhitespace,
			vstring.Text,
		),
		yav.Chain(
			"email", a.Email,
			vstring.Required,
			vstring.Between(6, 100),
			vstring.Email,
			vstring.Lowercase,
		),
		yav.Chain(
			"phone", a.Phone,
			vstring.Required,
			vstring.Between(8, 16),
			vstring.E164,
		),
		yav.Chain(
			"age", a.Age,
			vnumeric.OmitEmpty[uint8],
			vnumeric.GreaterThanOrEqual[uint8](18),
			vnumeric.LessThan[uint8](120),
		),
		yav.Chain(
			"avatars", a.Avatars,
			vmap.OmitEmpty[map[Size][]byte],
			vmap.Between[map[Size][]byte](3, 10),
			vmap.Keys[map[Size][]byte](
				vstruct.Validate[Size],
			),
			vmap.Values[map[Size][]byte](
				vbytes.Required,
			),
		),
		yav.Chain(
			"secret", a.Secret,
			vstring.Required,
			vstring.Equal("secure"),
		),
		yav.Chain(
			"promoCode", a.PromoCode,
			vstring.OmitEmpty,
			vstring.OneOf("BlackFriday2022", "BlackFriday2023"),
		),
		yav.Chain(
			"firstName", a.FirstName,
			vstring.OmitEmpty,
			vstring.Between(2, 30),
			vstring.Alpha,
			vstring.StartsWithUpperAlpha,
			vstring.EndsWithLowerAlpha,
		),
		yav.Chain(
			"lastName", a.LastName,
			vstring.OmitEmpty,
			vstring.Between(2, 30),
			vstring.Alpha,
			vstring.StartsWithUpperAlpha,
			vstring.EndsWithLowerAlpha,
		),
		yav.Chain(
			"displayName", a.DisplayName,
			vstring.RequiredWithoutAll(
				"FirstName LastName",
				yav.RequiredWithoutAll().String(a.FirstName).String(a.LastName),
			),
			vstring.OmitEmpty,
			vstring.Between(2, 50),
			vstring.Title,
			vstring.Alpha,
			vstring.Uppercase,
		),
	)
}

func (a Account) OzzoValidate() error {
	return ozzo.ValidateStruct(
		&a,
		ozzo.Field(
			&a.ID,
			ozzo.Required,
			is.UUID,
		),
		ozzo.Field(
			&a.Login,
			ozzo.Required,
			ozzo.Length(4, 20),
			is.Alphanumeric,
			is.LowerCase,
		),
		ozzo.Field(
			&a.Password,
			ozzo.Required,
			ozzo.When(ozzo.IsEmpty(a.Password), ozzo.Skip),
			ozzo.Length(8, 32),
			ozzoContainsLowerAlpha,
			ozzoContainsUpperAlpha,
		),
		ozzo.Field(
			&a.Email,
			ozzo.Required,
			ozzo.Length(6, 100),
			is.EmailFormat,
			is.LowerCase,
		),
		ozzo.Field(
			&a.Phone,
			ozzo.Required,
			ozzo.Length(8, 16),
			is.E164,
		),
		ozzo.Field(
			&a.Age,
			ozzo.When(ozzo.IsEmpty(a.Age), ozzo.Skip),
			ozzo.Min(18),
			ozzo.Max(119),
		),
		ozzo.Field(
			&a.Avatars,
			ozzo.When(ozzo.IsEmpty(a.Avatars), ozzo.Skip),
			ozzo.Min(3),
			ozzo.Max(10),
		),
		ozzo.Field(
			&a.Secret,
			ozzo.Required,
			ozzo.In("secure"),
		),
		ozzo.Field(
			&a.PromoCode,
			ozzo.When(ozzo.IsEmpty(a.PromoCode), ozzo.Skip),
			ozzo.In("BlackFriday2022", "BlackFriday2023"),
		),
		ozzo.Field(
			&a.FirstName,
			ozzo.When(ozzo.IsEmpty(a.FirstName), ozzo.Skip),
			ozzo.Length(2, 30),
			is.Alpha,
		),
		ozzo.Field(
			&a.LastName,
			ozzo.When(ozzo.IsEmpty(a.LastName), ozzo.Skip),
			ozzo.Length(2, 30),
			is.Alpha,
		),
		ozzo.Field(
			&a.DisplayName,
			ozzo.When(ozzo.IsEmpty(a.DisplayName), ozzo.Skip),
			ozzo.Length(2, 50),
			is.Alpha,
			is.UpperCase,
		),
	)
}

func ValidAccount() Account {
	return Account{
		ID:          "6a310c88-4698-4807-9578-f1f054a8b4ca",
		Login:       "yav123",
		Password:    "DasPassword#123",
		Email:       "yav+123@yav.yav",
		Phone:       "+1234567890",
		Secret:      "secure",
		DisplayName: "YAV",
	}
}
