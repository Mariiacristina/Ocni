package schema


//Schema del marciano! parte en Mayuscula pq es public
//tiene q tener el nombre de las columnas de la BDD
type Marciano struct {
    Id_marciano int `json:"id_marciano,omitempty"`
    Nombre string `json:"nombre,omitempty"`
}

type Nave struct{
  Id_nodriza int `json:"id_nodriza,omitempty"`
  Nombre string `json:"nombre,omitempty"`
}

type Aero struct {
  Id int `json:"id,omitempty"`
  Nombre string `json:"nombre,omitempty"`
  Nave_o string `json:"nave_o,omitempty"`
  Nave_d string `json:"nave_d,omitempty"`
  Cant_max int `json:"cant_max,omitempty"`
}

type Viaje struct {
  Id_Pasajero int `json:"id_pasajero,omitempty"`
  Id_Aero int `json:"id_aero,omitempty"`
  Fecha string `json:"fecha,omitempty"`
  Id_viaje int `json:"id_viaje,omitempty"`
  Estado int `json:"estado,omitempty"`
}

type Revision struct {
  Id_revision int `json:"id_revision,omitempty"`
  Nombre_r string `json:"nombre_r,omitempty"`
  Id_Aero int `json:"id_aero,omitempty"`
  Fecha string `json:"fecha,omitempty"`
}
