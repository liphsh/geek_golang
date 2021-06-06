import ( 
   "database/sql" 
   "fmt" 
 
   "github.com/pkg/errors" 
) 
 
func GetSql() error { 
   return errors.Wrap(sql.ErrNoRows, "GetSql failed") 
} 
 
func main() { 
   err := Call() 
   if errors.Cause(err) == sql.ErrNoRows { 
      fmt.Printf("data not found, %v\n", err) 
      fmt.Printf("%+v\n", err) 
      return 
   } 
   if err != nil { 
      // unknown error 
   } 
} 
