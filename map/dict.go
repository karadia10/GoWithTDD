package dict

type Dictionary map[string]string
const(
ErrNotFound=DictErr("key not found") 
ErrWordExist=DictErr("key already exists")
)
type DictErr string

func (d DictErr) Error() string {
  return string(d)
}

func (d Dictionary) Search(key string) (string, error) {
  def, ok := d[key]
  if !ok {
    return "", ErrNotFound 
  }
  return def, nil
}

func (d Dictionary) Add(key, value string) error {
  _, err:=d.Search(key)
  switch err{
  case ErrNotFound:
    d[key]=value
  case nil:
    return ErrWordExist
  default:
  return err
  }
  return nil
  
  
}

func (d Dictionary) Update(key, value string) error {
  _, err := d.Search(key)
  switch err {
  case ErrNotFound:
    return err
  default:
    d[key] = value
  }
  return nil
}

func (d Dictionary) Delete(key string) error {
  _, err := d.Search(key)
  if err != nil {
    return ErrNotFound
  }
  delete(d, key)
  return nil
}
