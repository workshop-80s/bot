package repository

import (
	"log"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"bot/domain/subscribe/infrastructure/model"
	"bot/infrastructure/storage/database"
)

func generateFakeArticle() []model.Article {
	data := []model.Article{}

	for i := range 10 {
		t := model.Article{
			ID:          i + 1,
			Url:         gofakeit.URL(),
			Title:       gofakeit.Sentence(6),
			Sapo:        gofakeit.Paragraph(3, 4, 8, "\n"),
			OriginalUrl: gofakeit.URL(),
		}

		data = append(data, t)
	}

	return data
}

type articleTestSuite struct {
	suite.Suite
	storage *gorm.DB
	data    []model.Article
}

// run once, before test suite methods
func (s *articleTestSuite) SetupSuite() {
	log.Println("SetupSuite()")

	s.storage = database.Connect()
}

// run once, after test suite methods
func (s *articleTestSuite) TearDownSuite() {
	log.Println("TearDownSuite()")

	database.Disconnect(s.storage)
}

// run before each test
func (s *articleTestSuite) SetupTest() {
	log.Println("SetupTest()")

	data := generateFakeArticle()

	s.storage.Exec("truncate articles;")
	for _, t := range data {
		s.storage.Create(t)
	}

	s.data = data
}

// run after each test
func (s *articleTestSuite) TearDownTest() {
	log.Println("TearDownTest()")
	s.storage.Exec("truncate articles;")
}

// run before each test
func (s *articleTestSuite) BeforeTest(suiteName, testName string) {
	log.Println("BeforeTest()", suiteName, testName)
}

// run after each test
func (s *articleTestSuite) AfterTest(suiteName, testName string) {
	log.Println("AfterTest()", suiteName, testName)
}

func (t *articleTestSuite) TestFind() {
	r := NewArticle(t.storage)

	data := r.Find()

	t.Run("verify number of records", func() {
		got := len(data)
		want := len(t.data)

		assert.Equal(t.T(), want, got, "should be 2 records")
	})
}

func TestArticleTestSuite(t *testing.T) {
	suite.Run(t, new(articleTestSuite))
}

func TestCastToArticle(t *testing.T) {
	m := model.NewArticle(
		1,
		"url",
		"title",
		"sapo",
		"ou",
	)

	e := castToArticle(m)

	actual := e.Title()
	expected := "title"

	assert.Equal(t, expected, actual)
}
