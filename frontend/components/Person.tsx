import React from 'react'

export default function Person(props) {
 const { index, dragPerson, draggedOverPerson, person, handleSort } = props
 return (
  <div className="relative flex space-x-3 border rounded p-2 bg-gray-100"
   draggable
   onDragStart={() => (dragPerson.current = index)}
   onDragEnter={() => {
    console.log(index)
    draggedOverPerson.current = index
   }
   }
   onDragEnd={handleSort}
   onDragOver={(e) => e.preventDefault()}
  >
   <p>{person.name}</p>
  </div>
 )
}

