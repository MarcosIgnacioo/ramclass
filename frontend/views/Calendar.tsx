import React, { useState } from 'react'
import { Week, calendar, currentMonth, currentYear, getWeekDay, months } from '../classes/Calendar'
import "../css/calendar.css"
import FowardArrow from '../components/FowardArrow'
import BackwardArrow from '../components/BackwardArrow'
import Day from '../components/Day'
import updateCurrentLocation from '../functions/location'
import Title from '../components/Title'

export default function Calendar() {
 updateCurrentLocation()
 const [year, setYear] = useState(currentYear)
 const [month, setMonth] = useState(currentMonth)

 let calendarRows: Array<Week> = []
 calendarRows.push(getEmptyWeek("Do", "Lu", "Ma", "Mi",
  "Ju", "Vi", "Sá"))

 if (calendar[year] != undefined) {
  if (calendar[year][months[month]] !== undefined) {
   const daysInMonth = calendar[year][months[month]]
   let week: Week = getEmptyWeek()
   daysInMonth.forEach((day, index) => {
    const weekDay = getWeekDay(year, day.day, month)
    if (weekDay === "sunday" && index !== 0) {
     calendarRows.push(week)
     week = getEmptyWeek()
    }
    week[weekDay] = (<Day day={day.day} events={day.events} type={day.type} />)
    if ((index + 1) === daysInMonth.length) {
     calendarRows.push({ ...week })
    }
   });
  }
 }

 const rows = calendarRows.map(row => (
  <div className='week'>
   {row.sunday}
   {row.monday}
   {row.tuesday}
   {row.wednesday}
   {row.thursday}
   {row.friday}
   {row.saturday}
  </div>
 ))

 return (
  <main className='calendar-container'>
   <Title title='Calendario' to='#' />
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
   {...rows}
   <div className='simbology-container'>
    <h1 id='symbology-title'>Simbología</h1>
    <div className='simbology'>
     <div id='symbology-left'>
      <h1 className='begin'>Inicio de clases</h1>
      <h1 className='holyday'>Días festivos</h1>
      <h1 className='vacational'>Período vacacional</h1>
      <h1 className='blank'>Sin evento</h1>
      <h1 className='academic'>Asueto Académico</h1>
      <h1 className='school'>Día lectivo</h1>
      <h1 className='log_records'>Último día para capturas de actas</h1>
      <h1 className='end'>Fin de clases</h1>
     </div>
     <div className='symbology-right'>
      <h1 className='administrative'>Asueto de personal administrativo</h1>
      <h1 className='graduation'>Ceremonia de egreso</h1>
      <h1 className='ordinaries'>Exámenes ordinarios</h1>
      <h1 className='extraordinaries'>Exámenes extraordinarios</h1>
      <h1 className='inscriptions'>Inscripciones de nuevo ingreso</h1>
      <h1 className='reinscriptions'>Reinscripciones <br />/ Cambios de situacion escolar</h1>
      <h1 className='new_admission_call'>Convocatoria de nuevo ingreso</h1>
     </div>
    </div>
   </div>
  </main>
 )
}


function getEmptyWeek(sunday = "", monday = "", tuesday = "", wednesday = "", thursday = "", friday = "", saturday = "") {
 return {
  sunday: <Day day={sunday} events={[]} type="week-day-container" />,
  monday: <Day day={monday} events={[]} type="week-day-container" />,
  tuesday: <Day day={tuesday} events={[]} type="week-day-container" />,
  wednesday: <Day day={wednesday} events={[]} type="week-day-container" />,
  thursday: <Day day={thursday} events={[]} type="week-day-container" />,
  friday: <Day day={friday} events={[]} type="week-day-container" />,
  saturday: <Day day={saturday} events={[]} type="week-day-container" />,
 }
}
