package main

import "os"
import "fmt"

type AutoGenerated struct {
	Packages struct {
		Bootstrap struct {
			Default struct {
				Download []struct {
					URL  string `json:"url"`
					Type string `json:"type"`
				} `json:"download"`
			} `json:"default"`
		} `json:"bootstrap"`
	} `json:"packages"`
}

func main() {

	//obtener argumentos enviados por consola
    argumentos := os.Args
	
	switch accion:= argumentos[0]{
		//creamos las carpetas y archivos necesarios para el proyecto
		case "init":
			
		default:
			
	}
}

func downloadFile(filepath string, url string) (err error) {

  // Crea el archivo
  out, err := os.Create(filepath)
  if err != nil  {
    return err
  }
  defer out.Close()

  // Obtiene al archivo
  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  // Verifica la rspuesta del servidor
  if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("bad status: %s", resp.Status)
  }

  // Escribe el archivo con el contenido obtenido
  _, err = io.Copy(out, resp.Body)
  if err != nil  {
    return err
  }

  return nil
}