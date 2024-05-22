import React from 'react'
import Message from './Message'

export default function MessagePopUp(props) {
 const { savingTask, successMessageRef, errorMessageRef } = props
 if (savingTask.isSuccess) {
  return <Message ref={successMessageRef} message="Se guardó con éxito" class="message success" />
 }
 if (savingTask.isError) {
  return <Message ref={errorMessageRef} message="Ocurrió un error innesperado" class="message error" />
 }
}

