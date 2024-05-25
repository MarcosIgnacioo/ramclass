import React, { useState } from 'react'
import { Week, calendar, currentMonth, currentYear, getWeekDay, months, monthsEng, weekDays, weekDaysTraslation } from '../classes/Calendar'
import "../css/calendar.css"
import FowardArrow from '../components/FowardArrow'
import BackwardArrow from '../components/BackwardArrow'
import Day from '../components/Day'
import updateCurrentLocation from '../functions/location'

export default function Calendar() {
 updateCurrentLocation()
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

 // pasara de una interfa z a una clase paro
 let calendarRows: Array<Week> = []
 calendarRows.push({
  sunday: <Day day="Do" events={[]} type="week-title" />,
  monday: <Day day="Lu" events={[]} type="week-title" />,
  tuesday: <Day day="Ma" events={[]} type="week-title" />,
  wednesday: <Day day="Mi" events={[]} type="week-title" />,
  thursday: <Day day="Ju" events={[]} type="week-title" />,
  friday: <Day day="Vi" events={[]} type="week-title" />,
  saturday: <Day day="Sá" events={[]} type="week-title" />,
 }
 )

 if (calendar[year] != undefined) {
  if (calendar[year][months[month]] !== undefined) {
   // Obtenemos los dias del mes en el que estamos
   const daysInMonth = calendar[year][months[month]]
   // Los recorremos
   let obj: Week = {
    sunday: <Day day="" events={[]} type="" />,
    monday: <Day day="" events={[]} type="" />,
    tuesday: <Day day="" events={[]} type="" />,
    wednesday: <Day day="" events={[]} type="" />,
    thursday: <Day day="" events={[]} type="" />,
    friday: <Day day="" events={[]} type="" />,
    saturday: <Day day="" events={[]} type="" />,
   }

   daysInMonth.forEach((day, index) => {
    const weekDay = getWeekDay(year, day.day, month)
    if (weekDay === "sunday" && index !== 0) {
     calendarRows.push(obj)
     obj = {
      sunday: <Day day="" events={[]} type="" />,
      monday: <Day day="" events={[]} type="" />,
      tuesday: <Day day="" events={[]} type="" />,
      wednesday: <Day day="" events={[]} type="" />,
      thursday: <Day day="" events={[]} type="" />,
      friday: <Day day="" events={[]} type="" />,
      saturday: <Day day="" events={[]} type="" />,
     }
    }
    obj[weekDay] = (<Day day={day.day} events={day.events} type={day.type} />)
    if ((index + 1) === daysInMonth.length) {
     calendarRows.push({ ...obj })
    }
    daysWeek[weekDay].push(<Day day={day.day} events={day.events} type={day.type} />)
   });
  }
 }

 let daysContainer = <div className='month-days'>
  {weekDays.map((weekDay) => (
   <div className={`week-days ${weekDay}`}>
    <h1 className='week-day-title'>{weekDaysTraslation[weekDay]}</h1>
    {...daysWeek[weekDay] as Array<React.JSX.Element>}
   </div>
  ))}
 </div>
 // console.log(daysWeek)

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

 // // tener un arreglo con un objketop que tenga la interfaz week
 // //
 // daysContainer = calendarRows.map(row => {
 //  return (<div className={`week-days ${weekDay}`}>
 //   <div className='monday'>
 //    row[monday]
 //   </div>
 //   <div className='tuesday'>
 //    row[tuesday]
 //   </div>
 //   y asi se crearan cuadrados vacios en vasos d q no haya
 //   <h1 className='week-day-title'>{weekDaysTraslation[weekDay]}</h1>
 //   {...daysWeek[weekDay] as Array<React.JSX.Element>}
 //  </div>)
 // })

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


function getEmptyWeek() {
 return {
  sunday: <Day day="" events={[]} type="" />,
  monday: <Day day="" events={[]} type="" />,
  tuesday: <Day day="" events={[]} type="" />,
  wednesday: <Day day="" events={[]} type="" />,
  thursday: <Day day="" events={[]} type="" />,
  friday: <Day day="" events={[]} type="" />,
  saturday: <Day day="" events={[]} type="" />,
 }
}
