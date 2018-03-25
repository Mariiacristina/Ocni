package schema


//Schema del marciano! parte en Mayuscula pq es public
type Marciano struct {
    Id int   `json:"id,omitempty"`
    Nombre string `json:"nombre,omitempty"`
}
