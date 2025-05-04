package repository

import (
	"log"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"bot/infrastructure/storage/database"
	"bot/lib"

	"bot/domain/data_warehouse/article/infrastructure/model"
)

func generateFakeArticle() []model.Article {
	data := []model.Article{}

	for i := range 10 {
		t := model.Article{
			ID:           i + 1,
			Mode:         1,
			Title:        gofakeit.Sentence(6),
			Sapo:         gofakeit.Paragraph(3, 4, 8, "\n"),
			Content:      gofakeit.Paragraph(3, 4, 8, "\n"),
			Image:        gofakeit.URL(),
			Origin:       gofakeit.URL(),
			ArticleHubID: 1,
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

	config := lib.GetEnvConfigMap("db", "test")
	s.storage = database.Connect(config)
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

	s.storage.Exec("truncate article;")
	for _, t := range data {
		s.storage.Create(t)
	}

	s.data = data
}

// run after each test
func (s *articleTestSuite) TearDownTest() {
	log.Println("TearDownTest()")
	s.storage.Exec("truncate article;")
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

		assert.Equal(t.T(), want, got, "should be 10 records")
	})
}

func TestArticleTestSuite(t *testing.T) {
	suite.Run(t, new(articleTestSuite))
}

func TestCastToArticle(t *testing.T) {
	e := model.Article{
		ID:           1,
		Mode:         1,
		Title:        "title",
		Sapo:         gofakeit.Paragraph(3, 4, 8, "\n"),
		Content:      gofakeit.Paragraph(3, 4, 8, "\n"),
		Image:        gofakeit.URL(),
		Origin:       gofakeit.URL(),
		ArticleHubID: 1,
	}

	actual := e.Title
	expected := "title"

	assert.Equal(t, expected, actual)
}
