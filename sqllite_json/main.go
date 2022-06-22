package main
import (
   "fmt"
   "log"
   "database/sql"
   "encoding/json"
   "errors"
   "github.com/mattn/go-sqlite3"
)

const dbFileName = "bmw_sqllite.db"

var (
    ErrDuplicate    = errors.New("record already exists")
    ErrNotExists    = errors.New("row not exists")
    ErrUpdateFailed = errors.New("update failed")
    ErrDeleteFailed = errors.New("delete failed")
)

type Person struct {
   ID int64 `json:"id"`
   Name string `json:"name"`
   Age int64 `json:"age"`
}

type SQLiteRepository struct {
   db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
   return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) Migrate() error {
   query := `
   CREATE TABLE IF NOT EXISTS people(
      id INTEGER PRIMARY KEY,
      person TEXT NOT NULL
   );
   `
   _, err := r.db.Exec(query)
   return err
}

func (r *SQLiteRepository) Create(person Person) (*Person, error) {
   personJson,_ := json.Marshal(person)
   _, err := r.db.Exec("INSERT INTO people(id,person) values (?,?)",
                         person.ID, string(personJson))

   if err != nil {
        var sqliteErr sqlite3.Error
        if errors.As(err, &sqliteErr) {
            if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
                return nil, ErrDuplicate
            }
        }
        return nil, err
    }

    return &person, nil
}

func (r *SQLiteRepository) All() ([]Person, error) {
    rows, err := r.db.Query("SELECT * FROM people")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var all []Person
    for rows.Next() {
        var person Person
        var id int64
        var personJson string
        if err := rows.Scan(&id, &personJson); err != nil {
            return nil, err
        }
        err = json.Unmarshal([]byte(personJson), &person)
        all = append(all, person)
    }
    return all, nil
}

func main() {
   fmt.Println("*** SQLLITE TEST ***");

   db, err := sql.Open("sqlite3", dbFileName);
   if err != nil {
      log.Fatal(err)
   }

   peopleRepository :=  NewSQLiteRepository(db)

   if err := peopleRepository.Migrate(); err != nil {
      log.Fatal(err)
   }

   person1 := Person{ID: 10, Name: "Williams", Age: 56}
   person2 := Person{ID: 20, Name: "Eliasson", Age: 52}

   createdPerson1, err := peopleRepository.Create(person1)
   if err != nil {
      log.Println(err)
   } else {
      fmt.Printf("createdPerson1: %+v\n", createdPerson1)
   }

   createdPerson2, err := peopleRepository.Create(person2)
   if err != nil {
      log.Println(err)
   } else {
      fmt.Printf("createdPerson2: %+v\n", createdPerson2)
   }

   all, err := peopleRepository.All()
   if err != nil {
      log.Fatal(err)
   }

   fmt.Printf("\nAll people:\n")
   for _, person := range all {
      fmt.Printf("person: %+v\n", person)
   }



}
