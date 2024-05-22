import React from 'react'
import { deleteTask } from '../functions/todoFunctions'

export default function DeleteButton(props) {
 return (
  <svg onClick={() => deleteTask(props.props)
  } className='trash-button' fill="#000000" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" width="24px" height="24px" viewBox="0 0 485 485" ><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <g> <g> <rect x="67.224" width="350.535" height="71.81"></rect> <path d="M417.776,92.829H67.237V485h350.537V92.829H417.776z M165.402,431.447h-28.362V146.383h28.362V431.447z M256.689,431.447 h-28.363V146.383h28.363V431.447z M347.97,431.447h-28.361V146.383h28.361V431.447z"></path> </g> </g> </g></svg >
 )
}

