import React, { useState } from 'react'
import { calendar, currentMonth, currentYear, getWeekDay, months, weekDays, weekDaysTraslation } from '../classes/Calendar'
import "../css/calendar.css"
import FowardArrow from '../components/FowardArrow'
import BackwardArrow from '../components/BackwardArrow'
import Day from '../components/Day'
import updateCurrentLocation from '../functions/location'

export default function Calendar() {
 updateCurrentLocation()
 const asdf = new Date()
 const [year, setYear] = useState(currentYear)
 const [month, setMonth] = useState(currentMonth)
 const daysWeek = {
  sunday: [],
  monday: [],
  tuesday: [],
  wednesday: [],
  thursday: [],
  friday: [],
  saturday: [],
 }

 // let days = [<h1>No hay dias pa</h1>]
 if (calendar[year] != undefined) {
  if (calendar[year][months[month]] !== undefined) {
   const daysInMonth = calendar[year][months[month]]
   daysInMonth.forEach(day => {
    const weekDay = getWeekDay(year, day.day, month)
    daysWeek[weekDay].push(<Day day={day.day} events={day.events} type={day.type} />)
   });
  }
 }

 // const daysHTML = daysInMonth.map(day => {
 //  console.log()
 //  return <Day day={day.day} events={day.events} type={day.type} />
 // })
 // days = (daysHTML)

 const daysContainer = <div className='month-days'>
  {weekDays.map((weekDay) => (
   <div className={weekDay}>
    <h1>{weekDaysTraslation[weekDay]}</h1>
    {...daysWeek[weekDay] as Array<React.JSX.Element>}
   </div>
  ))}
 </div>

 return (
  <main className='calendar-container'>
   <div className='year-container'>
    <BackwardArrow thing={year} setThing={setYear} high={9999} low={0} />
    <h2>{year}</h2>
    <FowardArrow thing={year} setThing={setYear} high={9999} low={0} />
   </div>
   <div className='month-container'>
    <BackwardArrow thing={month} setThing={setMonth} high={11} low={0} />
    <h2>{months[month]}</h2>
    <FowardArrow thing={month} setThing={setMonth} high={11} low={0} />
   </div>
   {daysContainer}
   <div className='simbology'>
    <h1 id='symbology-title'>Simbología</h1>
    <h1 className='begin'>Inicio de clases</h1>
    <h1 className='holyday'>Días festivos</h1>
    <h1 className='vacational'>Período vacacional</h1>
    <h1 className='blank'>Sin evento</h1>
    <h1 className='academic'>Asueto Académico</h1>
    <h1 className='school'>Día lectivo</h1>
    <h1 className='administrative'>Asueto de personal administrativo</h1>
    <h1 className='graduation'>Ceremonia de egreso</h1>
    <h1 className='ordinaries'>Exámenes ordinarios</h1>
    <h1 className='extraordinaries'>Exámenes extraordinarios</h1>
    <h1 className='inscriptions'>Inscripciones de nuevo ingreso</h1>
    <h1 className='reinscriptions'>Reinscripciones / Cambios de situacion escolar</h1>
    <h1 className='new_admission_call'>Convocatoria de nuevo ingreso</h1>
    <h1 className='log_records'>Último día para capturas de actas</h1>
    <h1 className='end'>Fin de clases</h1>
   </div>
  </main>
 )
}

