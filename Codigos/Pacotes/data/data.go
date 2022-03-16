package data

import(
  "errors"
)

type Date struct{ // letra minuscula = private, letra maiuscula = public
  year int
  month int
  day int
}
//Getters
func (d *Date)Year()int{
  return d.year
}

func (d *Date)Month()int{
  return d.month
}

func (d *Date)Day()int{
  return d.day
}

//Setters
func (d *Date) DateConstrutor(year int, month int, day int) error{
  d.SetYear(year)
  err := d.SetMonth(month)
  if err != nil{
    return err
  }
  d.SetDay(day)

  return nil
}

func (d *Date) SetYear(year int){
  d.year = year
}

func (d *Date) SetMonth(month int)error{
  if(month > 12 || month < 1){
    return errors.New("Mês inválido.")
  }
  d.month = month
  return nil
}

func (d *Date) SetDay(day int){
  d.day = day
}