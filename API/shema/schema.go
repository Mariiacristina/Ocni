package schema


//Schema del marciano! parte en Mayuscula pq es public
type Marciano struct {
    id int   `json:"id,omitempty"`
    nombre string `json:"nombre,omitempty"`
}
