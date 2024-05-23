import React, { useState } from 'react'
import { calendar, currentMonth, currentYear, months } from '../classes/Calendar'
import "../css/calendar.css"
import FowardArrow from '../components/FowardArrow'
import BackwardArrow from '../components/BackwardArrow'

export default function Calendar() {
 const [year, setYear] = useState(currentYear)
 const [month, setMonth] = useState(currentMonth)
 let days = [<h1>No hay dias pa</h1>]
 if (calendar[year] != undefined) {
  if (calendar[year][months[month]] !== undefined) {
   const daysInMonth = calendar[year][months[month]]
   const daysHTML = daysInMonth.map(day => (
    <div className={day.type}>
     {day.day}
     {day.events.map(event => (<h1>{event}</h1>))}
    </div>
   ))
   days = (daysHTML)
  }
 }

 const daysContainer = <div className='month-days'>
  {days.map(day => (day))}
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
  </main>
 )
}

