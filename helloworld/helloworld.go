package main

const ( 
  english = "Hello, "
  spanish = "Hola, "
)
func Hello(name, language string) string {
  if name == "" {
    name = "world"
  } 
  
  return getPrefix(language)+ name 
}

func getPrefix(language string) string {
  switch language {
    case "spanish":
      return spanish
    default:
      return english
  }
}
