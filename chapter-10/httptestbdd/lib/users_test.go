package lib_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/shijuvar/go-web/chapter-10/httptestbdd/lib"
)

var _ = Describe("Users", func() {
	userRepository := NewFakeUserRepo()
	var r *mux.Router
	var w *httptest.ResponseRecorder

	BeforeEach(func() {
		r = mux.NewRouter() 
	})

	Describe("Get Users", func() {
		Context("Get all Users", func() {
			 //providing mocked data of 3 users 	
			It("should get list of Users", func() {
				r.Handle("/users", GetUsers(userRepository)).Methods("GET")
				req, err := http.NewRequest("GET", "/users", nil)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(200))
				var users []User
				json.Unmarshal(w.Body.Bytes(), &users)
				//Verifying mocked data of 3 users 
				Expect(len(users)).To(Equal(3))
			})
		})
	})

	Describe("Post a new User", func() {
		Context("Provide a valid User data", func() {
			It("should create a new User and get HTTP Status: 201", func() {
				r.Handle("/users", CreateUser(userRepository)).Methods("POST")
				userJson := `{"firstname": "Alex", "lastname": "John", "email": "alex@xyz.com"}`

				req, err := http.NewRequest(
					"POST",
					"/users",
					strings.NewReader(userJson),
				)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(201))
			})
		})
		Context("Provide a User data that contains duplicate email id", func() {
			It("should get HTTP Status: 400", func() {
				r.Handle("/users", CreateUser(userRepository)).Methods("POST")
				userJson := `{"firstname": "Alex", "lastname": "John", "email": "alex@xyz.com"}`

				req, err := http.NewRequest(
					"POST",
					"/users",
					strings.NewReader(userJson),
				)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(400))
			})
		})
	})
})

type FakeUserRepository struct {
	DataStore []User
}

func (repo *FakeUserRepository) GetAll() []User {

	return repo.DataStore
}
func (repo *FakeUserRepository) Create(user User) error {
	err := repo.Validate(user)
	if err != nil {
		return err
	}
	repo.DataStore = append(repo.DataStore, user)
	return nil
}
func (repo *FakeUserRepository) Validate(user User) error {
	for _, u := range repo.DataStore {
		if u.Email == user.Email {
			return errors.New("The Email is already exists")
		}
	}
	return nil
}
func NewFakeUserRepo() *FakeUserRepository {
	return &FakeUserRepository{
		DataStore: []User{
			User{"Shiju", "Varghese", "shiju@xyz.com"},
			User{"Rosmi", "Shiju", "rose@xyz.com"},
			User{"Irene", "Rose", "irene@xyz.com"},
		},
	}
}
