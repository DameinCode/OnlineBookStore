package store

import (
	"sort"
	"sync"

	"github.com/DameinCode/Golang-spring-2022/Uni/Final/OnlineBookStore/Book/rpc"
	"github.com/pkg/errors"
)

type BookStore interface {
	// Retrieve a book by ID from the store.
	Book(id string) (*bookrpc.Book, error)
	// Retrieve all books in the store, sorted by id asc.
	Books() []*bookrpc.Book
}

type bookStore struct {
	lock sync.RWMutex
	m    map[string]*bookrpc.Book
}

func NewBookStore() BookStore {
	return &bookStore{
		m: map[string]*bookrpc.Book{
			"First-book": {
				Id:    "First-book",
				Type: "Roman",
				Name:  "Love hurts",
				Author: "Gulbina Bolatbekkyzy",
				Year: 2022,
				Annotation: "Okay. Love hurts",
				Content: "Love Hurts is a song written and composed by the American songwriter Boudleaux Bryant. First recorded by the Everly Brothers in July 1960, the song is most well known from the 1974 international hit version by Scottish hard rock band Nazareth and 1975 Top 5 hit in the UK by English singer Jim Capaldi.",
				Price: 25.0,
			},
			"Second-book": {
				Id:    "Second-book",
				Type: "Roman",
				Name:  "Love heals",
				Author: "Nodoby Nobody",
				Year: 2023,
				Annotation: "Okay. Love heals. Dont beleave to Gulbina",
				Content: "LOVE HEALS follows the journey of Dana, a chronic pain sufferer in search of healing. Her partner, Krisanna, is a filmmaker, and together they travel the country to understand how this ancient principle has helped so many heal and to see what’s possible for those experiencing these practices for the first time. What does it mean to truly heal? And does this wisdom have the power to change our world?",
				Price: 40.0,
			},
			"Third-book": {
				Id:    "Third-book",
				Type: "Science",
				Name:  "About science",
				Author: "Adolf Hm",
				Year: 1999,
				Annotation: "Okay. Love is enough. Lets talk about science",
				Content: "Science (from Latin scientia knowledge) is a systematic enterprise that builds and organizes knowledge in the form of testable explanations and predictions about the universe.",
				Price: 50.0,
			},
			"Fourth-book": {
				Id:    "Fourth-book",
				Type: "Programming",
				Name:  "Golang",
				Author: "Jaime Enrique Garcia Lopez",
				Year: 2013,
				Annotation: "Build fast, reliable, and efficient software at scale",
				Content: "At the time, no single team member knew Go, but within a month, everyone was writing in Go and we were building out the endpoints. It was the flexibility, how easy it was to use, and the really cool concept behind Go (how Go handles native concurrency, garbage collection, and of course safety+speed.) that helped engage us during the build. Also, who can beat that cute mascot!",
				Price: 40.0,
			},
			"Fifth-book": {
				Id:    "Fifth-book",
				Type: "Story",
				Name:  "A Bronx Tale",
				Author: "Jane Rosentha",
				Year: 1993,
				Annotation: "In 1960, Lorenzo works as an MTA bus driver in Belmont, a working-class Italian-American neighborhood in The Bronx, with his wife Rosina and their nine-year-old son Calogero. Calogero becomes enamored with the criminal life and Mafia presence in his neighborhood, led by Sonny. ",
				Content: "Two men are shot in the head, and a lot of blood is shown. There's a bar fight where motorcycle bikers are beaten by gangsters with baseball bats and bottles. There are kids selling guns on the street. there is a lot of blood. illed bottles, a car explodes and the white kids all burn to death: there are some very gruesome close-ups of burned flesh and limbs.",
				Price: 30.0,
			},
		},
	}
}

// Retrieve all books in the store, sorted by id asc.
func (s *bookStore) Products() []*bookrpc.Book {
	s.lock.RLock()
	defer s.lock.RUnlock()

	pts := make([]*bookrpc.Book, 0, len(s.m))
	for _, o := range s.m {
		pts = append(pts, o)
	}

	sort.Slice(pts, func(i, j int) bool {
		return pts[i].GetId() < pts[j].GetId()
	})

	return pts
}

// Retrieve a books by ID from the store.
func (s *bookStore) Book(id string) (*bookrpc.Book, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	p, ok := s.m[id]
	if !ok {
		return nil, errors.New("book not found")
	}

	return p, nil
}


// Retrieve a books by ID from the store.
func (s *bookStore) Book(id string) (*bookrpc.Book, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	p, ok := s.m[id]
	if !ok {
		return nil, errors.New("book not found")
	}

	return p, nil
}

// Create a book 

// func (s *bookStore) Book(id string, typ string, name string, author string, year int32, annotation string, content string, price float32) (*bookrpc.Book, error) {
// 	s.lock.RLock()
// 	defer s.lock.RUnlock()

// 	p, ok := s.m[id]
// 	if !ok {
// 		return nil, errors.New("book not found")
// 	}

// 	return p, nil
// }